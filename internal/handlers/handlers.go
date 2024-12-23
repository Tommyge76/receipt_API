package handlers

import (
	"log"
	"encoding/json"
	"net/http"
	"io"
	"process_receipts/internal/request_models"
	"process_receipts/internal/database"
	"database/sql"
	"strconv"
	"process_receipts/internal/utils"
)

func HandleAddReceipt(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var receipt request_models.Receipt
		if err := json.Unmarshal(body, &receipt); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if !utils.ValidateReceipt(receipt) {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{
				"description": "The receipt is invalid",
			}
		
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)			
			return
		}

		id := database.AddReceipt(db, receipt)

		response := map[string] string {
			"id": id,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func HandleGetReceiptById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/receipts/") : len(r.URL.Path)-len("/points")]
		log.Println("id: ", id)
		receipt, err := database.GetReceiptById(db, id)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			response := map[string]string{
				"description": "No receipt found for that id",
			}
		
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)			
			return
		}

		response := map[string]string{
			"points": strconv.Itoa(receipt.Points),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
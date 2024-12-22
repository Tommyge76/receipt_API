package database

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"process_receipts/internal/request_models"
	"process_receipts/internal/utils"
	"process_receipts/internal/response_models"
)

func CreateDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./receipts.db")
	if err != nil {
		log.Fatal("Error creating database: ", err)
	}

	CreateReceiptsTable(db)
	CreateItemsTable(db)
	return db
}

func CreateReceiptsTable(db *sql.DB) {
	createTable := `
		CREATE TABLE IF NOT EXISTS receipts (
			id TEXT PRIMARY KEY,
			retailer TEXT,
			purchase_date TEXT,
			purchase_time TEXT,
			total TEXT,
			points INTEGER
		)
	`
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Receipts table created successfully")
}

func CreateItemsTable(db *sql.DB) {
	createTable := `
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			receipt_id TEXT,
			short_description TEXT,
			price TEXT,
			FOREIGN KEY (receipt_id) REFERENCES receipts(id)
		)
	`
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Items table created successfully")
}

func AddReceipt(db *sql.DB, receipt request_models.Receipt) string { 
	id := utils.GenerateUUID()
	points := utils.CalculatePoints(receipt)
	insertReceipt := `INSERT INTO receipts (id, retailer, purchase_date, purchase_time, total, points) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(insertReceipt, id, receipt.Retailer, receipt.PurchaseDate, receipt.PurchaseTime, receipt.Total, points)
	for _, item := range receipt.Items {
		AddItem(db, item, id)
	}

	if err != nil {
		log.Println("Error inserting receipt: ", err)
	}

	return id
}

func AddItem(db *sql.DB, item request_models.Item, receiptID string) { 
	insertItem := `INSERT INTO items (receipt_id, short_description, price) VALUES (?, ?, ?)`
	_, err := db.Exec(insertItem, receiptID, item.ShortDescription, item.Price)

	if err != nil {
		log.Println("Error inserting item: ", err)
	}
}

func GetReceiptById(db *sql.DB, id string) (response_models.Receipt, error) {
	getReceipt := `SELECT * FROM receipts WHERE id = ?`
	row := db.QueryRow(getReceipt, id)

	var receipt response_models.Receipt
	err := row.Scan(&receipt.Id, &receipt.Retailer, &receipt.PurchaseDate, &receipt.PurchaseTime, &receipt.Total, &receipt.Points)
	if err != nil {
		log.Println("Error getting receipt by id: ", err)
		return receipt, err
	}
	return receipt, nil
}

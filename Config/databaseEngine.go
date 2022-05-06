package Config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func(c ConfigSettingSql) InitDB () {
	db, err := sql.Open(string(DATABASE_MAIN), "FITA.db")
	if err != nil {
		log.Println(err)
	}

	if db != nil {
		SqlConnection = db
		err = migrate(db)
		if err != nil{
			log.Println(err)
		}
		log.Println("ENGINE "+ DATABASE_MAIN +" start....")
	}
}

func migrate(db *sql.DB) (err error) {
	var execTable []string
	queryTCart := `CREATE TABLE IF NOT EXISTS t_cart(
    id_cart INTEGER NOT NULL PRIMARY KEY,
	sku VARCHAR(50),
    name VARCHAR(50),
    qty INTEGER)`

	query := `CREATE TABLE IF NOT EXISTS t_items(
    sku VARCHAR NOT NULL PRIMARY KEY, 
    name VARCHAR(50), 
    price NUMERIC, 
    inventory_qty INTEGER);

	INSERT INTO t_items(sku, name, price, inventory_qty) VALUES("120P90", "Google Home", 49.99, 10);
	INSERT INTO t_items(sku, name, price, inventory_qty) VALUES("43N23P", "Macbook Pro", 5399.99, 5);
	INSERT INTO t_items(sku, name, price, inventory_qty) VALUES("A304SD", "Alexa Speaker", 109.50, 10);
	INSERT INTO t_items(sku, name, price, inventory_qty) VALUES("234234", "Raspberry Pi B", 30.00, 2)`

	execTable = append(execTable, queryTCart, query)
	for _, v := range execTable {
		_, err = db.Exec(v)
		if err != nil {
			return
		}
	}
	return
}

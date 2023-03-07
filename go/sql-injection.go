package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}
	sts := `
DROP TABLE IF EXISTS cars;
CREATE TABLE cars(id INTEGER PRIMARY KEY, name TEXT, price INT);
INSERT INTO cars(name, price) VALUES('Audi',52642);
INSERT INTO cars(name, price) VALUES('Mercedes',57127);
INSERT INTO cars(name, price) VALUES('Skoda',9000);
INSERT INTO cars(name, price) VALUES('Volvo',29000);
INSERT INTO cars(name, price) VALUES('Bentley',350000);
INSERT INTO cars(name, price) VALUES('Citroen',21000);
INSERT INTO cars(name, price) VALUES('Hummer',41400);
INSERT INTO cars(name, price) VALUES('Volkswagen',21600);
`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table cars created")
	fmt.Println()

	defer db.Close()

	fmt.Println("Simple Query - All data")
	rows, err := db.Query("SELECT * FROM cars")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id int
		var name string
		var price int

		err = rows.Scan(&id, &name, &price)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %s %d\n", id, name, price)
	}

	defer db.Close()

	fmt.Println()
	fmt.Println("Vulnerable Query - Get all data instead 1")
	nameInj := "1 OR 1=1"
	stmInj, err := db.Query("SELECT * FROM cars WHERE name = " + nameInj)

	if err != nil {
		log.Fatal(err)
	}

	defer stmInj.Close()
	for stmInj.Next() {

		var id int
		var name string
		var price int

		err = stmInj.Scan(&id, &name, &price)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %s %d\n", id, name, price)
	}

	defer db.Close()
	fmt.Println()

	fmt.Println("Not Vulnerable Query - Get right data")
	stm, err := db.Query("SELECT * FROM cars WHERE name = ?", "0 OR 1=1")

	if err != nil {
		log.Fatal(err)
	}

	defer stm.Close()
	for stm.Next() {

		var id int
		var name string
		var price int

		err = stm.Scan(&id, &name, &price)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %s %d\n", id, name, price)
	}

	defer db.Close()

}

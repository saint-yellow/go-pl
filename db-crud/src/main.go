package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	server = "localhost"
	port = 3306
	user = "root"
	password = "saint-yellow"
	database = "beauty_photography"
)

func connect() *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", user, password, server, port, database)
	fmt.Println(connectionString)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}

	context := context.Background()

	err = db.PingContext(context)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Connected!")

	return db
}

func disconnect(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Disconnected.")
}

type photoSet struct {
	id int
	title string
	author string
	datetimePublished string
	description string
}

func query(db *sql.DB) {
	results, err := db.Query("select id, title, author, description, datetime_published from photo_sets limit 10;")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var p photoSet
		err = results.Scan(&p.id, &p.title, &p.author, &p.description, &p.datetimePublished)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(p)
	}
}

func main() {
	db := connect()
	query(db)
	disconnect(db)
}
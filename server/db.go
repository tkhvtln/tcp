package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
	dataSource = "root:ilnur@/contacts"
)

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatalf("error pinging database: %v", err)
	}

	return nil
}

func getUserInfo(idUser int) (string, error) {
	userFind := user{}

	err := db.QueryRow("select * from contacts.users where id = ?", idUser).Scan(&userFind.id, &userFind.name, &userFind.phoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no user found with id: %v", idUser)
		}
		return "", fmt.Errorf("error querying database: %v", err)
	}

	userInfo := fmt.Sprintf("%v - %v\n", userFind.name, userFind.phoneNumber)
	return userInfo, nil
}

func getAllUsers() (string, error) {
	rows, err := db.Query("select id, name from contacts.users")
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("database is empty")
		}

		return "", fmt.Errorf("error querying database %v", err)
	}
	defer rows.Close()

	userList := ""
	for rows.Next() {
		u := user{}
		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			log.Printf("error scanning row: %v\n", err)
			continue
		}

		userList += fmt.Sprintf("%v. %v\n", u.id, u.name)
	}

	if rows.Err() != nil {
		return "", fmt.Errorf("error reading rows: %v", rows.Err())
	}

	return userList, nil
}

func closeDB() {
	err := db.Close()
	if err != nil {
		log.Fatalf("error closing database: %v", err)
	}
}

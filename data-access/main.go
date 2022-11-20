package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	ID           int64
	Name         string
	WishList     []string
	RejectPerson string
}

func UserByName(name string, db *sql.DB) ([]User, error) {
	// A Users slice to hold data from returned rows.
	var users []User
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected again")

	rows, err := db.Query("SELECT * FROM secretSanta WHERE name = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, pq.Array(&user.WishList), &user.RejectPerson); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return users, nil
}

func addUser(user User, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO secretSanta (name, wishList, rejectPerson) VALUES ($1, $2, $3)", user.Name, pq.Array(user.WishList), user.RejectPerson)
	if err != nil {
		return fmt.Errorf("addUser: %v", err)
	}
	if err != nil {
		return fmt.Errorf("addUser: %v", err)
	}
	return nil
}

func main() {
	var db *sql.DB
	// Capture connection properties.
	connStr := "user=tonyprestifilippo dbname=tonyprestifilippo sslmode=disable"
	//Open connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	Users, err := UserByName("Mario", db)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range Users {
		fmt.Println(user.ID, user.Name, user.WishList)
	}
	Jessica := User{
		Name:         "Jessica",
		WishList:     []string{"dog", "boots"},
		RejectPerson: "Antonio",
	}
	if err := addUser(Jessica, db); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Album added: %v\n", id)
	jessica, err := UserByName("Jessica", db)
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range jessica {
		fmt.Println(j)
	}
}

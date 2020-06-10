package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/common/log"
)

func main() {
	db, err := sql.Open("mysql", "root:Aa123456@/dbusers?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new user
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	db.Exec(`insert into users (username, password, created_at) 
		values(?, ?, ?)`, username, password, createdAt)

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := `select id, username, password, created_at
			from users
			where id=?`

		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`
			select id, username, password, created_at
			from users`)

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{ // Delete
		_, err := db.Exec(`
			delete 
			from users 
			where id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}

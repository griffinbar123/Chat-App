package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	// var err error
	Db, _ = sql.Open("mysql", "root:rootboot@tcp(127.0.0.1:3306)/chat")
}

func (s *uService) InsertUser(newUser *FUsers) error {
	_, err := Db.Exec("insert into FUsers (Email, Username, PasswordH) values (?, ?, ?)", newUser.Email, newUser.Username, newUser.PasswordH)
	return err
}

func (s *uService) CheckUser(newUser *Users) (string, error) {
	var x string
	row := Db.QueryRow("select passwordh from FUsers where email = ?", newUser.Email)
	err := row.Scan(&x)
	return x, err
}

func (msg *Message) CreateMessage() error {
	_, err := Db.Exec("insert into Message (Username, Content) values (?, ?)", msg.Username, msg.Content)
	for c := range Connections {
		ch := Connections[c]
		ch.Writer(msg)
	}

	return err
}
func CheckEmail(email string) (err error) {
	var x string
	row := Db.QueryRow("select email from FUsers where email = ?", email)
	err = row.Scan(&x)
	return
}

func GetAllMessages() []Message {
	var msgs []Message
	rows, err := Db.Query("select Username, Content from Message")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	msg := Message{}
	for rows.Next() {
		err := rows.Scan(&msg.Username, &msg.Content)
		if err != nil {
			fmt.Println(err)
		}
		msgs = append(msgs, msg)
	}
	return msgs

}

package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

func Checksession(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session")
	email, ok := session.Values["Email"]
	err := CheckEmail(email.(string))
	if err != nil {
		ok = false
	}
	return ok
}
func FindUsername(r *http.Request) (string, bool) {
	session, _ := store.Get(r, "session")
	user, ok := session.Values["Username"]
	return user.(string), ok
}

func Messages(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("templates/other.html")
	if r.Method == "GET" {
		ok := Checksession(w, r)
		if !ok {
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
		}
	}
	temp.Execute(w, "msgs")
}

func remove(s []*Connection, i int) []*Connection {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func reader(conn *Connection, r *http.Request) {
	for {
		_, p, err := conn.ws.ReadMessage()
		if err != nil {
			fmt.Println("error reading message: ", err)
			return
		}
		if string(p) == "close" {
			for i, con := range Connections {
				if con == conn {
					Connections = remove(Connections, i)
				}
			}
		}
		usrname, _ := FindUsername(r)
		// fmt.Println(usrname, string(p))
		msg := Message{
			Username: usrname,
			Content:  string(p),
		}
		err = msg.CreateMessage()
		if err != nil {
			fmt.Println(err)
		}

	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Connection struct {
	// The websocket connection.
	ws *websocket.Conn
}

func (wss *Connection) Writer(msg *Message) {
	conn := wss.ws
	conn.WriteJSON(*msg)
}

var Connections []*Connection

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
	}

	ws.WriteJSON(Message{
		Id: -1,
	})
	msgs := GetAllMessages()
	for _, msg := range msgs {
		ws.WriteJSON(msg)
	}

	fmt.Println("connected")

	conn := Connection{
		ws: ws,
	}
	Connections = append(Connections, &conn)
	go reader(&conn, r)
	// go Writer(ws)
}

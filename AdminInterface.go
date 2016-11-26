package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("pages/admin.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}
func wsAdminHandler(c chan tickerMessage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received Admin WS request")
		if r.Header.Get("Origin") != "http://"+r.Host {
			http.Error(w, "Origin not allowed", 403)
			return
		}
		conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}
		sendInitData(conn)
		go receiveFromUser(conn, c)
	}
}
func sendInitData(conn *websocket.Conn) {
	if err := conn.WriteJSON(tickerMessage{Type: "State", State: currentTicker.State}); err != nil {
		fmt.Println(err)
	}
}
func receiveFromUser(conn *websocket.Conn, c chan tickerMessage) {
	var m tickerMessage
	for {
		if err := conn.ReadJSON(&m); err != nil {
			if ce, ok := err.(*websocket.CloseError); ok {
				switch ce.Code {
				case websocket.CloseNormalClosure,
					websocket.CloseGoingAway,
					websocket.CloseNoStatusReceived:
					fmt.Println(err)
					return
				}
			}
		}
		c <- m
	}
}

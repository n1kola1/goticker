package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

func wsTickerHandler(c chan tickerMessage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received WS request")
		if r.Header.Get("Origin") != "http://"+r.Host {
			http.Error(w, "Origin not allowed", 403)
			return
		}
		conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}

		sendInitData(conn)

		go sendToTicker(conn, c)
	}
}

func sendToTicker(conn *websocket.Conn, c chan tickerMessage) {
	for {
		m := <-c
		fmt.Printf("Received message %v from admin\n", m)
		if err := conn.WriteJSON(m); err != nil {
			if ce, ok := err.(*websocket.CloseError); ok {
				switch ce.Code {
				case websocket.CloseNormalClosure,
					websocket.CloseGoingAway,
					websocket.CloseNoStatusReceived:
					fmt.Println(err)
					return
				}
				fmt.Println(err)
			}
		}
		fmt.Println("written")
		currentTicker.State.PreviousEvents = append(currentTicker.State.PreviousEvents, m.Event)
		saveTicker(currentTicker)
	}
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("pages/index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

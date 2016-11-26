package main

import (
	"net/http"

	"gokik"
)

var (
	currentTicker *ticker
)

func main() {

	/*ad := gokik.AuthData{User: "rcaticker", Pass: "694494c6-a5e2-412f-88a1-98ea99626877"}
	bc := gokik.BotConfig{Webhook: "http://rcatickertest.hopto.org/incoming",
		Features: gokik.BotFeatures{
			ReceiveReadReceipts:      false,
			ReceiveIsTyping:          false,
			ManuallySendReadReceipts: false,
			ReceiveDeliveryReceipts:  false}}
	gokik.Configure(ad, bc)*/
	currentTicker = loadTicker(1)
	c := make(chan tickerMessage)
	http.HandleFunc("/tickerws", wsTickerHandler(c))
	http.HandleFunc("/adminws", wsAdminHandler(c))
	http.HandleFunc("/config", configHandler)
	http.HandleFunc("/incoming", gokik.MessageHandler)
	http.HandleFunc("/", webHandler)
	fs := http.FileServer(http.Dir("pages"))
	http.Handle("/pages/", http.StripPrefix("/pages/", fs))

	http.ListenAndServe(":8080", nil)

}

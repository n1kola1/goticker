package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	tickerDir = ".\\ticker\\"
)

type (
	tickerMessage struct {
		Type  string
		State currentState
		Event tickerEvent
	}

	currentState struct {
		StartTime            time.Time
		PreviousEvents       []tickerEvent
		HomeTeam, AwayTeam   string
		HomeScore, AwayScore uint
		CurrentHalf          uint
		RunningTime          time.Duration
	}

	tickerEvent struct {
		Name      string
		Text      string
		Timestamp int64
	}
	ticker struct {
		State    currentState
		TickerID uint
		File     *os.File
	}
)

func newTicker() *ticker {
	if !exists(tickerDir) {
		os.Mkdir(tickerDir, os.ModeDir)
	}
	i := uint(1)
	for ; exists(fmt.Sprintf("%s\\%v.txt", tickerDir, i)); i++ {
	}
	f, err := os.Create(fmt.Sprintf("%s\\%v.txt", tickerDir, i))
	if err != nil {
		panic(err)
	}
	return &ticker{
		File:     f,
		TickerID: i,
		State:    currentState{PreviousEvents: make([]tickerEvent, 0)}}
}

func loadTicker(id uint) *ticker {
	var t ticker
	f, err := os.Open(fmt.Sprintf("%s\\%v.txt", tickerDir, id))
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(f)

	if err = jsonParser.Decode(&t); err != nil {
		panic("error parsing ticker file " + err.Error())
	}
	t.File = f
	return &t
}

func saveTicker(t *ticker) {
	f, err := os.OpenFile(fmt.Sprintf("%s\\%v.txt", tickerDir, t.TickerID), os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	jsonEncoder := json.NewEncoder(f)
	if err = jsonEncoder.Encode(t); err != nil {
		panic("error parsing ticker file " + err.Error())
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

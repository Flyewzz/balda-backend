package main

import (
	"errors"

	"github.com/gorilla/websocket"
)

type Player struct {
	Nickname     string `json:"nickname"`
	connection   *websocket.Conn
	room         *Room
	chanMessages chan Message
}

func newPlayer(nickName string, ws *websocket.Conn) *Player {
	return &Player{
		Nickname:     nickName,
		connection:   ws,
		room:         nil,
		chanMessages: make(chan Message),
	}
}

func SendDataToPlayer(player *Player, message []byte) error {
	if player.connection == nil {
		return errors.New("Player is not connected")
	}
	return player.connection.WriteMessage(websocket.TextMessage, message)
}

package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gorilla/websocket"
)

type Player struct {
	Nickname     string `json:"nickname"`
	connection   *websocket.Conn
	room         *Room
	chanMessages chan Message
	finish       chan struct{}
}

func newPlayer(nickName string, ws *websocket.Conn) *Player {
	return &Player{
		Nickname:     nickName,
		connection:   ws,
		room:         nil,
		chanMessages: make(chan Message),
		finish:       make(chan struct{}),
	}
}

func SendDataToPlayer(player *Player, message []byte) error {
	if player.connection == nil {
		return errors.New("Player is not connected")
	}
	return player.connection.WriteMessage(websocket.TextMessage, message)
}

func (p *Player) DisconnectPlayer() {
	defer func() {
		if r := recover(); r != nil {
			//...)
		}
	}()
	p.finish <- struct{}{}
}

func (p *Player) Listen() {
	defer p.DisconnectPlayer()
	go func() {
		for {
			select {
			case <-p.finish:
				close(p.finish)
				p.room.RemovePlayer(p)
				return
			case message := <-p.chanMessages:
				log.Printf("Got message with type '%s' from player '%s'\n", message.Type, p.Nickname)
			}
		}
	}()

	for {
		_, data, err := p.connection.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		var msg Message
		err = json.Unmarshal(data, &msg)
		if err != nil {
			log.Println("error with unmarshalling:", err)
			break
		}
		p.chanMessages <- msg
	}

}

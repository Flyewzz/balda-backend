package main

import (
	"errors"
	"log"
	"math/rand"
	"sync"
)

type Room struct {
	FieldSize  int32
	turn       string
	Players    map[string]*Player
	register   chan *Player
	unregister chan *Player
	init       chan struct{}
	isStarted  bool
	finish     chan struct{}
	mtx        *sync.Mutex
}

func NewRoom(size int32) *Room {
	return &Room{
		Players:    make(map[string]*Player),
		register:   make(chan *Player),
		unregister: make(chan *Player),
		init:       make(chan struct{}, 1),
		finish:     make(chan struct{}, 1),
		FieldSize:  size,
		isStarted:  false,
		mtx:        &sync.Mutex{},
	}
}

func (room *Room) AddPlayer(player *Player) {
	room.register <- player
}

func (room *Room) RemovePlayer(player *Player) {
	if player.room == nil {
		log.Printf("Player '%s' is not in any rooms\n", player.Nickname)
		return
	}
	log.Printf("Removing player '%s'...\n", player.Nickname)
	room.unregister <- player
	log.Printf("Player '%s' was removed!\n", player.Nickname)
}

func (r *Room) GetRandomTurn() (string, error) {
	if len(r.Players) == 0 {
		return "", errors.New("The room is empty")
	}
	var nicknames []string

	for nick := range r.Players {
		nicknames = append(nicknames, nick)
	}

	return nicknames[rand.Intn(len(nicknames))], nil

}

func (r *Room) SetTurn(nickname string) {
	r.turn = nickname
}

func (room *Room) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic was occured:", r)
		}
	}()
	for {
		select {
		case <-room.finish:
			for _, player := range room.Players {
				SendDataToPlayer(player, []byte("Game was finished!"))
			}
			room.Deactivate()
			log.Println("Game was finished!")
		case <-room.init:
			// TODO Initial game
			if !room.isStarted {
				room.Activate()
				log.Println("Game was started")
			} else {
				log.Println("Room is already started")
			}
		case player := <-room.register:
			switch len(room.Players) {
			case 0:
				room.mtx.Lock()
				room.Players[player.Nickname] = player
				player.room = room
				room.mtx.Unlock()
			case 1:
				room.mtx.Lock()
				room.Players[player.Nickname] = player
				player.room = room
				room.mtx.Unlock()
				room.init <- struct{}{}
			}
		case player := <-room.unregister:
			room.mtx.Lock()
			delete(room.Players, player.Nickname)
			player.room = nil
			player.DisconnectPlayer()
			room.mtx.Unlock()
			if room.isStarted {
				room.finish <- struct{}{}
			}
		default:
			continue
		}
	}
}

//******** Room start/finish control *******
func (r *Room) Activate() {
	r.isStarted = true
	for _, p := range r.Players {
		go p.Listen()
	}
}

func (r *Room) Deactivate() {
	r.isStarted = false
	for _, p := range r.Players {
		p.connection.Close()
	}
}

//*******************************************

func (r *Room) StartGame() {

}

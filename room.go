package main

import (
	"log"
	"sync"
)

type Room struct {
	FieldSize  int32
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
	room.unregister <- player
}

func RemovePlayer(player *Player) {

	if player.room == nil {
		log.Printf("Player '%s' is not in any rooms\n", player.Nickname)
		return
	}
	log.Printf("Removing player '%s'...\n", player.Nickname)
	// player.messagesClose <- struct{}{}
	player.room.unregister <- player
	log.Printf("Player '%s' was removed!\n", player.Nickname)
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
				room.RemovePlayer(player)
			}
		case <-room.init:
			// TODO Initial game
			if !room.isStarted {
				go func() {

					<-room.finish
				}()
				room.isStarted = true
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
		default:
			continue
		case player := <-room.unregister:
			room.mtx.Lock()
			delete(room.Players, player.Nickname)
			if room.isStarted {
				room.finish <- struct{}{}
			}
			player.room = nil
			room.mtx.Unlock()
		}
	}
}
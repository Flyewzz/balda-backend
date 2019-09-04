package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Game struct {
	MaxRooms uint
	rooms    []*Room
	mtx      sync.Mutex
	register chan *Player
}

func NewGame(maxRooms uint) *Game {
	return &Game{
		register: make(chan *Player),
	}
}

func (g *Game) Run() {
	defer func() {
		if e := recover(); e != nil {
			log.Println("Error at game.Run was occured (function Run)", e)
		}
	}()
	log.Println("Main loop started")
LOOP:
	for {
		player := <-g.register

		for _, room := range g.rooms {
			if len(room.Players) < 2 && !room.isStarted { // A room must not be started
				g.mtx.Lock()
				room.AddPlayer(player)
				g.mtx.Unlock()
				continue LOOP
			}
		}

		room := NewRoom(5) // Field 5x5
		g.mtx.Lock()
		g.AddRoom(room)
		g.mtx.Unlock()
		go room.Run()
		room.AddPlayer(player)
	}
}

func (g *Game) AddPlayer(player *Player) {
	log.Printf("Player '%s' added to the game\n", player.Nickname)
	// if viper.ConfigFileUsed() != "../config/test.yml" {
	// 	// go player.Listen()
	// }
	g.register <- player
}

func (g *Game) AddRoom(room *Room) {
	g.rooms = append(g.rooms, room)
}

func (g *Game) RemoveRoom(room *Room) error {

	rooms := &g.rooms
	g.mtx.Lock()
	lastIndex := len(*rooms) - 1
	for index, r := range g.rooms {
		if r == room {

			(*rooms)[index] = (*rooms)[lastIndex]
			g.rooms = (*rooms)[:lastIndex]

			fmt.Println("The room was deleted")
			g.mtx.Unlock()
			return nil
		}
	}
	g.mtx.Unlock()
	return errors.New("The room is not found")
}

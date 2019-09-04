package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options
var game *Game

var nickCounter int = 0

func echo(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	log.Println("*NEW CONNECTION*")
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	// defer connection.Close()
	// Create a new player and give him a new connection
	player := newPlayer(strconv.Itoa(nickCounter), connection)
	nickCounter++
	err = connection.WriteMessage(websocket.TextMessage, []byte("Your new nickname: "+player.Nickname))
	if err != nil {
		log.Println("write:", err)
		return
	}

	game.AddPlayer(player)
}

func main() {

	// viper.SetConfigFile(os.Args[1])
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Println("Cannot read config", err)
	// 	return
	// }
	// port := viper.GetString("game.port")
	// // hostAuth := viper.GetString("authsrv.host") + ":" + viper.GetString("authsrv.port")

	// game = NewGame(viper.GetUint("game.countRoom"))

	rand.Seed(time.Now().UnixNano())
	port := "8080"

	game = NewGame(1000)
	http.HandleFunc("/ws", echo)

	go game.Run()
	http.ListenAndServe(":"+port, nil)
}

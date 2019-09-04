package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
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
	defer connection.Close()
	// Create a new player and give him a new connection
	player := newPlayer(strconv.Itoa(nickCounter), connection)
	nickCounter++
	err = connection.WriteMessage(websocket.TextMessage, []byte("Your new nickname: "+player.Nickname))
	if err != nil {
		log.Println("write:", err)
		return
	}

	game.AddPlayer(player)

	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
	}
}

func main() {

	viper.SetConfigFile(os.Args[1])
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Cannot read config", err)
		return
	}
	port := viper.GetString("game.port")
	// hostAuth := viper.GetString("authsrv.host") + ":" + viper.GetString("authsrv.port")

	http.HandleFunc("/ws", echo)

	game = NewGame(viper.GetUint("game.countRoom"))
	go game.Run()
	http.ListenAndServe(":"+port, nil)
}

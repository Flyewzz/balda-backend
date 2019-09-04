package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	dict "github.com/Flyewzz/balda-backend/pkg/dictionary"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

var upgrader = websocket.Upgrader{} // use default options
var game *Game

var nickCounter int = 0

var Dict *dict.DictionaryClient

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

	conn, err := grpc.Dial(
		"127.0.0.1:8090",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Cannot connect to gRPC Dictionary server!")
	}

	defer conn.Close()

	DictManager := dict.NewDictionaryClient(conn)
	result, _ := DictManager.CheckWord(context.Background(), &dict.Word{
		Title: "слово",
	})
	if result.Status {
		fmt.Println("Слово есть в базе!")
	} else {
		fmt.Println("Слова нет в базе!")
	}
	go game.Run()
	http.ListenAndServe(":"+port, nil)
}

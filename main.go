package main

import (
	"log"
	"net/http"

	"github.com/AyushLalShrestha/go-redis-chat-app/config"
	"github.com/AyushLalShrestha/go-redis-chat-app/message"
	"github.com/AyushLalShrestha/go-redis-chat-app/rediscli"
	"github.com/AyushLalShrestha/go-redis-chat-app/websocket"
)

func main() {

	cnf := config.NewConfig()

	log.Printf("Initializing app with config:\n%+v\n", cnf)
	redisCli := rediscli.NewRedis(cnf.RedisAddress, cnf.RedisPassword)
	messageController := message.NewController(redisCli)

	http.Handle("/ws", websocket.Handler(redisCli, messageController))
	http.HandleFunc("/links", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(`{"github":"https://github.com/AyushLalShrestha/go-redis-chat-app"}`))
	})
	http.Handle("/", http.FileServer(http.Dir(cnf.ClientLocation)))
	log.Fatal(http.ListenAndServe(cnf.ServerAddress, nil))
}

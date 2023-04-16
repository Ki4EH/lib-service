package main

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Message struct {
	Message string
	Details []int
}

func main() {
	// Задаём адрес сервера.
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8888", Path: "/"}

	// Создаём подключение к серверу.
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Ошибка подключения к серверу:", err)
	}
	defer conn.Close()

	// Гненерируем сообщение и передаём его на сервер.
	message := Message{"Zubenko", []int{1, 2, 2, 1, 1}}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatal("Ошибка при маршалинге JSON-сообщения:", err)
	}

	// Отправляем json файл на сервер.
	err = conn.WriteMessage(websocket.TextMessage, jsonMessage)
	if err != nil {
		log.Fatal("Ошибка при отправке JSON-сообщения на сервер:", err)
	}

	defer conn.Close()
}

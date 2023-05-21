package handlers

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

func start_server_sender() (*websocket.Conn, error) {
	// Задаём адрес сервера.
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8888", Path: "/"}

	// Создаём подключение к серверу.
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к серверу: %w", err)
	}

	return conn, nil
}

func send_json(conn *websocket.Conn, message entities.json_t) error {
	// Кодируем сообщение в JSON.
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("ошибка при маршалинге JSON-сообщения: %w", err)
	}

	// Отправляем сообщение на сервер.
	err = conn.WriteMessage(websocket.TextMessage, jsonMessage)
	if err != nil {
		return fmt.Errorf("ошибка при отправке JSON-сообщения на сервер: %w", err)
	}

	return nil
}

/*
func main() {
	// Создаём подключение к серверу.
	conn, err := start_server_sender()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Гненерируем сообщение и передаём его на сервер.
	message := entities.json_t{"Zubenko", []int{1, 2, 2, 1, 1}}

	// Отправляем сообщение на сервер.
	err = send_json(conn, message)
	if err != nil {
		log.Fatal(err)
	}
}
*/

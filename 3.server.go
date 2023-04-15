package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type json_t struct {
	Message string `json:"Message"`
	Details []int  `json:"Details"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  65536,
	WriteBufferSize: 65536,
}

func main() {
	// Задаём адрес сервера.
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8899", nil))
}

// Обработчик подключений.
func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Прослушивание всех сообщений, полученных сервером.
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		// Обработка сырого JSON.
		var msg json_t
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println(err)
			break
		}

		// Вывод сырого сообщения, которое было получено.
		msgJSON, err := json.MarshalIndent(msg, "", "  ")
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Сообщение:\n%v\n\n", string(msgJSON))
	}

}

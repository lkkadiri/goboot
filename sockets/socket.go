package socket

import
(
  "fmt"
	"log"
	"net/http"
  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool { return true },
}

var conn *websocket.Conn

type Notification struct{
  Id  int `json:"id"`
  Uuid  string `json:"uuid"`
  AssetId  int `json:"assetId"`
  Type  string `json:"type"`
  EventType     string `json:"eventType"`
  DateCreated     string `json:"dateCreated"`
  DateUpdated     string `json:"dateUpdated"`
  EventPayload struct {
      Body string `json:"body"`
    } `json:"eventPayload"`
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	for {
    m := Notification{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", m)

		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}	}
}

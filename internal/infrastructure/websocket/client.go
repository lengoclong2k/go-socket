package websocket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins in development
	},
}

type Client struct {
	ID     uuid.UUID
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
	UserId uuid.UUID
}

type ClientMessage struct {
	Type   string      `json:"type"`
	Data   interface{} `json:"data"`
	RoomID uuid.UUID   `json:"room_id,omitempty"`
}

func NewClient(hub *Hub, conn *websocket.Conn, UserId uuid.UUID) *Client {
	return &Client{
		ID:     uuid.New(),
		Hub:    hub,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		UserId: UserId,
	}
}


func (c *Client) ReadPump(){
	defer func(){
		c.Hub.
	}
}

package websocket

import (
	"github.com/google/uuid"
	"go-chat-app/internal/domain/entities"
	"log"
	"sync"
)

type Hub struct {
	clients    map[uuid.UUID]*Client
	rooms      map[uuid.UUID]map[uuid.UUID]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan *BroadcastMessage
	mutex      sync.RWMutex
}

type BroadcastMessage struct {
	RoomID  uuid.UUID         `json:"room_id"`
	Message *entities.Message `json:"message"`
	Type    string            `json:"type"`
}

type WebSocketMessage struct {
	Type   string      `json:"type"`
	Data   interface{} `json:"data"`
	RoomID uuid.UUID   `json:"room_id,omitempty"`
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[uuid.UUID]*Client),
		rooms:      make(map[uuid.UUID]map[uuid.UUID]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *BroadcastMessage),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client.ID] = client
			h.mutex.Unlock()
			log.Printf("Client %s connected", client.ID)

		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client.ID]; ok {
				delete(h.clients, client.ID)
				close(client.Send)

				// Remove from all rooms
				for roomID, roomClients := range h.rooms {
					if _, exists := roomClients[client.ID]; exists {
						delete(roomClients, client.ID)
						if len(roomClients) == 0 {
							delete(h.rooms, roomID)
						}
					}
				}
			}
			h.mutex.Unlock()
			log.Printf("Client %s disconnected", client.ID)

		case message := <-h.broadcast:
			h.mutex.RLock()
			if roomClients, ok := h.rooms[message.RoomID]; ok {
				messageData, _ := json.Marshal(WebSocketMessage{
					Type:   message.Type,
					Data:   message.Message,
					RoomID: message.RoomID,
				})

				for clientID, client := range roomClients {
					select {
					case client.Send <- messageData:
					default:
						delete(roomClients, clientID)
						close(client.Send)
					}
				}
			}
			h.mutex.RUnlock()
		}
	}
}

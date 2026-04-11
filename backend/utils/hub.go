package forum

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type       string `json:"type"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
	Timestamp  string `json:"timestamp"`
}

type Client struct {
	ID   string
	Conn *websocket.Conn
}

type Hub struct {
	Db  *sql.DB
	mu      sync.Mutex
	Clients map[string]*websocket.Conn
}

func NewHub(db *sql.DB) *Hub {
	return &Hub{Clients: make(map[string]*websocket.Conn),Db: db}
}

func (hub *Hub) Register(client *Client) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.Clients[client.ID] = client.Conn
}

func (hub *Hub) Unregister(client *Client) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	delete(hub.Clients, client.ID)
}

func (hub *Hub) SendMessage(message *Message) {
 err := hub.SaveMessageToDB(message)
 if err != nil {
	fmt.Println(" Error saving message:", err)
	return 
 }
	hub.mu.Lock()
	defer hub.mu.Unlock()

	conn, exist := hub.Clients[message.ReceiverID]

	if exist {
		err := conn.WriteJSON(message)
		if err != nil {
			fmt.Println("Failed to send, closing connection")
			conn.Close()
			hub.Unregister(&Client{ID: message.ReceiverID})
		}
	} 
}
	

 func IsvalidTime(s string) bool {
	_,err := time.Parse("15:04",s) 

	 return err == nil 
 }

  func (hub *Hub) SaveMessageToDB(message *Message)error{
	query :=  "INSERT INTO messages (sender_id,receiver_id ,content,message_type)VALUES(?,?,?,?)"
    now := time.Now().Format("2006-01-02 15:04:05")
    message.Timestamp = now
	tx , err :=   hub.Db.Begin()
	if err != nil  {
		return err
	}
	_, err = tx.Exec(query,message.SenderID,message.ReceiverID,message.Content,message.Type)
	if err !=  nil  {
		tx.Rollback()
		return err
	}
   
	err  = tx.Commit()
	if err != nil{
		return err
	}
	
	return nil
  }

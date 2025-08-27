package models

//
//import (
//	"github.com/gorilla/websocket"
//	"sync"
//)
//
//type Client struct {
//	ID     string
//	Conn   *websocket.Conn
//	Send   chan models.Message
//	Game   *Game
//	Player *models.PlayerData
//}
//
//type Game struct {
//	ID         string
//	Clients    map[string]*Client
//	State      *models.GameStateData
//	mu         sync.RWMutex
//	colorIndex int
//}
//
//type Server struct {
//	games    map[string]*Game
//	upgrader websocket.Upgrader
//	mu       sync.RWMutex
//}

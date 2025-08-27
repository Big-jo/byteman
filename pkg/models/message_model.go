package models

import "time"

// MessageType represents the type of message being sent
type MessageType string

const (
	PlayerMove  MessageType = "player_move"
	GameState   MessageType = "game_state"
	PlayerJoin  MessageType = "player_join"
	PlayerLeave MessageType = "player_leave"
	PelletEaten MessageType = "pellet_eaten"
	GameOver    MessageType = "game_over"
	GameStart   MessageType = "game_start"
	PlayerReady MessageType = "player_ready"
	ChatMessage MessageType = "chat_message"
	Ping        MessageType = "ping"
	Pong        MessageType = "pong"
)

// Message represents a WebSocket message between client and server
type Message struct {
	Type      MessageType `json:"type"`
	PlayerID  string      `json:"player_id"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

// MoveData represents player movement data
type MoveData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// PelletData represents pellet interaction data
type PelletData struct {
	X     int `json:"x"`
	Y     int `json:"y"`
	Score int `json:"score"`
}

// ChatData represents chat message data
type ChatData struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

// ErrorData represents error information
type ErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ReadyData represents player ready state
type ReadyData struct {
	Ready bool `json:"ready"`
}

// NewMessage creates a new message with timestamp
func NewMessage(msgType MessageType, playerID string, data interface{}) Message {
	return Message{
		Type:      msgType,
		PlayerID:  playerID,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// NewMoveMessage creates a move message
func NewMoveMessage(playerID string, dx, dy float64) Message {
	return NewMessage(PlayerMove, playerID, MoveData{X: dx, Y: dy})
}

// NewPelletMessage creates a pellet eaten message
func NewPelletMessage(playerID string, x, y, score int) Message {
	return NewMessage(PelletEaten, playerID, PelletData{
		X:     x,
		Y:     y,
		Score: score,
	})
}

// NewChatMessage creates a chat message
func NewChatMessage(playerID, username, message string) Message {
	return NewMessage(ChatMessage, playerID, ChatData{
		Message:  message,
		Username: username,
	})
}

// NewErrorMessage creates an error message
func NewErrorMessage(playerID string, code int, message string) Message {
	return NewMessage(GameOver, playerID, ErrorData{
		Code:    code,
		Message: message,
	})
}

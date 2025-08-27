package models

import (
	"byteman/pkg/world"
	"time"
)

// PlayerStatus represents the current status of a player
type PlayerStatus string

const (
	StatusConnected    PlayerStatus = "connected"
	StatusReady        PlayerStatus = "ready"
	StatusPlaying      PlayerStatus = "playing"
	StatusDisconnected PlayerStatus = "disconnected"
)

// PlayerColor represents available player colors
type PlayerColor string

const (
	ColorYellow  PlayerColor = "yellow"
	ColorCyan    PlayerColor = "cyan"
	ColorMagenta PlayerColor = "magenta"
	ColorGreen   PlayerColor = "green"
	ColorRed     PlayerColor = "red"
	ColorBlue    PlayerColor = "blue"
)

// PlayerData represents a player's game data
type PlayerData struct {
	ID          string       `json:"id"`
	Username    string       `json:"username"`
	Pos         world.Vec2   `json:"pos"`
	Score       int          `json:"score"`
	Color       PlayerColor  `json:"color"`
	Status      PlayerStatus `json:"status"`
	Lives       int          `json:"lives"`
	PowerUpTime time.Time    `json:"power_up_time,omitempty"`
	LastSeen    time.Time    `json:"last_seen"`
	JoinedAt    time.Time    `json:"joined_at"`
}

// PlayerStats represents detailed player statistics
type PlayerStats struct {
	PlayerID     string        `json:"player_id"`
	TotalScore   int           `json:"total_score"`
	PelletsEaten int           `json:"pellets_eaten"`
	PowerUpsUsed int           `json:"power_ups_used"`
	TimeAlive    time.Duration `json:"time_alive"`
	GamesPlayed  int           `json:"games_played"`
	GamesWon     int           `json:"games_won"`
	HighScore    int           `json:"high_score"`
	LastPlayed   time.Time     `json:"last_played"`
}

// PlayerPreferences represents player settings and preferences
type PlayerPreferences struct {
	PlayerID         string      `json:"player_id"`
	PreferredColor   PlayerColor `json:"preferred_color"`
	PreferredName    string      `json:"preferred_name"`
	SoundEnabled     bool        `json:"sound_enabled"`
	ShowOtherPlayers bool        `json:"show_other_players"`
	AutoReady        bool        `json:"auto_ready"`
}

// NewPlayerData creates a new player with default values
func NewPlayerData(id, username string, spawnPos world.Vec2, color PlayerColor) *PlayerData {
	return &PlayerData{
		ID:       id,
		Username: username,
		Pos:      spawnPos,
		Score:    0,
		Color:    color,
		Status:   StatusConnected,
		Lives:    3, // Default lives
		LastSeen: time.Now(),
		JoinedAt: time.Now(),
	}
}

// IsAlive returns true if the player has lives remaining
func (p *PlayerData) IsAlive() bool {
	return p.Lives > 0
}

// IsActive returns true if the player is actively playing
func (p *PlayerData) IsActive() bool {
	return p.Status == StatusPlaying && p.IsAlive()
}

// HasPowerUp returns true if the player currently has a power-up active
func (p *PlayerData) HasPowerUp() bool {
	return !p.PowerUpTime.IsZero() && time.Since(p.PowerUpTime) < 10*time.Second
}

// AddScore adds points to the player's score
func (p *PlayerData) AddScore(points int) {
	p.Score += points
}

// LoseLife decrements the player's lives
func (p *PlayerData) LoseLife() {
	if p.Lives > 0 {
		p.Lives--
	}
	if p.Lives == 0 {
		p.Status = StatusDisconnected
	}
}

// ActivatePowerUp activates a power-up for the player
func (p *PlayerData) ActivatePowerUp() {
	p.PowerUpTime = time.Now()
}

// UpdateLastSeen updates the player's last seen timestamp
func (p *PlayerData) UpdateLastSeen() {
	p.LastSeen = time.Now()
}

// SetReady sets the player's ready status
func (p *PlayerData) SetReady(ready bool) {
	if ready {
		p.Status = StatusReady
	} else {
		p.Status = StatusConnected
	}
}

// StartPlaying sets the player's status to playing
func (p *PlayerData) StartPlaying() {
	p.Status = StatusPlaying
}

// Disconnect sets the player's status to disconnected
func (p *PlayerData) Disconnect() {
	p.Status = StatusDisconnected
}

// GetAvailableColors returns all available player colors
func GetAvailableColors() []PlayerColor {
	return []PlayerColor{
		ColorYellow,
		ColorCyan,
		ColorMagenta,
		ColorGreen,
		ColorRed,
		ColorBlue,
	}
}

// IsValidColor checks if a color is valid
func IsValidColor(color PlayerColor) bool {
	for _, c := range GetAvailableColors() {
		if c == color {
			return true
		}
	}
	return false
}

// PlayerListData represents a list of players for client updates
type PlayerListData struct {
	Players    map[string]*PlayerData `json:"players"`
	MaxPlayers int                    `json:"max_players"`
	GameID     string                 `json:"game_id"`
}

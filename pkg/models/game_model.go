package models

import (
	"byteman/pkg/world"
	"time"
)

// GameMode represents different game modes
type GameMode string

const (
	ModeSinglePlayer GameMode = "single_player"
	ModeMultiplayer  GameMode = "multiplayer"
	ModeCooperative  GameMode = "cooperative"
	ModeTimeAttack   GameMode = "time_attack"
	ModeSurvival     GameMode = "survival"
)

// GameStatus represents the current state of the game
type GameStatus string

const (
	StatusWaiting  GameStatus = "waiting"  // Waiting for players
	StatusStarting GameStatus = "starting" // Game is about to start
	StatusActive   GameStatus = "active"   // Game is running
	StatusPaused   GameStatus = "paused"   // Game is paused
	StatusEnded    GameStatus = "ended"    // Game has finished
	StatusAborted  GameStatus = "aborted"  // Game was terminated
)

// GameConfig represents game configuration settings
type GameConfig struct {
	MaxPlayers      int           `json:"max_players"`
	MapWidth        int           `json:"map_width"`
	MapHeight       int           `json:"map_height"`
	GameMode        GameMode      `json:"game_mode"`
	TimeLimit       time.Duration `json:"time_limit,omitempty"`  // 0 means no time limit
	ScoreLimit      int           `json:"score_limit,omitempty"` // 0 means no score limit
	AllowSpectators bool          `json:"allow_spectators"`
	EnablePowerUps  bool          `json:"enable_power_ups"`
	RespawnEnabled  bool          `json:"respawn_enabled"`
	FriendlyFire    bool          `json:"friendly_fire"`
}

// GameStateData represents the complete game state
type GameStateData struct {
	GameID           string                 `json:"game_id"`
	Status           GameStatus             `json:"status"`
	Mode             GameMode               `json:"mode"`
	Players          map[string]*PlayerData `json:"players"`
	Spectators       map[string]*PlayerData `json:"spectators,omitempty"`
	Tiles            [][]*world.Tile        `json:"tiles"`
	Config           GameConfig             `json:"config"`
	StartTime        time.Time              `json:"start_time,omitempty"`
	EndTime          time.Time              `json:"end_time,omitempty"`
	TimeLeft         time.Duration          `json:"time_left,omitempty"`
	TotalPellets     int                    `json:"total_pellets"`
	RemainingPellets int                    `json:"remaining_pellets"`
	Winner           string                 `json:"winner,omitempty"`
}

// GameResult represents the final result of a game
type GameResult struct {
	GameID         string                 `json:"game_id"`
	Mode           GameMode               `json:"mode"`
	Duration       time.Duration          `json:"duration"`
	Winner         string                 `json:"winner,omitempty"`
	Players        map[string]*PlayerData `json:"players"`
	FinalScores    map[string]int         `json:"final_scores"`
	TotalPellets   int                    `json:"total_pellets"`
	PelletsEaten   int                    `json:"pellets_eaten"`
	CompletionRate float64                `json:"completion_rate"`
	EndReason      string                 `json:"end_reason"` // "completed", "timeout", "disconnected", etc.
	Timestamp      time.Time              `json:"timestamp"`
}

// GameSettings represents client-side game settings
type GameSettings struct {
	ServerAddress string        `json:"server_address"`
	GameID        string        `json:"game_id"`
	PlayerID      string        `json:"player_id"`
	PlayerName    string        `json:"player_name"`
	GameMode      GameMode      `json:"game_mode"`
	AutoReconnect bool          `json:"auto_reconnect"`
	Timeout       time.Duration `json:"timeout"`
}

// PowerUp represents a power-up in the game
type PowerUp struct {
	Type      PowerUpType   `json:"type"`
	Pos       world.Vec2    `json:"pos"`
	Duration  time.Duration `json:"duration"`
	SpawnTime time.Time     `json:"spawn_time"`
	Active    bool          `json:"active"`
}

// PowerUpType represents different types of power-ups
type PowerUpType string

const (
	PowerUpSpeed  PowerUpType = "speed"  // Increases movement speed
	PowerUpGhost  PowerUpType = "ghost"  // Can pass through walls briefly
	PowerUpMagnet PowerUpType = "magnet" // Attracts nearby pellets
	PowerUpMulti  PowerUpType = "multi"  // Double points for a time
	PowerUpFreeze PowerUpType = "freeze" // Freezes other players briefly
	PowerUpShield PowerUpType = "shield" // Protects from losing a life
)

// LobbyData represents lobby information
type LobbyData struct {
	Games        map[string]*GameStateData `json:"games"`
	TotalPlayers int                       `json:"total_players"`
	ServerInfo   ServerInfo                `json:"server_info"`
}

// ServerInfo represents server information
type ServerInfo struct {
	Version     string        `json:"version"`
	Uptime      time.Duration `json:"uptime"`
	MaxGames    int           `json:"max_games"`
	MaxPlayers  int           `json:"max_players_per_game"`
	ActiveGames int           `json:"active_games"`
	Features    []string      `json:"features"`
}

// DefaultGameConfig returns a default game configuration
func DefaultGameConfig() GameConfig {
	return GameConfig{
		MaxPlayers:      4,
		MapWidth:        64,
		MapHeight:       32,
		GameMode:        ModeMultiplayer,
		TimeLimit:       0, // No time limit
		ScoreLimit:      0, // No score limit
		AllowSpectators: true,
		EnablePowerUps:  true,
		RespawnEnabled:  false,
		FriendlyFire:    false,
	}
}

// SinglePlayerConfig returns a configuration for single player mode
func SinglePlayerConfig() GameConfig {
	config := DefaultGameConfig()
	config.MaxPlayers = 1
	config.GameMode = ModeSinglePlayer
	config.MapWidth = 128
	config.MapHeight = 128
	config.AllowSpectators = false
	return config
}

// TimeAttackConfig returns a configuration for time attack mode
func TimeAttackConfig(timeLimit time.Duration) GameConfig {
	config := DefaultGameConfig()
	config.GameMode = ModeTimeAttack
	config.TimeLimit = timeLimit
	config.RespawnEnabled = true
	return config
}

// SurvivalConfig returns a configuration for survival mode
func SurvivalConfig() GameConfig {
	config := DefaultGameConfig()
	config.GameMode = ModeSurvival
	config.RespawnEnabled = false
	config.EnablePowerUps = true
	return config
}

// IsTimeUp checks if the game time limit has been reached
func (g *GameStateData) IsTimeUp() bool {
	if g.Config.TimeLimit == 0 {
		return false
	}
	return time.Since(g.StartTime) >= g.Config.TimeLimit
}

// GetTimeRemaining returns the remaining game time
func (g *GameStateData) GetTimeRemaining() time.Duration {
	if g.Config.TimeLimit == 0 {
		return 0
	}
	elapsed := time.Since(g.StartTime)
	if elapsed >= g.Config.TimeLimit {
		return 0
	}
	return g.Config.TimeLimit - elapsed
}

// IsGameComplete checks if all pellets have been collected
func (g *GameStateData) IsGameComplete() bool {
	return g.RemainingPellets == 0
}

// GetLeadingPlayer returns the player with the highest score
func (g *GameStateData) GetLeadingPlayer() *PlayerData {
	var leader *PlayerData
	maxScore := -1

	for _, player := range g.Players {
		if player.Score > maxScore {
			maxScore = player.Score
			leader = player
		}
	}

	return leader
}

// GetActivePlayers returns only the active players
func (g *GameStateData) GetActivePlayers() map[string]*PlayerData {
	active := make(map[string]*PlayerData)
	for id, player := range g.Players {
		if player.IsActive() {
			active[id] = player
		}
	}
	return active
}

// CanStart checks if the game can be started
func (g *GameStateData) CanStart() bool {
	if g.Status != StatusWaiting {
		return false
	}

	activePlayers := len(g.GetActivePlayers())
	return activePlayers > 0 && activePlayers <= g.Config.MaxPlayers
}

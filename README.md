# Byteman
Terminal based mutiplayer pacman


## Project Structure
```
byteman/
├── cmd/
│   └── server/
│       └── main.go          # Server entry point
├── pkg/
│   ├── game/
│   │   └── game.go          # Updated game logic with multiplayer support
│   ├── render/
│   │   └── renderer.go      # Updated renderer for multiplayer
│   ├── server/
│   │   └── server.go        # WebSocket server implementation
│   └── world/
│       └── world.go         # World/map generation (unchanged)
├── main.go                  # Client entry point
├── go.mod                   # Go module dependencies
└── README.md                # This file
```


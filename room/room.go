package room

import (
	"SyncCinemaNet_WebGo/auth"
)

type Room struct {
	ID       int64
	Name     string
	Private  bool
	Author   auth.User
	userlist map[string]auth.User
}

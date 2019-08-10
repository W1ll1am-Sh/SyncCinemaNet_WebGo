package main

import (
	"SyncCinemaNet_WebGo/api"
	"SyncCinemaNet_WebGo/pg"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api.Initialize(r)

	err := pg.Initialize()
	defer pg.CloseDb()
	if err != nil {
		panic(err.Error())
	}

}

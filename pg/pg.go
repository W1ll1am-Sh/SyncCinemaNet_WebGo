package pg

import (
	"SyncCinemaNet_WebGo/auth"
	"SyncCinemaNet_WebGo/chat"
	"SyncCinemaNet_WebGo/room"

	"github.com/jinzhu/gorm"

	// blank import needed by gorm package
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Initialize - Open db and return open error
func Initialize() error {
	var err error
	db, err = gorm.Open("postgres", "host="+pgHost+" port="+pgPort+" user="+pgUser+" dbname="+pgDbName+" password="+pgPassword+" sslmode="+pgSslMode)
	if err != nil {
		return err
	}

	createTableIfDoesntExist(&auth.User{})
	createTableIfDoesntExist(&auth.TokenAuth{})
	createTableIfDoesntExist(&room.Room{})
	createTableIfDoesntExist(&chat.Chat{})
	createTableIfDoesntExist(&chat.Message{})

	return err
}

// CloseDb - Closing db and return closing error
func CloseDb() error {
	return db.Close()
}

func createTableIfDoesntExist(table interface{}) {
	if !db.HasTable(table) {
		db.AutoMigrate(table)
	}
}

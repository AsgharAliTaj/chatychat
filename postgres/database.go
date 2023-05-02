package postgres

import (
	"log"

	"github.com/asgharalitaj/chatychat"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
  var err error

  dsn := "host=localhost user=root password=toor dbname=chatychatdb port=5432 sslmode=disable TimeZone=Asia/Karachi"
  
  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal("error connecting to db", err)
  } else {
    log.Println("connected to database")
  }
  db.AutoMigrate(&chatychat.Thread{}, &chatychat.Post{}, &chatychat.Comment{})
}

func GETDB() *gorm.DB {
  return db
}



package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  "model"
)

func Init() {
  db, err := gorm.Open("sqlite3", "link.sqlite3")
  if err != nil {
          panic("failed to connect database\n")
  }

  db.AutoMigrate(&model.Link{})
}

func Create(name string, path string) {
  db, err := gorm.Open("sqlite3", "link.sqlite3")
  if err != nil {
          panic("failed to connect database\n")
  }

  db.Create(&model.Link{Name: name, Path: path})
}

func GetAll() []model.Link {
  db, err := gorm.Open("sqlite3", "link.sqlite3")
  if err != nil {
          panic("failed to connect database\n")
  }

  var links []model.Link
  db.Find(&links)
  return links
}

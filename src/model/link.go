package model

import (
  "github.com/jinzhu/gorm"
)

type Link struct {
    gorm.Model
    Name string
    Path string
}

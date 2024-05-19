package schemas

import (
  "gorm.io/gorm"
  "time"
)

type Escala struct {
  gorm.Model
  Ministry string
  Month int64
  Year int64
  Data string
}


type EscalaResponse struct {
  ID uint `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt time.Time `json:"deleted_at"`
  Ministry string `json:"ministry"`
  Month int64 `json:"month"`
  Year int64 `json:"year"`
  Data string `json:"data"`
}



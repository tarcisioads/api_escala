package schemas

import (
  "gorm.io/gorm"
  "time"
)

type Ministry struct {
  gorm.Model
  Name string
}

type MinistryResponse struct {
  ID uint `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt time.Time `json:"deleted_at"`
  Name string `json:"name"`
}

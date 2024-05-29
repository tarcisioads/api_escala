package schemas

import (
  "gorm.io/gorm"
  "time"
)

type User struct {
  gorm.Model
  Username string
  Password string
}

type UserResponse struct {
  ID uint `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt time.Time `json:"deleted_at"`
  Username string `json:"username"`
  Password string `json:"password"`
}

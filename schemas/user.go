package schemas

import (
  "gorm.io/gorm"
  "time"
)

type User struct {
  gorm.Model
  Email string `gorm:"unique"`
  Password string
}

type UserResponse struct {
  ID uint `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt time.Time `json:"deleted_at"`
  Email string `json:"email"`
  Password string `json:"password"`
}

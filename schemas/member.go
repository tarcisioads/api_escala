package schemas

import (
  "gorm.io/gorm"
  "time"
)

type Member struct {
  gorm.Model
  Name string
  Ministries []*Ministry `gorm:"many2many:members_ministries;"`
}

type MemberResponse struct {
  ID uint `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt time.Time `json:"deleted_at"`
  Name string `json:"name"`
  Ministries []MinistryResponse `json:"ministries"`
}

package entity

import "time"

type Product struct {
	Id          int       `gorm:"primaryKey;not null" json:"id"`
	Title       string    `gorm:"unique;not null;type:varchar(255)" json:"email"`
	Description string    `gorm:"type:text[];not null" json:"description"`
	UserId      int       `gorm:"foreignKey:UserId;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"userId"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
}

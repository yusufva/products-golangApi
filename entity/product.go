package entity

import (
	"time"
	"tugas-sesi12/dto"
)

type Product struct {
	Id          int       `gorm:"primaryKey;not null" json:"id"`
	Title       string    `gorm:"not null;type:varchar(255)" json:"email"`
	Description string    `gorm:"type:text;not null" json:"description"`
	UserId      int       `gorm:"foreignKey:UserId;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"userId"`
	CreatedAt   time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now()" json:"updated_at"`
}

func (p *Product) EntityToProductResponseDto() dto.ProductResponse {
	return dto.ProductResponse{
		Id:          p.Id,
		Title:       p.Title,
		Description: p.Description,
		UserId:      p.UserId,
		Created_At:  p.CreatedAt,
		Updated_At:  p.UpdatedAt,
	}
}

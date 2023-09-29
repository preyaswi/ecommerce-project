package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         int    `gorm:"primarykey" bson:"_id,omitempty" json:"-"`
	FirstName  string `json:"firstname" bson:"firstname" binding:"required"`
	LastName   string `json:"lastname" bson:"lastname" binding:"required"`
	Email      string `json:"email" bson:"email" binding:"required"`
	Phone      string `json:"phone" bson:"phone" binding:"required"`
	Password   string `json:"-" bson:"password" binding:"required"`
	Wallet     int    `json:"wallet"`
	Permission bool   `gorm:"not null;default:true" json:"-"`
}
type OtpKey struct {
	gorm.Model
	Key   string `json:"key"`
	Phone string `json:"phone"`
}


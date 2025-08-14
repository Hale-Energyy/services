package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Address represents user's address details
type Address struct {
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	StreetName string `json:"street_name"`
	Pincode    string `json:"pincode"`
}

type User struct {
	ID        string  `gorm:"primaryKey;size:20" json:"id"`
	FirstName string  `json:"name"`
	LastName  string  `json:"last_name"`
	Username  *string `gorm:"unique;default:null" json:"username"` // nullable & unique
	Email     string  `json:"email" gorm:"unique"`
	Age       int     `json:"age"`
	Gender    string
	Height    float64 `json:"height"`
	Weight    float64 `json:"weight"`
	BloodType string  `json:"blood_type"`
	Address   Address `gorm:"embedded;embeddedPrefix:address_" json:"address"`
	// FitnessProfile UserFitnessProfile `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"fitness_profile"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var lastUser User

	// Find last created user by sorting ID in descending order
	if err := tx.Order("id DESC").First(&lastUser).Error; err != nil {
		// If no user found (first user), set ID to HALE0001
		if err == gorm.ErrRecordNotFound {
			u.ID = "HALEU0001"
			return nil
		}
		return err // Return other DB errors
	}

	// Extract numeric part from last ID (e.g., "HALE0007" -> 7)
	var lastNumber int
	fmt.Sscanf(lastUser.ID, "HALEU%04d", &lastNumber)

	// Increment and format as HALE000X with leading zeros
	newNumber := lastNumber + 1
	u.ID = fmt.Sprintf("HALEU%04d", newNumber)

	return nil
}

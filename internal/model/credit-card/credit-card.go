package credit_card

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model `json:"gorm_._model"`
	Number     string `json:"number"`
	UserID     uint   `json:"user_id"`
}

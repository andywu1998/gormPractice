package reposity

import (
	"database/sql"
	"gormPractice/models"
	"time"
)

var users = []models.User{
	{
		Name:         "xxx",
		Email:        nil,
		Age:          0,
		Birthday:     nil,
		MemberNumber: sql.NullString{},
		ActivatedAt:  sql.NullTime{},
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		Location:     models.Location{},
		CreditCard:   models.CreditCard{},
	},
	{
		Name:         "xxname",
		Email:        nil,
		Age:          0,
		Birthday:     nil,
		MemberNumber: sql.NullString{},
		ActivatedAt:  sql.NullTime{},
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		Location:     models.Location{},
		CreditCard:   models.CreditCard{},
	},
}

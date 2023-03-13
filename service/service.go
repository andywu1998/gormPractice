package service

import (
	"gormPractice/db/reposity"
)

func InsertViaMap() error {
	err := reposity.CreateUserFromMap(map[string]interface{}{
		"Name": "jinzhumap", "Age": 38,
	})
	if err != nil {
		return err
	}
	return nil
}

package reposity

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	dblib "gormPractice/db"
	"gormPractice/models"
)

var db *gorm.DB

func init() {
	db = dblib.Db()
}

func AutoMigrate() {
	db.AutoMigrate(models.User{})
}

func CreateBlog(blog models.Blog) error {
	ret := db.Create(&blog)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

func CreateUserFromMap(m map[string]interface{}) error {
	// 用map的坏处就是，自动填充的字段没有填充。
	ret := db.Model(&models.User{}).Create(m)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

func CreateUserWithLocaltion() error {
	ret := db.Model(models.User{}).Create(map[string]interface{}{
		"Name": "jinzhu",
		"loc":  clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 200)"}},
	})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

func CreateUserWithLocaltionViaStruct() error {
	ret := db.Create(&models.User{
		Name:     "jinzhu",
		Location: models.Location{X: 100, Y: 300},
	})

	return ret.Error
}

func CreateRelation() error {
	ret := db.Create(&models.User{
		Name:       "jinzhu",
		CreditCard: models.CreditCard{Number: "411111111111"},
	})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

func CreateWithOmit() error {
	ret := db.Omit("CreditCard").Create(&models.User{
		Name:       "jinzhu",
		CreditCard: models.CreditCard{Number: "411111111112"},
	})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

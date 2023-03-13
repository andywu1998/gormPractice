package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gormPractice/models"
	"log"
)

var db *gorm.DB

func InitDB() {
	dsn := "root:rootpassword@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_",
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("init db error", err)
	}
}

func SyncBuyRecordAndMaterialPackage() {
	InitDB()
	pkgs := make([]models.CashTaskMaterialPackage, 0)
	db.Find(&pkgs)
	for _, pkg := range pkgs {
		db.Model(models.CashTaskMaterialPackageBuyRecord{}).Where("cash_task_material_package_id=?", pkg.Id).Update("cat_food_num", pkg.CatFoodNum*100)
	}
}

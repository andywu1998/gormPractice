package reposity

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gormPractice/models"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	db.Delete(&models.User{})
}

var email = "jinzhu@gorm.com"

// TestCreateUser 单行创建
func TestCreateUser(t *testing.T) {
	birthday := time.Date(1998, 11, 29, 1, 1, 1, 1, time.Local)

	user := models.User{
		Name:     "TestCreateUser",
		Email:    &email,
		Age:      23,
		Birthday: &birthday,
		MemberNumber: sql.NullString{
			String: "123",
			Valid:  false,
		},
		ActivatedAt: sql.NullTime{},
		Location:    models.Location{},
		CreditCard:  models.CreditCard{},
	}
	result := db.Create(&user) // 通过数据的指针来创建
	if result.Error != nil {
		t.Fail()
	}
}

// TestCreateUsers 批量创建
func TestCreateUsers(t *testing.T) {
	// INSERT INTO `tbl_user` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`,`location`) VALUES ('xxx',NULL,0,NULL,NULL,NULL,'2023-03-13 16:18:49.645','2023-03-13 16:18:49.645',ST_PointFromText('POINT(0 0)')),('xxname',NULL,0,NULL,NULL,NULL,'2023-03-13 16:18:49.645','2023-03-13 16:18:49.645',ST_PointFromText('POINT(0 0)'))
	result := db.Create(&users) // 通过数据的指针来创建
	if result.Error != nil {
		t.Fail()
	}
}

// TestCreateUserFromMap 用map创建
func TestCreateUserFromMap(t *testing.T) {
	ret := db.Model(models.User{}).Create(map[string]interface{}{
		"Name":     "TestCreateUserFromMa",
		"Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 200)"}},
		"age":      19,
	})
	// INSERT INTO `tbl_user` (`location`,`name`) VALUES (ST_PointFromText('POINT(100 200)'),'TestCreateUserFromMa')
	if ret.Error != nil {
		t.Fail()
	}
}

// TestCreateRelation 创建关系
func TestCreateRelation(t *testing.T) {
	// 分开插入两行 INSERT INTO `tbl_credit_card` INSERT INTO `tbl_user`
	// 没有用事务
	ret := db.Create(&models.User{
		Name:       "jinzhu",
		CreditCard: models.CreditCard{Number: "411111111111"},
	})

	if ret.Error != nil {
		t.Fail()
	}
}

// TestInsertOnConflict 插入冲突时更新让age+1
func TestInsertOnConflict(t *testing.T) {
	user := models.User{
		Name:     "user1",
		Email:    &email,
		Age:      19,
		Birthday: nil,
		MemberNumber: sql.NullString{
			String: "123",
			Valid:  false,
		},
		ActivatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: false,
		},
		Location: models.Location{
			X: 1,
			Y: 2,
		},
	}
	db.Create(&user)
	// INSERT INTO `tbl_user` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`,`location`) VALUES ('user1','jinzhu@gorm.com',19,NULL,NULL,NULL,'2023-03-13 16:55:16.41','2023-03-13 16:55:16.41',ST_PointFromText('POINT(1 2)'))
	fmt.Println(user.Age)
	// 19
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"age": gorm.Expr("age + 1")}),
	}).Create(&user)
	// INSERT INTO `tbl_user` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`,`location`,`id`) VALUES ('user1','jinzhu@gorm.com',19,NULL,NULL,NULL,'2023-03-13 16:55:16.41','2023-03-13 16:55:16.41',ST_PointFromText('POINT(1 2)'),6) ON DUPLICATE KEY UPDATE `age`=age + 1
	db.Where("id=?", user.ID).First(&user)
	fmt.Println(user.Age)
	// 20
}

// TestInsertOnConflict2 遇到冲突时更新一部分值
func TestInsertOnConflict2(t *testing.T) {
	user := models.User{
		Name:     "user1",
		Email:    &email,
		Age:      19,
		Birthday: nil,
		MemberNumber: sql.NullString{
			String: "123",
			Valid:  false,
		},
		ActivatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: false,
		},
		Location: models.Location{
			X: 1,
			Y: 2,
		},
	}
	db.Create(&user)
	// INSERT INTO `tbl_user` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`,`location`) VALUES ('user1','jinzhu@gorm.com',19,NULL,NULL,NULL,'2023-03-13 17:00:38.829','2023-03-13 17:00:38.829',ST_PointFromText('POINT(1 2)'))
	user.Name = "updateName"
	user.Age = 199
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	}).Create(&user)
	// INSERT INTO `tbl_user` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`,`location`,`id`) VALUES ('updateName','jinzhu@gorm.com',199,NULL,NULL,NULL,'2023-03-13 17:00:38.829','2023-03-13 17:00:38.829',ST_PointFromText('POINT(1 2)'),9) ON DUPLICATE KEY UPDATE `name`=VALUES(`name`),`age`=VALUES(`age`)
}

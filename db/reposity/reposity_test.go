package reposity

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/hints"
	dblib "gormPractice/db"
	"gormPractice/models"
	"testing"
)

func init() {
	db = dblib.Db()
}

func TestCreateUserWithLocaltionViaStruct(t *testing.T) {
	CreateUserWithLocaltionViaStruct()
}

func TestCreateWithOmit(t *testing.T) {
	CreateWithOmit()
}

func TestLimit(t *testing.T) {
	users1 := make([]models.User, 0)
	users2 := make([]models.User, 0)
	db.Limit(3).Find(&users1).Limit(-1).Find(&users2)
	// 翻译成两行sql
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)
	for _, user := range users1 {
		fmt.Println(user)
	}
	fmt.Println("~~~")
	for _, user := range users2 {
		fmt.Println(user)
	}
}

func TestGroup(t *testing.T) {
	type result struct {
		Name  string
		Total int
	}
	res := []result{}
	db.Model(&models.User{}).Select("name, sum(age) as total").Where("name LIKE ?", "j%").Group("name").Find(&res)
	fmt.Println(res)
}

func TestJoin(t *testing.T) {
	type result struct {
		Name    string
		Upvotes int32
	}

	ret := []result{}

	db.Model(models.User{}).Select("tbl_user.name as name, tbl_blog.upvotes as upvotes").Joins("JOIN tbl_blog ON tbl_blog.author_id = tbl_user.id").Scan(&ret)
	for _, r := range ret {
		fmt.Println(r)
	}
}

func TestJoinAndGroup(t *testing.T) {
	type result struct {
		Name string
		// 这里一定要int64，因为int32聚合之后就不够装了。
		A int64
	}

	ret := []result{}

	db.Model(models.User{}).Select("tbl_user.name as name, sum(tbl_blog.upvotes) as A").Joins("JOIN tbl_blog ON tbl_blog.author_id = tbl_user.id").Group("name").Scan(&ret)
	// SELECT tbl_user.name as name, sum(tbl_blog.upvotes) as A FROM `tbl_user` JOIN tbl_blog ON tbl_blog.author_id = tbl_user.id GROUP BY `name`
	for _, r := range ret {
		fmt.Println(r)
	}
}

func TestForUpdate(t *testing.T) {
	db.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id=1").Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
}

func TestForShare(t *testing.T) {
	db.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Where("id=1").Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
}

func TestNoWait(t *testing.T) {
	// todo: 去了解一下NOWAIT关键字
	db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Options:  "NOWAIT",
	}).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
}

func TestSubQuery(t *testing.T) {
	orders := []models.Orders{}
	res := make(map[string]interface{})
	// 查询平均值
	db.Table("tbl_orders").Select("AVG(amount)").Find(&res)
	for k, v := range res {
		fmt.Println(k, v)
	}

	db.Where("amount > (?)", db.Table("tbl_orders").Select("AVG(amount)")).Find(&orders)
	// 取出所有大于平均值的
	// SELECT * FROM `tbl_orders` WHERE amount > (SELECT AVG(amount) FROM `tbl_orders`) AND `tbl_orders`.`deleted_at` IS NULL
	for _, order := range orders {
		fmt.Println(order)
	}
}

func TestInWithMultipleColumns(t *testing.T) {
	db.Where("(name, age, role) IN ?", [][]interface{}{{"jinzhu", 18, "admin"}, {"jinzhu2", 19, "user"}}).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
}

func TestToSQL(t *testing.T) {
	sql := db.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Locking{
			Strength: "UPDATE",
			Options:  "NOWAIT",
		}).Find(&users)
	})
	fmt.Println(sql)
}

func TestFirstOrInit(t *testing.T) {
	var user = models.User{}
	db.Where("id=?", 20).FirstOrInit(models.User{Name: "non_existing"})
	fmt.Println(user)
	user = models.User{}
	db.Where("id=?", 21).FirstOrInit(&user, models.User{Name: "non_existing"})
	fmt.Println(user)
}

func TestIndex(t *testing.T) {
	users = []models.User{}
	// SELECT * FROM `tbl_user` USE INDEX (`idx_user_name`) WHERE name='jinzhu'
	db.Clauses(hints.UseIndex("idx_user_name")).Where("name=?", "jinzhu").Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
}

func TestUpdateSingleColumn(t *testing.T) {
	user := models.User{ID: 19, Name: "jinzhu2"}
	ret := db.Model(&user).Update("name", "hello")
	// UPDATE `tbl_user` SET `name`='hello',`updated_at`='2023-03-13 15:38:44.995' WHERE `id` = 19
	// 虽然user里填了Name，但是实际上去update的时候不会用
	// 并且update_at是自动填充上去的
	if ret.Error != nil {
		t.Fail()
	}
}

func TestUpdateSingleColumn2(t *testing.T) {
	ret := db.Model(&models.User{}).Where("name", "jinzhu2").Update("name", "uujinzhu")
	// UPDATE `tbl_user` SET `name`='uujinzhu',`updated_at`='2023-03-13 15:41:38.193' WHERE `name` = 'jinzhu2'
	// where和update的组合
	if ret.Error != nil {
		t.Fail()
	}
	// 为了不影响后续测试，把值改回来
	ret = db.Model(&models.User{}).Where("name", "uujinzhu").Update("name", "jinzhu2")
	if ret.Error != nil {
		t.Fail()
	}
}

// TestUpdateSingleColumn3 在已有值更新
func TestUpdateSingleColumn3(t *testing.T) {
	ret := db.Model(&models.User{}).Where("name", "jinzhu2").Update("age", gorm.Expr("age*3"))
	// UPDATE `tbl_user` SET `age`=age*3,`updated_at`='2023-03-13 15:46:55.419' WHERE `name` = 'jinzhu2'
	// where和update的组合
	if ret.Error != nil {
		t.Fail()
	}
}

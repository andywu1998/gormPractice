package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	gorm.Model
	AuthorID uint
	Upvotes  int32
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Location     Location
	CreditCard   CreditCard
}

type Location struct {
	X, Y int
}

// Scan 方法实现了 sql.Scanner 接口
func (loc *Location) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	return nil
}

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("beforeCreate")
	if u.Name == "jinzhu1" {
		return errors.New("invalid role")
	}
	return
}

type Orders struct {
	gorm.Model
	Amount float64
}

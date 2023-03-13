package main

import (
	"gormPractice/db"
	"gormPractice/db/reposity"
	"gormPractice/models"
	"log"
	"reflect"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func testService(fs []func() error) {
	for _, f := range fs {
		err := f()
		if err != nil {
			log.Fatalln(GetFunctionName(f), "err", err)
			return
		}
	}
}

func main() {
	x, _ := db.Db().DB()
	err := x.Ping()
	if err != nil {
		log.Println(err)
	}
	db.Db().AutoMigrate(models.User{})

	testService([]func() error{
		//service.BatchInsert,
		//service.Insert,
		//service.InsertViaMap,
		reposity.CreateUserWithLocaltion,
	})
}

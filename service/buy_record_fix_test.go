package service

import (
	"fmt"
	"gormPractice/models"
	"testing"
)

func TestSyncBuyRecordAndMaterialPackage(t *testing.T) {
	SyncBuyRecordAndMaterialPackage()
}

func TestAddMaterialPackage(t *testing.T) {
	InitDB()
	db.Create(&models.CashTaskMaterialPackage{
		Id:             0,
		CashTaskId:     0,
		UserId:         0,
		User:           models.User{},
		Cover:          "",
		Name:           "",
		CatFoodNum:     0,
		MaterialNum:    0,
		FolderNum:      0,
		MaterialUseNum: 0,
		ThirteenUseNum: 0,
		ImgNum:         0,
		VideoNum:       0,
		MusicNum:       0,
		Status:         0,
		PublicStatus:   0,
		Reason:         "",
		ApproveTime:    0,
		Quality:        0,
		AvgScore:       0,
		ScoreUserNum:   0,
		CreateAt:       0,
		UpdateAt:       0,
		IsFree:         true,
	})
	pkgs := make([]models.CashTaskMaterialPackage, 0)
	db.Find(&pkgs)
	for _, pkg := range pkgs {
		fmt.Println(pkg)
	}
}

package repo

import (
	"context"
	"fmt"
	model "nextts_react_go/internal/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func TB_Import_Error_Main() []model.TB_Import_Error {
	dsn := "sqlserver://sa@hmmta-ppm?database=New_Kanban_F3"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	fmt.Println(dsn)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()

	// UserList, err := gorm.G[model.User](db.Where("_ID =?", 14)).First(ctx)
	var UserList []model.TB_Import_Error
	if err := db.WithContext(ctx).Find(&UserList).Error; err != nil {
		//return  err.Error();
		fmt.Println(err)
	}
	if len(UserList) == 0 {
		return nil
	}

	return UserList
}

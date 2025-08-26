package repo

import (
	"context"
	"fmt"
	model "nextts_react_go/internal/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func UserMain() model.User {
	dsn := "sqlserver://sa@hmmta-ppm?database=New_Kanban_F3"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	//fmt.Println(dsn)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()

	// UserList, err := gorm.G[model.User](db.Where("_ID =?", 14)).First(ctx)
	var User model.User
	//UserList = db.Find(&UserList).Where("_ID = ?", 14)
	if err := db.WithContext(ctx).Select("Code", "Name", "Email", "Surname").Where("_ID = ?", 14).First(&User).Error; err != nil {
		//return  err.Error();
		fmt.Println("db Error => ", err)
	}
	if User == (model.User{}) {
		return model.User{}
	}
	// for _, user := range UserList {
	// 	fmt.Println(user.Name)
	// }
	fmt.Printf("%s %s \n", *User.Name, *User.Surname)
	fmt.Println(*User.Name + " " + *User.Surname + " 34")
	//fmt.Println(User)
	//fmt.Sprintf("%s %s 34\n", *User.Name, *User.Surname)
	return User
}

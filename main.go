package main

import (
	"fmt"
	"os"

	"github.com/harukitosa/ddd_sample/application"
	"github.com/harukitosa/ddd_sample/domain/model"
	"github.com/harukitosa/ddd_sample/infrastructure/datastore"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// データベースを準備
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}

	// マイグレーション
	db.AutoMigrate(&model.User{})

	//　依存性の解決
	userRepository := datastore.NewUserRepositoryImpliment(db)
	userService := application.NewUserService(userRepository)

	name := "とさ"

	user, err := userService.CreateUser(name)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	} else {
		fmt.Println(user)
	}
}

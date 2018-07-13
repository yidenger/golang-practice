package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gwp")

	if err != nil {
		panic(err)
	}

	fmt.Println("init db successful")

	defer db.Close()

	db.AutoMigrate(&Product{})
}
package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3305)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		return
	}
	db.LogMode(true)
	/*
		db.AutoMigrate(&Product{})
		p1 := Product{
			Code:  "sgsexfes",
			Price: 110,
		}
		err = db.Create(&p1).Error
		if err!= nil {
			log.Println(err)
			return
		}

	*/
	var list []int
	rows, _ := db.Raw("select `id` from products where price =?", 110).Rows()
	for rows.Next() {

	}
	log.Println(list)
}

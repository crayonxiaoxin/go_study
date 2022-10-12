// package main

// import (
// 	"fmt"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func init() {
// 	dsn := "root:New4you!@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db = d
// }

// func test1() {
// 	type Company struct {
// 		ID   int
// 		Name string
// 	}

// 	// 模型
// 	type User struct {
// 		gorm.Model
// 		Name      string
// 		Age       int
// 		Birthday  time.Time
// 		CompanyID int
// 		Company   Company
// 	}

// 	db.AutoMigrate(&Company{}, &User{})
// }

// // 多对一
// func n2one() {
// 	type Toy struct {
// 		ID        int
// 		Name      string
// 		OwnerID   int
// 		OwnerType string
// 	}

// 	type Cat struct {
// 		ID   int
// 		Name string
// 		Toy  Toy `gorm:"polymorphic:Owner;"`
// 	}

// 	type Dog struct {
// 		ID   int
// 		Name string
// 		Toy  Toy `gorm:"polymorphic:Owner;"`
// 	}

// 	db.AutoMigrate(&Toy{}, &Cat{}, &Dog{})

// 	// db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
// 	// db.Create(&Cat{Name: "cat1", Toy: Toy{Name: "toy2"}})
// 	var toy Toy
// 	db.First(&toy)
// 	fmt.Printf("toy: %v\n", toy)
// }

// // 一对多
// func one2n() {
// 	type CreditCard struct {
// 		gorm.Model
// 		Number string
// 		UserID uint
// 	}

// 	type User struct {
// 		gorm.Model
// 		CreditCards []CreditCard
// 	}

// 	db.AutoMigrate(&User{}, &CreditCard{})
// }

// func one2n_2() {
// 	type Toy struct {
// 		ID        int
// 		Name      string
// 		OwnerID   int
// 		OwnerType string
// 	}

// 	type Dog struct {
// 		ID   int
// 		Name string
// 		Toys []Toy `gorm:"polymorphic:Owner;"`
// 	}

// 	db.AutoMigrate(&Dog{}, &Toy{})
// 	// db.Create(&Dog{Name: "dog1", Toys: []Toy{{Name: "toy1"}, {Name: "toy2"}}})
// 	var dogs []Dog
// 	db.Model(&Dog{}).Preload("Toys").Find(&dogs)
// 	fmt.Printf("dogs: %v\n", dogs)
// }

// func main() {
// 	// test1()
// 	// n2one()
// 	// one2n()
// 	one2n_2()
// }

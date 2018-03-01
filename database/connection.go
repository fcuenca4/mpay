package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	if DB == nil {
		db, err := gorm.Open("mysql", "root:debezium@tcp(127.0.0.1:3306)/mpay?charset=utf8&parseTime=true&loc=Local")
		DB = db
		if err != nil {
			fmt.Println("db err: ", err)
		}
	}
	return DB
	//db.LogMode(true)
}

//Error Handler
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

// This function will create a temporarily database for running testing cases
/* func TestDBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:debezium@tcp(127.0.0.1:3306)/mpay?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	test_db.DB().SetMaxIdleConns(3)
	test_db.LogMode(true)
	DB = test_db
	return DB
}

// Delete the database after running testing cases.
func TestDBFree(test_db *gorm.DB) error {
	test_db.Close()
	err := os.Remove("mysql", "root:debezium@tcp(127.0.0.1:3306)/mpay?charset=utf8&parseTime=True&loc=Local")
	return err
}
*/

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

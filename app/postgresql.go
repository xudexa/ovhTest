package app

import (
	"fmt"
	"log"
	"ovhTest/app/functions"
	"proto/app/db"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func OpenDbPostgres(host string, port string, dbName string, user string, password string) (*gorm.DB, error) {
	dbport, _ := strconv.Atoi(port)
	dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		host,
		dbport,
		dbName,
		user,
		password)

	db, err := gorm.Open("postgres", dbURI)

	return db, err
}

func AddTodoInPostgresql(t Todo) {
	var err error

	t.ID = functions.NewUUID()
	t.CreatedAt = time.Now()
	t.Completed = false

	dbTransac := server.Database.Begin()
	if err = dbTransac.Create(&t).Error; err == nil {
		dbTransac.Commit()
	} else {

		dbTransac.Rollback()
	}
}

func CompleteTodoInPostgresql(id string) {
	var todo *Todo

	var err error

	if db.GetInstance().Where("id = ?", id).First(&todo).RecordNotFound() {
		log.Fatal("Todo " + id + " not found.")
	}

	dbTransac := db.GetInstance().Begin()

	todo.Completed = !todo.Completed
	todo.CompletedAt = time.Now()

	err = dbTransac.Model(&todo).Where("id = ?", id).Updates(todo).Error
	if err != nil {
		dbTransac.Rollback()
	} else {
		dbTransac.Commit()
	}

}

func RemoveTodoInPostgresql(id string) {
	var todo *Todo

	var err error

	if db.GetInstance().Where("id = ?", id).First(&todo).RecordNotFound() {
		log.Fatal("Todo " + id + " not found.")
	}

	dbTransac := db.GetInstance().Begin()

	err = dbTransac.Model(&todo).Where("id = ?", id).Delete(todo).Error
	if err != nil {
		dbTransac.Rollback()
	} else {
		dbTransac.Commit()
	}
}

package database

import (
	"fmt"
	"time"

	"github.com/PraserX/fullstack-rookie/pkg/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Database struct {
	orm *gorm.DB
}

func New(opts ...Option) (*Database, error) {
	var err error
	var db = Database{}

	var options = &Options{}
	options.Path = "temp.db"

	for _, opt := range opts {
		opt(options)
	}

	db.orm, err = gorm.Open(sqlite.Open(options.Path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cannot create database")
	}

	db.orm.AutoMigrate(&model.User{}, &model.Comment{})

	return &db, nil
}

func (db *Database) AddUser(nickname string, email string) {
	// TODO: Check if user exists

	user := model.User{
		Nickname: nickname,
		Email:    email,
	}

	db.orm.Create(&user)
}

func (db *Database) GetUser(email string) (model.User, error) {
	var user model.User

	if result := db.orm.Where("Email = ?", email).First(&user); result.Error != nil {
		return model.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}

func (db *Database) GetUsers() []model.User {
	var users []model.User

	if result := db.orm.Find(&users); result.Error != nil {
		return nil
	}

	return users
}

func (db *Database) UserExists(email string) bool {
	var user model.User

	if result := db.orm.Where("Email = ?", email).First(&user); result.RowsAffected == 0 {
		return false
	}

	return true
}

func (db *Database) AddComment(text string, user model.User) {
	comment := model.Comment{
		Comment:   text,
		Timestamp: time.Now(),
		UserID:    user.ID,
		User:      user,
	}

	db.orm.Create(&comment)
}

func (db *Database) GetComments() []model.Comment {
	var comments []model.Comment

	if result := db.orm.Preload(clause.Associations).Find(&comments); result.Error != nil {
		return nil
	}

	return comments
}

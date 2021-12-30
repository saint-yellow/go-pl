package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/saint-yellow/go-pl/toys/memorandum/cache"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model

	UserName string `gorm:"unique"`
	PasswordDigest string
}

type Task struct {
	gorm.Model

	UserID int
	Title string
	Status int `gorm:"default:'0'"`
	Content string `gorm:"type:longtext"`
	StartTime int64 `gorm:"default:'0'"`
	EndTime int64 `gorm:"default:'0'"`
}

const (
	passwordEncryptionCost = 8
)

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordEncryptionCost)
	if err != nil {
		return err
	}

	u.PasswordDigest = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

// 获取任务被浏览的次数
func (t *Task) BrowseCount() uint64 {
	countString, _ := cache.RedisClient.Get(cache.TaskViewKey(t.ID)).Result()
	count, _ := strconv.ParseUint(countString, 10, 64)
	return count
}

// 任务被浏览的次数+1
func (t *Task) AddBrowseCount() {
	cache.RedisClient.Incr(cache.TaskViewKey(t.ID))
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(t.ID)))
}
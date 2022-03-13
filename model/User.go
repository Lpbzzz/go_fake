package model

import (
	"encoding/base64"
	"fake_twitter/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20);not null" json:"username"`
	Password string `gorm:"type: varchar(20);not null" json:"password"`
	Role     int    `gorm:"type: int;not null" json:"role"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPW(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// ScryptPW 密码加密
func ScryptPW(password string) string {
	const KenLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 34, 45, 56, 67, 78, 89}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KenLen)
	if err != nil {
		log.Fatal(err)
	}
	Fpw := base64.StdEncoding.EncodeToString(HashPw)
	return Fpw
}

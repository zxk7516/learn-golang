package models

import (
	"time"
	"zxk/go-fiber-demo/common"

	"github.com/alexandrevicenzi/unchained"
)

type User struct {
	ID        uint      `gorm:"column:id;primary_key" json:"id"`
	Password  string    `gorm:"column:password" json:"-"`
	LastLogin time.Time `gorm:"column:last_login" json:"last_login"`
	UID       string    `gorm:"column:uid" json:"uid"`
	Username  string    `gorm:"column:username" json:"username"`
	Email     string    `gorm:"column:email" json:"email"`
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	Role      string    `gorm:"column:role" json:"role"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Activated int       `gorm:"column:activated" json:"-"`
	TimeStamps
}

func (u *User) TableName() string {
	return "user"
}

func GetUser(condition interface{}) (User, error) {
	var user User
	err := common.DB.Where(condition).First(&user).Error
	return user, err
}

func GetUsers(condition interface{}) ([]User, error) {
	var users []User
	err := common.DB.Where(condition).Find(&users).Error
	return users, err
}

func (u *User) CheckPassword(password string) bool {
	valid, _ := unchained.CheckPassword(password, u.Password)
	if valid {
		return true
	}
	return false
}

func (u *User) HashPassword(password string) string {
	hash, err := unchained.MakePassword(password, unchained.GetRandomString(12), "default")
	if err != nil {
		return hash
	}
	return ""
}

func (u *User) GetOrExtendToken() AuthToken {
	token, err := GetToken(&AuthToken{UID: u.UID})
	if err != nil {
		return *createTokenForUser(u.UID)
	} else if token.IsGoingToExpired() {
		token.extendToken()
		return token
	} else {
		return token
	}
}

func (u *User) SetPassword(password string) error {
	u.Password = u.HashPassword(password)
	return common.DB.Save(u).Error
}

type AuthToken struct {
	Key       string    `gorm:"column:key;primary_key" json:"key"`
	UID       string    `gorm:"column:uid" json:"uid"`
	Created   time.Time `gorm:"column:created" json:"created"`
	ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at"`
}

func (a *AuthToken) TableName() string {
	return "auth_token"
}
func GetToken(condition interface{}) (AuthToken, error) {
	var token AuthToken
	err := common.DB.Where(condition).First(&token).Error
	return token, err
}

func createTokenForUser(uid string) *AuthToken {

	key := unchained.GetRandomString(20)
	now := time.Now()
	token := &AuthToken{
		Key:       key,
		UID:       uid,
		Created:   now,
		ExpiredAt: now.AddDate(0, 0, common.GetConfig().TokenExpiredDays),
	}
	common.DB.Create(token)
	return token
}

func (a *AuthToken) extendToken() {
	now := time.Now()
	conf := common.GetConfig()
	common.DB.Delete(&a)
	a.Key = unchained.GetRandomString(20)
	a.ExpiredAt = now.AddDate(0, 0, conf.TokenExpiredDays)
	common.DB.Create(a)
}

func (a *AuthToken) HasExpired() bool {
	now := time.Now()
	return a.ExpiredAt.Before(now)
}

func (a *AuthToken) IsGoingToExpired() bool {
	conf := common.GetConfig()
	then := time.Now()
	then.AddDate(0, 0, conf.TokenRefreshDays)
	return a.ExpiredAt.Before(then)
}

func ValidteToken(key string) (User, error) {
	token, err := GetToken(&AuthToken{Key: key})
	if err == nil && !token.HasExpired() {
		var user, err = GetUser(&User{UID: token.UID})
		return user, err
	}
	return User{}, err
}

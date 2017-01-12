package models

import (
	"net/http"
	"errors"
	"ServerManagementSystem/util"
	"github.com/astaxie/beego/validation"
	"ServerManagementSystem/logs"
	"strconv"
	"encoding/json"
	"time"
	"strings"
	"fmt"
)

const USER_INFO_REDIS = "UserInfo"

const (
	USER_TYPE_NORMAL = 1 << iota
	USER_TYPE_SUPER_ADMIN
	USER_TYPE_NORMAL_ADMIN
)

type User struct {
	Id          int64  `description:"id" json:"id" xorm:"pk autoincr"`
	NickName    string `description:"昵称 最长50字符" json:"nickName" valid:"MaxSize(50)"`
	UserName    string `description:"登录名" json:"userName"`
	Password    string `description:"密码,加密存储" json:"password"`
	Salt        string `description:"加密后缀" json:"salt"`
	Created     int64  `description:"注册时间" xorm:"created"`
	Updated     int64  `description:"修改时间" xorm:"updated"`
	UserType    int    `description:"用户类型,1为超级管理员，0为普通用户" json:"userType"`
	Email       string `description:"邮箱地址" json:"email"`
	Mobile      string `description:"手机号" json:"mobile"`
	MobileCheck int    `description:"手机是否验证" json:"mobileCheck"`
	Avatar      string `description:"头像" json:"avatar"`
	UserState   int    `description:"用户状态" json:"userState"`
}

type UserDTO struct {
	Id          int64  `description:"id" json:"id" xorm:"pk autoincr"`
	NickName    string `description:"昵称 最长50字符" json:"nickName" valid:"MaxSize(50)"`
	UserName    string `description:"登录名" json:"userName"`
	Created     int64  `description:"注册时间" xorm:"created"`
	Mobile      string `description:"手机号" json:"mobile"`
	UserState   int    `description:"用户状态" json:"userState"`
}

type BaseUser interface {
	GetUserByUserName(userName string) (has bool, error error, errCode int)
	IUserSalt() string
	IUserPassword() string
	SetUserType(userType int)
	SetPassword(password string)
	SetUserName(userName string)
	SetSalt(salt string)
	AddUser() error
}

func User_ValidateLoginPwd(userName string, password string, user BaseUser, cacheUser BaseUser) (error, int) {
	if cacheUser == nil {
		has, _, _ := user.GetUserByUserName(userName)
		if !has {
			return errors.New("用户名或密码错误，请重新输入！"), SERVER_PHONE_NUM_OR_PASSWORD_ERROR
		}
	} else {
		user = cacheUser
	}
	hashPassword, pwdErr := util.EncryptPasswordWithSalt(password, user.IUserSalt())
	if pwdErr != nil {
		return pwdErr, http.StatusInternalServerError
	}
	if hashPassword != user.IUserPassword() {
		return errors.New("用户名或密码错误，请重新输入"), SERVER_PHONE_NUM_OR_PASSWORD_ERROR
	}
	return nil, 0

}

func User_Login(userName string, userPassword string, baseUser BaseUser) (err error, errCode int) {
	return User_ValidateLoginPwd(userName, userPassword, baseUser, nil)
}

func User_GetByUid(uid int64, isNeedCache bool) (user User, err error) {
	returnUser := User{}
	var has bool
	// 是否使用缓存
	if isNeedCache {
		redisTemp := RedisCache.Get(USER_INFO_REDIS + "_" + strconv.FormatInt(uid, 10))
		if redisTemp != nil {
			err = json.Unmarshal(redisTemp.([]byte), &returnUser)
			if err != nil {
				return returnUser, err
			}
			return returnUser, nil
		}
	}

	has, err = Orm.Where("id = ?", uid).Get(&returnUser)
	jsonByte, _ := json.Marshal(&returnUser)
	RedisCache.Put(USER_INFO_REDIS + "_" + strconv.FormatInt(returnUser.Id, 10), jsonByte, 60 * 60 * 24 * time.Second)
	if err != nil {
		return returnUser, err
	}

	if has {
		return returnUser, nil
	}
	return returnUser, errors.New("用户不存在")
}

// admin操作内容
func User_Regist(userName string, userPassword string, baseuser BaseUser) (err error, errCode int) {
	has, err, eCode := baseuser.GetUserByUserName(userName)
	if eCode == http.StatusInternalServerError {
		return err, eCode
	}
	if has {
		return errors.New("该用户已存在"), SERVER_USER_EXIST
	}

	baseuser.SetUserType(USER_TYPE_NORMAL)
	baseuser.SetPassword(userPassword)
	baseuser.SetUserName(userName)
	err, eCode = Validate(baseuser)
	if err != nil {
		return err, eCode
	}

	hashedPwd, salt, err := util.EncryptPassword(userPassword)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	baseuser.SetSalt(salt)
	baseuser.SetPassword(hashedPwd)

	err = baseuser.AddUser()

	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, 0
}

func User_GetListForPage(queryStr string) {// ([]User){
	var querySql = "";
	if strings.EqualFold("", queryStr){
		fmt.Print(querySql)
	}
}

func (this *User) GetUserByUserName(userName string) (has bool, error error, errCode int) {
	has, err := Orm.Where("user_name = ?", userName).Get(this)
	if err != nil {
		return false, err, http.StatusInternalServerError
	}

	if has {
		return true, nil, 0
	}
	return false, errors.New("用户不存在"), ServerUserNotExist
}

func (this *User) IUserSalt() string {
	return this.Salt
}

func (this *User) IUserPassword() string {
	return this.Password
}

func (this *User) SetUserType(userType int) {
	this.UserType = userType
}

func (this *User) SetPassword(password string) {
	this.Password = password
}

func (this *User) SetUserName(userName string) {
	this.UserName = userName
}

func (this *User) SetSalt(salt string) {
	this.Salt = salt
}

func (this *User) AddUser() error {
	_, err := Orm.InsertOne(this)
	if err != nil {
		logs.Logger.Error(err.Error())
		return errors.New(err.Error())
	}
	return nil
}

func Validate(data interface{}) (error error, errCode int) {
	valid := validation.Validation{}
	b, err := valid.Valid(data)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if !b {
		vErr := valid.Errors[0]
		logs.Logger.Errorf("message: %s Field: %s Key: %s Name: %s", vErr.Message, vErr.Field, vErr.Key, vErr.Name)
		return errors.New(vErr.Field + ":" + vErr.Message), http.StatusBadRequest
	}
	return nil, 0
}
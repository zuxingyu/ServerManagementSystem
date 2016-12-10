package util

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	"log"
	"net/http"
)


//获取盐
func SaltString() (salt string, error error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return fmt.Sprintf("%x", b), nil
}

//对密码进行sha256
func EncryptPasswordWithSalt(password, salt string) (hashedPwd string, error error) {
	sha_256 := sha256.New()
	_, err := sha_256.Write([]byte(password + salt))
	if err != nil {
		return "", errors.New(err.Error())
	}
	return fmt.Sprintf("%x", sha_256.Sum(nil)), nil
}

//加密密码 返回加密结果 以及使用的盐
func EncryptPassword(password string) (hashedPwd string, salt string, error error) {
	saltStr, err := SaltString()
	if err != nil {
		return "", "", errors.New("server err")
	}

	hashedPd, err := EncryptPasswordWithSalt(password, saltStr)
	if err != nil {
		return "", "", errors.New("server err")
	}
	return hashedPd, saltStr, nil
}

//使用beego的验证功能 验证输入
func Validate(data interface{}) (error error, errCode int) {
	valid := validation.Validation{}
	b, err := valid.Valid(data)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if !b {
		vErr := valid.Errors[0]
		log.Println(vErr.Message, vErr.Field, vErr.Key, vErr.Name)
		return errors.New(vErr.Field + ":" + vErr.Message), http.StatusBadRequest
	}
	return nil, 0
}
package service

import (
	"errors"
	"myDemo/config"
	"myDemo/model"

	"golang.org/x/crypto/bcrypt"
)

// Register 注册逻辑
func Register(username, password string) error {
	// 1. 检查用户名是否存在
	var count int64
	config.DB.Model(&model.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return errors.New("用户名已存在")
	}

	// 2. 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. 创建用户
	user := model.User{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Login 登录逻辑
func Login(username, password string) (string, error) {
	var user model.User

	// 1. 查询用户
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("用户不存在")
	}

	// 2. 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("密码错误")
	}

	// 这里简单返回成功信息，实际项目中通常返回 JWT Token
	return "登录成功，欢迎 " + user.Username, nil
}

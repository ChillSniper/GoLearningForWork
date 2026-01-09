package model

import "gorm.io/gorm"

// User 实体类
type User struct {
	gorm.Model        // 自动包含 ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"type:varchar(64);not null;unique" json:"username"` // 用户名唯一
	Password   string `gorm:"type:varchar(128);not null" json:"password"`       // 加密后的密码
}

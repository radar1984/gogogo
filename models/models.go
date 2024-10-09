package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Phone    string `gorm:"unique;not null"`
	Nickname string
	Age      int
	Avatar   string
}

type Character struct {
	gorm.Model
	UserID  uint
	Content string
	RefImg  string
	RefType int
	Cfg     string
	Tags    []Tag `gorm:"many2many:character_tags;"`
}

type Task struct {
	gorm.Model
	UserID    uint
	WorkerID  uint
	Type      int
	Status    int
	RefID     uint
	Cfg       string
	Config    TaskCfg `gorm:"embedded"`
}

type Tag struct {
	gorm.Model
	GroupID uint
	Name    string
	Key     string
	Value   string
}

type TaskCfg struct {
	Content string
	Extra   string
	Num     int
	RefImg  int
	RefType int
	Role    TaskCfgRole `gorm:"embedded"`
	Tags    []string    `gorm:"-"`
}

type TaskCfgRole struct {
	Content string
	Image   string
	Tags    []string `gorm:"-"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Type     int    `json:"type" binding:"required"`
}

type RegisterReq struct {
	Name          string `json:"name" binding:"required,min=5,max=12"`
	Nickname      string `json:"nickname" binding:"required,min=2,max=10"`
	Password      string `json:"passwd" binding:"required,min=8,max=20"`
	ConfirmPasswd string `json:"confirm_passwd" binding:"required,eqfield=Password"`
	Email         string `json:"email" binding:"required,email"`
	Phone         string `json:"phone" binding:"required"`
	Age           int    `json:"age" binding:"min=0,max=130"`
	Avatar        string `json:"avatar"`
}

type CreateCharacterReq struct {
	Content string   `json:"content"`
	RefImg  string   `json:"ref_img"`
	RefType int      `json:"ref_type"`
	Cfg     string   `json:"cfg"`
	Tags    []uint   `json:"tags"`
}

type UpdateWorkerTaskReq struct {
	TaskID   uint   `json:"task_id" binding:"required"`
	ImageURL string `json:"image_url" binding:"required"`
}
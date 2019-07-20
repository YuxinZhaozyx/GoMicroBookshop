package basic

import (
	"github.com/YuxinZhaozyx/GoMicroBookshop/basic/config"
	"github.com/YuxinZhaozyx/GoMicroBookshop/basic/db"
	"github.com/YuxinZhaozyx/GoMicroBookshop/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}

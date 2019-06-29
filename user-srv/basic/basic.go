package basic

import (
	"github.com/YuxinZhaozyx/GoMicroBookshop/user-srv/basic/config"
	"github.com/YuxinZhaozyx/GoMicroBookshop/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}

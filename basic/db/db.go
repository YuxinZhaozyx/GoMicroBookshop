package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/YuxinZhaozyx/GoMicroBookshop/basic/config"
	"github.com/micro/go-micro/util/log"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

// Init 初始化数据库
func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过了")
		log.Logf(err.Error())
		return
	}

	// 如果配置声明使用mysql
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}

	inited = true
}

// GetDB 获取数据库
func GetDB() *sql.DB {
	return mysqlDB
}

package mysql

import (
	"fmt"
	"web_app/settings"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init 初始化MySQL连接
func Init(cfg *settings.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("mysql connect error", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
	return
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}

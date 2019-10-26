package db

import (
	"fmt"
	"time"

	"github.com/spf13/viper"

	"database/sql"
	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// NativeDB ..
type NativeDB struct {
	Self *sql.DB // ci_institution_user_center 个人中心

}

// MyDB ...
var MyDB *NativeDB

func openDB(username, password, addr, name string) *sql.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := sql.Open("mysql", config)
	if err != nil {
		panic(err)

		// log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *sql.DB) {

	db.SetMaxOpenConns(viper.GetInt("db.max_open_conns")) //最大连接数 default 0
	db.SetMaxIdleConns(viper.GetInt("db.max_idle_conns")) //用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。default 2
	db.SetConnMaxLifetime(time.Hour * 1)                  //长连接有效期一个小时

}

// Init ...
func (db *NativeDB) Init() {
	MyDB = &NativeDB{
		Self: InitSelfDB(),
	}
}

// Close ...
func (db *NativeDB) Close() {
	MyDB.Self.Close()

}

// InitSelfDB ... used for cli
func InitSelfDB() *sql.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

package user_db_client

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/logger"
	"log"
)

// const (
// 	MysqlUsersUsername = "MysqlUsersUsername"
// 	MysqlUsersPassword = "MysqlUsersPassword"
// 	MysqlUsersHost     = "MysqlUsersHost"
// 	MysqlUsersSchema   = "MysqlUsersSchema"
// 	MysqlUsersPort     = "MysqlUsersPort"
// )

var (
	username = "root"           //os.Getenv(MysqlUsersUsername)
	password = ""               //os.Getenv(MysqlUsersPassword)
	host     = "localhost:3306" //os.Getenv(MysqlUsersHost)
	database = "db_go"          //os.Getenv(MysqlUsersSchema)
	// port     = ":3306"          // os.Getenv(MysqlUsersPort)
	ClientDb *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, database,
	)
	// datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC", "root", "", "localhost", 3306, "db_go")
	// datasourceName := "root:@/db_go?charset=utf8"
	var err error
	ClientDb, err = sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}
	if err = ClientDb.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger(logger.GetLogger())
	log.Println("Database succesfully configuration")

}

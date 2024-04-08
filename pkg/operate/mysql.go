package operate

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Mericusta/go-stp"
	"github.com/go-sql-driver/mysql"
)

var (
	SQL_QUERY_ERROR_FORMAT = "ERROR: execute SQL query occurs error,"
)

func connectMySQL(urlString string) (*sql.DB, error) {
	mysqlCfg, err := mysql.ParseDSN(urlString)
	if err != nil {
		return nil, err
	}

	mysqlCfg.ParseTime = true
	mysqlCfg.MultiStatements = true

	dbSession, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	return dbSession, err
}

func OperateMySQL(argURL, option, argFilepath string, args string) {
	dbSession, err := connectMySQL(argURL)
	if err != nil {
		fmt.Println("ERROR: connect mysql with url", argURL, "occurs error,", err.Error())
		return
	}

	switch option {
	case "import":
		importToMySQL(dbSession, argFilepath)
	case "truncate":
		truncateMySQL(dbSession, strings.Split(args, ",")...)
	}
}

func importToMySQL(dbSession *sql.DB, sqlFilePath string) {
	err := stp.ReadFileLineOneByOne(sqlFilePath, func(s string, l int) bool {
		_, err := dbSession.Exec(s)
		if err != nil {
			fmt.Println(SQL_QUERY_ERROR_FORMAT, err.Error())
		} else {
			fmt.Println("INFO: executed SQL,", s)
		}
		return true
	})
	if err != nil {
		fmt.Println("ERROR: read file line occurs error,", err.Error())
	}
}

func truncateMySQL(dbSession *sql.DB, tables ...string) {
	for _, table := range tables {
		truncateSQL := fmt.Sprintf("truncate `%v`;", table)
		_, err := dbSession.Exec(truncateSQL)
		if err != nil {
			fmt.Println(SQL_QUERY_ERROR_FORMAT, err.Error())
		} else {
			fmt.Println("INFO: executed SQL,", truncateSQL)
		}
	}
}

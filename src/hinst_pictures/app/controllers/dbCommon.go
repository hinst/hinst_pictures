package controllers

import "database/sql"
import _ "github.com/nakagami/firebirdsql"
import "github.com/revel/revel"

const globalDatabaseType = "firebirdsql"
const globalDatabaseKey = "hinst_pictures:hinst_pictures@localhost/hinst_pictures"
const databaseSizeQuery = "SELECT (MON$PAGE_SIZE * MON$PAGES) from MON$DATABASE"

func dbConnect() *sql.DB {
	conn, result := sql.Open(globalDatabaseType, globalDatabaseKey)
	if result == nil {
		return conn
	} else {
		revel.ERROR.Print(result)
		return nil
	}
}

func OpenTransaction() Transaction {
	var result Transaction
	result.connection = dbConnect()
	result.transaction = nil
	if result.connection != nil {
		var transaction, beginTransactionResult = result.connection.Begin()
		if beginTransactionResult == nil {
			result.transaction = transaction
		} else {
			revel.ERROR.Print(beginTransactionResult)
		}
	}
	return result
}

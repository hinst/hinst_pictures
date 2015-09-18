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

func OpenTransaction() *TTransaction {
	var result *TTransaction = nil
	var connection = dbConnect()
	if connection != nil {
		var transaction, beginTransactionResult = connection.Begin()
		if beginTransactionResult == nil {
			result = &TTransaction{}
			result.Tx = transaction
			result.Connection = connection
		} else {
			revel.ERROR.Print(beginTransactionResult)
		}
	}
	return result
}

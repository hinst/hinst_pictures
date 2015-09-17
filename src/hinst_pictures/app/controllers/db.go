package controllers

import "database/sql"
import _ "github.com/nakagami/firebirdsql"
import "github.com/revel/revel"

const globalDatabaseType = "firebirdsql"
const globalDatabaseKey = "hinst_pictures:hinst_pictures@localhost:D:/Dev/hinst_pictures/data/hinst_pictures.fdb"
const databaseSizeQuery = "SELECT MON$DATABASE_NAME, (MON$PAGE_SIZE * MON$PAGES)"

type DB struct {
	*revel.Controller
}

func (c DB) TestDB() revel.Result {
	return c.Render()
}

func dbConnect() *sql.DB {
	conn, result := sql.Open(globalDatabaseType, globalDatabaseKey)
	if result == nil {
		return conn
	} else {
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
		}
	}
	return result
}

package controllers

import "github.com/revel/revel"

type DB struct {
	*revel.Controller
}

func (c DB) Test() revel.Result {
	var transaction = OpenTransaction()
	var databaseSizeRow = transaction.transaction.QueryRow(databaseSizeQuery)
	var databaseSize int
	databaseSizeRow.Scan(databaseSize)
	c.Args["databaseSize"] = formatInteger(databaseSize)
	transaction.Close()
	return c.Render()
}

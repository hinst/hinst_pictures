package controllers

import "github.com/revel/revel"

type DB struct {
	*revel.Controller
}

func (c DB) Test() revel.Result {
	var transaction = OpenTransaction()
	var databaseSizeRow = transaction.transaction.QueryRow(databaseSizeQuery)
	var databaseSize int
	databaseSizeRow.Scan(&databaseSize)
	var databaseSizeText = formatInteger(databaseSize)
	c.RenderArgs["databaseSize"] = databaseSizeText
	transaction.Close()
	return c.Render()
}

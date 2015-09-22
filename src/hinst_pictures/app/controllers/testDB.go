package controllers

import "github.com/revel/revel"

func (this *App) TestDB() revel.Result {
	var databaseSizeRow = this.Transaction.QueryRow(databaseSizeQuery)
	var databaseSize int
	databaseSizeRow.Scan(&databaseSize)
	var databaseSizeText = formatInteger(databaseSize)
	this.RenderArgs["databaseSize"] = databaseSizeText
	this.RenderArgs["usersTableExists"] = this.GetUsersTable().CheckTableExists()
	return this.Render()
}

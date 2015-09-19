package controllers

import "github.com/revel/revel"
import db "hinst_db"

type App struct {
	*revel.Controller
	Transaction *db.TTransaction
}

func (this *App) Index() revel.Result {
	return this.Render()
}

func (this *App) prepareDB() revel.Result {
	this.Transaction = OpenTransaction()
	if this.Transaction != nil {
		return nil
	} else {
		revel.ERROR.Println("Could not obtain database transaction")
		return this.RenderTemplate("dbError.html")
	}
}

func (this *App) prepare() revel.Result {
	var result = this.prepareDB()
	if result != nil {
		revel.ERROR.Println("Database error")
		return result
	}
	return result
}

func (this *App) finalize() revel.Result {
	if this.Transaction != nil {
		this.Transaction.Close()
		this.Transaction = nil
	}
	return nil
}

func init() {
	revel.InterceptMethod((*App).prepare, revel.BEFORE)
	revel.InterceptMethod((*App).finalize, revel.AFTER)
}

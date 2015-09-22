package controllers

import "github.com/revel/revel"
import "hinst_db"

type App struct {
	*revel.Controller
	Transaction *hinst_db.TTransaction
	User        *TUserRow
	userTable   *TUserTable
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

func (this *App) checkUser() {
	var cookies = this.Request.Cookies()
	for i := range cookies {
		var cookie = cookies[i]
		if cookie.Name == UserNameCookieName {

		}
	}
}

func (this *App) GetUsersTable() *TUserTable {
	if nil == this.userTable {
		this.userTable = CreateUserTable()
		this.userTable.Transaction = this.Transaction.Tx
	}
	return this.userTable
}

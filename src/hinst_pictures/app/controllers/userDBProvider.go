package controllers

import "hinst_db"

type TUserTable struct {
	hinst_db.TTable
}

func CreateUserTable() (result *TUserTable) {
	result = &TUserTable{}
	result.RowConstructor = CreateUserRowAsIRow
	result.TableName = "Users"
	return
}

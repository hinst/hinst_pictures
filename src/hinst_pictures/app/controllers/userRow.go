package controllers

import "hinst_db"

import "time"

type TUserRow struct {
	Name         string
	Password     string
	Admin        int
	CreationDate time.Time
}

func CreateUserRow() *TUserRow {
	return &TUserRow{}
}

func (this *TUserRow) GetFields() []hinst_db.TField {
	return []hinst_db.TField{
		hinst_db.TField{Field: &this.Name, Name: "Name"},
		hinst_db.TField{Field: &this.Password, Name: "Password"},
		hinst_db.TField{Field: &this.Admin, Name: "Admin"},
		hinst_db.TField{Field: &this.CreationDate, Name: "CreationDate"}}
}

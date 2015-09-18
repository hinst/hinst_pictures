package controllers

import "database/sql"

type TTransaction struct {
	*sql.Tx
	Connection *sql.DB
}

func (this *TTransaction) Close() {
	this.Rollback()
	if this.Connection != nil {
		this.Connection.Close()
		this.Connection = nil
	}
}

func (this *TTransaction) Commit() {
	this.Commit()
	if this.Connection != nil {
		this.Connection.Close()
		this.Connection = nil
	}
}

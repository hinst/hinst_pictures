package controllers

import "database/sql"

type Transaction struct {
	connection  *sql.DB
	transaction *sql.Tx
}

func (this *Transaction) Close() {
	if this.transaction != nil {
		this.transaction.Rollback()
		this.transaction = nil
	}
	if this.connection != nil {
		this.connection.Close()
		this.connection = nil
	}
}

func (this *Transaction) Commit() {
	if this.transaction != nil {
		this.transaction.Commit()
		this.transaction = nil
	}
	if this.connection != nil {
		this.connection.Close()
		this.connection = nil
	}
}

func (this *Transaction) Valid() bool {
	return (this.connection != nil) && (this.transaction != nil)
}

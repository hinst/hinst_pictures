package hinst_db

import "database/sql"

type TTable struct {
	RowConstructor func() IRow
	Transaction    *sql.Tx
	TableName      string
	Where          string
	Rows           []IRow
}

func (this *TTable) Load() {
	this.Rows = nil
	var query string = "select " + GetFieldsString(this.RowConstructor()) + " from '" + this.TableName + "'"
	if len(this.Where) > 0 {
		query = query + " where " + this.Where
	}
	var rows, queryResult = this.Transaction.Query(query)
	if queryResult == nil {
		defer rows.Close()
		for rows.Next() {
			var row = this.RowConstructor()
			row.GetFields()
		}
	}
}

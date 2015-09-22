package hinst_db

import "database/sql"
import "errors"

type TTable struct {
	RowConstructor func() IRow
	Transaction    *sql.Tx
	TableName      string
	Where          string
	Rows           []IRow
}

func (this *TTable) assertTransactionAssigned() {
	if this.Transaction == nil {
		panic(errors.New("Transaction not assigned"))
	}
}

func (this *TTable) assertTableNameAssigned() {
	if this.TableName == "" {
		panic(errors.New(""))
	}
}

func (this *TTable) Load() {
	this.Rows = []IRow{}
	if this.RowConstructor == nil {
		panic(errors.New("RowConstructor not assigned"))
	}
	this.assertTransactionAssigned()
	this.assertTableNameAssigned()
	var query = "select " + GetFieldsStringFromRow(this.RowConstructor()) + " from \"" + this.TableName + "\""
	if len(this.Where) > 0 {
		query = query + " where " + this.Where
	}
	var rows, queryResult = this.Transaction.Query(query)
	if queryResult == nil {
		defer rows.Close()
		for rows.Next() {
			var row = this.RowConstructor()
			var fields = row.GetFields()
			var fieldFaces = GetScanInterfacesFromFields(fields)
			rows.Scan(fieldFaces...)
			this.Rows = append(this.Rows, row)
		}
	}
}

// Using Firebird "update or insert" SQL feature.
func (this *TTable) Save() {
	if len(this.Rows) > 0 {
		var firstRow = this.Rows[0]
		var statementText = "update or insert into \"" + this.TableName + "\" " +
			"(" + GetFieldsStringFromRow(firstRow) + ") " +
			"values (" + GetValuesTemplateStringFromRow(firstRow) + ")"
		var statement, statementResult = this.Transaction.Prepare(statementText)
		if statementResult == nil {
			defer statement.Close()
			for i := range this.Rows {
				var row = this.Rows[i]
				var fields = row.GetFields()
				var fieldFaces = GetScanInterfacesFromFields(fields)
				statement.Exec(fieldFaces...)
			}
		}
	}
}

// Using Firebird RDB$RELATIONS table.
func (this *TTable) CheckTableExists() bool {
	var result = false
	this.assertTableNameAssigned()
	this.assertTransactionAssigned()
	var query = "select 1 from RDB$RELATIONS where RDB$RELATION_NAME=\"" + this.TableName + "\""
	var rows, queryResult = this.Transaction.Query(query)
	if queryResult == nil {
		defer rows.Close()
		if rows.Next() {
			result = true
		}
	}
	return result
}

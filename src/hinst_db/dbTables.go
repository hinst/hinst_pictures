package hinst_db

import "database/sql"

type TTable struct {
	RowConstructor func() IRow
}

func (table *TTable) Load(table, where string) {
	var query string = "select"
}

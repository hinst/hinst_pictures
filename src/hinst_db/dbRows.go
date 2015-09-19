package hinst_db

import "database/sql"

type TField struct {
	Field *interface{}
	Name  string
}

type IRow interface {
	GetFields() []TField
}

func GetFieldsString(row IRow) string {
	var result = ""
	var fields = row.GetFields()
	for i := 0; i < len(fields); i++ {
		result = result + "'" + fields[i].Name + "'"
		if i < len(fields)-1 {
			result = result + ", "
		}
	}
	return result
}

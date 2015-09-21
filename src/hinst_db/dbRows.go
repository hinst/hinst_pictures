package hinst_db

type TField struct {
	Field interface{}
	Name  string
}

func GetScanInterfacesFromFields(source []TField) (result []interface{}) {
	var n = len(source)
	result = make([]interface{}, n)
	for i := range source {
		result[i] = source[i]
	}
	return
}

type IRow interface {
	GetFields() []TField
}

func GetFieldsStringFromRow(row IRow) string {
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

func GetValuesTemplateStringFromRow(row IRow) string {
	var result = ""
	var fields = row.GetFields()
	for i := 0; i < len(fields); i++ {
		result = result + "?"
		if i < len(fields)-1 {
			result = result + ", "
		}
	}
	return result
}

package row

type Row struct {
	data map[string]interface{}
}

func NewRow(data map[string]interface{}) *Row {
	return &Row{
		data: data,
	}
}

func (r *Row) GetData() map[string]interface{} {
	return r.data
}

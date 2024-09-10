package table

import (
	"errors"
	"fmt"
	c "inMemoryDB/internal/column"
	r "inMemoryDB/internal/row"
	"inMemoryDB/internal/utils"
)

type Table struct {
	id      string
	name    string
	columns map[string]*c.Column
	rows    []*r.Row
}

func NewTable(name string, columns map[string]*c.Column) *Table {
	table := &Table{
		id:      utils.GenerateUUID(),
		name:    name,
		columns: columns,
		rows:    make([]*r.Row, 0),
	}

	return table
}

func (table *Table) GetID() string {
	return table.id
}

func (table *Table) InsertRow(rowData map[string]interface{}) error {
	row := r.NewRow(rowData)

	err := table.ValidateRow(row)
	if err != nil {
		return err
	}

	table.rows = append(table.rows, row)
	return nil
}

func (table *Table) ValidateRow(row *r.Row) error {
	rowData := row.GetData()
	for columnID, column := range table.columns {

		val, ok := rowData[columnID]
		if !ok && column.IsRequired() {
			return errors.New("missing required column")
		}

		err := column.ValidateRowData(val)
		if err != nil {
			return err
		}
	}

	return nil
}

func (table *Table) PrintAllRows() {
	fmt.Printf("Table: %s\n", table.name)
	for _, col := range table.columns {
		fmt.Printf("%s\t", col.GetName())
	}
	fmt.Println()

	for _, row := range table.rows {
		for columnID := range table.columns {
			rowData := row.GetData()
			fmt.Printf("%v\t", rowData[columnID])
		}
		fmt.Println()
	}
}

func (table *Table) FilterRows(columnID string, value interface{}) {
	filteredRows := make([]*r.Row, 0)
	for _, row := range table.rows {
		rowData := row.GetData()
		if val, ok := rowData[columnID]; ok {
			if val == value {
				filteredRows = append(filteredRows, row)
			}
		}
	}

	if len(filteredRows) == 0 {
		fmt.Println("No matching rows found")
	}

	for _, row := range filteredRows {
		rowData := row.GetData()
		fmt.Printf("%v: %v\t", columnID, rowData[columnID])
	}
}

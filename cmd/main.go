package main

import (
	"fmt"
	c "inMemoryDB/internal/column"
	ct "inMemoryDB/internal/constraint"
	"inMemoryDB/internal/datatype"
	"inMemoryDB/internal/db"
	"inMemoryDB/internal/table"
)

func main() {
	// create DB
	db := db.NewDB("sql-like db")

	// create columns
	columns := make(map[string]*c.Column, 4)
	columns["name"] = c.NewColumn("name", datatype.String(8), ct.NotEmptyConstraint(), true)
	columns["rollNo"] = c.NewColumn("rollNo", datatype.Int(-1024, 1024), nil, true)
	columns["rank"] = c.NewColumn("rank", datatype.Int(-1024, 1024), nil, false)
	columns["percentage"] = c.NewColumn("percentage", datatype.Int(-1024, 1024), nil, false)

	// add table with columns
	db.AddTable("students", columns)

	table := db.GetTable("students")

	// insert rows
	insertRows(table, map[string]interface{}{
		"name": "abc", "rollNo": 12, "rank": 3, "percentage": 95,
	})
	insertRows(table, map[string]interface{}{
		"name": "def", "rollNo": 1025, "rank": 4, "percentage": 99,
	})
	insertRows(table, map[string]interface{}{
		"name": "ghi", "rollNo": -13, "rank": 5, "percentage": 67,
	})
	insertRows(table, map[string]interface{}{
		"rollNo": 13, "rank": 6, "percentage": 68,
	})
	insertRows(table, map[string]interface{}{
		"name": "", "rollNo": 13, "rank": 6, "percentage": 68,
	})

	// filter rows
	table.FilterRows("rank", 1)

}

func insertRows(table *table.Table, rowData map[string]interface{}) {
	err := table.InsertRow(rowData)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Row inserted successfully\n")
	}
}

package db

import (
	"errors"
	"inMemoryDB/internal/column"
	"inMemoryDB/internal/table"
	"inMemoryDB/internal/utils"
)

type DB struct {
	id     string
	name   string
	tables map[string]*table.Table
}

func NewDB(name string) *DB {
	db := &DB{
		id:     utils.GenerateUUID(),
		name:   name,
		tables: map[string]*table.Table{},
	}

	return db
}

func (db *DB) AddTable(name string, columns map[string]*column.Column) {
	// create Table
	table := table.NewTable(name, columns)

	db.tables[name] = table
}

func (db *DB) GetTable(name string) *table.Table {
	if table, ok := db.tables[name]; ok {
		return table
	}
	return nil
}

func (db *DB) DeleteTable(name string) error {
	if _, ok := db.tables[name]; ok {
		return errors.New("table does not exisit")
	}
	delete(db.tables, name)
	return nil
}
package column

import (
	"fmt"
	"inMemoryDB/internal/utils"

	ct "inMemoryDB/internal/constraint"
	dt "inMemoryDB/internal/datatype"
)

type Column struct {
	id         string
	name       string
	dataType   dt.DataType
	constraint ct.Constraint
	required   bool
}

func NewColumn(name string, dataType dt.DataType, constraint ct.Constraint, required bool) *Column {
	return &Column{
		id:         utils.GenerateUUID(),
		name:       name,
		dataType:   dataType,
		constraint: constraint,
		required:   required,
	}
}

func (c *Column) IsRequired() bool {
	return c.required
}

func (c *Column) GetName() string {
	return c.name
}

func (c *Column) ValidateRowData(val interface{}) error {
	err := c.dataType.Validate(val)
	if err != nil {
		return err
	}

	// validate constraint if present
	if c.constraint != nil {
		err := c.constraint.Validate(val)
		if err != nil {
			return fmt.Errorf("constraint failed: %v", err)
		}
	}

	return nil
}

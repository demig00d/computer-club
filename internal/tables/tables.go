package tables

type Tables struct {
	PricePerHour int
	tables       []*Table
}

func NewTables(pricePerHour, maxTables int) Tables {
	return Tables{
		PricePerHour: pricePerHour,
		tables:       CreateN(maxTables),
	}
}

func (t Tables) ForEach(f func(*Table)) {
	for _, table := range t.tables {
		f(table)
	}
}

func (t Tables) GetTable(tableId int) *Table {
	var resultTable *Table
	t.ForEach(func(table *Table) {
		if table.Id == tableId {
			resultTable = table
		}
	})

	return resultTable
}

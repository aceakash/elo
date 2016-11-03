package elo

// Save persists the table to the provided ITableStore.
type TableStore interface {
	Save(*Table) error
	Load() (*Table, error)
}

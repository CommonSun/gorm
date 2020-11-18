package gorm

// Instrumenter interface to help instrument queries
type Instrumenter interface {
	Instrument(db *DB) *DB
}

// Instrument allows an instrumenter to manipulate the db object
func (s *DB) Instrument(i Instrumenter) *DB {
	return i.Instrument(s.clone())
}

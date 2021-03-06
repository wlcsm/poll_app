package db

// Constant errors https://dave.cheney.net/tag/errors
type Error string

const (
	DBFull   = Error("Database is full")
	NotFound = Error("Data not found")
)

func (e Error) Error() string {
	return string(e)
}

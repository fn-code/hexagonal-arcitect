package psql

import "errors"

var (
	// ErrDataNotFound is when data in tabel dosent exit
	ErrDataNotFound = errors.New("data is not found")
)

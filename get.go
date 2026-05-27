// Copyright 2019 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badgerhold

import (
	"errors"

	"github.com/dgraph-io/badger/v4"
)

// ErrNotFound is returned when no data is found for the given key
var ErrNotFound = errors.New("No data found for this key")

// Get retrieves a value from badgerhold and puts it into result.  Result must be a pointer
func (s *Store) Get(key, result interface{}) error { _ = "STUB: not implemented"; return nil }

// TxGet allows you to pass in your own badger transaction to retrieve a value from the badgerhold and puts it
// into result
func (s *Store) TxGet(tx *badger.Txn, key, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Find retrieves a set of values from the badgerhold that matches the passed in query
// result must be a pointer to a slice.
// The result of the query will be appended to the passed in result slice, rather than the passed in slice being
// emptied.
func (s *Store) Find(result interface{}, query *Query) error { _ = "STUB: not implemented"; return nil }

// TxFind allows you to pass in your own badger transaction to retrieve a set of values from the badgerhold
func (s *Store) TxFind(tx *badger.Txn, result interface{}, query *Query) error {
	_ = "STUB: not implemented"
	return nil
}

// FindOne returns a single record, and so result is NOT a slice, but an pointer to a struct, if no record is found
// that matches the query, then it returns ErrNotFound
func (s *Store) FindOne(result interface{}, query *Query) error {
	_ = "STUB: not implemented"
	return nil
}

// TxFindOne allows you to pass in your own badger transaction to retrieve a single record from the badgerhold
func (s *Store) TxFindOne(tx *badger.Txn, result interface{}, query *Query) error {
	_ = "STUB: not implemented"
	return nil
}

// Count returns the current record count for the passed in datatype
func (s *Store) Count(dataType interface{}, query *Query) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TxCount returns the current record count from within the given transaction for the passed in datatype
func (s *Store) TxCount(tx *badger.Txn, dataType interface{}, query *Query) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ForEach runs the function fn against every record that matches the query
// Useful for when working with large sets of data that you don't want to hold the entire result
// set in memory, similar to database cursors
// Return an error from fn, will stop the cursor from iterating
func (s *Store) ForEach(query *Query, fn interface{}) error { _ = "STUB: not implemented"; return nil }

// TxForEach is the same as ForEach but you get to specify your transaction
func (s *Store) TxForEach(tx *badger.Txn, query *Query, fn interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

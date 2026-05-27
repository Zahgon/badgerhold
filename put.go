// Copyright 2019 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badgerhold

import (
	"errors"

	"github.com/dgraph-io/badger/v4"
)

// ErrKeyExists is the error returned when data is being Inserted for a Key that already exists
var ErrKeyExists = errors.New("This Key already exists in badgerhold for this type")

// ErrUniqueExists is the error thrown when data is being inserted for a unique constraint value that already exists
var ErrUniqueExists = errors.New("This value cannot be written due to the unique constraint on the field")

// sequence tells badgerhold to insert the key as the next sequence in the bucket
type sequence struct{}

// NextSequence is used to create a sequential key for inserts
// Inserts a uint64 as the key
// store.Insert(badgerhold.NextSequence(), data)
func NextSequence() interface{} {
	_ = "STUB: not implemented"

	// Insert inserts the passed in data into the badgerhold
	//
	// If the key already exists in the badgerhold, then an ErrKeyExists is returned
	// If the data struct has a field tagged as `badgerholdKey` and it is the same type
	// as the Insert key, AND the data struct is passed by reference, AND the key field
	// is currently set to the zero-value for that type, then that field will be set to
	// the value of the insert key.
	//
	// To use this with badgerhold.NextSequence() use a type of `uint64` for the key field.
	return nil
}

func (s *Store) Insert(key, data interface{}) error { _ = "STUB: not implemented"; return nil }

// TxInsert is the same as Insert except it allows you to specify your own transaction
func (s *Store) TxInsert(tx *badger.Txn, key, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// insert data

// insert any indexes

// Update updates an existing record in the badgerhold
// if the Key doesn't already exist in the store, then it fails with ErrNotFound
func (s *Store) Update(key interface{}, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// TxUpdate is the same as Update except it allows you to specify your own transaction
func (s *Store) TxUpdate(tx *badger.Txn, key interface{}, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// delete any existing indexes

// put data

// insert any new indexes

// Upsert inserts the record into the badgerhold if it doesn't exist.  If it does already exist, then it updates
// the existing record
func (s *Store) Upsert(key interface{}, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// TxUpsert is the same as Upsert except it allows you to specify your own transaction
func (s *Store) TxUpsert(tx *badger.Txn, key interface{}, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// existing entry found
// delete any existing indexes

// existing entry not found

// put data

// insert any new indexes

// UpdateMatching runs the update function for every record that match the passed in query
// Note that the type  of record in the update func always has to be a pointer
func (s *Store) UpdateMatching(dataType interface{}, query *Query, update func(record interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

// TxUpdateMatching does the same as UpdateMatching, but allows you to specify your own transaction
func (s *Store) TxUpdateMatching(tx *badger.Txn, dataType interface{}, query *Query,
	update func(record interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Copyright 2019 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badgerhold

import (
	"github.com/dgraph-io/badger/v4"
)

const indexPrefix = "_bhIndex"

// size of iterator keys stored in memory before more are fetched
const iteratorKeyMinCacheSize = 100

// Index is a function that returns the indexable, encoded bytes of the passed in value
type Index struct {
	IndexFunc func(name string, value interface{}) ([]byte, error)
	Unique    bool
}

// adds an item to the index
func (s *Store) indexAdd(storer Storer, tx *badger.Txn, key []byte, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// removes an item from the index
// be sure to pass the data from the old record, not the new one
func (s *Store) indexDelete(storer Storer, tx *badger.Txn, key []byte, originalData interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// adds or removes a specific index on an item
func (s *Store) indexUpdate(typeName, indexName string, index Index, tx *badger.Txn, key []byte, value interface{},
	delete bool) error {
	_ = "STUB: not implemented"
	return nil
}

// indexKeyPrefix returns the prefix of the badger key where this index is stored
func indexKeyPrefix(typeName, indexName string) []byte { _ = "STUB: not implemented"; return nil }

// newIndexKey returns the badger key where this index is stored
func newIndexKey(typeName, indexName string, value []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// KeyList is a slice of unique, sorted keys([]byte) such as what an index points to
type KeyList [][]byte

func (v *KeyList) add(key []byte) { _ = "STUB: not implemented"; return }

// already added

func (v *KeyList) remove(key []byte) { _ = "STUB: not implemented"; return }

func (v *KeyList) in(key []byte) bool { _ = "STUB: not implemented"; return false }

func indexExists(it *badger.Iterator, typeName, indexName string) bool {
	_ = "STUB: not implemented"
	return false
}

// test if any data exists for type

// store is empty for this data type so the index could possibly exist
// we don't want to fail on a "bad index" because they could simply be running a query against
// an empty dataset

// test if an index exists

type iterator struct {
	keyCache [][]byte
	nextKeys func(*badger.Iterator) ([][]byte, error)
	iter     *badger.Iterator
	bookmark *iterBookmark
	lastSeek []byte
	tx       *badger.Txn
	err      error
}

// iterBookmark stores a seek location in a specific iterator
// so that a single RW iterator can be shared within a single transaction
type iterBookmark struct {
	iter    *badger.Iterator
	seekKey []byte
}

func (s *Store) newIterator(tx *badger.Txn, typeName string, query *Query, bookmark *iterBookmark) *iterator {
	_ = "STUB: not implemented"
	return nil
}

// can't use indexes on matchFuncs as the entire record isn't available for testing in the passed
// in function

// Key field or index not specified - test key against criteria (if it exists) or return everything

// nothing to check return key for value testing

// indexed field, get keys from index

// no currentRow on indexes as it refers to multiple rows
// remove index prefix for matching

// append the slice of keys stored in the index

func (i *iterator) createBookmark() *iterBookmark { _ = "STUB: not implemented"; return nil }

// Next returns the next key value that matches the iterators criteria
// If no more kv's are available the return nil, if there is an error, they return nil
// and iterator.Error() will return the error
func (i *iterator) Next() (key []byte, value []byte) { _ = "STUB: not implemented"; return nil, nil }

// Error returns the last error, iterator.Next() will not continue if there is an error present
func (i *iterator) Error() error { _ = "STUB: not implemented"; return nil }

func (i *iterator) Close() { _ = "STUB: not implemented"; return }

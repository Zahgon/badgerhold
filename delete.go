// Copyright 2019 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badgerhold

import (
	"github.com/dgraph-io/badger/v4"
)

// Delete deletes a record from the badgerhold, datatype just needs to be an example of the type stored so that
// the proper bucket and indexes are updated
func (s *Store) Delete(key, dataType interface{}) error { _ = "STUB: not implemented"; return nil }

// TxDelete is the same as Delete except it allows you to specify your own transaction
func (s *Store) TxDelete(tx *badger.Txn, key, dataType interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// delete data

// remove any indexes

// DeleteMatching deletes all the records that match the passed in query
func (s *Store) DeleteMatching(dataType interface{}, query *Query) error {
	_ = "STUB: not implemented"
	return nil
}

// TxDeleteMatching does the same as DeleteMatching, but allows you to specify your own transaction
func (s *Store) TxDeleteMatching(tx *badger.Txn, dataType interface{}, query *Query) error {
	_ = "STUB: not implemented"
	return nil
}

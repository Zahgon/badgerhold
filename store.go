// Copyright 2019 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badgerhold

import (
	"reflect"
	"sync"

	"github.com/dgraph-io/badger/v4"
)

const (
	// BadgerHoldIndexTag is the struct tag used to define a field as indexable for a badgerhold
	BadgerHoldIndexTag = "badgerholdIndex"

	// BadgerholdKeyTag is the struct tag used to define a field as a key for use in a Find query
	BadgerholdKeyTag = "badgerholdKey"

	// badgerholdPrefixTag is the prefix for an alternate (more standard) version of a struct tag
	badgerholdPrefixTag         = "badgerhold"
	badgerholdPrefixIndexValue  = "index"
	badgerholdPrefixKeyValue    = "key"
	badgerholdPrefixUniqueValue = "unique"
)

// Store is a badgerhold wrapper around a badger DB
type Store struct {
	db               *badger.DB
	sequenceBandwith uint64
	sequences        *sync.Map

	encode EncodeFunc
	decode DecodeFunc
}

// Options allows you set different options from the defaults
// For example the encoding and decoding funcs which default to Gob
type Options struct {
	Encoder          EncodeFunc
	Decoder          DecodeFunc
	SequenceBandwith uint64
	badger.Options
}

// DefaultOptions are a default set of options for opening a BadgerHold database
// Includes badgers own default options
var DefaultOptions = Options{
	Options:          badger.DefaultOptions(""),
	Encoder:          DefaultEncode,
	Decoder:          DefaultDecode,
	SequenceBandwith: 100,
}

// Open opens or creates a badgerhold file.
func Open(options Options) (*Store, error) { _ = "STUB: not implemented"; return nil, nil }

// Badger returns the underlying Badger DB the badgerhold is based on
func (s *Store) Badger() *badger.DB {
	_ = "STUB: not implemented"

	// Close closes the badger db
	return nil
}

func (s *Store) Close() error { _ = "STUB: not implemented"; return nil }

/*
	NOTE: Not going to implement ReIndex and Remove index
	I had originally created these to make the transition from a plain bolt or badger DB easier
	but there is too much chance for lost data, and it's probably better that any conversion be
	done by the developer so they can directly manage how they want data to be migrated.
	If you disagree, feel free to open an issue and we can revisit this.
*/

// Storer is the Interface to implement to skip reflect calls on all data passed into the badgerhold
type Storer interface {
	Type() string              // used as the badgerdb index prefix
	Indexes() map[string]Index //[indexname]indexFunc
}

// anonType is created from a reflection of an unknown interface
type anonStorer struct {
	rType   reflect.Type
	indexes map[string]Index
}

// Type returns the name of the type as determined from the reflect package
func (t *anonStorer) Type() string { _ = "STUB: not implemented"; return "" }

// Indexes returns the Indexes determined by the reflect package on this type
func (t *anonStorer) Indexes() map[string]Index {
	_ = "STUB: not implemented"

	// newStorer creates a type which satisfies the Storer interface based on reflection of the passed in dataType
	// if the Type doesn't meet the requirements of a Storer (i.e. doesn't have a name) it panics
	// You can avoid any reflection costs, by implementing the Storer interface on a type
	return nil
}

func (s *Store) newStorer(dataType interface{}) Storer {
	_ = "STUB: not implemented"
	return *new(Storer)
}

// indexName is stored canonically as the field name NOT the name in the tag

func (s *Store) getSequence(typeName string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func typePrefix(typeName string) []byte { _ = "STUB: not implemented"; return nil }

func getKeyField(tp reflect.Type) (reflect.StructField, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.StructField), false
}

func newElemType(datatype interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// makes sure that interface your working with is not a pointer
func getElem(value interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// Copyright 2019 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badgerhold

// EncodeFunc is a function for encoding a value into bytes
type EncodeFunc func(value interface{}) ([]byte, error)

// DecodeFunc is a function for decoding a value from bytes
type DecodeFunc func(data []byte, value interface{}) error

// DefaultEncode is the default encoding func for badgerhold (Gob)
func DefaultEncode(value interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// DefaultDecode is the default decoding func for badgerhold (Gob)
func DefaultDecode(data []byte, value interface{}) error { _ = "STUB: not implemented"; return nil }

// encodeKey encodes key values with a type prefix which allows multiple different types
// to exist in the badger DB
func (s *Store) encodeKey(key interface{}, typeName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decodeKey decodes the key value and removes the type prefix
func (s *Store) decodeKey(data []byte, key interface{}, typeName string) error {
	_ = "STUB: not implemented"
	return nil
}

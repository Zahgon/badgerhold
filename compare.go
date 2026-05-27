// Copyright 2019 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badgerhold

// ErrTypeMismatch is the error thrown when two types cannot be compared
type ErrTypeMismatch struct {
	Value interface{}
	Other interface{}
}

func (e *ErrTypeMismatch) Error() string { _ = "STUB: not implemented"; return "" }

// Comparer compares a type against the encoded value in the store. The result should be 0 if current==other,
// -1 if current < other, and +1 if current > other.
// If a field in a struct doesn't specify a comparer, then the default comparison is used (convert to string and compare)
// this interface is already handled for standard Go Types as well as more complex ones such as those in time and big
// an error is returned if the type cannot be compared
// The concrete type will always be passed in, not a pointer
type Comparer interface {
	Compare(other interface{}) (int, error)
}

func (c *Criterion) compare(rowValue, criterionValue interface{}, currentRow interface{}) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func compare(value, other interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

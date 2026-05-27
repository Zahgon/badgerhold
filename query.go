// Copyright 2019 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badgerhold

import (
	"reflect"
	"regexp"

	"github.com/dgraph-io/badger/v4"
)

const (
	eq    = iota // ==
	ne           // !=
	gt           // >
	lt           // <
	ge           // >=
	le           // <=
	in           // in
	re           // regular expression
	fn           // func
	isnil        // test's for nil
	sw           // string starts with
	ew           // string ends with
	hk           // match map keys

	contains // slice only
	any      // slice only
	all      // slice only
)

// Key is shorthand for specifying a query to run again the Key in a badgerhold, simply returns ""
// Where(badgerhold.Key).Eq("testkey")
const Key = ""

// Query is a chained collection of criteria of which an object in the badgerhold needs to match to be returned
// an empty query matches against all records
type Query struct {
	index         string
	currentField  string
	fieldCriteria map[string][]*Criterion
	ors           []*Query

	badIndex bool
	dataType reflect.Type
	tx       *badger.Txn
	writable bool
	subquery bool
	bookmark *iterBookmark

	limit   int
	skip    int
	sort    []string
	reverse bool
}

// Slice turns a slice of any type into []interface{} by copying the slice values so it can be easily passed
// into queries that accept variadic parameters.
// Will panic if value is not a slice
func Slice(value interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

// panics if value is not slice, array or map

// IsEmpty returns true if the query is an empty query
// an empty query matches against everything
func (q *Query) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Criterion is an operator and a value that a given field needs to match on
type Criterion struct {
	query    *Query
	operator int
	value    interface{}
	values   []interface{}
}

func hasMatchFunc(criteria []*Criterion) bool { _ = "STUB: not implemented"; return false }

// Field allows for referencing a field in structure being compared
type Field string

// Where starts a query for specifying the criteria that an object in the badgerhold needs to match to
// be returned in a Find result
/*
Query API Example

	s.Find(badgerhold.Where("FieldName").Eq(value).And("AnotherField").Lt(AnotherValue).
		Or(badgerhold.Where("FieldName").Eq(anotherValue)

Since Gobs only encode exported fields, this will panic if you pass in a field with a lower case first letter
*/
func Where(field string) *Criterion { _ = "STUB: not implemented"; return nil }

// And creates another set of criterion the needs to apply to a query
func (q *Query) And(field string) *Criterion { _ = "STUB: not implemented"; return nil }

// Skip skips the number of records that match all the rest of the query criteria, and does not return them
// in the result set.  Setting skip multiple times, or to a negative value will panic
func (q *Query) Skip(amount int) *Query { _ = "STUB: not implemented"; return nil }

// Limit sets the maximum number of records that can be returned by a query
// Setting Limit multiple times, or to a negative value will panic
func (q *Query) Limit(amount int) *Query { _ = "STUB: not implemented"; return nil }

// Contains tests if the current field is a slice that contains the passed in value
func (c *Criterion) Contains(value interface{}) *Query { _ = "STUB: not implemented"; return nil }

// ContainsAll tests if the current field is a slice that contains all of the passed in values.  If any of the
// values are NOT contained in the slice, then no match is made
func (c *Criterion) ContainsAll(values ...interface{}) *Query {
	_ = "STUB: not implemented"
	return nil
}

// ContainsAny tests if the current field is a slice that contains any of the passed in values.  If any of the
// values are contained in the slice, then a match is made
func (c *Criterion) ContainsAny(values ...interface{}) *Query {
	_ = "STUB: not implemented"
	return nil
}

// HasKey tests if the field has a map key matching the passed in value
func (c *Criterion) HasKey(value interface{}) *Query { _ = "STUB: not implemented"; return nil }

// SortBy sorts the results by the given fields name
// Multiple fields can be used
func (q *Query) SortBy(fields ...string) *Query { _ = "STUB: not implemented"; return nil }

// Reverse will reverse the current result set
// useful with SortBy
func (q *Query) Reverse() *Query { _ = "STUB: not implemented"; return nil }

// Index specifies the index to use when running this query
func (q *Query) Index(indexName string) *Query { _ = "STUB: not implemented"; return nil }

// NOTE: I may reconsider this in the future

func (q *Query) validateIndex(data interface{}) error { _ = "STUB: not implemented"; return nil }

// no field name or custom index name found

// Or creates another separate query that gets unioned with any other results in the query
// Or will panic if the query passed in contains a limit or skip value, as they are only
// allowed on top level queries
func (q *Query) Or(query *Query) *Query { _ = "STUB: not implemented"; return nil }

// Matches returns whether the provided data matches the query.
// Will match all field criteria, including nested OR queries, but ignores limits, skips, sort orders, etc.
func (q *Query) Matches(s *Store, data interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (q *Query) matches(s *Store, key []byte, value reflect.Value, data interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (q *Query) matchesAllFields(s *Store, key []byte, value reflect.Value, currentRow interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// already handled by index Iterator

func fieldValue(value reflect.Value, field string) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (c *Criterion) op(op int, value interface{}) *Query { _ = "STUB: not implemented"; return nil }

// Eq tests if the current field is Equal to the passed in value
func (c *Criterion) Eq(value interface{}) *Query { _ = "STUB: not implemented"; return nil }

// Ne test if the current field is Not Equal to the passed in value
func (c *Criterion) Ne(value interface{}) *Query { _ = "STUB: not implemented"; return nil }

// Gt test if the current field is Greater Than the passed in value
func (c *Criterion) Gt(value interface{}) *Query { _ = "STUB: not implemented"; return nil }

// Lt test if the current field is Less Than the passed in value
func (c *Criterion) Lt(value interface{}) *Query { _ = "STUB: not implemented"; return nil }

// Ge test if the current field is Greater Than or Equal To the passed in value
func (c *Criterion) Ge(value interface{}) *Query { _ = "STUB: not implemented"; return nil }

// Le test if the current field is Less Than or Equal To the passed in value
func (c *Criterion) Le(value interface{}) *Query { _ = "STUB: not implemented"; return nil }

// In test if the current field is a member of the slice of values passed in
func (c *Criterion) In(values ...interface{}) *Query { _ = "STUB: not implemented"; return nil }

// RegExp will test if a field matches against the regular expression
// The Field Value will be converted to string (%s) before testing
func (c *Criterion) RegExp(expression *regexp.Regexp) *Query { _ = "STUB: not implemented"; return nil }

// IsNil will test if a field is equal to nil
func (c *Criterion) IsNil() *Query { _ = "STUB: not implemented"; return nil }

// HasPrefix will test if a field starts with provided string
func (c *Criterion) HasPrefix(prefix string) *Query { _ = "STUB: not implemented"; return nil }

// HasSuffix will test if a field ends with provided string
func (c *Criterion) HasSuffix(suffix string) *Query { _ = "STUB: not implemented"; return nil }

// MatchFunc is a function used to test an arbitrary matching value in a query
type MatchFunc func(ra *RecordAccess) (bool, error)

// RecordAccess allows access to the current record, field or allows running a sub-query within a
// MatchFunc
type RecordAccess struct {
	record interface{}
	field  interface{}
	query  *Query
	store  *Store
}

// Field is the current field being queried
func (r *RecordAccess) Field() interface{} {
	_ = "STUB: not implemented"

	// Record is the complete record for a given row in badgerhold
	return nil
}

func (r *RecordAccess) Record() interface{} {
	_ = "STUB: not implemented"

	// SubQuery allows you to run another query in the same transaction for each
	// record in a parent query
	return nil
}

func (r *RecordAccess) SubQuery(result interface{}, query *Query) error {
	_ = "STUB: not implemented"
	return nil
}

// SubAggregateQuery allows you to run another aggregate query in the same transaction for each
// record in a parent query
func (r *RecordAccess) SubAggregateQuery(query *Query, groupBy ...string) ([]*AggregateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MatchFunc will test if a field matches the passed in function
func (c *Criterion) MatchFunc(match MatchFunc) *Query { _ = "STUB: not implemented"; return nil }

// test if the criterion passes with the passed in value
func (c *Criterion) test(s *Store, testValue interface{}, encoded bool, keyType string, currentRow interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// value is a slice of values, use c.values

// used with keys

// make slice containing recordValue

// c.operator == all {

// comparison operators

func (s *Store) matchesAllCriteria(criteria []*Criterion, value interface{}, encoded bool, keyType string,
	currentRow interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func startsUpper(str string) bool { _ = "STUB: not implemented"; return false }

func (q *Query) String() string { _ = "STUB: not implemented"; return "" }

// remove last AND

func (c *Criterion) String() string { _ = "STUB: not implemented"; return "" }

type record struct {
	key   []byte
	value reflect.Value
}

func (s *Store) runQuery(tx *badger.Txn, dataType interface{}, query *Query, retrievedKeys KeyList, skip int,
	action func(r *record) error) error {
	_ = "STUB: not implemented"
	return nil
}

// don't check this record if it's already been retrieved

// track that this key's entry has been added to the result list

// runQuerySort runs the query without sort, skip, or limit, then applies them to the entire result set
func (s *Store) runQuerySort(tx *badger.Txn, dataType interface{}, query *Query, action func(r *record) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Run query without sort, skip or limit
// apply sort, skip and limit to entire dataset

func getSkipAndLimitRange(query *Query, recordsLen int) (startIndex, endIndex int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func sortFunction(query *Query, first, second reflect.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// shouldn't happen due to field check above

// shouldn't happen due to field check above

// if for some reason there is an error on compare, fallback to a lexicographic compare

func validateSortFields(query *Query) error { _ = "STUB: not implemented"; return nil }

func (s *Store) findQuery(tx *badger.Txn, result interface{}, query *Query) error {
	_ = "STUB: not implemented"
	return nil
}

func isFindByIndexQuery(query *Query) bool { _ = "STUB: not implemented"; return false }

func (s *Store) deleteQuery(tx *badger.Txn, dataType interface{}, query *Query) error {
	_ = "STUB: not implemented"
	return nil
}

// remove any indexes

func (s *Store) updateQuery(tx *badger.Txn, dataType interface{}, query *Query, update func(record interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

// delete any existing indexes bad on original value

// insert any new indexes

func (s *Store) aggregateQuery(tx *badger.Txn, dataType interface{}, query *Query, groupBy ...string) ([]*AggregateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if group part is equal, compare the next group part

// group already exists, append results to reduction

// group  not found, create another grouping at i

func (s *Store) findOneQuery(tx *badger.Txn, result interface{}, query *Query) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) forEach(tx *badger.Txn, query *Query, fn interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) countQuery(tx *badger.Txn, dataType interface{}, query *Query) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Store) findByIndexQuery(tx *badger.Txn, resultSlice reflect.Value, query *Query) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) fetchIndexValues(tx *badger.Txn, query *Query, typeName string, indexKeys ...interface{}) (KeyList, error) {
	_ = "STUB: not implemented"
	return *new(KeyList), nil
}

func (s *Store) setKeyField(data []byte, key reflect.Value, keyField reflect.StructField, typeName string) error {
	_ = "STUB: not implemented"
	return nil
}

func dereference(value reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

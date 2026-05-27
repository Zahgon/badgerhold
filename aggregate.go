// Copyright 2019 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package badgerhold

import (
	"reflect"

	"github.com/dgraph-io/badger/v4"
)

// AggregateResult allows you to access the results of an aggregate query
type AggregateResult struct {
	reduction []reflect.Value // always pointers
	group     []reflect.Value
	sortby    string
}

// Group returns the field grouped by in the query
func (a *AggregateResult) Group(result ...interface{}) { _ = "STUB: not implemented"; return }

// Reduction is the collection of records that are part of the AggregateResult Group
func (a *AggregateResult) Reduction(result interface{}) { _ = "STUB: not implemented"; return }

type aggregateResultSort AggregateResult

func (a *aggregateResultSort) Len() int      { _ = "STUB: not implemented"; return 0 }
func (a *aggregateResultSort) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (a *aggregateResultSort) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// reduction values are always pointers
	return false
}

// Sort sorts the aggregate reduction by the passed in field in ascending order
// Sort is called automatically by calls to Min / Max to get the min and max values
func (a *AggregateResult) Sort(field string) { _ = "STUB: not implemented"; return }

// already sorted

// Max Returns the maxiumum value of the Aggregate Grouping, uses the Comparer interface
func (a *AggregateResult) Max(field string, result interface{}) { _ = "STUB: not implemented"; return }

// Min returns the minimum value of the Aggregate Grouping, uses the Comparer interface
func (a *AggregateResult) Min(field string, result interface{}) { _ = "STUB: not implemented"; return }

// Avg returns the average float value of the aggregate grouping
// panics if the field cannot be converted to an float64
func (a *AggregateResult) Avg(field string) float64 { _ = "STUB: not implemented"; return 0 }

// Sum returns the sum value of the aggregate grouping
// panics if the field cannot be converted to an float64
func (a *AggregateResult) Sum(field string) float64 { _ = "STUB: not implemented"; return 0 }

// Count returns the number of records in the aggregate grouping
func (a *AggregateResult) Count() uint64 { _ = "STUB: not implemented"; return 0 }

// FindAggregate returns an aggregate grouping for the passed in query
// groupBy is optional
func (s *Store) FindAggregate(dataType interface{}, query *Query, groupBy ...string) ([]*AggregateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TxFindAggregate is the same as FindAggregate, but you specify your own transaction
// groupBy is optional
func (s *Store) TxFindAggregate(tx *badger.Txn, dataType interface{}, query *Query,
	groupBy ...string) ([]*AggregateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tryFloat(val reflect.Value) float64 { _ = "STUB: not implemented"; return 0 }

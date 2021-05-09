// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testFlywaySchemaHistories(t *testing.T) {
	t.Parallel()

	query := FlywaySchemaHistories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFlywaySchemaHistoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FlywaySchemaHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFlywaySchemaHistoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := FlywaySchemaHistories().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FlywaySchemaHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFlywaySchemaHistoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FlywaySchemaHistorySlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FlywaySchemaHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFlywaySchemaHistoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FlywaySchemaHistoryExists(tx, o.InstalledRank)
	if err != nil {
		t.Errorf("Unable to check if FlywaySchemaHistory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FlywaySchemaHistoryExists to return true, but got false.")
	}
}

func testFlywaySchemaHistoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	flywaySchemaHistoryFound, err := FindFlywaySchemaHistory(tx, o.InstalledRank)
	if err != nil {
		t.Error(err)
	}

	if flywaySchemaHistoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFlywaySchemaHistoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = FlywaySchemaHistories().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testFlywaySchemaHistoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := FlywaySchemaHistories().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFlywaySchemaHistoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	flywaySchemaHistoryOne := &FlywaySchemaHistory{}
	flywaySchemaHistoryTwo := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, flywaySchemaHistoryOne, flywaySchemaHistoryDBTypes, false, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, flywaySchemaHistoryTwo, flywaySchemaHistoryDBTypes, false, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = flywaySchemaHistoryOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = flywaySchemaHistoryTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FlywaySchemaHistories().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFlywaySchemaHistoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	flywaySchemaHistoryOne := &FlywaySchemaHistory{}
	flywaySchemaHistoryTwo := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, flywaySchemaHistoryOne, flywaySchemaHistoryDBTypes, false, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, flywaySchemaHistoryTwo, flywaySchemaHistoryDBTypes, false, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = flywaySchemaHistoryOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = flywaySchemaHistoryTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FlywaySchemaHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func flywaySchemaHistoryBeforeInsertHook(e boil.Executor, o *FlywaySchemaHistory) error {
	*o = FlywaySchemaHistory{}
	return nil
}

func flywaySchemaHistoryAfterInsertHook(e boil.Executor, o *FlywaySchemaHistory) error {
	*o = FlywaySchemaHistory{}
	return nil
}

func flywaySchemaHistoryAfterSelectHook(e boil.Executor, o *FlywaySchemaHistory) error {
	*o = FlywaySchemaHistory{}
	return nil
}

func flywaySchemaHistoryBeforeUpdateHook(e boil.Executor, o *FlywaySchemaHistory) error {
	*o = FlywaySchemaHistory{}
	return nil
}

func flywaySchemaHistoryAfterUpdateHook(e boil.Executor, o *FlywaySchemaHistory) error {
	*o = FlywaySchemaHistory{}
	return nil
}

func flywaySchemaHistoryBeforeDeleteHook(e boil.Executor, o *FlywaySchemaHistory) error {
	*o = FlywaySchemaHistory{}
	return nil
}

func flywaySchemaHistoryAfterDeleteHook(e boil.Executor, o *FlywaySchemaHistory) error {
	*o = FlywaySchemaHistory{}
	return nil
}

func flywaySchemaHistoryBeforeUpsertHook(e boil.Executor, o *FlywaySchemaHistory) error {
	*o = FlywaySchemaHistory{}
	return nil
}

func flywaySchemaHistoryAfterUpsertHook(e boil.Executor, o *FlywaySchemaHistory) error {
	*o = FlywaySchemaHistory{}
	return nil
}

func testFlywaySchemaHistoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &FlywaySchemaHistory{}
	o := &FlywaySchemaHistory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory object: %s", err)
	}

	AddFlywaySchemaHistoryHook(boil.BeforeInsertHook, flywaySchemaHistoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	flywaySchemaHistoryBeforeInsertHooks = []FlywaySchemaHistoryHook{}

	AddFlywaySchemaHistoryHook(boil.AfterInsertHook, flywaySchemaHistoryAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	flywaySchemaHistoryAfterInsertHooks = []FlywaySchemaHistoryHook{}

	AddFlywaySchemaHistoryHook(boil.AfterSelectHook, flywaySchemaHistoryAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	flywaySchemaHistoryAfterSelectHooks = []FlywaySchemaHistoryHook{}

	AddFlywaySchemaHistoryHook(boil.BeforeUpdateHook, flywaySchemaHistoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	flywaySchemaHistoryBeforeUpdateHooks = []FlywaySchemaHistoryHook{}

	AddFlywaySchemaHistoryHook(boil.AfterUpdateHook, flywaySchemaHistoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	flywaySchemaHistoryAfterUpdateHooks = []FlywaySchemaHistoryHook{}

	AddFlywaySchemaHistoryHook(boil.BeforeDeleteHook, flywaySchemaHistoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	flywaySchemaHistoryBeforeDeleteHooks = []FlywaySchemaHistoryHook{}

	AddFlywaySchemaHistoryHook(boil.AfterDeleteHook, flywaySchemaHistoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	flywaySchemaHistoryAfterDeleteHooks = []FlywaySchemaHistoryHook{}

	AddFlywaySchemaHistoryHook(boil.BeforeUpsertHook, flywaySchemaHistoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	flywaySchemaHistoryBeforeUpsertHooks = []FlywaySchemaHistoryHook{}

	AddFlywaySchemaHistoryHook(boil.AfterUpsertHook, flywaySchemaHistoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	flywaySchemaHistoryAfterUpsertHooks = []FlywaySchemaHistoryHook{}
}

func testFlywaySchemaHistoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FlywaySchemaHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFlywaySchemaHistoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(flywaySchemaHistoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := FlywaySchemaHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFlywaySchemaHistoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFlywaySchemaHistoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FlywaySchemaHistorySlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testFlywaySchemaHistoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FlywaySchemaHistories().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	flywaySchemaHistoryDBTypes = map[string]string{`InstalledRank`: `int`, `Version`: `varchar`, `Description`: `varchar`, `Type`: `varchar`, `Script`: `varchar`, `Checksum`: `int`, `InstalledBy`: `varchar`, `InstalledOn`: `timestamp`, `ExecutionTime`: `int`, `Success`: `tinyint`}
	_                          = bytes.MinRead
)

func testFlywaySchemaHistoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(flywaySchemaHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(flywaySchemaHistoryAllColumns) == len(flywaySchemaHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FlywaySchemaHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFlywaySchemaHistoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(flywaySchemaHistoryAllColumns) == len(flywaySchemaHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FlywaySchemaHistory{}
	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FlywaySchemaHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, flywaySchemaHistoryDBTypes, true, flywaySchemaHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(flywaySchemaHistoryAllColumns, flywaySchemaHistoryPrimaryKeyColumns) {
		fields = flywaySchemaHistoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			flywaySchemaHistoryAllColumns,
			flywaySchemaHistoryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := FlywaySchemaHistorySlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFlywaySchemaHistoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(flywaySchemaHistoryAllColumns) == len(flywaySchemaHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLFlywaySchemaHistoryUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := FlywaySchemaHistory{}
	if err = randomize.Struct(seed, &o, flywaySchemaHistoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FlywaySchemaHistory: %s", err)
	}

	count, err := FlywaySchemaHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, flywaySchemaHistoryDBTypes, false, flywaySchemaHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FlywaySchemaHistory struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FlywaySchemaHistory: %s", err)
	}

	count, err = FlywaySchemaHistories().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
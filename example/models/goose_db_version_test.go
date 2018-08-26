// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testGooseDBVersions(t *testing.T) {
	t.Parallel()

	query := GooseDBVersions(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testGooseDBVersionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = gooseDBVersion.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := GooseDBVersions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGooseDBVersionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = GooseDBVersions(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := GooseDBVersions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGooseDBVersionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GooseDBVersionSlice{gooseDBVersion}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := GooseDBVersions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testGooseDBVersionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := GooseDBVersionExists(tx, gooseDBVersion.ID)
	if err != nil {
		t.Errorf("Unable to check if GooseDBVersion exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GooseDBVersionExistsG to return true, but got false.")
	}
}
func testGooseDBVersionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	gooseDBVersionFound, err := FindGooseDBVersion(tx, gooseDBVersion.ID)
	if err != nil {
		t.Error(err)
	}

	if gooseDBVersionFound == nil {
		t.Error("want a record, got nil")
	}
}
func testGooseDBVersionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = GooseDBVersions(tx).Bind(gooseDBVersion); err != nil {
		t.Error(err)
	}
}

func testGooseDBVersionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := GooseDBVersions(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGooseDBVersionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersionOne := &GooseDBVersion{}
	gooseDBVersionTwo := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersionOne, gooseDBVersionDBTypes, false, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}
	if err = randomize.Struct(seed, gooseDBVersionTwo, gooseDBVersionDBTypes, false, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = gooseDBVersionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := GooseDBVersions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGooseDBVersionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	gooseDBVersionOne := &GooseDBVersion{}
	gooseDBVersionTwo := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersionOne, gooseDBVersionDBTypes, false, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}
	if err = randomize.Struct(seed, gooseDBVersionTwo, gooseDBVersionDBTypes, false, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = gooseDBVersionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GooseDBVersions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func gooseDBVersionBeforeInsertHook(e boil.Executor, o *GooseDBVersion) error {
	*o = GooseDBVersion{}
	return nil
}

func gooseDBVersionAfterInsertHook(e boil.Executor, o *GooseDBVersion) error {
	*o = GooseDBVersion{}
	return nil
}

func gooseDBVersionAfterSelectHook(e boil.Executor, o *GooseDBVersion) error {
	*o = GooseDBVersion{}
	return nil
}

func gooseDBVersionBeforeUpdateHook(e boil.Executor, o *GooseDBVersion) error {
	*o = GooseDBVersion{}
	return nil
}

func gooseDBVersionAfterUpdateHook(e boil.Executor, o *GooseDBVersion) error {
	*o = GooseDBVersion{}
	return nil
}

func gooseDBVersionBeforeDeleteHook(e boil.Executor, o *GooseDBVersion) error {
	*o = GooseDBVersion{}
	return nil
}

func gooseDBVersionAfterDeleteHook(e boil.Executor, o *GooseDBVersion) error {
	*o = GooseDBVersion{}
	return nil
}

func gooseDBVersionBeforeUpsertHook(e boil.Executor, o *GooseDBVersion) error {
	*o = GooseDBVersion{}
	return nil
}

func gooseDBVersionAfterUpsertHook(e boil.Executor, o *GooseDBVersion) error {
	*o = GooseDBVersion{}
	return nil
}

func testGooseDBVersionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &GooseDBVersion{}
	o := &GooseDBVersion{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, gooseDBVersionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion object: %s", err)
	}

	AddGooseDBVersionHook(boil.BeforeInsertHook, gooseDBVersionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	gooseDBVersionBeforeInsertHooks = []GooseDBVersionHook{}

	AddGooseDBVersionHook(boil.AfterInsertHook, gooseDBVersionAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	gooseDBVersionAfterInsertHooks = []GooseDBVersionHook{}

	AddGooseDBVersionHook(boil.AfterSelectHook, gooseDBVersionAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	gooseDBVersionAfterSelectHooks = []GooseDBVersionHook{}

	AddGooseDBVersionHook(boil.BeforeUpdateHook, gooseDBVersionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	gooseDBVersionBeforeUpdateHooks = []GooseDBVersionHook{}

	AddGooseDBVersionHook(boil.AfterUpdateHook, gooseDBVersionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	gooseDBVersionAfterUpdateHooks = []GooseDBVersionHook{}

	AddGooseDBVersionHook(boil.BeforeDeleteHook, gooseDBVersionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	gooseDBVersionBeforeDeleteHooks = []GooseDBVersionHook{}

	AddGooseDBVersionHook(boil.AfterDeleteHook, gooseDBVersionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	gooseDBVersionAfterDeleteHooks = []GooseDBVersionHook{}

	AddGooseDBVersionHook(boil.BeforeUpsertHook, gooseDBVersionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	gooseDBVersionBeforeUpsertHooks = []GooseDBVersionHook{}

	AddGooseDBVersionHook(boil.AfterUpsertHook, gooseDBVersionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	gooseDBVersionAfterUpsertHooks = []GooseDBVersionHook{}
}
func testGooseDBVersionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GooseDBVersions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGooseDBVersionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx, gooseDBVersionColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := GooseDBVersions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGooseDBVersionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = gooseDBVersion.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testGooseDBVersionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GooseDBVersionSlice{gooseDBVersion}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testGooseDBVersionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := GooseDBVersions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	gooseDBVersionDBTypes = map[string]string{`ID`: `integer`, `IsApplied`: `boolean`, `Tstamp`: `timestamp without time zone`, `VersionID`: `bigint`}
	_                     = bytes.MinRead
)

func testGooseDBVersionsUpdate(t *testing.T) {
	t.Parallel()

	if len(gooseDBVersionColumns) == len(gooseDBVersionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GooseDBVersions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	if err = gooseDBVersion.Update(tx); err != nil {
		t.Error(err)
	}
}

func testGooseDBVersionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(gooseDBVersionColumns) == len(gooseDBVersionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	gooseDBVersion := &GooseDBVersion{}
	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GooseDBVersions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, gooseDBVersion, gooseDBVersionDBTypes, true, gooseDBVersionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(gooseDBVersionColumns, gooseDBVersionPrimaryKeyColumns) {
		fields = gooseDBVersionColumns
	} else {
		fields = strmangle.SetComplement(
			gooseDBVersionColumns,
			gooseDBVersionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(gooseDBVersion))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := GooseDBVersionSlice{gooseDBVersion}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testGooseDBVersionsUpsert(t *testing.T) {
	t.Parallel()

	if len(gooseDBVersionColumns) == len(gooseDBVersionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	gooseDBVersion := GooseDBVersion{}
	if err = randomize.Struct(seed, &gooseDBVersion, gooseDBVersionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = gooseDBVersion.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert GooseDBVersion: %s", err)
	}

	count, err := GooseDBVersions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &gooseDBVersion, gooseDBVersionDBTypes, false, gooseDBVersionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GooseDBVersion struct: %s", err)
	}

	if err = gooseDBVersion.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert GooseDBVersion: %s", err)
	}

	count, err = GooseDBVersions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
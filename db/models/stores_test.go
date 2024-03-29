// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testStores(t *testing.T) {
	t.Parallel()

	query := Stores()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStoresDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Stores().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStoresQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Stores().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Stores().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStoresSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StoreSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Stores().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStoresExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StoreExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Store exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StoreExists to return true, but got false.")
	}
}

func testStoresFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	storeFound, err := FindStore(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if storeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStoresBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Stores().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStoresOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Stores().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStoresAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	storeOne := &Store{}
	storeTwo := &Store{}
	if err = randomize.Struct(seed, storeOne, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}
	if err = randomize.Struct(seed, storeTwo, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = storeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = storeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Stores().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStoresCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	storeOne := &Store{}
	storeTwo := &Store{}
	if err = randomize.Struct(seed, storeOne, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}
	if err = randomize.Struct(seed, storeTwo, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = storeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = storeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stores().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testStoresInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stores().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStoresInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(storeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Stores().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStoreToOneOrgUsingOrg(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Store
	var foreign Org

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, storeDBTypes, false, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, orgDBTypes, false, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OrgID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Org().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := StoreSlice{&local}
	if err = local.L.LoadOrg(ctx, tx, false, (*[]*Store)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Org == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Org = nil
	if err = local.L.LoadOrg(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Org == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStoreToOneSetOpOrgUsingOrg(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Store
	var b, c Org

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, storeDBTypes, false, strmangle.SetComplement(storePrimaryKeyColumns, storeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, orgDBTypes, false, strmangle.SetComplement(orgPrimaryKeyColumns, orgColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, orgDBTypes, false, strmangle.SetComplement(orgPrimaryKeyColumns, orgColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Org{&b, &c} {
		err = a.SetOrg(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Org != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Stores[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrgID != x.ID {
			t.Error("foreign key was wrong value", a.OrgID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrgID))
		reflect.Indirect(reflect.ValueOf(&a.OrgID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrgID != x.ID {
			t.Error("foreign key was wrong value", a.OrgID, x.ID)
		}
	}
}

func testStoresReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStoresReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StoreSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStoresSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Stores().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	storeDBTypes = map[string]string{`ID`: `bigint`, `OrgID`: `bigint`, `Name`: `character varying`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_            = bytes.MinRead
)

func testStoresUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(storePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(storeAllColumns) == len(storePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stores().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, storeDBTypes, true, storePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStoresSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(storeAllColumns) == len(storePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Store{}
	if err = randomize.Struct(seed, o, storeDBTypes, true, storeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stores().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, storeDBTypes, true, storePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(storeAllColumns, storePrimaryKeyColumns) {
		fields = storeAllColumns
	} else {
		fields = strmangle.SetComplement(
			storeAllColumns,
			storePrimaryKeyColumns,
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

	slice := StoreSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStoresUpsert(t *testing.T) {
	t.Parallel()

	if len(storeAllColumns) == len(storePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Store{}
	if err = randomize.Struct(seed, &o, storeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Store: %s", err)
	}

	count, err := Stores().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, storeDBTypes, false, storePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Store struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Store: %s", err)
	}

	count, err = Stores().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

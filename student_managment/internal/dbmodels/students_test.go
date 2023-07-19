// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"bytes"
	"context"
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

func testStudents(t *testing.T) {
	t.Parallel()

	query := Students()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStudentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
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

	count, err := Students().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStudentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Students().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Students().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStudentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StudentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Students().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStudentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StudentExists(ctx, tx, o.StudentID)
	if err != nil {
		t.Errorf("Unable to check if Student exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StudentExists to return true, but got false.")
	}
}

func testStudentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	studentFound, err := FindStudent(ctx, tx, o.StudentID)
	if err != nil {
		t.Error(err)
	}

	if studentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStudentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Students().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStudentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Students().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStudentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	studentOne := &Student{}
	studentTwo := &Student{}
	if err = randomize.Struct(seed, studentOne, studentDBTypes, false, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}
	if err = randomize.Struct(seed, studentTwo, studentDBTypes, false, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = studentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = studentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Students().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStudentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	studentOne := &Student{}
	studentTwo := &Student{}
	if err = randomize.Struct(seed, studentOne, studentDBTypes, false, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}
	if err = randomize.Struct(seed, studentTwo, studentDBTypes, false, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = studentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = studentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Students().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func studentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Student) error {
	*o = Student{}
	return nil
}

func studentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Student) error {
	*o = Student{}
	return nil
}

func studentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Student) error {
	*o = Student{}
	return nil
}

func studentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Student) error {
	*o = Student{}
	return nil
}

func studentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Student) error {
	*o = Student{}
	return nil
}

func studentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Student) error {
	*o = Student{}
	return nil
}

func studentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Student) error {
	*o = Student{}
	return nil
}

func studentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Student) error {
	*o = Student{}
	return nil
}

func studentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Student) error {
	*o = Student{}
	return nil
}

func testStudentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Student{}
	o := &Student{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, studentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Student object: %s", err)
	}

	AddStudentHook(boil.BeforeInsertHook, studentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	studentBeforeInsertHooks = []StudentHook{}

	AddStudentHook(boil.AfterInsertHook, studentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	studentAfterInsertHooks = []StudentHook{}

	AddStudentHook(boil.AfterSelectHook, studentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	studentAfterSelectHooks = []StudentHook{}

	AddStudentHook(boil.BeforeUpdateHook, studentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	studentBeforeUpdateHooks = []StudentHook{}

	AddStudentHook(boil.AfterUpdateHook, studentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	studentAfterUpdateHooks = []StudentHook{}

	AddStudentHook(boil.BeforeDeleteHook, studentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	studentBeforeDeleteHooks = []StudentHook{}

	AddStudentHook(boil.AfterDeleteHook, studentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	studentAfterDeleteHooks = []StudentHook{}

	AddStudentHook(boil.BeforeUpsertHook, studentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	studentBeforeUpsertHooks = []StudentHook{}

	AddStudentHook(boil.AfterUpsertHook, studentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	studentAfterUpsertHooks = []StudentHook{}
}

func testStudentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Students().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStudentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(studentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Students().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStudentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
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

func testStudentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StudentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStudentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Students().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	studentDBTypes = map[string]string{`StudentID`: `integer`, `LastName`: `character varying`, `FirstName`: `character varying`}
	_              = bytes.MinRead
)

func testStudentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(studentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(studentAllColumns) == len(studentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Students().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, studentDBTypes, true, studentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStudentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(studentAllColumns) == len(studentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Student{}
	if err = randomize.Struct(seed, o, studentDBTypes, true, studentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Students().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, studentDBTypes, true, studentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(studentAllColumns, studentPrimaryKeyColumns) {
		fields = studentAllColumns
	} else {
		fields = strmangle.SetComplement(
			studentAllColumns,
			studentPrimaryKeyColumns,
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

	slice := StudentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStudentsUpsert(t *testing.T) {
	t.Parallel()

	if len(studentAllColumns) == len(studentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Student{}
	if err = randomize.Struct(seed, &o, studentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Student: %s", err)
	}

	count, err := Students().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, studentDBTypes, false, studentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Student struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Student: %s", err)
	}

	count, err = Students().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

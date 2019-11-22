// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Store is an object representing the database table.
type Store struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	OrgID     int64     `boil:"org_id" json:"orgID" toml:"orgID" yaml:"orgID"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *storeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L storeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StoreColumns = struct {
	ID        string
	OrgID     string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	OrgID:     "org_id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var StoreWhere = struct {
	ID        whereHelperint64
	OrgID     whereHelperint64
	Name      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"stores\".\"id\""},
	OrgID:     whereHelperint64{field: "\"stores\".\"org_id\""},
	Name:      whereHelperstring{field: "\"stores\".\"name\""},
	CreatedAt: whereHelpertime_Time{field: "\"stores\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"stores\".\"updated_at\""},
}

// StoreRels is where relationship names are stored.
var StoreRels = struct {
	Org string
}{
	Org: "Org",
}

// storeR is where relationships are stored.
type storeR struct {
	Org *Org
}

// NewStruct creates a new relationship struct
func (*storeR) NewStruct() *storeR {
	return &storeR{}
}

// storeL is where Load methods for each relationship are stored.
type storeL struct{}

var (
	storeAllColumns            = []string{"id", "org_id", "name", "created_at", "updated_at"}
	storeColumnsWithoutDefault = []string{"org_id", "name"}
	storeColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	storePrimaryKeyColumns     = []string{"id"}
)

type (
	// StoreSlice is an alias for a slice of pointers to Store.
	// This should generally be used opposed to []Store.
	StoreSlice []*Store

	storeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	storeType                 = reflect.TypeOf(&Store{})
	storeMapping              = queries.MakeStructMapping(storeType)
	storePrimaryKeyMapping, _ = queries.BindMapping(storeType, storeMapping, storePrimaryKeyColumns)
	storeInsertCacheMut       sync.RWMutex
	storeInsertCache          = make(map[string]insertCache)
	storeUpdateCacheMut       sync.RWMutex
	storeUpdateCache          = make(map[string]updateCache)
	storeUpsertCacheMut       sync.RWMutex
	storeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single store record from the query using the global executor.
func (q storeQuery) OneG(ctx context.Context) (*Store, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single store record from the query.
func (q storeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Store, error) {
	o := &Store{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stores")
	}

	return o, nil
}

// AllG returns all Store records from the query using the global executor.
func (q storeQuery) AllG(ctx context.Context) (StoreSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Store records from the query.
func (q storeQuery) All(ctx context.Context, exec boil.ContextExecutor) (StoreSlice, error) {
	var o []*Store

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Store slice")
	}

	return o, nil
}

// CountG returns the count of all Store records in the query, and panics on error.
func (q storeQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Store records in the query.
func (q storeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stores rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q storeQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q storeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stores exists")
	}

	return count > 0, nil
}

// Org pointed to by the foreign key.
func (o *Store) Org(mods ...qm.QueryMod) orgQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.OrgID),
	}

	queryMods = append(queryMods, mods...)

	query := Orgs(queryMods...)
	queries.SetFrom(query.Query, "\"orgs\"")

	return query
}

// LoadOrg allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (storeL) LoadOrg(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStore interface{}, mods queries.Applicator) error {
	var slice []*Store
	var object *Store

	if singular {
		object = maybeStore.(*Store)
	} else {
		slice = *maybeStore.(*[]*Store)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &storeR{}
		}
		args = append(args, object.OrgID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &storeR{}
			}

			for _, a := range args {
				if a == obj.OrgID {
					continue Outer
				}
			}

			args = append(args, obj.OrgID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`orgs`), qm.WhereIn(`orgs.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Org")
	}

	var resultSlice []*Org
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Org")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for orgs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for orgs")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Org = foreign
		if foreign.R == nil {
			foreign.R = &orgR{}
		}
		foreign.R.Stores = append(foreign.R.Stores, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrgID == foreign.ID {
				local.R.Org = foreign
				if foreign.R == nil {
					foreign.R = &orgR{}
				}
				foreign.R.Stores = append(foreign.R.Stores, local)
				break
			}
		}
	}

	return nil
}

// SetOrgG of the store to the related item.
// Sets o.R.Org to related.
// Adds o to related.R.Stores.
// Uses the global database handle.
func (o *Store) SetOrgG(ctx context.Context, insert bool, related *Org) error {
	return o.SetOrg(ctx, boil.GetContextDB(), insert, related)
}

// SetOrg of the store to the related item.
// Sets o.R.Org to related.
// Adds o to related.R.Stores.
func (o *Store) SetOrg(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Org) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stores\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"org_id"}),
		strmangle.WhereClause("\"", "\"", 2, storePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrgID = related.ID
	if o.R == nil {
		o.R = &storeR{
			Org: related,
		}
	} else {
		o.R.Org = related
	}

	if related.R == nil {
		related.R = &orgR{
			Stores: StoreSlice{o},
		}
	} else {
		related.R.Stores = append(related.R.Stores, o)
	}

	return nil
}

// Stores retrieves all the records using an executor.
func Stores(mods ...qm.QueryMod) storeQuery {
	mods = append(mods, qm.From("\"stores\""))
	return storeQuery{NewQuery(mods...)}
}

// FindStoreG retrieves a single record by ID.
func FindStoreG(ctx context.Context, iD int64, selectCols ...string) (*Store, error) {
	return FindStore(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindStore retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStore(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Store, error) {
	storeObj := &Store{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stores\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, storeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stores")
	}

	return storeObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Store) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Store) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stores provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(storeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	storeInsertCacheMut.RLock()
	cache, cached := storeInsertCache[key]
	storeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			storeAllColumns,
			storeColumnsWithDefault,
			storeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(storeType, storeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(storeType, storeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"stores\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"stores\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into stores")
	}

	if !cached {
		storeInsertCacheMut.Lock()
		storeInsertCache[key] = cache
		storeInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Store record using the global executor.
// See Update for more documentation.
func (o *Store) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Store.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Store) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	storeUpdateCacheMut.RLock()
	cache, cached := storeUpdateCache[key]
	storeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			storeAllColumns,
			storePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update stores, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stores\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, storePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(storeType, storeMapping, append(wl, storePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update stores row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for stores")
	}

	if !cached {
		storeUpdateCacheMut.Lock()
		storeUpdateCache[key] = cache
		storeUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q storeQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q storeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for stores")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for stores")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StoreSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StoreSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), storePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"stores\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, storePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in store slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all store")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Store) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Store) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stores provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(storeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	storeUpsertCacheMut.RLock()
	cache, cached := storeUpsertCache[key]
	storeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			storeAllColumns,
			storeColumnsWithDefault,
			storeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			storeAllColumns,
			storePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert stores, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(storePrimaryKeyColumns))
			copy(conflict, storePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"stores\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(storeType, storeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(storeType, storeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert stores")
	}

	if !cached {
		storeUpsertCacheMut.Lock()
		storeUpsertCache[key] = cache
		storeUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Store record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Store) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Store record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Store) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Store provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), storePrimaryKeyMapping)
	sql := "DELETE FROM \"stores\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from stores")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for stores")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q storeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no storeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stores")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stores")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o StoreSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StoreSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), storePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"stores\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, storePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from store slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stores")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Store) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Store provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Store) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStore(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StoreSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty StoreSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StoreSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StoreSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), storePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"stores\".* FROM \"stores\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, storePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StoreSlice")
	}

	*o = slice

	return nil
}

// StoreExistsG checks if the Store row exists.
func StoreExistsG(ctx context.Context, iD int64) (bool, error) {
	return StoreExists(ctx, boil.GetContextDB(), iD)
}

// StoreExists checks if the Store row exists.
func StoreExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"stores\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stores exists")
	}

	return exists, nil
}
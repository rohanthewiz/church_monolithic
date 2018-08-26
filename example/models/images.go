// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"github.com/vattle/sqlboiler/types"
	"gopkg.in/nullbio/null.v6"
)

// Image is an object representing the database table.
type Image struct {
	ID         int64             `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt  null.Time         `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt  null.Time         `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	UpdatedBy  string            `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`
	Published  null.Bool         `boil:"published" json:"published,omitempty" toml:"published" yaml:"published,omitempty"`
	Title      string            `boil:"title" json:"title" toml:"title" yaml:"title"`
	Slug       null.String       `boil:"slug" json:"slug,omitempty" toml:"slug" yaml:"slug,omitempty"`
	Summary    null.String       `boil:"summary" json:"summary,omitempty" toml:"summary" yaml:"summary,omitempty"`
	Categories types.StringArray `boil:"categories" json:"categories,omitempty" toml:"categories" yaml:"categories,omitempty"`
	LargePath  null.String       `boil:"large_path" json:"large_path,omitempty" toml:"large_path" yaml:"large_path,omitempty"`
	SmallB64   null.String       `boil:"small_b64" json:"small_b64,omitempty" toml:"small_b64" yaml:"small_b64,omitempty"`
	ThumbB64   null.String       `boil:"thumb_b64" json:"thumb_b64,omitempty" toml:"thumb_b64" yaml:"thumb_b64,omitempty"`
	ImageType  null.String       `boil:"image_type" json:"image_type,omitempty" toml:"image_type" yaml:"image_type,omitempty"`

	R *imageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L imageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ImageColumns = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	UpdatedBy  string
	Published  string
	Title      string
	Slug       string
	Summary    string
	Categories string
	LargePath  string
	SmallB64   string
	ThumbB64   string
	ImageType  string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	UpdatedBy:  "updated_by",
	Published:  "published",
	Title:      "title",
	Slug:       "slug",
	Summary:    "summary",
	Categories: "categories",
	LargePath:  "large_path",
	SmallB64:   "small_b64",
	ThumbB64:   "thumb_b64",
	ImageType:  "image_type",
}

// imageR is where relationships are stored.
type imageR struct {
}

// imageL is where Load methods for each relationship are stored.
type imageL struct{}

var (
	imageColumns               = []string{"id", "created_at", "updated_at", "updated_by", "published", "title", "slug", "summary", "categories", "large_path", "small_b64", "thumb_b64", "image_type"}
	imageColumnsWithoutDefault = []string{"created_at", "updated_at", "updated_by", "published", "title", "slug", "summary", "categories", "large_path", "small_b64", "thumb_b64", "image_type"}
	imageColumnsWithDefault    = []string{"id"}
	imagePrimaryKeyColumns     = []string{"id"}
)

type (
	// ImageSlice is an alias for a slice of pointers to Image.
	// This should generally be used opposed to []Image.
	ImageSlice []*Image
	// ImageHook is the signature for custom Image hook methods
	ImageHook func(boil.Executor, *Image) error

	imageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	imageType                 = reflect.TypeOf(&Image{})
	imageMapping              = queries.MakeStructMapping(imageType)
	imagePrimaryKeyMapping, _ = queries.BindMapping(imageType, imageMapping, imagePrimaryKeyColumns)
	imageInsertCacheMut       sync.RWMutex
	imageInsertCache          = make(map[string]insertCache)
	imageUpdateCacheMut       sync.RWMutex
	imageUpdateCache          = make(map[string]updateCache)
	imageUpsertCacheMut       sync.RWMutex
	imageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var imageBeforeInsertHooks []ImageHook
var imageBeforeUpdateHooks []ImageHook
var imageBeforeDeleteHooks []ImageHook
var imageBeforeUpsertHooks []ImageHook

var imageAfterInsertHooks []ImageHook
var imageAfterSelectHooks []ImageHook
var imageAfterUpdateHooks []ImageHook
var imageAfterDeleteHooks []ImageHook
var imageAfterUpsertHooks []ImageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Image) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range imageBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Image) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range imageBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Image) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range imageBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Image) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range imageBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Image) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range imageAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Image) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range imageAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Image) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range imageAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Image) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range imageAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Image) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range imageAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddImageHook registers your hook function for all future operations.
func AddImageHook(hookPoint boil.HookPoint, imageHook ImageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		imageBeforeInsertHooks = append(imageBeforeInsertHooks, imageHook)
	case boil.BeforeUpdateHook:
		imageBeforeUpdateHooks = append(imageBeforeUpdateHooks, imageHook)
	case boil.BeforeDeleteHook:
		imageBeforeDeleteHooks = append(imageBeforeDeleteHooks, imageHook)
	case boil.BeforeUpsertHook:
		imageBeforeUpsertHooks = append(imageBeforeUpsertHooks, imageHook)
	case boil.AfterInsertHook:
		imageAfterInsertHooks = append(imageAfterInsertHooks, imageHook)
	case boil.AfterSelectHook:
		imageAfterSelectHooks = append(imageAfterSelectHooks, imageHook)
	case boil.AfterUpdateHook:
		imageAfterUpdateHooks = append(imageAfterUpdateHooks, imageHook)
	case boil.AfterDeleteHook:
		imageAfterDeleteHooks = append(imageAfterDeleteHooks, imageHook)
	case boil.AfterUpsertHook:
		imageAfterUpsertHooks = append(imageAfterUpsertHooks, imageHook)
	}
}

// OneP returns a single image record from the query, and panics on error.
func (q imageQuery) OneP() *Image {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single image record from the query.
func (q imageQuery) One() (*Image, error) {
	o := &Image{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for images")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Image records from the query, and panics on error.
func (q imageQuery) AllP() ImageSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Image records from the query.
func (q imageQuery) All() (ImageSlice, error) {
	var o []*Image

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Image slice")
	}

	if len(imageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Image records in the query, and panics on error.
func (q imageQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Image records in the query.
func (q imageQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count images rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q imageQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q imageQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if images exists")
	}

	return count > 0, nil
}

// ImagesG retrieves all records.
func ImagesG(mods ...qm.QueryMod) imageQuery {
	return Images(boil.GetDB(), mods...)
}

// Images retrieves all the records using an executor.
func Images(exec boil.Executor, mods ...qm.QueryMod) imageQuery {
	mods = append(mods, qm.From("\"images\""))
	return imageQuery{NewQuery(exec, mods...)}
}

// FindImageG retrieves a single record by ID.
func FindImageG(id int64, selectCols ...string) (*Image, error) {
	return FindImage(boil.GetDB(), id, selectCols...)
}

// FindImageGP retrieves a single record by ID, and panics on error.
func FindImageGP(id int64, selectCols ...string) *Image {
	retobj, err := FindImage(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindImage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindImage(exec boil.Executor, id int64, selectCols ...string) (*Image, error) {
	imageObj := &Image{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"images\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(imageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from images")
	}

	return imageObj, nil
}

// FindImageP retrieves a single record by ID with an executor, and panics on error.
func FindImageP(exec boil.Executor, id int64, selectCols ...string) *Image {
	retobj, err := FindImage(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Image) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Image) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Image) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Image) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no images provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	if o.UpdatedAt.Time.IsZero() {
		o.UpdatedAt.Time = currTime
		o.UpdatedAt.Valid = true
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(imageColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	imageInsertCacheMut.RLock()
	cache, cached := imageInsertCache[key]
	imageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			imageColumns,
			imageColumnsWithDefault,
			imageColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(imageType, imageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(imageType, imageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"images\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"images\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into images")
	}

	if !cached {
		imageInsertCacheMut.Lock()
		imageInsertCache[key] = cache
		imageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Image record. See Update for
// whitelist behavior description.
func (o *Image) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Image record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Image) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Image, and panics on error.
// See Update for whitelist behavior description.
func (o *Image) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Image.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Image) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	imageUpdateCacheMut.RLock()
	cache, cached := imageUpdateCache[key]
	imageUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			imageColumns,
			imagePrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update images, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"images\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, imagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(imageType, imageMapping, append(wl, imagePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update images row")
	}

	if !cached {
		imageUpdateCacheMut.Lock()
		imageUpdateCache[key] = cache
		imageUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q imageQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q imageQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for images")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ImageSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ImageSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ImageSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ImageSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), imagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"images\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, imagePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in image slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Image) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Image) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Image) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Image) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no images provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(imageColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
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
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	imageUpsertCacheMut.RLock()
	cache, cached := imageUpsertCache[key]
	imageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			imageColumns,
			imageColumnsWithDefault,
			imageColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			imageColumns,
			imagePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert images, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(imagePrimaryKeyColumns))
			copy(conflict, imagePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"images\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(imageType, imageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(imageType, imageMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert images")
	}

	if !cached {
		imageUpsertCacheMut.Lock()
		imageUpsertCache[key] = cache
		imageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Image record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Image) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Image record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Image) DeleteG() error {
	if o == nil {
		return errors.New("models: no Image provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Image record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Image) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Image record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Image) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Image provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), imagePrimaryKeyMapping)
	sql := "DELETE FROM \"images\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from images")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q imageQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q imageQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no imageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from images")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ImageSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ImageSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Image slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ImageSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ImageSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Image slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(imageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), imagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"images\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, imagePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from image slice")
	}

	if len(imageAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Image) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Image) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Image) ReloadG() error {
	if o == nil {
		return errors.New("models: no Image provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Image) Reload(exec boil.Executor) error {
	ret, err := FindImage(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ImageSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ImageSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ImageSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ImageSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ImageSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	images := ImageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), imagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"images\".* FROM \"images\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, imagePrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&images)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ImageSlice")
	}

	*o = images

	return nil
}

// ImageExists checks if the Image row exists.
func ImageExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"images\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if images exists")
	}

	return exists, nil
}

// ImageExistsG checks if the Image row exists.
func ImageExistsG(id int64) (bool, error) {
	return ImageExists(boil.GetDB(), id)
}

// ImageExistsGP checks if the Image row exists. Panics on error.
func ImageExistsGP(id int64) bool {
	e, err := ImageExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ImageExistsP checks if the Image row exists. Panics on error.
func ImageExistsP(exec boil.Executor, id int64) bool {
	e, err := ImageExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

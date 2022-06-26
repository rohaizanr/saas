// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/go-saas/saas/examples/ent/tenant/ent/post"
	"github.com/go-saas/saas/examples/ent/tenant/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePost = "Post"
)

// PostMutation represents an operation that mutates the Post nodes in the graph.
type PostMutation struct {
	config
	op            Op
	typ           string
	id            *int
	tenant_id     **sql.NullString
	title         *string
	description   *string
	dsn           *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Post, error)
	predicates    []predicate.Post
}

var _ ent.Mutation = (*PostMutation)(nil)

// postOption allows management of the mutation configuration using functional options.
type postOption func(*PostMutation)

// newPostMutation creates new mutation for the Post entity.
func newPostMutation(c config, op Op, opts ...postOption) *PostMutation {
	m := &PostMutation{
		config:        c,
		op:            op,
		typ:           TypePost,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPostID sets the ID field of the mutation.
func withPostID(id int) postOption {
	return func(m *PostMutation) {
		var (
			err   error
			once  sync.Once
			value *Post
		)
		m.oldValue = func(ctx context.Context) (*Post, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Post.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPost sets the old Post of the mutation.
func withPost(node *Post) postOption {
	return func(m *PostMutation) {
		m.oldValue = func(context.Context) (*Post, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PostMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PostMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Post entities.
func (m *PostMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PostMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PostMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Post.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTenantID sets the "tenant_id" field.
func (m *PostMutation) SetTenantID(ss *sql.NullString) {
	m.tenant_id = &ss
}

// TenantID returns the value of the "tenant_id" field in the mutation.
func (m *PostMutation) TenantID() (r *sql.NullString, exists bool) {
	v := m.tenant_id
	if v == nil {
		return
	}
	return *v, true
}

// OldTenantID returns the old "tenant_id" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldTenantID(ctx context.Context) (v *sql.NullString, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTenantID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTenantID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTenantID: %w", err)
	}
	return oldValue.TenantID, nil
}

// ClearTenantID clears the value of the "tenant_id" field.
func (m *PostMutation) ClearTenantID() {
	m.tenant_id = nil
	m.clearedFields[post.FieldTenantID] = struct{}{}
}

// TenantIDCleared returns if the "tenant_id" field was cleared in this mutation.
func (m *PostMutation) TenantIDCleared() bool {
	_, ok := m.clearedFields[post.FieldTenantID]
	return ok
}

// ResetTenantID resets all changes to the "tenant_id" field.
func (m *PostMutation) ResetTenantID() {
	m.tenant_id = nil
	delete(m.clearedFields, post.FieldTenantID)
}

// SetTitle sets the "title" field.
func (m *PostMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *PostMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *PostMutation) ResetTitle() {
	m.title = nil
}

// SetDescription sets the "description" field.
func (m *PostMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *PostMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ClearDescription clears the value of the "description" field.
func (m *PostMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[post.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the "description" field was cleared in this mutation.
func (m *PostMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[post.FieldDescription]
	return ok
}

// ResetDescription resets all changes to the "description" field.
func (m *PostMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, post.FieldDescription)
}

// SetDsn sets the "dsn" field.
func (m *PostMutation) SetDsn(s string) {
	m.dsn = &s
}

// Dsn returns the value of the "dsn" field in the mutation.
func (m *PostMutation) Dsn() (r string, exists bool) {
	v := m.dsn
	if v == nil {
		return
	}
	return *v, true
}

// OldDsn returns the old "dsn" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldDsn(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDsn is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDsn requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDsn: %w", err)
	}
	return oldValue.Dsn, nil
}

// ClearDsn clears the value of the "dsn" field.
func (m *PostMutation) ClearDsn() {
	m.dsn = nil
	m.clearedFields[post.FieldDsn] = struct{}{}
}

// DsnCleared returns if the "dsn" field was cleared in this mutation.
func (m *PostMutation) DsnCleared() bool {
	_, ok := m.clearedFields[post.FieldDsn]
	return ok
}

// ResetDsn resets all changes to the "dsn" field.
func (m *PostMutation) ResetDsn() {
	m.dsn = nil
	delete(m.clearedFields, post.FieldDsn)
}

// Where appends a list predicates to the PostMutation builder.
func (m *PostMutation) Where(ps ...predicate.Post) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *PostMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Post).
func (m *PostMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PostMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.tenant_id != nil {
		fields = append(fields, post.FieldTenantID)
	}
	if m.title != nil {
		fields = append(fields, post.FieldTitle)
	}
	if m.description != nil {
		fields = append(fields, post.FieldDescription)
	}
	if m.dsn != nil {
		fields = append(fields, post.FieldDsn)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PostMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case post.FieldTenantID:
		return m.TenantID()
	case post.FieldTitle:
		return m.Title()
	case post.FieldDescription:
		return m.Description()
	case post.FieldDsn:
		return m.Dsn()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PostMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case post.FieldTenantID:
		return m.OldTenantID(ctx)
	case post.FieldTitle:
		return m.OldTitle(ctx)
	case post.FieldDescription:
		return m.OldDescription(ctx)
	case post.FieldDsn:
		return m.OldDsn(ctx)
	}
	return nil, fmt.Errorf("unknown Post field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) SetField(name string, value ent.Value) error {
	switch name {
	case post.FieldTenantID:
		v, ok := value.(*sql.NullString)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTenantID(v)
		return nil
	case post.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case post.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case post.FieldDsn:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDsn(v)
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PostMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PostMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Post numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PostMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(post.FieldTenantID) {
		fields = append(fields, post.FieldTenantID)
	}
	if m.FieldCleared(post.FieldDescription) {
		fields = append(fields, post.FieldDescription)
	}
	if m.FieldCleared(post.FieldDsn) {
		fields = append(fields, post.FieldDsn)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PostMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PostMutation) ClearField(name string) error {
	switch name {
	case post.FieldTenantID:
		m.ClearTenantID()
		return nil
	case post.FieldDescription:
		m.ClearDescription()
		return nil
	case post.FieldDsn:
		m.ClearDsn()
		return nil
	}
	return fmt.Errorf("unknown Post nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PostMutation) ResetField(name string) error {
	switch name {
	case post.FieldTenantID:
		m.ResetTenantID()
		return nil
	case post.FieldTitle:
		m.ResetTitle()
		return nil
	case post.FieldDescription:
		m.ResetDescription()
		return nil
	case post.FieldDsn:
		m.ResetDsn()
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PostMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PostMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PostMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PostMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PostMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PostMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PostMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Post unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PostMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Post edge %s", name)
}

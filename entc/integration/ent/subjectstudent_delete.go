// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/entc/integration/ent/subjectstudent"
	"entgo.io/ent/schema/field"
)

// SubjectStudentDelete is the builder for deleting a SubjectStudent entity.
type SubjectStudentDelete struct {
	config
	hooks    []Hook
	mutation *SubjectStudentMutation
}

// Where appends a list predicates to the SubjectStudentDelete builder.
func (ssd *SubjectStudentDelete) Where(ps ...predicate.SubjectStudent) *SubjectStudentDelete {
	ssd.mutation.Where(ps...)
	return ssd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ssd *SubjectStudentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ssd.sqlExec, ssd.mutation, ssd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ssd *SubjectStudentDelete) ExecX(ctx context.Context) int {
	n, err := ssd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ssd *SubjectStudentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(subjectstudent.Table, sqlgraph.NewFieldSpec(subjectstudent.FieldID, field.TypeUUID))
	if ps := ssd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ssd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ssd.mutation.done = true
	return affected, err
}

// SubjectStudentDeleteOne is the builder for deleting a single SubjectStudent entity.
type SubjectStudentDeleteOne struct {
	ssd *SubjectStudentDelete
}

// Where appends a list predicates to the SubjectStudentDelete builder.
func (ssdo *SubjectStudentDeleteOne) Where(ps ...predicate.SubjectStudent) *SubjectStudentDeleteOne {
	ssdo.ssd.mutation.Where(ps...)
	return ssdo
}

// Exec executes the deletion query.
func (ssdo *SubjectStudentDeleteOne) Exec(ctx context.Context) error {
	n, err := ssdo.ssd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{subjectstudent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ssdo *SubjectStudentDeleteOne) ExecX(ctx context.Context) {
	if err := ssdo.Exec(ctx); err != nil {
		panic(err)
	}
}
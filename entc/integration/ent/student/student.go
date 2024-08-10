// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package student

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the student type in the database.
	Label = "student"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeSubjects holds the string denoting the subjects edge name in mutations.
	EdgeSubjects = "subjects"
	// EdgeSubjectStudents holds the string denoting the subject_students edge name in mutations.
	EdgeSubjectStudents = "subject_students"
	// Table holds the table name of the student in the database.
	Table = "students"
	// SubjectsTable is the table that holds the subjects relation/edge. The primary key declared below.
	SubjectsTable = "subject_students"
	// SubjectsInverseTable is the table name for the Subject entity.
	// It exists in this package in order to avoid circular dependency with the "subject" package.
	SubjectsInverseTable = "subjects"
	// SubjectStudentsTable is the table that holds the subject_students relation/edge.
	SubjectStudentsTable = "subject_students"
	// SubjectStudentsInverseTable is the table name for the SubjectStudent entity.
	// It exists in this package in order to avoid circular dependency with the "subjectstudent" package.
	SubjectStudentsInverseTable = "subject_students"
	// SubjectStudentsColumn is the table column denoting the subject_students relation/edge.
	SubjectStudentsColumn = "student_id"
)

// Columns holds all SQL columns for student fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// SubjectsPrimaryKey and SubjectsColumn2 are the table columns denoting the
	// primary key for the subjects relation (M2M).
	SubjectsPrimaryKey = []string{"student_id", "subject_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Student queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySubjectsCount orders the results by subjects count.
func BySubjectsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubjectsStep(), opts...)
	}
}

// BySubjects orders the results by subjects terms.
func BySubjects(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubjectsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubjectStudentsCount orders the results by subject_students count.
func BySubjectStudentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubjectStudentsStep(), opts...)
	}
}

// BySubjectStudents orders the results by subject_students terms.
func BySubjectStudents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubjectStudentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSubjectsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubjectsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SubjectsTable, SubjectsPrimaryKey...),
	)
}
func newSubjectStudentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubjectStudentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, SubjectStudentsTable, SubjectStudentsColumn),
	)
}

// comment from another template.
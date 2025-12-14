package dbx

import "fmt"

// FbsqlBuilder is the builder for Firebird databases.
type FbsqlBuilder struct {
	*BaseBuilder
	qb *BaseQueryBuilder
}

var _ Builder = &FbsqlBuilder{}

func NewFbsqlBuilder(db *DB, executor Executor) Builder {
	return &FbsqlBuilder{
		NewBaseBuilder(db, executor),
		NewBaseQueryBuilder(db),
	}
}

// Select returns a new SelectQuery objecxt that can be used to build a SELECT statement.
// The parameters to this method should be the list of column names to be selected
// A column name may have an optional alias name. For example, Select("id", "my_name As name").
func (b *FbsqlBuilder) Select(cols ...string) *SelectQuery {
	return NewSelectQuery(b, b.db).Select(cols...)
}

func (b *FbsqlBuilder) Model(model interface{}) *ModelQuery {
	return NewModelQuery(model, b.db.FieldMapper, b.db, b)
}

// GeneratePlaceholder generates an anonymous parameter placeholder with the given parameter Id.
func (b *FbsqlBuilder) GeneratePlaceholder(i int) string {
	return fmt.Sprintf("$%v", i)
}

// QueryBuilder returns the query builder supporting the current DB.
func (b *FbsqlBuilder) QueryBuilder() QueryBuilder {
	return b.qb
}

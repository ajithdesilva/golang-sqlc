// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package querytest

import (
	"context"
)

const schemaScopedUpdate = `-- name: SchemaScopedUpdate :exec
UPDATE foo.bar SET name = ? WHERE id = ?
`

type SchemaScopedUpdateParams struct {
	Name string
	ID   int64
}

func (q *Queries) SchemaScopedUpdate(ctx context.Context, arg SchemaScopedUpdateParams) error {
	_, err := q.db.ExecContext(ctx, schemaScopedUpdate, arg.Name, arg.ID)
	return err
}

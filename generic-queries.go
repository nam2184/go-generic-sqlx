package queries

import (
  "github.com/nam2184/generic-queries/model"
	"github.com/jmoiron/sqlx"
)

// Query struct which holds any struct representing a table
type Query[T QueryTypes, I any] struct {
    A     []T
    Tx    *sqlx.Tx
    Q     []string
    Args  []interface{}
    ID    []I
}

func NewQueryMany[T QueryTypes, I any](a []T, tx *sqlx.Tx) *Query[T, I] {
    return &Query[T, I]{
        A:  a,
        Tx: tx,
    }
}

type QueryTypes interface {
    model.TableNamer 
}

type QueryHandlerFunc[T QueryTypes, I any] func(*Query[T, I])


func (f QueryHandlerFunc[T, I]) HandleQuery(q *Query[T, I]) {
    f(q)
}

type QueryHandler[T QueryTypes, I any] interface {
    HandleQuery(q *Query[T, I])
}

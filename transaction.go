package queries

import "github.com/jmoiron/sqlx"

type Transaction[T QueryTypes, I any] struct {
  Handler QueryHandler[T, I]
  Tx *sqlx.Tx
}

func NewTransaction[T QueryTypes, I any](handler QueryHandler[T, I], tx *sqlx.Tx) *Transaction[T, I] {
  return &Transaction[T, I] { Handler : handler, Tx: tx}
}

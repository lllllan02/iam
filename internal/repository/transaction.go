package repository

import "context"

type Transaction interface {
	InTx(context.Context, func(context.Context) error) error
}

func NewTransaction(d *Repo) Transaction {
	return d
}

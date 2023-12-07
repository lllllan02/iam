package repository

import "context"

type MockRepo struct{}

func (MockRepo) Transaction(context.Context, func(c context.Context) error) error {
	return nil
}

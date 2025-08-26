package mps

import (
	"errors"
	"fmt"
	pointrserrs "testing_test/7-pointrs_errs"
)

type Obj map[string]pointrserrs.Entity

func (obj Obj) Search(q string) (pointrserrs.Entity, error) {
	item := obj[q]
	if item.Name == "" {
		errMsg := fmt.Sprintf("No entity with the name: %s", q)
		return item, errors.New(errMsg)
	}
	return item, nil
}

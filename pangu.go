package id

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependency(newGenerator)
}

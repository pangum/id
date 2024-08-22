package id

import (
	"github.com/pangum/id/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	creator := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		creator.New,
	).Build().Build().Apply()
}

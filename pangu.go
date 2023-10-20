package id

import (
	"github.com/pangum/id/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Build().Provide(new(plugin.Creator).New)
}

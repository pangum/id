package id

import (
	"github.com/goexl/id"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
)

type Generator struct {
	id.Generator
}

func newGenerator(config *pangu.Config, logger *logging.Logger) (generator *Generator, err error) {
	panguConfig := new(panguConfig)
	if err = config.Load(panguConfig); nil != err {
		return
	}

	generator = new(Generator)
	_config := panguConfig.Id
	builder := id.New(id.Logger(logger))
	if nil != _config.Autoincrement {
		generator.Generator = builder.Autoincrement().From(_config.Autoincrement.From).Build()
	} else if nil != _config.Snowflake {
		generator.Generator = builder.Snowflake(_config.Snowflake.Node).Build()
	}

	return
}

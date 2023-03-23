package pid

import (
	"github.com/goexl/id"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
)

func newGenerator(config *pangu.Config, logger logging.Logger) (generator Generator, err error) {
	wrap := new(wrapper)
	if err = config.Load(wrap); nil != err {
		return
	}

	conf := wrap.Id
	builder := id.New(id.Logger(logger))
	if nil != conf.Autoincrement {
		generator = builder.Autoincrement().From(conf.Autoincrement.From).Build()
	} else if nil != conf.Snowflake {
		generator = builder.Snowflake(conf.Snowflake.Node).Build()
	}

	return
}

package plugin

import (
	"github.com/goexl/id"
	"github.com/goexl/log"
	"github.com/pangum/pangu"
)

type Constructor struct {
	// 解决命名空间问题
}

func (c *Constructor) New(config *pangu.Config, logger log.Logger) (generator id.Generator, err error) {
	wrap := new(Wrapper)
	if ge := config.Build().Get(wrap); nil != ge {
		err = ge
	} else {
		generator, err = c.new(wrap.Id, logger)
	}

	return
}

func (c *Constructor) new(config *Config, logger log.Logger) (generator id.Generator, err error) {
	builder := id.New().Logger(logger)
	if nil != config.Autoincrement {
		generator = builder.Autoincrement().From(config.Autoincrement.From).Build().Build()
	} else if nil != config.Snowflake {
		generator = builder.Snowflake().Node(config.Snowflake.Node).Build().Build()
	}

	return
}

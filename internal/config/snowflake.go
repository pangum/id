package config

type Snowflake struct {
	// 节点
	Node int `default:"1" json:"node" yaml:"node" xml:"node" toml:"node" validate:"required,min=1,max=1024"`
}

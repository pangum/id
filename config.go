package id

type config struct {
	// 自增算法配置
	Autoincrement *autoincrement `json:"autoincrement" yaml:"autoincrement" xml:"autoincrement" toml:"autoincrement"`
	// 雪花算法配置
	Snowflake *snowflake `json:"snowflake" yaml:"snowflake" xml:"snowflake" toml:"snowflake"`
}

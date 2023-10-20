package plugin

import (
	"github.com/pangum/id/internal/config"
)

type Config struct {
	// 自增算法配置
	Autoincrement *config.Autoincrement `json:"autoincrement" yaml:"autoincrement" xml:"autoincrement" toml:"autoincrement"`
	// 雪花算法配置
	Snowflake *config.Snowflake `json:"snowflake" yaml:"snowflake" xml:"snowflake" toml:"snowflake"`
}

package id

type autoincrement struct {
	// 起始值
	From int64 `default:"1" json:"from" yaml:"from" xml:"from" toml:"from"`
}

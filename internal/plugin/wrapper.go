package plugin

type Wrapper struct {
	Id *Config `json:"id,omitempty" yaml:"id" xml:"id" toml:"id" validate:"required"`
}

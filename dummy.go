package dummy

// Generator provides methods for generating random structures.
type Generator struct {
	config Config
}

// Generate generates a single random structure.
func (g *Generator) Generate() interface{} {
	panic("unimplemented")
}

// New constructs a new Generator with some given configuration options. If no
// options are supplied then the defaults are used.
func New(options ...ConfigOption) *Generator {
	config := DefaultConfig()
	for _, option := range options {
		option(&config)
	}
	return &Generator{config: config}
}

// NewWithConfig constructs a new Generator directly from a Config structures.
func NewWithConfig(config Config) *Generator {
	return New(WithConfig(config))
}

// Generate generates a single random structure with some given configuration
// options. If no options are supplied then the defaults are used.
func Generate(options ...ConfigOption) interface{} {
	return New(options...).Generate()
}

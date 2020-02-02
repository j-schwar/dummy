package dummy

// Config holds configuration data which determines how to generate structures.
type Config struct {
	MaxFieldCount    uint // Maximum number of fields to generate for structures. (default 16)
	MaxNestingDepth  uint // Maximum nesting depth for structures/slices. (default 8)
	GenerateIntegers bool // Should we generate integer values? (default true)
	GenerateFloats   bool // Should we generate floating point values? (default true)
	GenerateNil      bool // Should we generate nil values? (default true)
	GenerateSlices   bool // Should we generate slices? (default true)

	NameGen  NameGenerator  // The field name generator to use.
	ValueGen ValueGenerator // The value generator to use.
}

// ConfigOption is the type alias for option functions which modify a configuration.
type ConfigOption func(cfg *Config)

// DefaultConfig returns the default configuration structure.
func DefaultConfig() Config {
	return Config{
		MaxFieldCount:    16,
		MaxNestingDepth:  8,
		GenerateIntegers: true,
		GenerateFloats:   true,
		GenerateNil:      true,
		GenerateSlices:   true,
		NameGen:          nil, // TODO: populate DefaultConfig::NameGen
		ValueGen:         nil, // TODO: populate DefaultConfig::ValueGen
	}
}

// WithConfig is a config option which overwrites the current configuration
// with a supplied one.
//
// Panics if any configuration fields are nil. Mandatory fields are NameGen and
// ValueGen.
func WithConfig(config Config) ConfigOption {
	if config.NameGen == nil {
		panic("NameGen cannot be nil")
	}
	if config.ValueGen == nil {
		panic("ValueGen cannot be nil")
	}
	return func(cfg *Config) {
		*cfg = config
	}
}

// WithMaxFieldCount is a config option which sets the maximum number of fields
// that a structure can have. Note that this option is also constrained by the
// number of unique names that the name generator can generate. The actual
// maximum value will be the minimum of the two.
//
// Panics if count == 0.
func WithMaxFieldCount(count uint) ConfigOption {
	if count == 0 {
		panic("MaxFieldCount must be > 0")
	}
	return func(cfg *Config) {
		cfg.MaxFieldCount = count
	}
}

// WithMaxNestingDepth is a config option with sets the maximum nesting depth
// for generated structures.
//
// Panics if depth == 0.
func WithMaxNestingDepth(depth uint) ConfigOption {
	if depth == 0 {
		panic("MaxNestingDepth must be > 0")
	}
	return func(cfg *Config) {
		cfg.MaxNestingDepth = depth
	}
}

// ShouldGenerateIntegers is a config option which determines whether integer
// fields can be generated or not. The default if not supplied is true.
func ShouldGenerateIntegers(b bool) ConfigOption {
	return func(cfg *Config) {
		cfg.GenerateIntegers = b
	}
}

// ShouldGenerateFloats is a config option which determines whether floating
// point fields can be generated or not. The default if not supplied is true.
func ShouldGenerateFloats(b bool) ConfigOption {
	return func(cfg *Config) {
		cfg.GenerateFloats = b
	}
}

// ShouldGenerateNil is a config option which determines whether nil values can
// be generated or not. The default if not supplied is true.
func ShouldGenerateNil(b bool) ConfigOption {
	return func(cfg *Config) {
		cfg.GenerateNil = b
	}
}

// ShouldGenerateSlices is a config option which determines whether slice
// values can be generated or not. The default if not supplied is true.
func ShouldGenerateSlices(b bool) ConfigOption {
	return func(cfg *Config) {
		cfg.GenerateSlices = b
	}
}

package dummy

import (
	"math/rand"
	"time"
)

type nameState uint8

const (
	nameNotUsed nameState = iota + 1
	nameUsed
)

// NameGenerator is responsible for generating field names for structures.
type NameGenerator interface {
	// SetSeed sets the seed for this generator.
	SetSeed(seed int64)

	// Count returns how many names this generator can generate until it has
	// to be reset.
	//
	// 0 should be returned iff this generator can run indefinitely. This implies
	// that a generator that can generate no names is invalid.
	Count() int

	// Reset notifies this generator that it can generate names which it had
	// previously generated.
	Reset()

	// Generate generates and returns a random field name.
	//
	// Generated names must be valid exported field names (i.e., start with a
	// capital letter). Generated names must also be unique up until Reset is
	// called. For example, if the name "Foo" is returned, "Foo" cannot be
	// returned again until Reset is called.
	//
	// If unable to generate a name, false should be returned as the second
	// return value, otherwise true must be returned.
	Generate() (string, bool)
}

// ValueGenerator is responsible for generating values for fields.
type ValueGenerator interface {
	// SetSeed sets the seed for this generator.
	SetSeed(seed int)

	// Value generates and returns a random field value.
	Value() interface{}
}

// NewFixedNameGenerator constructs a NameGenerator which will randomly pick
// names from a supplied set. The default seed for the generator is the current
// time which should be sufficient for common use cases.
//
// Panics if names is empty.
func NewFixedNameGenerator(names ...string) NameGenerator {
	if len(names) == 0 {
		panic("cannot construct a fixed name generator from no names")
	}

	gen := new(fixedNameGenerator)
	gen.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	gen.names = make(map[string]nameState)
	for _, name := range names {
		gen.names[name] = nameNotUsed
	}

	return gen
}

type fixedNameGenerator struct {
	r     *rand.Rand
	names map[string]nameState
}

// SetSeed sets the seed for the generator.
func (g *fixedNameGenerator) SetSeed(seed int64) {
	g.r.Seed(seed)
}

// Count returns the number of names that can be generated before the generator
// needs to be reset.
func (g fixedNameGenerator) Count() int {
	return len(g.names)
}

// Reset clears all names names from this generator allowing them to be names
// again.
func (g *fixedNameGenerator) Reset() {
	for k := range g.names {
		g.names[k] = nameNotUsed
	}
}

// Generate picks an unused name and returns it then marks it as used so that
// all generated names will be unique until the generator is reset.
//
// Complexity is O(n) in the number of names registered with the generator.
func (g *fixedNameGenerator) Generate() (string, bool) {
	valid := make([]string, 0, g.Count())
	for k, v := range g.names {
		if v == nameNotUsed {
			valid = append(valid, k)
		}
	}

	// All names are in use, unable to generate a new one.
	if len(valid) == 0 {
		return "", false
	}

	i := g.r.Intn(len(valid))
	name := valid[i]
	g.names[name] = nameUsed
	return name, true
}

package parsers

type format struct {
	results string
	pattern string
	first   string
	second  string
}

func createFormat(
	results string,
	pattern string,
	first string,
	second string,
) Format {
	out := format{
		results: results,
		pattern: pattern,
		first:   first,
		second:  second,
	}

	return &out
}

// Results returns the results variableName
func (obj *format) Results() string {
	return obj.results
}

// Pattern returns the results identifier
func (obj *format) Pattern() string {
	return obj.pattern
}

// First returns the first identifier
func (obj *format) First() string {
	return obj.first
}

// Second returns the second identifier
func (obj *format) Second() string {
	return obj.second
}

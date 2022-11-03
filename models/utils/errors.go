// Various utilitaries used through the app.
package utils

type Errors struct {
	errors map[string][]any
}

// Create a new Errors object.
func NewErrors() *Errors {
	err := &Errors{}
	err.errors = make(map[string][]any)

	return err
}

// Getter.
// The error can only be "get" once.
func (e *Errors) Get(key string) []any {
	results := []any{}

	results, ok := e.errors[key]
	if ok {
		delete(e.errors, key)
	}

	return results
}

// Setter.
// If the key already exists, the val components are appended,
// otherwise simply set
func (e *Errors) Set(key string, val []any) {
	if dest, ok := e.errors[key]; ok {
		e.errors[key] = append(dest, val...)
	} else {
		e.errors[key] = append([]any{}, val...)
	}
}

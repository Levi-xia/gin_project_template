package gdao

// GetEndPoint Get选择器
type GetEndPoint[T any] struct {
	Model      *T
	Table      string
	Conditions map[string]any
	Appends    []string
	Options    []string
	Fields     []string
}

// SelectEndPoint Select选择器
type SelectEndPoint[T any] struct {
	Model      *[]T
	Table      string
	Conditions map[string]any
	Appends    []string
	Options    []string
	Fields     []string
}

// UpdateEndPoint Update选择器
type UpdateEndPoint[T any] struct {
	Table      string
	Rows       map[string]any
	Conditions map[string]any
	Appends    []string
	Options    []string
}

// InsertEndpoint Insert选择器
type InsertEndpoint[T any] struct {
	Table string
	Rows  map[string]any
}

// BatchInsertEndpoint BatchInsert选择器
type BatchInsertEndpoint[T any] struct {
	Table string
	Rows  []map[string]any
}

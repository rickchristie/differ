package differ

type ChangeMap[T comparable] map[T]*ChangeField

// ChangeField represents a field. The Before and After could be Go primitive types (int, string, float, bool, etc.),
// but it could also be representing non-primitive types:
//   - Structs are represented with ChangeMap.
//   - Maps are represented with ChangeMap.
//   - Lists are represented with ChangeMap, with the index changed to string with strconv.Itoa.
//
// When IsNew is true, it means this is a new item in ChangeList or ChangeMap.
type ChangeField struct {
	Key       any
	IsNew     bool
	IsChanged bool
	Before    any
	After     any
}

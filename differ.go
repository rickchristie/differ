package differ

import (
	"fmt"
	"reflect"
)

// Diff will return the list of changes. If the given values are primitive, then the returned ChangeMap will only
// consist of 1 field with the given key. When struct, map, or list are given, the given key will be ignored and
// ChangeMap will contain only keys that has changed within the given struct, map, or list.
func Diff[K comparable](
	key K,
	before any,
	after any,
	allFields bool,
) (
	hasChanges bool,
	changes ChangeMap[K],
	err error,
) {
	tBf := reflect.ValueOf(before)
	tAf := reflect.ValueOf(after)
	kind := tBf.Kind()
	kindAf := tAf.Kind()
	if kind != kindAf {
		// todo: can process comparing int and string, declaring them as changed!
		return false, nil, fmt.Errorf(
			"diff: on key %v, kind %s is not %s",
			key,
			kind.String(),
			kindAf.String(),
		)
	}

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vBf := tBf.Int()
		vAf := tAf.Int()
		return DiffComparable(key, vBf, vAf, allFields)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vBf := tBf.Uint()
		vAf := tAf.Uint()
		return DiffComparable(key, vBf, vAf, allFields)
	case reflect.Float32, reflect.Float64:
		vBf := tBf.Float()
		vAf := tAf.Float()
		return DiffComparable(key, vBf, vAf, allFields)
	case reflect.String:
		vBf := tBf.String()
		vAf := tAf.String()
		return DiffComparable(key, vBf, vAf, allFields)
	case reflect.Bool:
		vBf := tBf.Bool()
		vAf := tAf.Bool()
		return DiffComparable(key, vBf, vAf, allFields)
	case reflect.Struct:
		return DiffStruct(key, before, allFields)
	case reflect.Array, reflect.Slice:
	case reflect.Map:
	case reflect.Chan:
	case reflect.Func:
	case reflect.Invalid:
	case reflect.Uintptr:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.Interface:
	case reflect.Pointer:
	case reflect.UnsafePointer:
	}

	fmt.Println(tBf, tAf)

	return
}

func DiffComparable[K comparable, T comparable](
	key K,
	before T,
	after T,
	allFields bool,
) (
	hasChanges bool,
	changes ChangeMap[K],
	err error,
) {
	changes = make(ChangeMap[K])
	if before != after {
		changes[key] = &ChangeField{
			Key:       key,
			IsChanged: true,
			Before:    before,
			After:     after,
		}
		return true, changes, nil
	}

	if allFields {
		changes[key] = &ChangeField{
			Key:    key,
			Before: before,
			After:  after,
		}
	}
	return false, changes, nil
}

func DiffStruct[K comparable](key K, bf any, af any) (hasChanges bool, changes ChangeMap[K], err error) {
	panic("not implemented yet")
}

func DiffMap[K comparable, V any](bf map[K]V, af map[K]V) (hasChanges bool, changes ChangeMap[K], err error) {
	checked := make(map[any]struct{})
	changes = make(ChangeMap[K])

	// First check keys that exist bf.
	for k, vBf := range bf {
		checked[k] = struct{}{}

		// If key doesn't exist in af, add to ChangeMap.
		vAf, ok := af[k]
		if ok == false {
			changes[k] = &ChangeField{
				Key:    k,
				Before: vBf,
				After:  nil,
			}
			hasChanges = true
			continue
		}

		// Otherwise key exist, check if they are the same by diffing them.
		isChanged, ch, err := Diff[K](k, vBf, vAf, false)
		if err != nil {
			return false, nil, err
		}

		if isChanged {
			changes[k] = ch[k]
			hasChanges = true
			continue
		}
	}

	panic("not implemented yet")

	// Next check keys that exist on af, but skip keys we already checked.
	//for k, v := range af {
	//	if _, ok := checked[k]; ok {
	//		// Skip if already checked.
	//		continue
	//	}
	//
	//}
	//
	//return changes
}

func DiffList[T any](a []T, b []T) string {
	panic("not implemented yet")
}

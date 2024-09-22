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
	return diffVal(key, &tBf, &tAf, allFields)
}

func diffVal[K comparable](
	key K,
	before *reflect.Value,
	after *reflect.Value,
	allFields bool,
) (
	hasChanges bool,
	changes ChangeMap[K],
	err error,
) {
	kind := before.Kind()
	kindAf := after.Kind()
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
		vBf := before.Int()
		vAf := after.Int()
		return DiffComparable(key, vBf, vAf, allFields)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vBf := before.Uint()
		vAf := after.Uint()
		return DiffComparable(key, vBf, vAf, allFields)
	case reflect.Float32, reflect.Float64:
		vBf := before.Float()
		vAf := after.Float()
		return DiffComparable(key, vBf, vAf, allFields)
	case reflect.String:
		vBf := before.String()
		vAf := after.String()
		return DiffComparable(key, vBf, vAf, allFields)
	case reflect.Bool:
		vBf := before.Bool()
		vAf := after.Bool()
		return DiffComparable(key, vBf, vAf, allFields)
	case reflect.Struct:
		return DiffStruct(key, before, allFields)
	case reflect.Array, reflect.Slice:
	case reflect.Map:
		vBf := before.Convert(reflect.TypeOf(make(map[any]any)))
		vAf := after.Convert(reflect.TypeOf(make(map[any]any)))
		return DiffMap(key, &vBf, &vAf)
	case reflect.Chan:
	case reflect.Func:
	case reflect.Complex64:
	case reflect.Complex128:

	case reflect.Interface:
	case reflect.Pointer:

	case reflect.Invalid:
		return false, nil, fmt.Errorf("diff: unsupported type: Invalid")
	case reflect.UnsafePointer:
		return false, nil, fmt.Errorf("diff: unsupported type: UnsafePointer")
	case reflect.Uintptr:
		// Uintptr is used in unsafe black magic. Not supported.
		return false, nil, fmt.Errorf("diff: unsupported type: Uintptr")
	}

	fmt.Println(before, after)

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

func DiffMap[K comparable](
	key K,
	before *reflect.Value,
	after *reflect.Value,
) (
	hasChanges bool,
	changes ChangeMap[K],
	err error,
) {
	//checked := make(map[any]struct{})
	//changes = make(ChangeMap[K])
	//
	//// First check keys that exist before.
	//bfIter := before.MapRange()
	//for bfIter.Next() {
	//	k := bfIter.Key()
	//	valBf := bfIter.Value()
	//	checked[k] = struct{}{}
	//
	//	// If key doesn't exist in after, add to ChangeMap.
	//	valAf := after.MapIndex(k)
	//	if valAf.IsZero() {
	//		changes[k.Interface()] = &ChangeField{
	//			Key:    k,
	//			Before: before.Interface(),
	//			After:  nil,
	//		}
	//		hasChanges = true
	//		continue
	//	}
	//
	//	// Otherwise key exist, check if they are the same by diffing them.
	//	isChanged, ch, err := Diff[K](k, vBf, vAf, false)
	//	if err != nil {
	//		return false, nil, err
	//	}
	//
	//	if isChanged {
	//		changes[k] = ch[k]
	//		hasChanges = true
	//		continue
	//	}
	//}
	//
	//// Next check keys that exist on after, but skip keys we already checked.
	//for k, v := range after {
	//	if _, ok := checked[k]; ok {
	//		// Skip if already checked.
	//		continue
	//	}
	//
	//}
	//
	//return changes

	panic("not implemented yet")
}

func DiffList[T any](a []T, b []T) string {
	panic("not implemented yet")
}

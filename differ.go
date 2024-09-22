package differ

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//	- Hasn't passed json.Marshal yet, could be struct, could be map[string]int.
// After passing json.Marshal:
//	- All maps are map[any]any?
//	- All slices are []any?

// Diff will return the list of changes. If the given values are primitive, then the returned ChangeMap will only
// consist of 1 field with the given key. When struct, map, or list are given, the given key will be ignored and
// ChangeMap will contain only keys that has changed within the given struct, map, or list.
//
// Diff uses json.Marshal internally, it will return error anything in the value cannot be marshalled. This should cover
// most types where a diff would be useful. Because it uses json.Marshal, it will resolve pointers and handle
// interfaces.
func Diff[K comparable](
	key K,
	before any,
	after any,
) (
	hasChanges bool,
	changes ChangeMap[K],
	err error,
) {
	valueBefore := reflect.ValueOf(before)
	kindBefore := valueBefore.Kind()
	valueAfter := reflect.ValueOf(after)
	kindAfter := valueAfter.Kind()

	// If the value is a pointer, then we need to figure out the type of what is being pointed.
	// Unless if it's nil, then we don't care.
	// todo: check for truly nil interface value
	// todo: check for nil interface with concrete type already set, but value is a nil pointer.
	if kindBefore == reflect.Ptr && valueBefore.IsNil() == false {
		resolved := reflect.Indirect(valueBefore)
		kindBefore = resolved.Kind()
	}
	if kindAfter == reflect.Ptr && valueAfter.IsNil() == false {
		resolved := reflect.Indirect(valueAfter)
		kindAfter = resolved.Kind()
	}

	// Prepare value if it's a struct/map, convert it to the typical map[string]any.
	if kindBefore == reflect.Struct || kindBefore == reflect.Map {
		beforeBytes, err := json.Marshal(before)
		if err != nil {
			return false, nil, err
		}
		before = make(map[string]any)
		err = json.Unmarshal(beforeBytes, &before)
		if err != nil {
			return false, nil, err
		}
	}
	if kindAfter == reflect.Struct || kindAfter == reflect.Map {
		afterBytes, err := json.Marshal(after)
		if err != nil {
			return false, nil, err
		}
		after = make(map[string]any)
		err = json.Unmarshal(afterBytes, &after)
		if err != nil {
			return false, nil, err
		}
	}

	// Prepare value if it's a slice, convert it to the typical []any.
	if kindBefore == reflect.Slice || kindBefore == reflect.Array {
		beforeBytes, err := json.Marshal(before)
		if err != nil {
			return false, nil, err
		}
		before = make([]any, 0)
		err = json.Unmarshal(beforeBytes, &before)
		if err != nil {
			return false, nil, err
		}
	}
	if kindAfter == reflect.Slice || kindAfter == reflect.Array {
		afterBytes, err := json.Marshal(after)
		if err != nil {
			return false, nil, err
		}
		after = make([]any, 0)
		err = json.Unmarshal(afterBytes, &after)
		if err != nil {
			return false, nil, err
		}
	}

	// At this point, structs & maps are normalized into map[string]any.
	// Slices and arrays are normalized into []any.
	// The following diff functions no longer needs to check for other values.
	return diff(key, before, after)
}

func diff[K comparable](
	key K,
	before any,
	after any,
) (
	hasChanges bool,
	changes ChangeMap[K],
	err error,
) {
	changes = make(ChangeMap[K])

	if before == nil {
		if after == nil {
			// Both values are nil.
			return false, changes, nil
		}

		// Otherwise the change is a new value.
		changes[key] = &ChangeField{
			Key:       key,
			IsNew:     true,
			IsChanged: true,
			Before:    before,
			After:     after,
		}
		return true, changes, nil
	}

	// This catches when "before" is not nil but "after" is nil.
	if after == nil {
		changes[key] = &ChangeField{
			Key:       key,
			IsNew:     false,
			IsChanged: true,
			Before:    before,
			After:     after,
		}
		return true, changes, nil
	}

	// In most cases, the majority of values in after json.Marshal should be primitive values.
	if valBefore, ok := before.(int); ok {
		valAfter, ok := after.(int)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(bool); ok {
		valAfter, ok := after.(bool)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(string); ok {
		valAfter, ok := after.(string)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(float64); ok {
		valAfter, ok := after.(float64)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(int64); ok {
		valAfter, ok := after.(int64)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(uint64); ok {
		valAfter, ok := after.(uint64)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	// Next on the list, check for maps.
	if valBefore, ok := before.(map[string]any); ok {
		valAfter, ok := after.(map[string]any)
		// The after value is of different type.
		if ok == false {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise, we need to check the diff.
		var mapChanges = ChangeMap[string]{}
		hasChanges, mapChanges, err = diffMap(valBefore, valAfter)
		if err != nil {
			return false, nil, err
		}
		if hasChanges {
			changes[key] = &ChangeField{}
		}

		return false, changes, nil
	}

	// Next on the list, check for slices.
	// todo

	if valBefore, ok := before.(int32); ok {
		valAfter, ok := after.(int32)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(float32); ok {
		valAfter, ok := after.(float32)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(uint); ok {
		valAfter, ok := after.(uint)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(uint32); ok {
		valAfter, ok := after.(uint32)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(int16); ok {
		valAfter, ok := after.(int16)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(int8); ok {
		valAfter, ok := after.(int8)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(uint32); ok {
		valAfter, ok := after.(uint32)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(uint16); ok {
		valAfter, ok := after.(uint16)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	if valBefore, ok := before.(uint8); ok {
		valAfter, ok := after.(uint8)
		// Check for different type or different value.
		if ok == false || valBefore != valAfter {
			changes[key] = &ChangeField{
				Key:       key,
				IsNew:     false,
				IsChanged: true,
				Before:    before,
				After:     after,
			}
			return true, changes, nil
		}

		// Otherwise value is the same.
		return false, changes, nil
	}

	// If we reach this part it means it's a type we don't expect.
	return false, nil, fmt.Errorf("diff: unexpected type: %T %T", before, after)
}

func diffMap(
	before map[string]any,
	after map[string]any,
) (
	hasChanges bool,
	changes ChangeMap[string],
	err error,
) {

	//checked := make(map[any]struct{})
	changes = make(ChangeMap[any])

	// First check all keys on before.
	//for k, beforeVal := range beforeMap {
	//	afterVal, ok := afterMap[k]
	//	if ok == false {
	//		changes[k] = &ChangeField{
	//			Key:       k,
	//			IsChanged: true,
	//			Before:    beforeVal,
	//			After:     afterVal,
	//		}
	//		continue
	//	}
	//
	//	// Otherwise we must diff the before and after value of this key.
	//	// The value of the key could have different types.
	//	// Because we might be comparing between some "any" shit.
	//	// Is there a way to do this without resorting to using reflection again?
	//	//
	//
	//}
	panic("not implemented yet")
}

func DiffSlice[K comparable, T any](
	key K,
	before T,
	after T,
	allFields bool,
) (
	hasChanges bool,
	changes ChangeMap[K],
	err error,
) {
	beforeBytes, err := json.Marshal(before)
	if err != nil {
		return false, nil, err
	}
	beforeSlice := make([]any, 0)
	err = json.Unmarshal(beforeBytes, &beforeSlice)
	if err != nil {
		return false, nil, err
	}

	afterBytes, err := json.Marshal(after)
	if err != nil {
		return false, nil, err
	}
	afterSlice := make([]any, 0)
	err = json.Unmarshal(afterBytes, &afterSlice)
	if err != nil {
		return false, nil, err
	}

	panic("not implemented yet")
}

// todo: switch to implementation by reflection.
//func DiffSlice[K comparable](
//	key K,
//	before []any,
//	after []any,
//	allFields bool,
//) (
//	hasChanges bool,
//	changes ChangeMap[K],
//	err error,
//) {
//	kind := before.Kind()
//	kindAf := after.Kind()
//	if kind != kindAf {
//		return false, nil, fmt.Errorf(
//			"diff: on key %v, kind %s is not %s",
//			key,
//			kind.String(),
//			kindAf.String(),
//		)
//	}
//
//	switch kind {
//	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//		vBf := before.Int()
//		vAf := after.Int()
//		return DiffComparable(key, vBf, vAf, allFields)
//	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
//		vBf := before.Uint()
//		vAf := after.Uint()
//		return DiffComparable(key, vBf, vAf, allFields)
//	case reflect.Float32, reflect.Float64:
//		vBf := before.Float()
//		vAf := after.Float()
//		return DiffComparable(key, vBf, vAf, allFields)
//	case reflect.String:
//		vBf := before.String()
//		vAf := after.String()
//		return DiffComparable(key, vBf, vAf, allFields)
//	case reflect.Bool:
//		vBf := before.Bool()
//		vAf := after.Bool()
//		return DiffComparable(key, vBf, vAf, allFields)
//	case reflect.Struct:
//		return DiffStruct(key, before, allFields)
//	case reflect.Array, reflect.Slice:
//	case reflect.Map:
//		vBf := before.Convert(reflect.TypeOf(make(map[any]any)))
//		vAf := after.Convert(reflect.TypeOf(make(map[any]any)))
//		return diffMap(key, &vBf, &vAf)
//	case reflect.Chan:
//	case reflect.Func:
//	case reflect.Complex64:
//	case reflect.Complex128:
//
//	case reflect.Interface:
//	case reflect.Pointer:
//
//	case reflect.Invalid:
//		return false, nil, fmt.Errorf("diff: unsupported type: Invalid")
//	case reflect.UnsafePointer:
//		return false, nil, fmt.Errorf("diff: unsupported type: UnsafePointer")
//	case reflect.Uintptr:
//		// Uintptr is used in unsafe black magic. Not supported.
//		return false, nil, fmt.Errorf("diff: unsupported type: Uintptr")
//	}
//
//	fmt.Println(before, after)
//
//	return
//}
//
//func DiffComparable[K comparable, T comparable](
//	key K,
//	before T,
//	after T,
//	allFields bool,
//) (
//	hasChanges bool,
//	changes ChangeMap[K],
//	err error,
//) {
//	changes = make(ChangeMap[K])
//	if before != after {
//		changes[key] = &ChangeField{
//			Key:       key,
//			IsChanged: true,
//			Before:    before,
//			After:     after,
//		}
//		return true, changes, nil
//	}
//
//	if allFields {
//		changes[key] = &ChangeField{
//			Key:    key,
//			Before: before,
//			After:  after,
//		}
//	}
//	return false, changes, nil
//}
//
//func DiffStruct[K comparable](key K, bf any, af any) (hasChanges bool, changes ChangeMap[K], err error) {
//	panic("not implemented yet")
//}
//
//func diffMap[K comparable](
//	key K,
//	before *reflect.Value,
//	after *reflect.Value,
//) (
//	hasChanges bool,
//	changes ChangeMap[K],
//	err error,
//) {
//	//checked := make(map[any]struct{})
//	//changes = make(ChangeMap[K])
//	//
//	//// First check keys that exist before.
//	//bfIter := before.MapRange()
//	//for bfIter.Next() {
//	//	k := bfIter.Key()
//	//	valBf := bfIter.Value()
//	//	checked[k] = struct{}{}
//	//
//	//	// If key doesn't exist in after, add to ChangeMap.
//	//	valAf := after.MapIndex(k)
//	//	if valAf.IsZero() {
//	//		changes[k.Interface()] = &ChangeField{
//	//			Key:    k,
//	//			Before: before.Interface(),
//	//			After:  nil,
//	//		}
//	//		hasChanges = true
//	//		continue
//	//	}
//	//
//	//	// Otherwise key exist, check if they are the same by diffing them.
//	//	isChanged, ch, err := Diff[K](k, vBf, vAf, false)
//	//	if err != nil {
//	//		return false, nil, err
//	//	}
//	//
//	//	if isChanged {
//	//		changes[k] = ch[k]
//	//		hasChanges = true
//	//		continue
//	//	}
//	//}
//	//
//	//// Next check keys that exist on after, but skip keys we already checked.
//	//for k, v := range after {
//	//	if _, ok := checked[k]; ok {
//	//		// Skip if already checked.
//	//		continue
//	//	}
//	//
//	//}
//	//
//	//return changes
//
//	panic("not implemented yet")
//}
//
//func DiffList[T any](a []T, b []T) string {
//	panic("not implemented yet")
//}

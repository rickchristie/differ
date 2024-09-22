package differ

/*

Diff(parentKey, map, map) -->
{
  "key1": {
    before: xx,
    after: xx
  },
  "key2": {
    before: xx,
    after: XX
  }
}

Diff(parentKey, 6, 7) -->
{
  "parentKey": {
    before: xx,
    after: xx
  }
}


// What if everything was a slice?

Diff(parentKey, map, map) -->
[
  {
    "key": xx,
    "before": xx,
    "after": xx,
    "children": [
    ]
  }
]

// What do you want?

{
  "key": {
    "childKey": {
      primitiveDiff: {
        before: xx,
        after: yy
      },
      sliceKey: [
        5: {
          before: xx,
          after: xx,
        },
      ],
    }
  }
}

What's the algorithm on the diff, how does it recognize inserted items?

1  1
2  2
3  5
   3


*/

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
	Changes   ChangeMap[]

	Before any
	After  any
}

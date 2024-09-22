package differ

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimitive_EqualUnequal(t *testing.T) {
	type testRow struct {
		name      string
		key       string
		before    any
		after     any
		allFields bool

		expectError      bool
		expectHasChanges bool
		expectChanges    any
	}

	runRows := func(t *testing.T, rows []*testRow) {
		for _, r := range rows {
			t.Run(r.name, func(t *testing.T) {
				hasChanges, changes, err := Diff(r.key, r.before, r.after, r.allFields)
				if r.expectError {
					assert.NotNil(t, err)
					assert.Equal(t, false, hasChanges)
					assert.Equal(t, nil, changes)
					return
				}

				assert.Equal(t, r.expectHasChanges, hasChanges)
				assert.Equal(t, r.expectChanges, changes)
			})
		}
	}

	runRows(t, []*testRow{
		// int
		{
			name:             "int equal",
			key:              "int",
			before:           1,
			after:            1,
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "int not equal #1",
			key:              "int",
			before:           1,
			after:            2,
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(1),
					After:     int64(2),
				},
			},
		},
		{
			name:             "int not equal #2",
			key:              "int",
			before:           0,
			after:            2,
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(0),
					After:     int64(2),
				},
			},
		},
		{
			name:             "int not equal #3",
			key:              "int",
			before:           -1,
			after:            2,
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(-1),
					After:     int64(2),
				},
			},
		},

		// int8
		{
			name:             "int8 equal",
			key:              "int",
			before:           int8(1),
			after:            int8(1),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "int8 not equal #1",
			key:              "int",
			before:           int8(1),
			after:            int8(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(1),
					After:     int64(2),
				},
			},
		},
		{
			name:             "int8 not equal #2",
			key:              "int",
			before:           int8(0),
			after:            int8(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(0),
					After:     int64(2),
				},
			},
		},
		{
			name:             "int8 not equal #3",
			key:              "int",
			before:           int8(-2),
			after:            int8(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(-2),
					After:     int64(2),
				},
			},
		},

		// int16
		{
			name:             "int16 equal",
			key:              "int",
			before:           int16(1),
			after:            int16(1),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "int16 not equal #1",
			key:              "int",
			before:           int16(1),
			after:            int16(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(1),
					After:     int64(2),
				},
			},
		},
		{
			name:             "int16 not equal #2",
			key:              "int",
			before:           int16(0),
			after:            int16(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(0),
					After:     int64(2),
				},
			},
		},
		{
			name:             "int16 not equal #3",
			key:              "int",
			before:           int16(-1),
			after:            int16(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(-1),
					After:     int64(2),
				},
			},
		},

		// int32
		{
			name:             "int32 equal",
			key:              "int",
			before:           int32(1),
			after:            int32(1),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "int32 not equal #1",
			key:              "int",
			before:           int32(1),
			after:            int32(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(1),
					After:     int64(2),
				},
			},
		},
		{
			name:             "int32 not equal #2",
			key:              "int",
			before:           int32(0),
			after:            int32(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(0),
					After:     int64(2),
				},
			},
		},
		{
			name:             "int32 not equal #3",
			key:              "int",
			before:           int32(-1),
			after:            int32(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(-1),
					After:     int64(2),
				},
			},
		},

		// int64
		{
			name:             "int64 equal",
			key:              "int",
			before:           int64(1),
			after:            int64(1),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "int64 not equal #1",
			key:              "int",
			before:           int64(1),
			after:            int64(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(1),
					After:     int64(2),
				},
			},
		},
		{
			name:             "int64 not equal #2",
			key:              "int",
			before:           int64(0),
			after:            int64(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"int": {
					Key:       "int",
					IsChanged: true,
					Before:    int64(0),
					After:     int64(2),
				},
			},
		},

		// uint
		{
			name:             "uint equal",
			key:              "uint",
			before:           uint(1),
			after:            uint(1),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "uint not equal #1",
			key:              "uint",
			before:           uint(1),
			after:            uint(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"uint": {
					Key:       "uint",
					IsChanged: true,
					Before:    uint64(1),
					After:     uint64(2),
				},
			},
		},
		{
			name:             "uint not equal #2",
			key:              "uint",
			before:           uint(0),
			after:            uint(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"uint": {
					Key:       "uint",
					IsChanged: true,
					Before:    uint64(0),
					After:     uint64(2),
				},
			},
		},

		// uint8
		{
			name:             "uint8 equal",
			key:              "uint",
			before:           uint8(1),
			after:            uint8(1),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "uint8 not equal #1",
			key:              "uint",
			before:           uint8(1),
			after:            uint8(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"uint": {
					Key:       "uint",
					IsChanged: true,
					Before:    uint64(1),
					After:     uint64(2),
				},
			},
		},
		{
			name:             "uint8 not equal #2",
			key:              "uint",
			before:           uint8(0),
			after:            uint8(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"uint": {
					Key:       "uint",
					IsChanged: true,
					Before:    uint64(0),
					After:     uint64(2),
				},
			},
		},

		// uint16
		{
			name:             "uint16 equal",
			key:              "uint",
			before:           uint16(1),
			after:            uint16(1),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "uint16 not equal #1",
			key:              "uint",
			before:           uint16(1),
			after:            uint16(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"uint": {
					Key:       "uint",
					IsChanged: true,
					Before:    uint64(1),
					After:     uint64(2),
				},
			},
		},
		{
			name:             "uint16 not equal #2",
			key:              "uint",
			before:           uint16(0),
			after:            uint16(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"uint": {
					Key:       "uint",
					IsChanged: true,
					Before:    uint64(0),
					After:     uint64(2),
				},
			},
		},

		// uint32
		{
			name:             "uint32 equal",
			key:              "uint",
			before:           uint32(1),
			after:            uint32(1),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "uint32 not equal #1",
			key:              "uint",
			before:           uint32(1),
			after:            uint32(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"uint": {
					Key:       "uint",
					IsChanged: true,
					Before:    uint64(1),
					After:     uint64(2),
				},
			},
		},
		{
			name:             "uint32 not equal #2",
			key:              "uint",
			before:           uint32(0),
			after:            uint32(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"uint": {
					Key:       "uint",
					IsChanged: true,
					Before:    uint64(0),
					After:     uint64(2),
				},
			},
		},

		// uint64
		{
			name:             "uint64 equal",
			key:              "uint",
			before:           uint64(1),
			after:            uint64(1),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "uint64 not equal #1",
			key:              "uint",
			before:           uint64(1),
			after:            uint64(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"uint": {
					Key:       "uint",
					IsChanged: true,
					Before:    uint64(1),
					After:     uint64(2),
				},
			},
		},
		{
			name:             "uint64 not equal #2",
			key:              "uint",
			before:           uint64(0),
			after:            uint64(2),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"uint": {
					Key:       "uint",
					IsChanged: true,
					Before:    uint64(0),
					After:     uint64(2),
				},
			},
		},

		// float32
		{
			name:             "float32 equal",
			key:              "float",
			before:           float32(1.2342),
			after:            float32(1.2342),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "float32 not equal #1",
			key:              "float",
			before:           float32(1.2342),
			after:            float32(1.2341),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"float": {
					Key:       "float",
					IsChanged: true,
					Before:    float64(float32(1.2342)),
					After:     float64(float32(1.2341)),
				},
			},
		},
		{
			name:             "float32 not equal #2",
			key:              "float",
			before:           float32(0.0),
			after:            float32(2.78),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"float": {
					Key:       "float",
					IsChanged: true,
					Before:    float64(float32(0.0)),
					After:     float64(float32(2.78)),
				},
			},
		},

		// float64
		{
			name:             "float64 equal",
			key:              "float",
			before:           float64(1.2342),
			after:            float64(1.2342),
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "float64 not equal #1",
			key:              "float",
			before:           float64(1.2342),
			after:            float64(1.2341),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"float": {
					Key:       "float",
					IsChanged: true,
					Before:    float64(1.2342),
					After:     float64(1.2341),
				},
			},
		},
		{
			name:             "float64 not equal #2",
			key:              "float",
			before:           float64(0.0),
			after:            float64(2.78),
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"float": {
					Key:       "float",
					IsChanged: true,
					Before:    float64(0.0),
					After:     float64(2.78),
				},
			},
		},

		// string
		{
			name:             "string equal",
			key:              "string",
			before:           "hello world!",
			after:            "hello world!",
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "string not equal #1",
			key:              "string",
			before:           "hello world!",
			after:            "hallo world!",
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"string": {
					Key:       "string",
					IsChanged: true,
					Before:    "hello world!",
					After:     "hallo world!",
				},
			},
		},
		{
			name:             "string not equal #2",
			key:              "string",
			before:           "hello world!",
			after:            "",
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"string": {
					Key:       "string",
					IsChanged: true,
					Before:    "hello world!",
					After:     "",
				},
			},
		},

		// bool
		{
			name:             "bool equal #1",
			key:              "bool",
			before:           true,
			after:            true,
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "bool equal #2",
			key:              "bool",
			before:           false,
			after:            false,
			expectHasChanges: false,
			expectChanges:    ChangeMap[string]{},
		},
		{
			name:             "bool not equal #1",
			key:              "bool",
			before:           true,
			after:            false,
			expectHasChanges: true,
			expectChanges: ChangeMap[string]{
				"bool": {
					Key:       "bool",
					IsChanged: true,
					Before:    true,
					After:     false,
				},
			},
		},
	})
}

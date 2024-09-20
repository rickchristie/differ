package differ

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimitive(t *testing.T) {
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

		//
	})
}

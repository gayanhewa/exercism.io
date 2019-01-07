package hierarchy

import (
	"testing"
)

func getTestInstance() *UserHierarchy {
	roles := []Role{
		Role{1, "System Admin", 0},
		Role{2, "Location Manager", 1},
		Role{3, "Supervisor", 2},
		Role{4, "Employee", 3},
		Role{5, "Trainer", 3},
	}

	users := []User{
		User{1, "Adam", 1},
		User{2, "Emily", 4},
		User{3, "Sam", 3},
		User{4, "Mary", 2},
		User{5, "Steve", 5},
		User{6, "Tim", 2},
		User{7, "Harry", 5},
		User{8, "Lance", 1},
	}

	return New(users, roles)
}
func TestGetSubordinates(t *testing.T) {
	var userHierarchy *UserHierarchy
	userHierarchy = getTestInstance()
	for _, tc := range hierarchyTestCases {
		users, err := userHierarchy.GetSubordinates(uint64(tc.input))
		if tc.err != "" {
			var _ error = err
			if err == nil {
				t.Fatalf("GetSubordinates(%d). error is nil.", tc.input)
			}
		} else {
			var exists bool
			for _, user := range users {
				exists = false
				for _, userId := range tc.expected {
					if userId == user.ID {
						exists = true
						break
					}
				}
				if exists == false {
					t.Fatalf("GetSubordinates(%d). expected results does not match.", tc.input)
				}
			}
		}
	}
}

func BenchmarkGetSubordinates(b *testing.B) {
	var userHierarchy *UserHierarchy
	userHierarchy = getTestInstance()
	for i := 0; i < b.N; i++ {
		for _, tc := range hierarchyTestCases {
			userHierarchy.GetSubordinates(uint64(tc.input))
		}
	}
}

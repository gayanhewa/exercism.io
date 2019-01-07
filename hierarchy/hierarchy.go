package hierarchy

import (
	"errors"
)

// UserHierarchy is an exported struct that will contain the list of users and roles.
type UserHierarchy struct {
	users []User
	roles []Role
}

// New is a factory for UserHierarchy struct for this package. Allows you to initialize the struct.
func New(users []User, roles []Role) *UserHierarchy {
	return &UserHierarchy{
		users,
		roles,
	}
}

// GetSubordinates returns a slice of users.
func (u *UserHierarchy) GetSubordinates(userID uint64) ([]User, error) {
	user, err := u.getUser(userID)
	if err != nil {
		return nil, err
	}
	var roles []Role
	roles = u.getChildRoles(user.Role, roles)
	return u.getUsersByRoles(roles), nil
}

// getChildRoles will return all the child roles and grand children.
func (u *UserHierarchy) getChildRoles(roleId uint64, roles []Role) []Role {
	var childRoles []Role
	for _, role := range u.roles {
		if roleId == role.Parent {
			roles = append(roles, role)
			childRoles = append(childRoles, role)
		}
	}
	if len(childRoles) == 0 {
		return roles
	}
	for _, i := range childRoles {
		return u.getChildRoles(i.ID, roles)
	}
	return nil
}

// getUsersByRoles is a helper function that returns all users that belong to the given slice of roles.
func (u *UserHierarchy) getUsersByRoles(roles []Role) (users []User) {
	for _, role := range roles {
		for _, user := range u.users {
			if user.Role == role.ID {
				users = append(users, user)
			}
		}
	}
	return users
}

// getUser helper will return the user struct by the userID
func (u *UserHierarchy) getUser(userID uint64) (user User, err error) {
	for _, user := range u.users {
		if user.ID == userID {
			return user, nil
		}
	}
	return user, errors.New("invalid user id")
}

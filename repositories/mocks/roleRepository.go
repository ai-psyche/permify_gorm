package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/ai-psyche/permify_gorm/collections"
	"github.com/ai-psyche/permify_gorm/models"
	"github.com/ai-psyche/permify_gorm/repositories/scopes"
)

// RoleRepository is an autogenerated mock type for the RoleRepository type
type RoleRepository struct {
	mock.Mock
}

// Migrate provides a mock function.
func (_m *RoleRepository) Migrate() (err error) {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRoleByID provides a mock function with given fields: ID
func (_m *RoleRepository) GetRoleByID(ID uint) (role models.Role, err error) {
	ret := _m.Called(ID)

	var r0 models.Role
	if rf, ok := ret.Get(0).(func(uint) models.Role); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(models.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleByIDWithPermissions provides a mock function with given fields: ID
func (_m *RoleRepository) GetRoleByIDWithPermissions(ID uint) (role models.Role, err error) {
	ret := _m.Called(ID)

	var r0 models.Role
	if rf, ok := ret.Get(0).(func(uint) models.Role); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(models.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleByGuardName provides a mock function with given fields: guardName
func (_m *RoleRepository) GetRoleByGuardName(guardName string) (role models.Role, err error) {
	ret := _m.Called(guardName)

	var r0 models.Role
	if rf, ok := ret.Get(0).(func(string) models.Role); ok {
		r0 = rf(guardName)
	} else {
		r0 = ret.Get(0).(models.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guardName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleByGuardNameWithPermissions provides a mock function with given fields: guardName
func (_m *RoleRepository) GetRoleByGuardNameWithPermissions(guardName string) (role models.Role, err error) {
	ret := _m.Called(guardName)

	var r0 models.Role
	if rf, ok := ret.Get(0).(func(string) models.Role); ok {
		r0 = rf(guardName)
	} else {
		r0 = ret.Get(0).(models.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guardName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoles provides a mock function with given fields: IDs
func (_m *RoleRepository) GetRoles(IDs []uint) (roles collections.Role, err error) {
	ret := _m.Called(IDs)

	var r0 collections.Role
	if rf, ok := ret.Get(0).(func([]uint) collections.Role); ok {
		r0 = rf(IDs)
	} else {
		r0 = ret.Get(0).(collections.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint) error); ok {
		r1 = rf(IDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRolesWithPermissions provides a mock function with given fields: IDs
func (_m *RoleRepository) GetRolesWithPermissions(IDs []uint) (roles collections.Role, err error) {
	ret := _m.Called(IDs)

	var r0 collections.Role
	if rf, ok := ret.Get(0).(func([]uint) collections.Role); ok {
		r0 = rf(IDs)
	} else {
		r0 = ret.Get(0).(collections.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint) error); ok {
		r1 = rf(IDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRolesByGuardNames provides a mock function with given fields: guardNames
func (_m *RoleRepository) GetRolesByGuardNames(guardNames []string) (roles collections.Role, err error) {
	ret := _m.Called(guardNames)

	var r0 collections.Role
	if rf, ok := ret.Get(0).(func([]string) collections.Role); ok {
		r0 = rf(guardNames)
	} else {
		r0 = ret.Get(0).(collections.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(guardNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRolesByGuardNamesWithPermissions provides a mock function with given fields: guardNames
func (_m *RoleRepository) GetRolesByGuardNamesWithPermissions(guardNames []string) (roles collections.Role, err error) {
	ret := _m.Called(guardNames)

	var r0 collections.Role
	if rf, ok := ret.Get(0).(func([]string) collections.Role); ok {
		r0 = rf(guardNames)
	} else {
		r0 = ret.Get(0).(collections.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(guardNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleIDs provides a mock function with given fields: pagination
func (_m *RoleRepository) GetRoleIDs(pagination scopes.GormPager) (roleIDs []uint, totalCount int64, err error) {
	ret := _m.Called(pagination)

	var r0 []uint
	if rf, ok := ret.Get(0).(func(scopes.GormPager) []uint); ok {
		r0 = rf(pagination)
	} else {
		r0 = ret.Get(0).([]uint)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(scopes.GormPager) int64); ok {
		r1 = rf(pagination)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(scopes.GormPager) error); ok {
		r2 = rf(pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRoleIDsOfUser provides a mock function with given fields: userID, pagination
func (_m *RoleRepository) GetRoleIDsOfUser(userID uint, pagination scopes.GormPager) (roleIDs []uint, totalCount int64, err error) {
	ret := _m.Called(userID, pagination)

	var r0 []uint
	if rf, ok := ret.Get(0).(func(uint, scopes.GormPager) []uint); ok {
		r0 = rf(userID, pagination)
	} else {
		r0 = ret.Get(0).([]uint)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(uint, scopes.GormPager) int64); ok {
		r1 = rf(userID, pagination)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint, scopes.GormPager) error); ok {
		r2 = rf(userID, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRoleIDsOfPermission provides a mock function with given fields: userID, pagination
func (_m *RoleRepository) GetRoleIDsOfPermission(permissionID uint, pagination scopes.GormPager) (roleIDs []uint, totalCount int64, err error) {
	ret := _m.Called(permissionID, pagination)

	var r0 []uint
	if rf, ok := ret.Get(0).(func(uint, scopes.GormPager) []uint); ok {
		r0 = rf(permissionID, pagination)
	} else {
		r0 = ret.Get(0).([]uint)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(uint, scopes.GormPager) int64); ok {
		r1 = rf(permissionID, pagination)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint, scopes.GormPager) error); ok {
		r2 = rf(permissionID, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FirstOrCreate provides a mock function with given fields: permission
func (_m *RoleRepository) FirstOrCreate(role *models.Role) error {
	ret := _m.Called(role)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Role) error); ok {
		r0 = rf(role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Updates provides a mock function with given fields: permission, updates
func (_m *RoleRepository) Updates(role *models.Role, updates map[string]interface{}) (err error) {
	ret := _m.Called(role, updates)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Role, map[string]interface{}) error); ok {
		r0 = rf(role, updates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: permission
func (_m *RoleRepository) Delete(role *models.Role) (err error) {
	ret := _m.Called(role)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Role) error); ok {
		r0 = rf(role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddPermissions provides a mock function with given fields: role, permissions
func (_m *RoleRepository) AddPermissions(role *models.Role, permissions collections.Permission) error {
	ret := _m.Called(role, permissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Role, collections.Permission) error); ok {
		r0 = rf(role, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReplacePermissions provides a mock function with given fields: role, permissions
func (_m *RoleRepository) ReplacePermissions(role *models.Role, permissions collections.Permission) error {
	ret := _m.Called(role, permissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Role, collections.Permission) error); ok {
		r0 = rf(role, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePermissions provides a mock function with given fields: role, permissions
func (_m *RoleRepository) RemovePermissions(role *models.Role, permissions collections.Permission) error {
	ret := _m.Called(role, permissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Role, collections.Permission) error); ok {
		r0 = rf(role, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearPermissions provides a mock function with given fields: role
func (_m *RoleRepository) ClearPermissions(role *models.Role) error {
	ret := _m.Called(role)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Role) error); ok {
		r0 = rf(role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasPermission provides a mock function with given fields: role, permissions
func (_m *RoleRepository) HasPermission(roles collections.Role, permission models.Permission) (b bool, err error) {
	ret := _m.Called(roles, permission)

	var r0 bool
	if rf, ok := ret.Get(0).(func(collections.Role, models.Permission) bool); ok {
		r0 = rf(roles, permission)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(collections.Role, models.Permission) error); ok {
		r1 = rf(roles, permission)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAllPermissions provides a mock function with given fields: roles, permissions
func (_m *RoleRepository) HasAllPermissions(roles collections.Role, permissions collections.Permission) (b bool, err error) {
	ret := _m.Called(roles, permissions)

	var r0 bool
	if rf, ok := ret.Get(0).(func(collections.Role, collections.Permission) bool); ok {
		r0 = rf(roles, permissions)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(collections.Role, collections.Permission) error); ok {
		r1 = rf(roles, permissions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAnyPermissions provides a mock function with given fields: roles, permissions
func (_m *RoleRepository) HasAnyPermissions(roles collections.Role, permissions collections.Permission) (b bool, err error) {
	ret := _m.Called(roles, permissions)

	var r0 bool
	if rf, ok := ret.Get(0).(func(collections.Role, collections.Permission) bool); ok {
		r0 = rf(roles, permissions)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(collections.Role, collections.Permission) error); ok {
		r1 = rf(roles, permissions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

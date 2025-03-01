
<h1 align="center">
    <img src="https://raw.githubusercontent.com/Permify/permify/master/assets/permify-logo.svg" alt="Permify logo" width="336px" /><br />
</h1>

[![Go Reference](https://pkg.go.dev/badge/github.com/ai-psyche/permify_gorm.svg)](https://pkg.go.dev/github.com/ai-psyche/permify_gorm)
[![Go Report Card](https://goreportcard.com/badge/github.com/ai-psyche/permify_gorm)](https://goreportcard.com/report/github.com/ai-psyche/permify_gorm)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ai-psyche/permify_gorm)
![GitHub](https://img.shields.io/github/license/Permify/go-role)
[![Twitter Follow](https://img.shields.io/twitter/follow/GetPermify?style=social)](https://twitter.com/GetPermify)

## Associate users with roles and permissions

This package allows you to manage user permissions and roles in your database.

## 👇 Setup

Install

```shell
go get github.com/ai-psyche/permify_gorm
```

Run All Tests

```shell
go test ./...
```

Get the database driver for gorm that you will be using

```shell
# mysql 
go get gorm.io/driver/mysql 
# or postgres
go get gorm.io/driver/postgres
# or sqlite
go get gorm.io/driver/sqlite
# or sqlserver
go get gorm.io/driver/sqlserver
# or clickhouse
go get gorm.io/driver/clickhouse
```

Import permify.

```go
import permify `github.com/ai-psyche/permify_gorm`
```

Initialize the new Permify.

```go
// initialize the database. (you can use all gorm's supported databases)
db, _ := gorm.Open(mysql.Open("user:password@tcp(host:3306)/db?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

// New initializer for Permify
// If migration is true, it generate all tables in the database if they don't exist.
permify, _ := permify.New(permify.Options{
	Migrate: true,
	DB: db,
})
```

## 🚲 Basic Usage

This package allows users to be associated with permissions and roles. Each role is associated with multiple permissions.

```go
// CreateRole create new role.
// Name parameter is converted to guard name. example: senior $#% associate -> senior-associate.
// If a role with the same name has been created before, it will not create it again. (FirstOrCreate)
// First parameter is role name, second parameter is role description.
err := permify.CreateRole("admin", "role description")

// CreatePermission create new permission.
// Name parameter is converted to guard name. example: create $#% contact -> create-contact.
// If a permission with the same name has been created before, it will not create it again. (FirstOrCreate)
err := permify.CreatePermission("edit user details", "")
```

Permissions can be added to a role using AddPermissionsToRole method in different ways:

```go
// first parameter is role id
err := permify.AddPermissionsToRole(1, "edit user details")
// or
err := permify.AddPermissionsToRole("admin", []string{"edit user details", "create contact"})
// or
err := permify.AddPermissionsToRole("admin", []uint{1, 3})
```

With using these methods you can remove and overwrite permissions:

```go
// overwrites the permissions of the role according to the permission names or ids.
err := permify.ReplacePermissionsToRole("admin", []string{"edit user details", "create contact"})

// remove permissions from role according to the permission names or ids.
err := permify.RemovePermissionsFromRole("admin", []string{"edit user details"})
```

Basic fetch queries:

```go
// Fetch all the roles. (with pagination option).
// If withPermissions is true, it will preload the permissions to the role.
// If pagination is nil, it returns without paging.
roles, totalCount, err := permify.GetAllRoles(options.RoleOption{
	WithPermissions: true,
	Pagination: &utils.Pagination{
		Page: 1,
		Limit: 1,
	},
})

// without paging.
roles, totalCount, err := permify.GetAllRoles(options.RoleOption{
    WithPermissions: false,
})

// The data returned is a collection of roles.
// Collections provides a fluent convenient wrapper for working with arrays of data.
fmt.Println(roles.IDs())
fmt.Println(roles.Names())
fmt.Println(roles.Permissions().Names())

// Fetch all permissions of the user that come with direct and roles.
permissions, _ := permify.GetAllPermissionsOfUser(1)

// Fetch all direct permissions of the user. (with pagination option)
permissions, totalCount, err := permify.GetDirectPermissionsOfUser(1, options.PermissionOption{
    Pagination: &utils.Pagination{
        Page: 1,
        Limit: 10,
    },
})
```

Controls

```go
// does the role or any of the roles have given permission?
can, err := permify.RoleHasPermission("admin", "edit user details")

// does the role or roles have any of the given permissions?
can, err := permify.RoleHasAnyPermissions([]string{"admin", "manager"}, []string{"edit user details", "create contact"})

// does the role or roles have all the given permissions?
can, err := permify.RoleHasAllPermissions("admin", []string{"edit user details", "create contact"})

// does the user have the given permission? (including the permissions of the roles)
can, err := permify.UserHasPermission(1, "edit user details")

// does the user have the given permission? (not including the permissions of the roles)
can, err := permify.UserHasDirectPermission(1, "edit user details")

// does the user have any of the given permissions? (including the permissions of the roles)
can, err := permify.UserHasAnyPermissions(1, []uint{1, 2})

// does the user have all the given roles?
can, err := permify.UserHasAllRoles(1, []string{"admin", "manager"})

// does the user have any of the given roles?
can, err := permify.UserHasAnyRoles(1, []string{"admin", "manager"})
```


## 🚘 Using permissions via roles

### Adding Role

Add roles to user according to the role names or ids:

```go
// add one role to user
err := permify.AddRolesToUser(1, "admin")

// you can also add multiple roles at once
err := permify.AddRolesToUser(1, []string{"admin", "manager"})
// or
err := permify.AddRolesToUser(1, []uint{1,2})
```

Replace the roles of the user according to the role names or ids:

```go
// remove all user roles and add admin role
err := permify.ReplaceRolesToUser(1, "admin")

// you can also replace multiple roles at once
err := permify.ReplaceRolesToUser(1, []string{"admin", "manager"})
// or
err := permify.RemoveRolesFromUser(1, []uint{1,2})
```

Remove the roles of the user according to the role names or ids:

```go
// remove one role to user
err := permify.RemoveRolesFromUser(1, "admin")

// you can also remove multiple roles at once
err := permify.RemoveRolesFromUser(1, []string{"admin", "manager"})
// or
err := permify.RemoveRolesFromUser(1, []uint{1,2})
```

Control Roles

```go
// does the user have the given role?
can, err := permify.UserHasRole(1, "admin")

// does the user have all the given roles?
can, err := permify.UserHasAllRoles(1, []string{"admin", "manager"})

// does the user have any of the given roles?
can, err := permify.UserHasAnyRoles(1, []string{"admin", "manager"})
```

Get User's Roles

```go
roles, totalCount, err := permify.GetRolesOfUser(1, options.RoleOption{
    WithPermissions: true, // preload role's permissions
    Pagination: &utils.Pagination{
        Page: 1,
        Limit: 1,
    },
})

// the data returned is a collection of roles. 
// Collections provides a fluent convenient wrapper for working with arrays of data.
fmt.Println(roles.IDs())
fmt.Println(roles.Names())
fmt.Println(roles.Len())
fmt.Println(roles.Permissions().Names())
```

Add Permissions to Roles

```go
// add one permission to role
// first parameter can be role name or id, second parameter can be permission name(s) or id(s).
err := permify.AddPermissionsToRole("admin", "edit contact details")

// you can also add multiple permissions at once
err := permify.AddPermissionsToRole("admin", []string{"edit contact details", "delete user"})
// or
err := permify.AddPermissionsToRole("admin", []uint{1, 2})
```

Remove Permissions from Roles

```go
// remove one permission to role
err := permify.RemovePermissionsFromRole("admin", "edit contact details")

// you can also add multiple permissions at once
err := permify.RemovePermissionsFromRole("admin", []string{"edit contact details", "delete user"})
// or
err := permify.RemovePermissionsFromRole("admin", []uint{1, 2})
```

Control Role's Permissions

```go
// does the role or any of the roles have given permission?
can, err := permify.RoleHasPermission([]string{"admin", "manager"}, "edit contact details")

// does the role or roles have all the given permissions?
can, err := permify.RoleHasAllPermissions("admin", []string{"edit contact details", "delete contact"})

// does the role or roles have any of the given permissions?
can, err := permify.RoleHasAnyPermissions(1, []string{"edit contact details", "delete contact"})
```

Get Role's Permissions

```go
permissions, totalCount, err := permify.GetPermissionsOfRoles([]string{"admin", "manager"}, options.PermissionOption{
    Pagination: &utils.Pagination{
        Page: 1,
        Limit: 1,
    },
})

// the data returned is a collection of permissions. 
// Collections provides a fluent convenient wrapper for working with arrays of data.
fmt.Println(permissions.IDs())
fmt.Println(permissions.Names())
fmt.Println(permissions.Len())
```

## 🚤 Direct Permissions

### Adding Direct Permissions

Add direct permission or permissions to user according to the permission names or ids.

```go
// add one permission to user
err := permify.AddPermissionsToUser(1, "edit contact details")

// you can also add multiple permissions at once
err := permify.AddPermissionsToUser(1, []string{"edit contact details", "create contact"})
// or
err := permify.AddPermissionsToUser(1, []uint{1,2})
```

Remove the roles of the user according to the role names or ids:

```go
// remove one role to user
err := permify.RemovePermissionsFromUser(1, "edit contact details")

// you can also remove multiple permissions at once
err := permify.RemovePermissionsFromUser(1, []string{"edit contact details", "create contact"})
// or
err := permify.RemovePermissionsFromUser(1, []uint{1,2})
```

Control Permissions

```go
// ALL PERMISSIONS

// does the user have the given permission? (including the permissions of the roles)
can, err := permify.UserHasPermission(1, "edit contact details")

// does the user have all the given permissions? (including the permissions of the roles)
can, err := permify.UserHasAllPermissions(1, []string{"edit contact details", "delete contact"})

// does the user have any of the given permissions? (including the permissions of the roles).
can, err := permify.UserHasAnyPermissions(1, []string{"edit contact details", "delete contact"})

// DIRECT PERMISSIONS

// does the user have the given permission? (not including the permissions of the roles)
can, err := permify.UserHasDirectPermission(1, "edit contact details")

// does the user have all the given permissions? (not including the permissions of the roles)
can, err := permify.UserHasAllDirectPermissions(1, []string{"edit contact details", "delete contact"})

// does the user have any of the given permissions? (not including the permissions of the roles)
can, err := permify.UserHasAnyDirectPermissions(1, []string{"edit contact details", "delete contact"})
```

Get User's All Permissions

```go
permissions, err := permify.GetAllPermissionsOfUser(1)

// the data returned is a collection of permissions. 
// Collections provides a fluent convenient wrapper for working with arrays of data.
fmt.Println(permissions.IDs())
fmt.Println(permissions.Names())
fmt.Println(permissions.Len())
```

Get User's Direct Permissions

```go
permissions, totalCount, err := permify.GetDirectPermissionsOfUser(1, options.PermissionOption{
    Pagination: &utils.Pagination{
        Page:  1,
        Limit: 10,
    },
})

// the data returned is a collection of permissions.
// Collections provides a fluent convenient wrapper for working with arrays of data.
fmt.Println(permissions.IDs())
fmt.Println(permissions.Names())
fmt.Println(permissions.Len())
```

## 🚀 Using your user model

You can create the relationships between the user and the role and permissions in this manner. In this way:

- You can manage user preloads
- You can create foreign key between users and pivot tables (user_roles, user_permissions).

```go
import (
    "gorm.io/gorm"
    models `github.com/ai-psyche/permify_gorm/models`
)

type User struct {
    gorm.Model
    Name string

    // permify
    Roles       []models.Role       `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
    Permissions []models.Permission `gorm:"many2many:user_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
```

## ⁉️ Error Handling

### ErrRecordNotFound

You can use error handling in the same way as gorm. for example:

```go
// check if returns RecordNotFound error
permission, err := permify.GetPermission(1)
if errors.Is(err, gorm.ErrRecordNotFound) {
	// record not found
}
```

### Errors

[Errors List](https://github.com/go-gorm/gorm/blob/master/errors.go)


Stargazers
-----------

[![Stargazers repo roster for @Permify/go-role](https://reporoster.com/stars/Permify/go-role)](https://github.com/ai-psyche/permify_gorm/stargazers)


## Community & Support
Join our [Discord channel](https://discord.gg/MJbUjwskdH) for issues, feature requests, feedbacks or anything else. We love to talk about authorization and access control :heart:

<p align="left">
<a href="https://discord.gg/MJbUjwskdH">
 <img height="70px" width="70px" alt="permify | Discord" src="https://user-images.githubusercontent.com/39353278/187209316-3d01a799-c51b-4eaa-8f52-168047078a14.png" />
</a>
<a href="https://twitter.com/GetPermify">
  <img height="70px" width="70px" alt="permify | Twitter" src="https://user-images.githubusercontent.com/39353278/187209323-23f14261-d406-420d-80eb-1aa707a71043.png"/>
</a>
<a href="https://www.linkedin.com/company/permifyco">
  <img height="70px" width="70px" alt="permify | Linkedin" src="https://user-images.githubusercontent.com/39353278/187209321-03293a24-6f63-4321-b362-b0fc89fdd879.png" />
</a>
</p>



package main

import (
	"fmt"
	"test/test-casbin/casbin"
	"testing"
)

func TestPromission(t *testing.T) {
	e := casbin.Casbin()
	ok := e.Enforce("rokey", "data11", "write")
	fmt.Printf("rokey data11 write %t\n", ok)
	ok = e.Enforce("rokey1", "data21", "read")
	fmt.Printf("rokey1 data21 read %t\n", ok)
	ok = e.Enforce("rokey1", "data31", "write")
	fmt.Printf("rokey1 data31 write %t\n", ok)
	ok = e.Enforce("rokey2", "data11", "write")
	fmt.Printf("rokey2 data11 write %t\n", ok)
	ok = e.Enforce("rokey3", "data21", "read")
	fmt.Printf("rokey3 data21 read %t\n", ok)
	ok = e.Enforce("rokey3", "data21", "write")
	fmt.Printf("rokey3 data21 write %t\n", ok)
	ok = e.Enforce("rokey4", "data11", "read")
	fmt.Printf("rokey4 data11 read %t\n", ok)
	ok = e.Enforce("rokey5", "data21", "write")
	fmt.Printf("rokey5 data21 write %t\n", ok)
}

func TestGetRolesForUser(t *testing.T) {
	e := casbin.Casbin()
	rokeyRoles := e.GetImplicitRolesForUser("rokey")
	fmt.Printf("rokey %+v\n", rokeyRoles)
}

func TestGetPromissionForUser(t *testing.T) {
	e := casbin.Casbin()
	rokeyPromission := e.GetImplicitPermissionsForUser("rokey")
	fmt.Printf("rokey %+v\n", rokeyPromission)
	rokey4Promission := e.GetImplicitPermissionsForUser("rokey4")
	fmt.Printf("rokey4 %+v\n", rokey4Promission)
}

func TestAddRoleForUser(t *testing.T) {
	e := casbin.Casbin()
	ok := e.AddRoleForUser("rokey", "node1_user")
	e.SavePolicy()
	fmt.Printf("add rokey node1_user %t\n", ok)
	TestGetRolesForUser(t)
}

func TestAddPromissionForUser(t *testing.T) {
	e := casbin.Casbin()
	ok := e.AddPermissionForUser("test_admin", "test_data, update|delete")
	e.SavePolicy()
	fmt.Printf("test_admin add permission %+v\n", ok)
}

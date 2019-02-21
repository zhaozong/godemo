package casbin

import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/rbac"
	"github.com/casbin/casbin/rbac/default-role-manager"
)

var rm rbac.RoleManager

type CasbinModel struct {
	Ptype    string `json:"ptype" bson:"ptype"`
	RoleName string `json:"rolename" bson:"v0"`
	Path     string `json:"path" bson:"v1"`
	Method   string `json:"method" bson:"v2"`
}

// 添加权限
func (c *CasbinModel) AddCasbin(cm CasbinModel) bool {
	e := Casbin()
	return e.AddPolicy(cm.RoleName, cm.Path, cm.Method)
}

func Casbin() *casbin.Enforcer {
	//var url = "mongodb://businessdb:6TQbbC9rrGuL3c@10.10.108.86:27017,10.10.108.74:27017,10.10.108.89:27017/businessdb?replicaSet=mymongodb"
	//a := mongodbadapter.NewAdapter(url)
	//e := casbin.NewEnforcer("test-casbin/model.conf", a)
	e := casbin.NewEnforcer("model.conf", "policy.csv")
	rm = defaultrolemanager.NewRoleManager(3)
	e.SetRoleManager(rm)
	e.AddFunction("my_func", RoleInheritFunc)
	e.LoadPolicy()
	return e
}

func RoleInherit(role1 string, role2 string) bool {
	if role1 == role2 {
		return true
	} else {
		if ok,_ :=rm.HasLink(role1, role2); ok {
			return true
		}
		return false
	}
}

func RoleInheritFunc(args ...interface{}) (interface{}, error) {
	role1 := args[0].(string)
	role2 := args[1].(string)
	return (bool)(RoleInherit(role1, role2)), nil
}

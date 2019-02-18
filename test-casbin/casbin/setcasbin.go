package casbin

import (
	"github.com/casbin/casbin"
	"github.com/casbin/mongodb-adapter"
)

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
	var url = "mongodb://businessdb:6TQbbC9rrGuL3c@10.10.108.86:27017,10.10.108.74:27017,10.10.108.89:27017/businessdb?replicaSet=mymongodb"
	a := mongodbadapter.NewAdapter(url)
	e := casbin.NewEnforcer("test-casbin/model.conf", a)
	e.LoadPolicy()
	return e
}


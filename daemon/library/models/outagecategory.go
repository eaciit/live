package models

import (
	//"fmt"
	"github.com/eaciit/orm"
	"strings"
	//"time"
)

type OutageCategory struct {
	orm.ModelBase
	Id       string `bson:"_id"`
	Keywords []string
}

func (o *OutageCategory) New() *OutageCategory {
	o.Keywords = []string{}
	return o
}

func (o *OutageCategory) Scan(s string) bool {
	s = strings.ToLower(s)
	for _, k := range o.Keywords {
		k = strings.Trim(strings.ToLower(k), " ")
		if strings.Contains(s, k) && k != "" {
			return true
		}
	}
	return false
}

func (o *OutageCategory) TableName() string {
	return "OutageCategories"
}

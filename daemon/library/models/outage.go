package models

import (
	//"github.com/eaciit/database/mongodb"
	// "fmt"
	//lq "github.com/eaciit/go-linq"
	"github.com/eaciit/orm"
	"time"
)

type Outage struct {
	orm.ModelBase `bson:"base"`
	PlantName     string
	PlantId       string
	UnitNo        string
	Month         int
	Year          int
	Hours         float64
	Count         int
	Valid         int
	ValidHours    float64
	Summaries     []OutageSummary
}

type OutageSummary struct {
	CategoryId string
	Count      int
	Valid      int
	Hours      float64
	ValidHours float64
	Outages    []OutageItem
}

type OutageItem struct {
	CategoryId          string `bson:"-"`
	DateFrom            time.Time
	DateTo              time.Time
	Hours               float64
	OutageType          string
	Reason              string
	ExpDate             time.Time
	SuggestedCategoryId string
	Valid               bool
}

func (o *Outage) Sync() {
	o.Count = 0
	o.Hours = 0
	o.Valid = 0
	o.ValidHours = 0
	for k, s := range o.Summaries {
		s.Sync()
		o.Summaries[k] = s
		o.Count += s.Count
		o.Hours += s.Hours
		o.ValidHours += s.ValidHours
		o.Valid += s.Valid
	}
}

func (o *OutageSummary) Sync() {
	o.Count = len(o.Outages)
	o.Valid = 0
	o.Hours = 0
	o.ValidHours = 0
	for k, s := range o.Outages {
		o.Hours += s.Hours
		if s.Valid {
			o.Valid += 1
			o.ValidHours += s.Hours
		}
		o.Outages[k] = s
	}
}

func (o *Outage) TableName() string {
	return "Outages"
}

// func (o *Outage) PrepareId() interface{} {
// 	o.Id = fmt.Sprintf("P%sU%s_%d", o.PlantName, o.UnitNo, o.Year*100+o.Month)
// 	return o.Id
// }

func (o *Outage) AddOutage(i OutageItem, syncNow bool) {
	var summary OutageSummary

	newelement := false
	idx := -1

	for k, v := range o.Summaries {
		if v.CategoryId == i.CategoryId {
			idx = k
			summary = v
		}
	}
	if idx == -1 {
		newelement = true
		summary = OutageSummary{}
		summary.CategoryId = i.CategoryId
		summary.Outages = []OutageItem{}
	}

	summary.Outages = append(summary.Outages, i)

	if newelement {
		o.Summaries = append(o.Summaries, summary)
	} else {
		o.Summaries[idx] = summary
	}

	if syncNow {
		o.Sync()
	}
}

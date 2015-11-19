package models

import (
	"fmt"
	"github.com/eaciit/database/base"
	"github.com/eaciit/database/mongodb"
	"github.com/eaciit/orm"
	tk "github.com/eaciit/toolkit"
	"strings"
	"testing"
	//"time"
)

var ctx base.IConnection

func connect() error {
	if ctx == nil {
		ctx = mongodb.NewConnection("localhost:27123", "", "", "ecsec")
	}
	return ctx.Connect()
}

func close() {
	if ctx != nil {
		ctx.Close()
	}
}

func TestLoad(t *testing.T) {
	t.Skip()
	if econnect := connect(); econnect != nil {
		t.Error("Unable to connect to database")
		return
	}
	defer close()

	q := ctx.Query().From("tmpOutages").Select("Start Date", "Start Time", "Total")
	c := q.Cursor(nil)
	defer c.Close()
	var m tk.M
	for b, _ := c.Fetch(&m); b; b, _ = c.Fetch(&m) {
		//go func(m tk.M) {
		dt := tk.MakeDate("1/2/06", m.Get("Start Date", "1/1/1980").(string)).UTC()
		tm := tk.MakeDate("03:04� ", m.Get("Start Time", "00:00").(string)).UTC()
		//tm := tk.MakeDate("03:04� ", m.Get("Start Time", "00:00").(string)).UTC().Sub(tk.MakeDate("03:04", "00:00").UTC())
		dt = tk.AddTime(dt, tm)
		fmt.Printf("Data %v has: %v\n",
			dt.Format("2-Jan-2006 03:04"),
			m.GetFloat64("Total"))
		//}(m)
	}
}

func TestModeler(t *testing.T) {
	t.Skip()
	//var e error

	if econnect := connect(); econnect != nil {
		t.Error("Unable to connect to database")
		return
	}
	defer close()

	ormConn := orm.New(ctx)
	defer ormConn.Close()
	ormConn.DeleteMany(new(Outage), nil)

	cats := make([]OutageCategory, 0)
	ormConn.Find(new(OutageCategory), nil).FetchAll(&cats, true)

	findCat := func(id string) *OutageCategory {
		for _, v := range cats {
			if v.Id == id {
				return &v
			}
		}
		return new(OutageCategory).New()
	}

	q := ctx.Query().From("tmpOutages")
	c := q.Cursor(nil)
	defer c.Close()
	var m tk.M
	for b, _ := c.Fetch(&m); b; b, _ = c.Fetch(&m) {
		//if m.Get("Start Date") == "9/30/12" {
		//go func(m tk.M) {
		dt := tk.MakeDate("1/2/06", m.Get("Start Date", "1/1/1980").(string)).UTC()
		tm := tk.MakeDate("15:04� ", m.Get("Start Time", "00:00").(string)).UTC()
		dt = tk.AddTime(dt, tm)
		//fmt.Println(tm)

		dt2 := tk.MakeDate("1/2/06", m.Get("Finish Date", "1/1/1980").(string)).UTC()
		tm2 := tk.MakeDate("15:04� ", m.Get("Finish Time", "00:00").(string)).UTC()
		dt2 = tk.AddTime(dt2, tm2)

		o := new(Outage)
		o.PlantName = m.Get("Plant Name").(string)

		if strings.Contains(o.PlantName, "Ghazlan") {
			o.PlantName = "Ghazlan"
		} else if strings.Contains(o.PlantName, "Rabigh") {
			o.PlantName = "Rabigh"
		}

		o.UnitNo = m.Get("Unit No").(string)
		o.Year = dt.Year()
		o.Month = int(dt.Month())
		o.PrepareId()

		o.Summaries = []OutageSummary{}

		b, e := ormConn.GetById(o, o.Id)
		if b {
			fmt.Printf("Data already exist for %s\n", o.Id)
		} else {
			fmt.Printf("Data is not exist for %s\n", o.Id)
		}

		i := OutageItem{}
		i.CategoryId = m.Get("Outage Type").(string)
		if strings.HasPrefix(i.CategoryId, "FO") {
			i.CategoryId = "FO"
		} else if i.CategoryId == "UC-" {
			i.CategoryId = "UC"
		}
		i.DateFrom = dt
		i.DateTo = dt2
		i.Reason = m.Get("Outage Reason").(string)
		i.Hours = m.GetFloat64("Total")

		cat := findCat(i.CategoryId)
		if cat != nil {
			i.Valid = cat.Scan(i.Reason)
		}
		if i.Valid == false {
			found := false
			for _, cv := range cats {
				found = cv.Scan(i.Reason)
				if found {
					i.SuggestedCategoryId = cv.Id
					break
				}
			}
		}

		o.AddOutage(i, true)

		if e = ormConn.Save(o); e != nil {
			t.Errorf("Unable to save %s => %s", o.Id, e.Error())
		}
		//}(m)
		//}
	}
}

func TestScan(t *testing.T) {
	var e error

	fmt.Println("Test Scan")
	if econnect := connect(); econnect != nil {
		t.Error("Unable to connect to database")
		return
	}
	defer close()

	ormConn := orm.New(ctx)
	defer ormConn.Close()

	cats := make([]OutageCategory, 0)
	e = ormConn.Find(new(OutageCategory), nil).FetchAll(&cats, true)
	if e != nil {
		t.Error(e.Error())
	}

	os := make([]*Outage, 0)
	e = ormConn.Find(new(Outage), nil).FetchAll(&os, true)
	if e != nil {
		t.Error(e.Error())
	}

	for _, o := range os {
		fmt.Printf("Before: %s \n\n", tk.JsonString(o))
		for k, sum := range o.Summaries {
			for k, i := range sum.Outages {
				i.Valid = false
				i.SuggestedCategoryId = ""
				for _, c := range cats {
					fmt.Printf("Evaluating keyword: %s ...", strings.Join(c.Keywords, ","))
					cok := c.Scan(i.Reason)
					fmt.Printf("%+v \n", cok)
					if cok {
						if i.Valid == false && sum.CategoryId == c.Id {
							fmt.Printf("Tag: %s as valid \n", c.Id)
							i.Valid = true
							i.SuggestedCategoryId = ""
						} else if i.Valid == false {
							fmt.Printf("Suggesting: %s \n", c.Id)
							i.SuggestedCategoryId = c.Id
						}
					}
				}
				sum.Outages[k] = i
			}
			o.Summaries[k] = sum
		}
		o.Sync()
		e = ormConn.Save(o)
		fmt.Printf("After: %s \n\n", tk.JsonString(o))
	}
}

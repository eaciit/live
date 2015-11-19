package main

import (
	. "eaciit/secsampeldata/Datahelper"
	// "fmt"
	// . "github.com/ahmetalpbalkan/go-linq"
	"github.com/metakeule/fmtdate"
	// "github.com/tealeg/xlsx"
	// "gopkg.in/mgo.v2/bson"
	//"bytes"
	//"encoding/csv"
	// "io/ioutil"
	// "math"
	//"os"
	"math/rand"
	"strconv"
	// "strings"
	"time"
)

func main() {
	// sampelAssetClass()
	// sampelAssetLevel()
	// SampleLocation()
	// SampleAssetType()
	// SamplePlant()
	sampelAsset()
	// sampleAssetPlacements()
	// sampleAssetPerformance()
	// sampleAssetMaintenance()
	// sampleAssetFinancial()
	// sampleAssetMaintenanceSchedule()
}

type ModelAssetClass struct {
	Code      string
	ClassName string
}

type ModelAssetLevel struct {
	Code      string
	LevelName string
}

type ModelSampleLocation struct {
	Country   string
	City      string
	Longitude float64
	Latitude  float64
}

type ModelSampleAssetType struct {
	Code     string
	TypeName string
}

type ModelSamplePlant struct {
	Name     string
	Code     string
	Location ModelSampleLocation
}

func sampelAssetClass() {
	//fmt.Println("Coba")
	for i := 1; i <= 5; i++ {
		var model ModelAssetClass
		model.Code = "Class" + strconv.Itoa(i)
		model.ClassName = "Class " + strconv.Itoa(i)
		Save("SampleAssetClass", model)
	}
}

func sampelAssetLevel() {
	for i := 1; i <= 5; i++ {
		var model ModelAssetLevel
		model.Code = "Level" + strconv.Itoa(i)
		model.LevelName = "Level  " + strconv.Itoa(i)
		Save("SampleAssetLevel", model)
	}
}

func SampleLocation() {
	for i := 1; i <= 5; i++ {
		var model ModelSampleLocation
		model.Country = "Country" + strconv.Itoa(i)
		model.City = "City" + strconv.Itoa(i)
		model.Longitude = 0
		model.Latitude = 0
		Save("SampleLocation", model)
	}
}

func SampleAssetType() {
	for i := 1; i <= 5; i++ {
		var model ModelSampleAssetType
		model.Code = "Type" + strconv.Itoa(i)
		model.TypeName = "Type " + strconv.Itoa(i)
		Save("SampleAssetType", model)
	}
}

func SamplePlant() {
	var ResultLoc []ModelSampleLocation
	PopulateAsObject(&ResultLoc, "SampleLocation", nil, 0, 0)
	for i := 1; i <= 5; i++ {
		var model ModelSamplePlant
		model.Name = "Plant " + strconv.Itoa(i)
		model.Code = "Plant" + strconv.Itoa(i)
		model.Location = ResultLoc[rand.Intn(len(ResultLoc))]
		Save("SamplePlant", model)
	}
}

type ModelAsset struct {
	Code              string
	Name              string
	Manufacturer      string
	ManufacturedYear  int
	Class             ModelAssetClass
	Type              ModelSampleAssetType
	Level             ModelAssetLevel
	NormalCapacity    int
	Notes             string
	PurchaseDate      time.Time
	PurchaseCost      float64
	PurchaseCondition string
	PurchaseVendor    string
	PurchaseDiscount  int
	EstimatedUsage    int
	AmortizationRate  int
	DepresiationRate  int
	LatestStatus      string
}

func sampelAsset() {
	var (
		ResultClass []ModelAssetClass
		ResultType  []ModelSampleAssetType
		ResultLevel []ModelAssetLevel
		dateRand    string
	)
	ManufacturedYearVal := [2]int{2015, 2014}
	CapacityVal := [8]int{15, 17, 10, 15, 22, 12, 14, 20}
	Monthstring := [12]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	CostVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	ConditionVal := [2]string{"New", "Used"}
	ManufacturerVal := [3]string{"eaciit", "pertamina", "telkom"}
	DiskonVal := [4]int{15, 20, 25, 50}
	StatusVal := [4]string{"Used", "Broken", "Available", "Other"}
	PopulateAsObject(&ResultClass, "SampleAssetClass", nil, 0, 0)
	PopulateAsObject(&ResultType, "SampleAssetType", nil, 0, 0)
	PopulateAsObject(&ResultLevel, "SampleAssetLevel", nil, 0, 0)
	for i := 1; i <= 100; i++ {
		var (
			model ModelAsset
		)
		model.Code = "Asset" + strconv.Itoa(i)
		model.Name = "Asset " + strconv.Itoa(i)
		model.Manufacturer = ManufacturerVal[rand.Intn(len(ManufacturerVal))]
		model.ManufacturedYear = ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]
		model.Class = ResultClass[rand.Intn(len(ResultClass))]
		model.Type = ResultType[rand.Intn(len(ResultType))]
		model.Level = ResultLevel[rand.Intn(len(ResultLevel))]
		model.NormalCapacity = CapacityVal[rand.Intn(len(CapacityVal))]
		model.Notes = "Note " + strconv.Itoa(i)
		dateRand = strconv.Itoa(ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]) + "/" + Monthstring[rand.Intn(len(Monthstring))] + "/" + strconv.Itoa(CapacityVal[rand.Intn(len(CapacityVal))]) + " 00:00:00"
		model.PurchaseDate, _ = fmtdate.Parse("YYYY/MM/DD hh:mm:ss", dateRand)
		model.PurchaseCost = CostVal[rand.Intn(len(CostVal))]
		model.PurchaseCondition = ConditionVal[rand.Intn(len(ConditionVal))]
		model.PurchaseVendor = "eaciit"
		model.PurchaseDiscount = DiskonVal[rand.Intn(len(DiskonVal))]
		// dateRand = strconv.Itoa(ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]) + "/" + Monthstring[rand.Intn(len(Monthstring))] + "/" + strconv.Itoa(CapacityVal[rand.Intn(len(CapacityVal))]) + " 00:00:00"
		model.EstimatedUsage = rand.Intn(12 + 1)
		if model.EstimatedUsage == 0 {
			model.EstimatedUsage = 1
		}
		model.AmortizationRate = DiskonVal[rand.Intn(len(DiskonVal))]
		model.DepresiationRate = DiskonVal[rand.Intn(len(DiskonVal))]
		model.LatestStatus = StatusVal[rand.Intn(len(StatusVal))]
		// fmt.Println(model.PurchaseDate)
		Save("SampleAsset", model)
	}
}

type ModelAssetPlacements struct {
	Plants             ModelSamplePlant
	Assets             ModelAsset
	PlacedDate         time.Time
	PlacedCondition    string
	ReturnDate         time.Time
	ReturnCondition    string
	LatestAvailability int
}

func sampleAssetPlacements() {
	var (
		ResultPlant []ModelSamplePlant
		ResultAsset []ModelAsset
		dateRand    string
	)
	ManufacturedYearVal := [2]int{2015, 2014}
	Monthstring := [12]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	CapacityVal := [8]int{15, 17, 10, 15, 22, 12, 14, 20}
	ConditionVal := [2]string{"New", "Used"}
	PopulateAsObject(&ResultPlant, "SamplePlant", nil, 0, 0)
	PopulateAsObject(&ResultAsset, "SampleAsset", nil, 0, 0)
	for i := 1; i <= 100; i++ {
		var model ModelAssetPlacements
		model.Plants = ResultPlant[rand.Intn(len(ResultPlant))]
		model.Assets = ResultAsset[rand.Intn(len(ResultAsset))]
		dateRand = strconv.Itoa(ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]) + "/" + Monthstring[rand.Intn(len(Monthstring))] + "/" + strconv.Itoa(CapacityVal[rand.Intn(len(CapacityVal))]) + " 00:00:00"
		model.PlacedDate, _ = fmtdate.Parse("YYYY/MM/DD hh:mm:ss", dateRand)
		model.PlacedCondition = ConditionVal[rand.Intn(len(ConditionVal))]
		dateRand = strconv.Itoa(ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]) + "/" + Monthstring[rand.Intn(len(Monthstring))] + "/" + strconv.Itoa(CapacityVal[rand.Intn(len(CapacityVal))]) + " 00:00:00"
		model.ReturnDate, _ = fmtdate.Parse("YYYY/MM/DD hh:mm:ss", dateRand)
		model.ReturnCondition = ConditionVal[rand.Intn(len(ConditionVal))]
		model.LatestAvailability = CapacityVal[rand.Intn(len(CapacityVal))]
		Save("SampleAssetPlacements", model)
	}
}

type ModelAssetPerformance struct {
	Assets           ModelAsset
	StartsTime       time.Time
	StopTime         time.Time
	Duration         int
	Availability     int
	UtilizedDuration int
	UtilizedPower    int
}

func sampleAssetPerformance() {
	var (
		ResultAsset []ModelAsset
		dateRand    string
	)
	ManufacturedYearVal := [2]int{2015, 2014}
	Monthstring := [12]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	CapacityVal := [8]int{15, 17, 10, 15, 22, 12, 14, 20}
	hourRand := [5]time.Duration{3, 5, 6, 9, 8}
	PopulateAsObject(&ResultAsset, "SampleAsset", nil, 0, 0)
	for i := 1; i <= 100; i++ {
		var (
			model     ModelAssetPerformance
			startrand time.Time
			hourVal   time.Duration
		)
		model.Assets = ResultAsset[rand.Intn(len(ResultAsset))]
		dateRand = strconv.Itoa(ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]) + "/" + Monthstring[rand.Intn(len(Monthstring))] + "/" + strconv.Itoa(CapacityVal[rand.Intn(len(CapacityVal))]) + " 00:00:00"
		// model.StartsTime, _ = fmtdate.Parse("YYYY/MM/DD hh:mm:ss", dateRand)
		startrand, _ = fmtdate.Parse("YYYY/MM/DD hh:mm:ss", dateRand)
		// dateRand = strconv.Itoa(ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]) + "/" + Monthstring[rand.Intn(len(Monthstring))] + "/" + strconv.Itoa(CapacityVal[rand.Intn(len(CapacityVal))]+3) + " 00:00:00"
		hourVal = hourRand[rand.Intn(len(hourRand))]
		model.StartsTime = startrand.Add(hourVal * time.Hour)
		hourVal = hourRand[rand.Intn(len(hourRand))]
		model.StopTime = model.StartsTime.Add(hourVal * time.Hour)
		model.Duration = model.StopTime.Hour() - model.StartsTime.Hour()
		// fmt.Println(rand.Intn(5))
		model.Availability = CapacityVal[rand.Intn(len(CapacityVal))]
		model.UtilizedDuration = rand.Intn(model.Duration + 1)
		if model.UtilizedDuration == 0 {
			model.UtilizedDuration = 1
		}
		model.UtilizedPower = CapacityVal[rand.Intn(len(CapacityVal))]
		Save("SampleAssetPerformance", model)
	}
}

type ModelAssetMaintenance struct {
	Assets              ModelAsset
	StartsTime          time.Time
	StopTime            time.Time
	Duration            int
	MaintenanceStatus   string
	MaintenanceBy       string
	MaintenanceGroup    string
	ChangedParts        string
	CausedBy            string
	ConditionBefore     string
	ConditionAfter      string
	ConditionPercentage int
	Notes               string
}

func sampleAssetMaintenance() {
	var (
		ResultAsset []ModelAsset
		dateRand    string
	)
	ManufacturedYearVal := [2]int{2015, 2014}
	Monthstring := [12]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	CapacityVal := [8]int{15, 17, 10, 15, 22, 12, 14, 20}
	StatusVal := [2]string{"Scheduled", "UnPlanned"}
	PersenVal := [7]int{15, 20, 25, 50, 70, 75, 80}
	hourRand := [5]time.Duration{3, 5, 6, 9, 8}
	GroupVal := [6]string{"Failure", "Renewal Overhauls", "Risk Limited", "Operating Time", "Condition Based", "Others"}
	PopulateAsObject(&ResultAsset, "SampleAsset", nil, 0, 0)
	for i := 1; i <= 100; i++ {
		var (
			model     ModelAssetMaintenance
			startrand time.Time
			hourVal   time.Duration
		)
		model.Assets = ResultAsset[rand.Intn(len(ResultAsset))]
		dateRand = strconv.Itoa(ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]) + "/" + Monthstring[rand.Intn(len(Monthstring))] + "/" + strconv.Itoa(CapacityVal[rand.Intn(len(CapacityVal))]) + " 00:00:00"
		// model.StartsTime, _ = fmtdate.Parse("YYYY/MM/DD hh:mm:ss", dateRand)
		startrand, _ = fmtdate.Parse("YYYY/MM/DD hh:mm:ss", dateRand)
		// dateRand = strconv.Itoa(ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]) + "/" + Monthstring[rand.Intn(len(Monthstring))] + "/" + strconv.Itoa(CapacityVal[rand.Intn(len(CapacityVal))]+3) + " 00:00:00"

		hourVal = hourRand[rand.Intn(len(hourRand))]
		model.StartsTime = startrand.Add(hourVal * time.Hour)
		hourVal = hourRand[rand.Intn(len(hourRand))]
		model.StopTime = model.StartsTime.Add(hourVal * time.Hour)
		model.Duration = model.StopTime.Hour() - model.StartsTime.Hour()

		model.MaintenanceStatus = StatusVal[rand.Intn(len(StatusVal))]
		model.MaintenanceBy = "eaciit"
		model.MaintenanceGroup = GroupVal[rand.Intn(len(GroupVal))]
		model.ChangedParts = "Parts " + strconv.Itoa(i)
		model.CausedBy = "Broken"
		model.ConditionBefore = "Good"
		model.ConditionAfter = "Damaged"
		model.ConditionPercentage = PersenVal[rand.Intn(len(PersenVal))]
		model.Notes = "Note " + strconv.Itoa(i)
		Save("SampleAssetMaintenance", model)
	}
}

type ModelAssetFinancial struct {
	Assets            ModelAsset
	Period            time.Time
	Revenues          float64
	OperationalCost   float64
	MaintenanceCost   float64
	InsuranceCost     float64
	OtherCost         float64
	SalvageValue      float64
	AcquisitionCost   float64
	SustainingCapital float64
}

func sampleAssetFinancial() {
	var (
		ResultAsset []ModelAsset
		dateRand    string
		Monthstring [6]string
		YearVal     int
	)
	ManufacturedYearVal := [2]int{2015, 2014}

	CapacityVal := [8]int{15, 17, 10, 15, 22, 12, 14, 20}
	RevenuesVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	OptCostVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	MaintCostVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	InsCostVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	OthCostVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	SalvageVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	AcqCostVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	SustCaptVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	PopulateAsObject(&ResultAsset, "SampleAsset", nil, 0, 0)
	for i := 1; i <= 100; i++ {
		var model ModelAssetFinancial
		model.Assets = ResultAsset[rand.Intn(len(ResultAsset))]

		YearVal = ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]
		if YearVal == 2014 {
			Monthstring = [6]string{"07", "08", "09", "10", "11", "12"}
		} else {
			Monthstring = [6]string{"01", "02", "03", "04", "05", "06"}
		}
		dateRand = strconv.Itoa(ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]) + "/" + Monthstring[rand.Intn(len(Monthstring))] + "/" + strconv.Itoa(CapacityVal[rand.Intn(len(CapacityVal))]) + " 00:00:00"
		model.Period, _ = fmtdate.Parse("YYYY/MM/DD hh:mm:ss", dateRand)
		model.Revenues = RevenuesVal[rand.Intn(len(RevenuesVal))]
		model.OperationalCost = OptCostVal[rand.Intn(len(OptCostVal))]
		model.MaintenanceCost = MaintCostVal[rand.Intn(len(MaintCostVal))]
		model.InsuranceCost = InsCostVal[rand.Intn(len(InsCostVal))]
		model.OtherCost = OthCostVal[rand.Intn(len(OthCostVal))]
		model.SalvageValue = SalvageVal[rand.Intn(len(SalvageVal))]
		model.AcquisitionCost = AcqCostVal[rand.Intn(len(AcqCostVal))]
		model.SustainingCapital = SustCaptVal[rand.Intn(len(SustCaptVal))]
		Save("SampleAssetFinancial", model)

	}
}

type ModelAssetMaintenanceSchedule struct {
	Assets                     ModelAsset
	ScheduledDate              time.Time
	EstimatedMaintenanceDetail float64
	EstimatedCost              float64
}

func sampleAssetMaintenanceSchedule() {
	var (
		ResultAsset []ModelAsset
		dateRand    string
	)
	ManufacturedYearVal := [2]int{2015, 2014}
	Monthstring := [12]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	CapacityVal := [8]int{15, 17, 10, 15, 22, 12, 14, 20}
	EstimatedDetailVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	EstimatedCostVal := [8]float64{1500, 1700, 1000, 1500, 2200, 1200, 1400, 2000}
	PopulateAsObject(&ResultAsset, "SampleAsset", nil, 0, 0)

	for i := 1; i <= 100; i++ {
		var model ModelAssetMaintenanceSchedule
		model.Assets = ResultAsset[rand.Intn(len(ResultAsset))]
		dateRand = strconv.Itoa(ManufacturedYearVal[rand.Intn(len(ManufacturedYearVal))]) + "/" + Monthstring[rand.Intn(len(Monthstring))] + "/" + strconv.Itoa(CapacityVal[rand.Intn(len(CapacityVal))]) + " 00:00:00"
		model.ScheduledDate, _ = fmtdate.Parse("YYYY/MM/DD hh:mm:ss", dateRand)
		model.EstimatedMaintenanceDetail = EstimatedDetailVal[rand.Intn(len(EstimatedDetailVal))]
		model.EstimatedCost = EstimatedCostVal[rand.Intn(len(EstimatedCostVal))]
		Save("SampleAssetMaintenanceSchedule", model)
	}
}

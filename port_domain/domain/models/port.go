package models

import (
	"gorm.io/gorm"
	"strings"
)

const sep = "#,#"

type Port struct {
	Id           int    `gorm:"column:id"`
	Code         string `gorm:"not null; unique"`
	Key          string `gorm:"not null"`
	Name         string
	City         string
	Country      string
	Province     string
	Timezone     string
	Alias        string   `json:"-"`
	AliasSlice   []string `gorm:"-" json:"alias"`
	Regions      string   `json:"-"`
	RegionsSlice []string `gorm:"-" json:"regions"`
	Unlocs       string   `json:"-"`
	UnlocsSlice  []string `gorm:"-" json:"unlocs"`
	Coordinates  `gorm:"embedded"`
	Timestamp    `gorm:"embedded"`
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func (m *Port) BeforeSave(tx *gorm.DB) (err error) {
	m.Alias = strings.Join(m.AliasSlice, sep)
	m.Regions = strings.Join(m.RegionsSlice, sep)
	m.Unlocs = strings.Join(m.UnlocsSlice, sep)
	return
}

func (m *Port) AfterFind(tx *gorm.DB) (err error) {
	m.AliasSlice = strings.Split(m.Alias, sep)
	m.RegionsSlice = strings.Split(m.Regions, sep)
	m.UnlocsSlice = strings.Split(m.Unlocs, sep)
	return
}

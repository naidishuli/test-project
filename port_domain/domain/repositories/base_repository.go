package repositories

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository interface {
	Create(value interface{}) *gorm.DB
	CreateInBatches(value interface{}, batchSize int) *gorm.DB
	First(dest interface{}, cond ...interface{}) *gorm.DB
	FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, cond ...interface{}) *gorm.DB
	FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB
	Pluck(column string, dest interface{}) *gorm.DB
	Count(*int64) *gorm.DB
	Save(value interface{}) *gorm.DB
	Update(column string, value interface{}) *gorm.DB
	Updates(values interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Scan(dest interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) BaseRepository
	Or(query interface{}, args ...interface{}) BaseRepository
	Order(value interface{}) BaseRepository
	Clauses(conds ...clause.Expression) BaseRepository
	Select(query interface{}, args ...interface{}) BaseRepository
	Omit(columns ...string) BaseRepository
	Limit(int) BaseRepository
	Group(name string) BaseRepository
	Distinct(args ...interface{}) BaseRepository
	Preload(query string, args ...interface{}) BaseRepository
	Joins(query string, args ...interface{}) BaseRepository
	Attrs(attrs ...interface{}) BaseRepository
	Assign(attrs ...interface{}) BaseRepository
	Session(s *gorm.Session) BaseRepository
	Model(value interface{}) BaseRepository
	Transaction(func(tx *gorm.DB) error, ...*sql.TxOptions) error
	New() BaseRepository
	Fresh() BaseRepository
	Table(t string) BaseRepository
	Debug() BaseRepository
}

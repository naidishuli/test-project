package postgres

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	. "port_domain/domain/repositories"
)

type baseRepo struct {
	*gorm.DB
}

func NewBaseRepo(db *gorm.DB) *baseRepo {
	return &baseRepo{db}
}

func (b *baseRepo) Create(value interface{}) *gorm.DB {
	cDb := b.DB.Create(value)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	cDb := b.DB.CreateInBatches(value, batchSize)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) First(dest interface{}, cond ...interface{}) *gorm.DB {
	cDb := b.DB.First(dest, cond...)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) FirstOrCreate(dest interface{}, cond ...interface{}) *gorm.DB {
	cDb := b.DB.FirstOrCreate(dest, cond...)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) Find(dest interface{}, cond ...interface{}) *gorm.DB {
	cDb := b.DB.Find(dest, cond...)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	cDb := b.DB.FindInBatches(dest, batchSize, fc)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) Pluck(column string, dest interface{}) *gorm.DB {
	cDb := b.DB.Pluck(column, dest)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) Count(int *int64) *gorm.DB {
	cDb := b.DB.Count(int)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) Save(value interface{}) *gorm.DB {
	cDb := b.DB.Save(value)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) Update(column string, value interface{}) *gorm.DB {
	cDb := b.DB.Update(column, value)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) Updates(value interface{}) *gorm.DB {
	cDb := b.DB.Updates(value)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	cDb := b.DB.Delete(value, conds...)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) Scan(dest interface{}) *gorm.DB {
	cDb := b.DB.Scan(dest)
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return cDb
}

func (b *baseRepo) Where(query interface{}, args ...interface{}) BaseRepository {
	b.DB = b.DB.Where(query, args...)
	return b
}

func (b *baseRepo) Or(query interface{}, args ...interface{}) BaseRepository {
	b.DB = b.DB.Or(query, args...)
	return b
}

func (b *baseRepo) Order(value interface{}) BaseRepository {
	b.DB = b.DB.Order(value)
	return b
}

func (b *baseRepo) Clauses(conds ...clause.Expression) BaseRepository {
	b.DB = b.DB.Clauses(conds...)
	return b
}

func (b *baseRepo) Select(query interface{}, args ...interface{}) BaseRepository {
	b.DB = b.DB.Select(query, args...)
	return b
}

func (b *baseRepo) Omit(columns ...string) BaseRepository {
	b.DB = b.DB.Omit(columns...)
	return b
}

func (b *baseRepo) Limit(i int) BaseRepository {
	b.DB = b.DB.Limit(i)
	return b
}

func (b *baseRepo) Group(name string) BaseRepository {
	b.DB = b.DB.Group(name)
	return b
}

func (b *baseRepo) Distinct(args ...interface{}) BaseRepository {
	b.DB = b.DB.Distinct(args...)
	return b
}

func (b *baseRepo) Preload(query string, args ...interface{}) BaseRepository {
	b.DB = b.DB.Preload(query, args...)
	return b
}

func (b *baseRepo) Joins(query string, args ...interface{}) BaseRepository {
	b.DB = b.DB.Joins(query, args...)
	return b
}

func (b *baseRepo) Attrs(attrs ...interface{}) BaseRepository {
	b.DB = b.DB.Attrs(attrs...)
	return b
}

func (b *baseRepo) Assign(attrs ...interface{}) BaseRepository {
	b.DB = b.DB.Assign(attrs...)
	return b
}

func (b *baseRepo) Session(s *gorm.Session) BaseRepository {
	b.DB = b.DB.Session(s)
	return b
}

func (b *baseRepo) Model(value interface{}) BaseRepository {
	b.DB = b.DB.Model(value)
	return b
}

func (b *baseRepo) New() BaseRepository {
	return NewBaseRepo(b.DB.Session(&gorm.Session{NewDB: true}))
}

func (b *baseRepo) Fresh() BaseRepository {
	b.DB = b.DB.Session(&gorm.Session{NewDB: true})
	return b
}

func (b *baseRepo) Table(t string) BaseRepository {
	b.DB = b.DB.Table(t)
	return b
}

func (b *baseRepo) Debug() BaseRepository {
	b.DB = b.DB.Debug()
	return b
}

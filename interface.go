package xormx

import (
	"context"
	"database/sql"
	"io"
	"xorm.io/xorm"
	"xorm.io/xorm/core"
)

type XormSQL interface {
	After(func(any)) *xorm.Session
	Alias(string) *xorm.Session
	AllCols() *xorm.Session
	Asc(...string) *xorm.Session
	Before(func(any)) *xorm.Session
	BufferSize(int) *xorm.Session
	Cascade(...bool) *xorm.Session
	Charset(string) *xorm.Session
	Close() error
	Cols(...string) *xorm.Session
	Context(context.Context) *xorm.Session
	Count(...any) (int64, error)
	CreateIndexes(any) error
	CreateUniques(any) error
	DB() *core.DB
	Decr(string, ...any) *xorm.Session
	Delete(...any) (int64, error)
	Desc(...string) *xorm.Session
	Distinct(...string) *xorm.Session
	DropIndexes(any) error
	Exec(...any) (sql.Result, error)
	Exist(...any) (bool, error)
	Find(any, ...any) error
	FindAndCount(any, ...any) (int64, error)
	Get(...any) (bool, error)
	GroupBy(string) *xorm.Session
	Having(string) *xorm.Session
	ID(any) *xorm.Session
	Import(io.Reader) ([]sql.Result, error)
	ImportFile(string) ([]sql.Result, error)
	In(string, ...any) *xorm.Session
	Incr(string, ...any) *xorm.Session
	Insert(...any) (int64, error)
	InsertOne(any) (int64, error)
	IsTableEmpty(any) (bool, error)
	IsTableExist(any) (bool, error)
	Iterate(any, xorm.IterFunc) error
	Join(string, any, string, ...any) *xorm.Session
	Limit(int, ...int) *xorm.Session
	MustCols(...string) *xorm.Session
	NoAutoCondition(...bool) *xorm.Session
	NoAutoTime() *xorm.Session
	NoCache() *xorm.Session
	NoCascade() *xorm.Session
	NotIn(string, ...any) *xorm.Session
	Nullable(...string) *xorm.Session
	Omit(...string) *xorm.Session
	OrderBy(any, ...any) *xorm.Session
	Ping() error
	PingContext(context.Context) error
	Prepare() *xorm.Session
	Query(...any) ([]map[string][]uint8, error)
	QueryInterface(...any) ([]map[string]any, error)
	QueryString(...any) ([]map[string]string, error)
	Rows(any) (*xorm.Rows, error)
	SQL(any, ...any) *xorm.Session
	Select(string) *xorm.Session
	SetExpr(string, any) *xorm.Session
	StoreEngine(string) *xorm.Session
	Sum(any, string) (float64, error)
	SumInt(any, string) (int64, error)
	Sums(any, ...string) ([]float64, error)
	SumsInt(any, ...string) ([]int64, error)
	Sync(...any) error
	Sync2(...any) error
	Table(any) *xorm.Session
	Unscoped() *xorm.Session
	Update(any, ...any) (int64, error)
	UseBool(...string) *xorm.Session
	Where(any, ...any) *xorm.Session
}

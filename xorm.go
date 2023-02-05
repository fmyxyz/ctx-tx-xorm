package xormx

import (
	"context"
	"database/sql"

	"github.com/fmyxyz/ctx-tx"

	"xorm.io/xorm"
)

type XormEngine struct {
	*xorm.Engine

	instance string
}

type XormSession struct {
	*xorm.Session
}

func (g *XormSession) SavePoint(name string) error {
	_, err := g.Session.Exec("SAVEPOINT " + name)
	return err
}

func (g *XormSession) RollbackTo(name string) error {
	_, err := g.Session.Exec("ROLLBACK TO SAVEPOINT " + name)
	return err
}

func (g *XormSession) Commit() error {
	return g.Session.Commit()
}

func (g *XormSession) Rollback() error {
	return g.Session.Rollback()
}

func (g *XormEngine) Name() string {
	return "xorm-" + g.instance
}

func (g *XormEngine) BeginTx(ctx context.Context, opts *sql.TxOptions) (tx.Tx, error) {
	session := g.Engine.NewSession().Context(ctx)
	err := session.Begin()
	return warp(session), err
}

func warp(db *xorm.Session) *XormSession {
	return &XormSession{Session: db}
}

const defaultInstance = "default"

func Register(db *xorm.Engine, opts ...XormDBOption) {
	xormDB := &XormEngine{Engine: db, instance: defaultInstance}
	for _, opt := range opts {
		opt(xormDB)
	}
	tx.Register(xormDB, tx.RegisterDefaultDB(xormDB.instance == defaultInstance))
}

type XormDBOption func(db *XormEngine)

func Instance(instance string) XormDBOption {
	return func(db *XormEngine) {
		db.instance = instance
	}
}

func FromContext(ctx context.Context, opts ...XormDBOption) XormSQL {
	gormDB := &XormEngine{instance: defaultInstance}
	for _, opt := range opts {
		opt(gormDB)
	}
	name := gormDB.Name()
	txManager := tx.GetTxManager(name)
	if txManager == nil {
		panic(name + " not register in txManagers")
	}
	tx0 := txManager.TxFromContext(ctx)
	if tx0 != nil {
		return tx0.(*XormSession)
	}
	db := txManager.DBFromContext(ctx)
	if db != nil {
		return db.(*XormEngine)
	}
	return nil
}

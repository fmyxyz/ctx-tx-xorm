package xormx

import (
	"context"
	"log"
	"testing"

	"github.com/fmyxyz/ctx-tx/test"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func TestWithTx(t *testing.T) {

	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	//db = db.Debug()
	Register(db)

	//resetAll(db)
	test.Update88 = update88
	test.Update99 = update99
	test.Update = update
	for _, tt := range test.Tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := tt.Fc(context.Background())
			eq88 := getValId(db, 88) == 88
			eq99 := getValId(db, 99) == 99
			resetAll(db)
			if eq88 != tt.Eq88 || eq99 != tt.Eq99 || (err != nil) != tt.WantErr {
				t.Errorf("WithTx() {eq88 = %v, want %v} {eq99 = %v, want %v} {error = %v, want has error %v}", eq88, tt.Eq88, eq99, tt.Eq99, err, tt.WantErr)
			}
			if err != nil {
				t.Log(err)
			}
		})
	}

	//resetAll(db)

}

type Item struct {
	ID  int `xorm:"id"`
	Qty int `xorm:"qty"`
}

func openDB() (db *xorm.Engine, err error) {

	db, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		return db, err
	}

	return db, nil
}

func getValId(db *xorm.Engine, id int) (val int) {
	dest := &Item{
		ID: id,
	}
	_, err := db.Table("item").Where("id=?", id).Get(dest)
	if err != nil {
		return 0
	}
	return dest.Qty
}

func reset(db *xorm.Engine, id int) error {
	dest := &Item{
		ID:  id,
		Qty: id + 1,
	}
	_, err := db.Table("item").Where("id=?", id).Update(dest)
	if err != nil {
		return err
	}
	return nil
}

func resetAll(db *xorm.Engine) error {
	log.Println("resetAll")
	reset(db, 88)
	reset(db, 99)
	return nil
}

func update(ctx context.Context, id, num int) error {
	db := FromContext(ctx)
	dest := &Item{
		ID:  id,
		Qty: num,
	}
	_, err := db.Table("item").Where("id=?", dest.ID).Update(dest)
	if err != nil {
		return err
	}
	return nil
}

func update88(ctx context.Context) error {
	db := FromContext(ctx)
	dest := &Item{
		ID:  88,
		Qty: 88,
	}
	_, err := db.Table("item").Where("id=?", dest.ID).Update(dest)
	if err != nil {
		return err
	}
	return nil
}

func update99(ctx context.Context) error {
	db := FromContext(ctx)
	dest := &Item{
		ID:  99,
		Qty: 99,
	}
	_, err := db.Table("item").Where("id=?", dest.ID).Update(dest)
	if err != nil {
		return err
	}
	return nil
}

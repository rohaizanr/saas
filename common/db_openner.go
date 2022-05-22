package common

import (
	"database/sql"
	"fmt"
	"sync"
)

type (
	// DbOpener open a real *sql.Driver instance by using connection string
	DbOpener interface {
		Open(s string) (*sql.DB, error)
	}

	dbOpener struct {
		mtx sync.Mutex
		db  map[string]*sql.DB
		d   DbOpener
	}

	DbOpenerFunc func(s string) (*sql.DB, error)
)

func (f DbOpenerFunc) Open(s string) (*sql.DB, error) {
	return f(s)
}

var _ DbOpener = (*dbOpener)(nil)

//NewCachedDbOpener wrap DbOpener with cache
func NewCachedDbOpener(d DbOpener) (DbOpener, func()) {
	ret := &dbOpener{
		db: make(map[string]*sql.DB),
		d:  d,
	}
	return ret, func() {
		ret.Close()
	}
}

func (d *dbOpener) Open(s string) (*sql.DB, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	db, ok := d.db[s]
	if ok {
		return db, nil
	}
	var err error
	db, err = d.d.Open(s)
	if err != nil {
		return nil, err
	}
	d.db[s] = db
	return db, nil
}

func (d *dbOpener) Close() error {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	var err error
	for _, d := range d.db {
		nerr := d.Close()
		if nerr != nil {
			if err == nil {
				err = nerr
			} else {
				err = fmt.Errorf("%w; ", err)
			}
		}
	}
	return err
}
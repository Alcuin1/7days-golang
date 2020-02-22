package geeorm

import (
	"database/sql"
	"log"
	"os"
)

var (
	ErrorLog = log.New(os.Stdout, "[error] ", log.LstdFlags|log.Lshortfile)
	InfoLog  = log.New(os.Stdout, "[info ] ", log.LstdFlags|log.Lshortfile)
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		ErrorLog.Println(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		ErrorLog.Println(err)
		return
	}
	e = &Engine{db: db}
	InfoLog.Println("Connect database success")
	return
}

func (engine *Engine) Close() (err error) {
	if err = engine.db.Close(); err == nil {
		InfoLog.Println("Close database success")
	}
	return
}

func (engine *Engine) NewSession() *Session {
	return &Session{engine: engine}
}

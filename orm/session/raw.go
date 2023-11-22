package session

import (
	"database/sql"
	"go-web-frame/orm/dialect"
	"go-web-frame/orm/log"
	"go-web-frame/orm/schema"
	"strings"
)

//DB Session

type Session struct {
	db        *sql.DB
	dialect   dialect.Dialect
	refTable  *schema.Schema
	sql       strings.Builder
	sqlParams []any
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlParams = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, sqlParam ...any) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlParams = append(s.sqlParams, sqlParam)
	return s
}

// Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlParams...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	return s.DB().QueryRow(s.sql.String(), s.sqlParams...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlParams)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlParams...); err != nil {
		log.Error(err)
	}
	return
}

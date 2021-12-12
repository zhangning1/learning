package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	db *sql.DB
)

func init() {
	tmpDb, err := sql.Open("mysql", "root:123456@(127.0.0.1)/test")
	if err != nil {
		panic(err)
	}
	db = tmpDb
}

type Student struct {
	Id   uint64
	Name string
	Age  uint32
}

func SearchStudent(name string) (*Student, error) {
	defer db.Close()
	var (
		row        *sql.Row
		resStudent Student
	)

	row = db.QueryRow("select * from student where name = ?", name)
	if err := row.Scan(&resStudent.Id, &resStudent.Name, &resStudent.Age); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Search Student im mysql name=%s", name))
	}

	return &resStudent, nil
}

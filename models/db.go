package models

import (
    // "database/sql"
    // _ "github.com/go-sql-driver/mysql"
    // "fmt"
)

// type database struct {
// 	dataSourceName string
// }
//
// func (d *database) connect() (*sql.DB, error) {
// 	db, err := sql.Open("mysql", d.dataSourceName)
// 	if err != nil {
// 	    return nil, err
// 	}
// 	if err = db.Ping(); err != nil {
// 	    return nil, err
// 	}
// 	return db, nil
// }
//
// type database interface {
//     Connect() *sql.DB
// }
//
// type H struct {
//     D *sql.DB
// }
//
// func (dir H) Connect() *sql.DB {
//     return dir.D
// }
//
// func Measure(d database) {
//     fmt.Println(d.Connect())
// }
//
// func DB(dataSourceName string) (*sql.DB, error) {
//     db, err := sql.Open("mysql", dataSourceName)
//     if err != nil {
//         return nil, err
//     }
//     if err = db.Ping(); err != nil {
//         return nil, err
//     }
//     return db, nil
// }

// db, err := models.DB("admin:manafzadeh@tcp(simurgh.ckyejmpx5kmy.us-west-2.rds.amazonaws.com:3306)/simurgh")
// if err != nil {
// 	log.Fatal(err)
// }
//
// r := models.H{D: db}
// models.Measure(r)
//
//
// fmt.Println(reflect.TypeOf(db))

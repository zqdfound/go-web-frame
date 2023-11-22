package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//func main() {
//	db, er := sql.Open("sqlite3", "/Users/zqdfound/Desktop/zqd/sqlite/testGO.db")
//	if er != nil {
//		log.Fatal(er.Error())
//	}
//	defer func() { db.Close() }()
//	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
//	_, _ = db.Exec("CREATE TABLE User(Name text);")
//	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
//	if err == nil {
//		affected, _ := result.RowsAffected()
//		log.Println(affected)
//	}
//	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
//	var name string
//	if err := row.Scan(&name); err == nil {
//		log.Println("======================")
//		log.Println(name)
//	}
//}

func main() {
	engine, _ := NewEngine("sqlite3", "/Users/zqdfound/Desktop/zqd/sqlite/testGO.db")
	defer engine.Close()
	s := engine.NewSession()
	//_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	//_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	//_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	//result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	//count, _ := result.RowsAffected()
	//fmt.Printf("Exec success, %d affected\n", count)

	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		log.Fatal("Failed to create table User")
	}
}

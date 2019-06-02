package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

const (
	username = "root"
	password = "wang5228"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "learning-go-db"
)

type User struct {
	id int
	name string
	age int
}

func (user *User) String() string {
	return fmt.Sprintf("User(id=%d, name=%s, age=%d)", user.id, user.name, user.age)
}

var db *sql.DB

// 连接数据库
func initDB() {
	dsn := strings.Join([]string{username, ":", password,
		"@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	log.Printf("dsn: %s\n", dsn)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(60 * time.Second)
	db.SetMaxIdleConns(30)
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Printf("mysql connect success!")
}

func selectUserById(userId int) *User {
	user := new(User)
	stmt, err := db.Prepare("select id, `name`, age from `user` where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(userId).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}

func selectAllUser() []User {
	stmt, err := db.Prepare("select id, `name`, age from `user`")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.id, &user.name, &user.name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}

func main() {
	initDB()
	user := selectUserById(1111)
	fmt.Println(user)
	users := selectAllUser()
	for _, user := range users {
		fmt.Println(&user)
	}
}










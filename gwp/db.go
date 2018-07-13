package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id int
	Content string
	Author string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gwp")
	if err != nil {
		panic(err)
	}

	err = Db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("db init successful")
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select i, content, author from posts limit $1", limit)

	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)

		if err != nil {
			return
		}

		posts = append(posts, post)
	}

	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = ?;", id).Scan(&post.Id, &post.Content, &post.Author)

	if err != nil {
		fmt.Println(err)
	}

	return
}

func (post *Post) Create() (err error) {
	fmt.Println("create")
	statement := "insert into posts (content, author) values (?, ?);"

	stmt, err := Db.Prepare(statement)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)

	if err != nil {
		fmt.Println(err)
	}
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)

	if err != nil {
		fmt.Println(err)
	}

	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post) 
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(1)
	fmt.Println(readPost)

	readPost.Content = "PHP is best language"
	readPost.Author = "Paul James"

	readPost.Update()

	readPost.Delete()
}零
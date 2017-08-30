package conductor

import (
	"fmt"
)

type Posts struct {
	Posts []Post

	Key     string
	ForumID int64
	TopicID int64
	UserID  int64
}

type Post struct {
	ID         int64
	AuthorID   int64
	AuthorName string
	TopicID    int64
	Title      string
	Text       string
}

func NewPosts() Posts {
	var ps Posts
	return ps
}

func NewPost() Post {
	var p Post
	return p
}

func (p Post) Set() {
	sql := `INSERT INTO posts (topic_id,user_id,text) VALUES(?,?,?)`
	fmt.Println(sql, p.TopicID, p.AuthorID, p.Text)
	DB.Exec(sql, p.TopicID, p.AuthorID, p.Text)

	UpdateTopicTime(p.TopicID)
}

func (ps *Posts) Get() {

	var params []interface{}
	w := NewWheres()

	if ps.TopicID != 0 {
		w.Add(`topic_id=?`)
		params = append(params, ps.TopicID)
	}

	sql := `SELECT p.id, u.name, p.text FROM posts p LEFT JOIN users u ON u.id=p.user_id  ` + w.Compile()
	fmt.Println(sql)
	rows, _ := DB.Query(sql, params...)
	defer rows.Close()

	for rows.Next() {
		var p Post
		rows.Scan(&p.ID, &p.AuthorName, &p.Text)

		ps.Posts = append(ps.Posts, p)
	}

	fmt.Println(ps)
}

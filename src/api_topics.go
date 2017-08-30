package conductor

import (
	"fmt"
)

type Topics struct {
	Topics []Topic

	Key   string
	Forum int64

	Order string
	Limit int64
}

type Topic struct {
	ID          int64
	Title       string
	Subtitle    string
	Key         string
	Description string
	Text        string
	CreatedDate string
	Created     string

	SiteID  int64
	ForumID int64

	Author    string
	AuthorKey string
	AuthorID  int64
}

func NewTopics() Topics {
	var ts Topics
	ts.Order = "t.updated DESC"
	return ts
}

func NewTopic() Topic {
	var t Topic
	return t
}

func (t *Topic) Set() {
	sql := `INSERT INTO topics (forum_id,user_id,title,subtitle,topic_key) VALUES(?,?,?,?,?)`
	ins, _ := DB.Exec(sql, t.ForumID, t.AuthorID, t.Title, t.Subtitle, t.Key)
	t.ID, _ = ins.LastInsertId()
}

func UpdateTopicTime(id int64) {
	sql := `UPDATE topics SET updated=CURRENT_TIMESTAMP WHERE id=?`

	DB.Exec(sql, id)
}

func (ts *Topics) Get() {
	var params []interface{}
	w := NewWheres()

	if ts.Forum != 0 {
		w.Add(`forum_id=?`)
		params = append(params, ts.Forum)
	}

	if ts.Key != "" {
		w.Add(`topic_key=?`)
		params = append(params, ts.Key)
	}

	sql := `SELECT t.id, t.title, t.topic_key, u.name, t.date FROM topics t LEFT JOIN users u ON u.id=t.user_id ` + w.Compile() + ` ORDER BY ` + ts.Order

	rows, _ := DB.Query(sql, params...)
	defer rows.Close()

	for rows.Next() {
		var t Topic
		rows.Scan(&t.ID, &t.Title, &t.Key, &t.Author, &t.CreatedDate)

		ts.Topics = append(ts.Topics, t)
	}

	fmt.Println(ts)
}

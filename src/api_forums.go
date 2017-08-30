package conductor

import (
	"fmt"
	_ "net/http"
)

type Forums struct {
	Forums []Forum

	Key    string
	SiteID int64
}

type Forum struct {
	ID          int64
	Name        string
	Key         string
	Description string
}

func NewForums() Forums {
	var fs Forums
	return fs
}

func (fs *Forums) Get() {
	var params []interface{}
	w := NewWheres()

	if fs.SiteID != 0 {
		w.Add(`site_id=?`)
		params = append(params, fs.SiteID)
	}

	if fs.Key != "" {
		w.Add(`forum_key=?`)
		params = append(params, fs.Key)
	}

	sql := `SELECT id, name, forum_key, description FROM forums f ` + w.Compile()

	rows, _ := DB.Query(sql, params...)
	defer rows.Close()

	for rows.Next() {
		var f Forum
		rows.Scan(&f.ID, &f.Name, &f.Key, &f.Description)

		fs.Forums = append(fs.Forums, f)
	}

	fmt.Println(fs)
}

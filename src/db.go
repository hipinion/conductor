package conductor

import (
	"strings"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

type Wheres struct {
	Wheres []interface{}
}

func (w *Wheres) Add(val interface{}) {
	w.Wheres = append(w.Wheres, val)
}

func Connect(cs string) error {
	var err error
	DB, err = sql.Open("mysql", `root:`+ConductorConf.Database.Password+`@(:3306)/conductor?collation=utf8mb4_unicode_ci`)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func NewWheres() Wheres {
	var w Wheres
	return w
}

func (w *Wheres) Compile() string {
	var out string
	var outs []string
	if len(w.Wheres) > 0 {
		out = " WHERE "
		for k := range w.Wheres {
			outs = append(outs, w.Wheres[k].(string))
		}
		out = out + strings.Join(outs, " AND ")
		return out
	}

	return out
}

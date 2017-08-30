package conductor

import "fmt"

type Users struct {
	Users []User

	Email string
	Name  string

	SessionKey string

	Count int64
}

type User struct {
	ID    int64
	Name  string
	Email string
	Phone string

	Salt          string
	Hash          string
	Password      string
	Authenticated bool
}

func NewUsers() Users {
	var us Users

	us.Count = 5
	return us
}

func NewUser() User {
	var u User
	return u
}

func (u User) Set() error {

	sql := `INSERT INTO users SET name=?, email=?, phone=?, salt=?, password=?`
	_, err := DB.Exec(sql, u.Name, u.Email, u.Phone, u.Salt, u.Hash)
	if err != nil {
		return err
	}

	return nil
}

func (us *Users) Get() error {
	var params []interface{}
	w := NewWheres()
	w.Add(`1 = 1`)

	if us.Name != "" {
		w.Add("name=?")
		params = append(params, us.Name)
	}

	if us.SessionKey != "" {
		w.Add("id IN (SELECT user_id FROM sessions WHERE id=?)")
		params = append(params, us.SessionKey)
	}

	sql := `SELECT id, name, salt, password FROM users` + w.Compile()
	fmt.Println(sql, params)
	rows, err := DB.Query(sql, params...)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Salt, &u.Password)
		us.Users = append(us.Users, u)
	}

	return nil
}

package conductor

import (
	"crypto/rand"
	"time"
)

// I generate a 64-bit random string
func GenerateSessionKey() (string, error) {
	key := make([]byte, 64)

	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return string(key), nil
}

type Sessions struct {
	Sessions []Session
}

type Session struct {
	ID   string
	User User
}

func NewSessions() Sessions {
	var ss Sessions
	return ss
}

func NewSession() Session {
	var s Session
	return s
}

func (s Session) Remove() {
	DB.Exec(`DELETE FROM sessions WHERE id=? LIMIT 1`, s.ID)
}

func (s Session) SetUser(id int64, key string) {
	timestamp := time.Now().Unix()
	DB.Exec(`UPDATE sessions SET user_id=?, updated=? WHERE id=?`, id, timestamp, key)
}

func (s Session) Set(key string) {
	timestamp := time.Now().Unix()
	DB.Exec(`INSERT INTO sessions SET id=?, started=?, updated=? ON DUPLICATE KEY UPDATE updated=?`, key, timestamp, timestamp, timestamp)
}

package conductor

import (
	_ "fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte("abcdabcdabcdabcd")
var blockKey = []byte("efghefghefghefgh")
var s = securecookie.New(hashKey, blockKey)
var cookieKey = "conductor_session"

func GetSessionCookies(r *http.Request) (bool, map[string]string, error) {

	sk := make(map[string]string)
	if cookie, err := r.Cookie(cookieKey); err == nil {
		if err = s.Decode(cookieKey, cookie.Value, &sk); err == nil {

			return true, sk, nil
		} else {

		}
	} else {

		return false, sk, nil
	}
	return false, sk, nil
}

func GetSessionUser(w http.ResponseWriter, r *http.Request) (Session, User) {

	var Sess Session
	var sessionKey string
	init, cs, err := GetSessionCookies(r)
	if err != nil {

	}
	if !init {
		newKey := MakeSession()
		sessionKey = newKey
		value := map[string]string{
			"session_id": newKey,
		}
		if encoded, err := s.Encode(cookieKey, value); err == nil {
			cookie := &http.Cookie{
				Name:  cookieKey,
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		} else {

		}
		Sess.ID = newKey
		Sess.Set(newKey)
	} else {
		sessionKey = cs["session_id"]
		Sess.ID = cs["session_id"]
		Sess.Set(cs["session_id"])
	}

	var u User
	var us Users
	us.SessionKey = sessionKey
	us.Get()

	if len(us.Users) > 0 {
		u = us.Users[0]
		u.Authenticated = true
	}

	return Sess, u
}

func SetSessionCookie() {

}

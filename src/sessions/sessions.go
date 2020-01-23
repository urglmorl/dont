package sessions

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

var Store *sessions.CookieStore

func Init() {
	sessionKey := []byte(os.Getenv("SESSION_KEY"))
	if len(sessionKey) == 0 {
		sessionKey := securecookie.GenerateRandomKey(32)
		_ = os.Setenv("SESSION_KEY", string(sessionKey))
	}
	Store = sessions.NewCookieStore(sessionKey)
}

func GetSession(w http.ResponseWriter, r *http.Request, sessionName string) (session *sessions.Session, err error) {
	session, err = Store.Get(r, sessionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func CreateSession(w http.ResponseWriter, r *http.Request, sessionName string) (session *sessions.Session, err error) {
	session, err = Store.Get(r, sessionName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = session.Save(r, w)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

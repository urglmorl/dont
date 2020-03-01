package auth

import (
	"constants"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(constants.LoginRoute, login).Methods("POST")
	router.HandleFunc(constants.RefreshTokenRoute, refreshToken).Methods("POST")
	return router
}

func login(w http.ResponseWriter, r *http.Request) {
	// TODO:
	//  генерируем два токена:
	//  1. accessToken - живёт 30 минут
	//  2. refreshToken - живёт 60 дней, сохранён со своим expired_in в БД
	//  Выдаем их клиенту. Клиент использует accessToken для запросов,
	//  а когда он умрёт - рефрешит токены с помощью refreshToken
	//_ = r.ParseForm()
	//username := r.PostForm.Get("login")
	//session, _ := sessions.Store.Get(r, "session")
	//session.Values[username] = utils.TokenGenerator()
	//_ = session.Save(r, w)
	//w.WriteHeader(http.StatusOK)
}

func refreshToken(w http.ResponseWriter, r *http.Request) {
	// TODO:
	//  1. Получаем refreshToken
	//  2. Декодируем его, получаем из него user_id
	//  3. По user_id ищем в БД refreshToken пользователя
	//  4. Проверяем два refreshToken'а на идентичность
	//  5. Если они идентичны - создаём новые accessToken и refreshToken и отправляем клиенту
}

package main

import (
	"net/http"

	gomniauthcommon "github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

type ChatUser interface {
	UniqueID() string
	AvatarURL() string
}

type chatUser struct {
	gomniauthcommon.User
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		h.next.ServeHTTP(w, r)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// segs := strings.Split(r.URL.Path, "/")
	// action := segs[2]
	// provider := segs[3]

	authCookieValue := objx.New(map[string]interface{}{
		"userid": "chat_user_unique_id",
		"name":   "user_name",
		// "avatar_url": avatarURL,
		// "email":      user.Email(),
	}).MustBase64()
	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/",
	})
	w.Header()["Location"] = []string{"/chat"}
	w.WriteHeader(http.StatusTemporaryRedirect)

	// switch action {

	// case "login":

	// 	provider, err := gomniauth.Provider(provider)
	// 	if err != nil {
	// 		log.Fatalln("Failed to get authentication provider: ", provider, "-", err)
	// 	}

	// 	loginUrl, err := provider.GetBeginAuthURL(nil, nil)
	// 	if err != nil {
	// 		log.Fatalln("An error has occured while calling GetBeginURL: ", provider, "-", err)
	// 	}

	// 	w.Header().Set("Location", loginUrl)
	// 	w.WriteHeader(http.StatusTemporaryRedirect)

	// case "callback":

	// 	provider, err := gomniauth.Provider(provider)
	// 	if err != nil {
	// 		log.Fatalln("Failed to get authentication provider: ", provider, "-", err)
	// 	}

	// 	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
	// 	if err != nil {
	// 		log.Fatalln("Failed to finish authentication: ", provider, "-", err)
	// 	}

	// 	user, err := provider.GetUser(creds)
	// 	if err != nil {
	// 		log.Fatalln("Failed to get user: ", provider, "-", err)
	// 	}

	// 	chatUser := &chatUser{User: user}
	// 	m := md5.New()
	// 	io.WriteString(m, strings.ToLower(user.Email()))
	// 	chatUser.uniqueID = fmt.Sprintf("%x", m.Sum(nil))
	// 	avatarURL, err := avatars.GetAvatarURL(chatUser)
	// 	if err != nil {
	// 		log.Fatalln("GetAvatarURL failed - ", err)
	// 	}
	// 	authCookieValue := objx.New(map[string]interface{}{
	// 		"userid":     chatUser.uniqueID,
	// 		"name":       user.Name(),
	// 		"avatar_url": avatarURL,
	// 		"email":      user.Email(),
	// 	}).MustBase64()
	// 	http.SetCookie(w, &http.Cookie{
	// 		Name:  "auth",
	// 		Value: authCookieValue,
	// 		Path:  "/",
	// 	})
	// 	w.Header()["Location"] = []string{"/chat"}
	// 	w.WriteHeader(http.StatusTemporaryRedirect)

	// default:
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Fprintf(w, "unsupported action: %s", action)

	// }
}

func logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	w.Header()["Location"] = []string{"/chat"}
	w.WriteHeader(http.StatusTemporaryRedirect)
}

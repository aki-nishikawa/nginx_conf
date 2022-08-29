package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/aki-nishikawa/go_tutorial/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

var avatars Avatar = TryAvatars{
	UseFilesystemAvatar,
	UseAuthAvatar,
	UseGravatarAvatar,
}

func main() {

	var addr = flag.String("addr", ":8080", "address")
	flag.Parse()

	gomniauth.SetSecurityKey("akihitoTestForGo")
	gomniauth.WithProviders(
		google.New(
			"498531321351-hb5tci5041uks3p2kkti2ciipvgbkg14.apps.googleusercontent.com",
			"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",
			"http://localhost:8080/auth/callback/google",
		),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/uploader", uploadHandler)
	http.HandleFunc("/logout", logout)
	http.Handle("/room", r)
	http.Handle("/avatars/", http.StripPrefix("/avatars/", http.FileServer(http.Dir("./avatars"))))

	// start chat room
	go r.run()

	// start Web server
	log.Println("Start Web Server. Port: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

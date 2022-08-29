package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/aki-nishikawa/go_tutorial/trace"
)

var avatars Avatar = TryAvatars{
	UseFilesystemAvatar,
	UseAuthAvatar,
	UseGravatarAvatar,
}

func main() {

	var addr = flag.String("addr", ":8080", "address")
	flag.Parse()
	
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

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

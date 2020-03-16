package main

import (
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// 탬플릿을 로그하고 컴파일하며 전달하는 구조체 타입
type templateHandler struct {
	once sync.Once
	filename string
	templ    *template.Template
}

// http에서 handleFunc에서 전달되는 함수
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 컴파일을 한번만 수행해서 lazy 초기화 방법을 수행
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data)
}

func main() {
	addr := flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()

	// gomniauth 설정
	gomniauth.SetSecurityKey("MySescretKeyTohash")
	gomniauth.WithProviders(
		// facebook.New("key", "secret", "http://localhost:8080/auth/callback/facebook"),
		// github.New("key", "secret", "http://localhost:8080/auth/callback/github"),
		google.New(os.Getenv("LETSGET23_GOOGLE_KEY"), os.Getenv("LETSGET23_GOOGLE_SECRET"), "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom(UseFileSystemAvatar)
	r.tracer = New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)
	http.Handle("/room", r)
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/avatars/", http.StripPrefix("/avatars", http.FileServer(http.Dir("avatars"))))
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	go r.run()


	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

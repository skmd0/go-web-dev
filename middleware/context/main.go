package main

import (
	"context"
	"html/template"
	"net/http"
)

var (
	indexView *template.Template
	testView  *template.Template
)

type User struct {
	ID uint
}

func ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user User
		if r.URL.Path == "/" {
			user.ID = 1
		} else {
			user.ID = 2
		}
		ctx := r.Context()
		ctx = WithUser(ctx, &user)
		r = r.WithContext(ctx)
		next(w, r)
	})
}

func WithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, "user", user)
}

func GetUser(ctx context.Context) *User {
	if temp := ctx.Value("user"); temp != nil {
		if user, ok := temp.(*User); ok {
			return user
		}
	}
	return nil
}

func index(w http.ResponseWriter, r *http.Request) {
	user := GetUser(r.Context())
	w.Header().Set("Content-Type", "text/html")
	err := indexView.ExecuteTemplate(w, "bulma", user.ID)
	if err != nil {
		panic(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	user := GetUser(r.Context())
	w.Header().Set("Content-Type", "text/html")
	err := testView.Execute(w, user.ID)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	dirPath := "/home/skmd/Projects/skamlic.com/go-web-dev/middleware/context/"
	indexView, err = template.ParseFiles(dirPath+"base.gohtml", dirPath+"index.gohtml")
	if err != nil {
		panic(err)
	}
	testView, err = template.ParseFiles(dirPath + "test.gohtml")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", ApplyFn(index))
	http.HandleFunc("/test", ApplyFn(test))
	http.ListenAndServe(":3000", nil)
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/getUser", handlers.GetUser)
// 	//start your server
// 	http.ListenAndServe(":8081", nil)
// }

// func main() {
// 	http.HandleFunc("/home", Mid(Mid1(HomePage)))
// 	http.ListenAndServe(":8081", nil)
// }

// func HomePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("home page invoked")
// 	fmt.Fprintln(w, "this is my home")
// }

// func Mid(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("middleware invoked")
// 		next(w, r)
// 	}
// }

// func Mid1(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("middleware invoked")
// 		next(w, r)
// 	}
// }

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type reqKey int

const RequestIDKey reqKey = 123

func main() {
	http.HandleFunc("/home", RequestIdMid(LoggingMid(Mid1(homePage))))
	http.ListenAndServe(":8081", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	log.Println("In home Page handler")
	fmt.Fprintln(w, "this is my home page")

}

func RequestIdMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.NewString()
		ctx := r.Context()
		ctx = context.WithValue(ctx, RequestIDKey, uuid)
		next(w, r.WithContext(ctx))

	}
}
func LoggingMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqId, ok := ctx.Value(RequestIDKey).(string)
		if !ok {
			reqId = "Unknown"
		}
		log.Printf("%s : started   : %s %s ",
			reqId,
			r.Method, r.URL.Path)
		defer log.Println("completed")
		next(w, r)
	}

}

func Mid1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqId, ok := ctx.Value(RequestIDKey).(string)
		log.Println(reqId, ok)
		next(w, r)
	}
}

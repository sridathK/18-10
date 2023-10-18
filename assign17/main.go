package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Grade string
}

func main() {
	http.HandleFunc("/studentDetails", LoggingMid(FetchStudentDetails))
	http.HandleFunc("/students", LoggingMid(FetchAllStudents))
	http.ListenAndServe(":8081", nil)
}

func FetchAllStudents(w http.ResponseWriter, r *http.Request) {
	uData, err := FetchAllUser()
	if err != nil {

		fetchError := map[string]string{"msg": "user not found"}

		errData, err := json.Marshal(fetchError)

		if err != nil {

			log.Println("Error while parsing fetchuser error conversion: ", err)

			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

			return

		}

		w.WriteHeader(http.StatusInternalServerError)

		w.Write(errData)

		return

	}

	userData, err := json.Marshal(uData)
	if err != nil {

		log.Println("Error while converting user data to json", err)

		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

		return

	}

	w.Write(userData)
}

func FetchStudentDetails(w http.ResponseWriter, r *http.Request) {
	userIdString := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseUint(userIdString, 10, 64)

	if err != nil {
		log.Println("Error: ", err)

		errorInConversion := map[string]string{"msg": "not a valid number"}

		jsonData, err := json.Marshal(errorInConversion)

		if err != nil {
			log.Println("Error while converting error to json", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		w.WriteHeader(http.StatusBadRequest)

		w.Write(jsonData)

		return

	}

	uData, err := FetchUser(userId)

	if err != nil {

		fetchError := map[string]string{"msg": "user not found"}

		errData, err := json.Marshal(fetchError)

		if err != nil {

			log.Println("Error while parsing fetchuser error conversion: ", err)

			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

			return

		}

		w.WriteHeader(http.StatusInternalServerError)

		w.Write(errData)

		return

	}
	userData, err := json.Marshal(uData)

	if err != nil {

		log.Println("Error while converting user data to json", err)

		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

		return

	}

	w.Write(userData)

}

var students = map[uint64]Student{
	123: {
		Name:  "Bob",
		Age:   23,
		Grade: "vII",
	},
	456: {
		Name:  "mm",
		Age:   25,
		Grade: "vIIi",
	},
}

func FetchUser(userId uint64) (Student, error) {
	u, ok := students[userId]
	if !ok {
		return Student{}, errors.New("user not there")
	}
	return u, nil

}

func FetchAllUser() ([]Student, error) {
	studentsSlice := []Student{}
	for _, v := range students {
		studentsSlice = append(studentsSlice, v)
	}
	return studentsSlice, nil
}

type reqKey int

const RequestIDKey reqKey = 123

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

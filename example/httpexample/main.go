package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	. "im-micro/example/httpexample/Services"
	"net/http"
)

func main(){
	user := UserService{}
	endp := GetUserEndpoint(user)

	serverHandler := httptransport.NewServer(endp, DecodeUserRequest, EncodeUserResponse)

	r := mymux.NewRouter()
	r.Handle(`/user/{uid:\d+}`, serverHandler)

	http.ListenAndServe(":8080", r)
}


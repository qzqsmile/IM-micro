package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	. "im-micro/example/httpexample/Services"
	"net/http"
)

func main(){
	user := UserService{}
	endp := GetUserEndpoint(user)

	serverHandler := httptransport.NewServer(endp, DecodeUserRequest, EncodeUserResponse)
	http.ListenAndServe(":8080", serverHandler)
}


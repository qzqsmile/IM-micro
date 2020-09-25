package main

import (
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	. "im-micro/example/httpexample/Services"
	"im-micro/example/httpexample/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	//servier->endpoint->transport
	user := UserService{}
	endp := GetUserEndpoint(user)

	serverHandler := httptransport.NewServer(endp, DecodeUserRequest, EncodeUserResponse)

	r := mymux.NewRouter()
	r.Handle(`/user/{uid:\d+}`, serverHandler)
	r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, r2 *http.Request){
		writer.Header().Set("Content-type", "application/json")
		writer.Write([]byte(`{"status": "ok"}`))
	})

	errChan := make(chan error)

	go (func() {
		util.RegService()
		err := http.ListenAndServe(":8080", r)
		if err != nil{
			log.Println(err)
			errChan <- err
		}
	})()

	go (func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c, syscall.SIGINT, syscall.SIGTERM)
		errChan<-fmt.Errorf("%s", <-sig_c)
	})()

	getErr := <-errChan
	util.Unregservice()
	log.Println(getErr)
}




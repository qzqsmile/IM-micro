package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	httptransport "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
	. "im-micro/examples/httpexample/ClientServices"
	"github.com/go-kit/kit/log"
	"io"
	"net/url"
	"os"
)

//func main(){
//	tgt, _:= url.Parse("http://localhost:8080")
//	client := httptransport.NewClient("GET", tgt, GetUserInfo_Request, GetUserInfo_Response)
//	getUserInfo := client.Endpoint()
//
//	ctx := context.Background()
//	res, err := getUserInfo(ctx, UserRequest{Uid: 102})
//	if err != nil{
//		fmt.Print(err)
//		os.Exit(1)
//	}
//
//	userinfo := res.(UserResponse)
//	fmt.Println(userinfo.Result)
//}

func main(){

	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	api_client, _ := consulapi.NewClient(config)
	client := consul.NewClient(api_client)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
	}
	{
		tags := []string{"primary"}
		instancer := consul.NewInstancer(client, logger, "userservice", tags, true)
		{
			factory := func(service_url string)(endpoint.Endpoint, io.Closer, error){
				tart, _ := url.Parse("http://"+service_url)
				return httptransport.NewClient("GET", tart, GetUserInfo_Request, GetUserInfo_Response).Endpoint(),
				nil, nil
			}
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			endpoints, _ := endpointer.Endpoints()
			fmt.Println("服务有", len(endpoints), "条")
			getUserInfo := endpoints[0]

			ctx := context.Background()
			res, err := getUserInfo(ctx, UserRequest{Uid: 102})
			if err != nil{
				fmt.Print(err)
				os.Exit(1)
			}

			userinfo := res.(UserResponse)
			fmt.Println(userinfo.Result)
		}
	}
}
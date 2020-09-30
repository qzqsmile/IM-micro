package util

import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

var ConsulClient * consulapi.Client

func init(){
	config := consulapi.DefaultConfig()
	config.Address="127.0.0.1:8500"

	client, err := consulapi.NewClient(config)
	if err != nil{
		log.Fatal(err)
	}
	ConsulClient = client
}

func RegService(){
	reg := consulapi.AgentServiceRegistration{}
	reg.ID = "userservice"
	reg.Name = "userservice"
	reg.Address = "127.0.0.1"
	reg.Port = 8080
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://127.0.0.1:8080/health"

	reg.Check = &check

	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil{
		log.Fatal(err)
	}
}


func Unregservice(){
	ConsulClient.Agent().ServiceDeregister("userservice")
}
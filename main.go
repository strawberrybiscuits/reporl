package main

import (
	"log"
	"fmt"
)

import(
	"reporl/config"
	"reporl/service"
)

func init(){
	
	config.Setup()

}

func main(){
	
	err := service.SaveRepoRssFromFile()
	if err != nil{
		
		log.Fatalf(fmt.Sprintf("error:%s",err))
		
	}
}
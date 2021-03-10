package main

import(
	"reporl/config"
	"reporl/util"
)

func init(){
	
	config.Setup()

}

func main(){
	
		util.TestReadRepos()
	
}
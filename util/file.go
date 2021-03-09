package util

import (
	"os"
	"sync"	
)

// 读取文件中仓库列表返回列表数组
func ReadFiles(filepaths []string) (arr []string){
	
	if 0 == len(filepaths) {
		
		return nil
		
	}
	waitGroup := sync.WaitGroup{}
	
	for _,v := range filepaths{
		
		waitGroup.Add(1)
		go func(array []string){
			
			arr = append(arr,ReadFile(v)...)
			defer waitGroup.Done()
			
		}(arr)
		
	}
	
	waitGroup.Wait()
	
	return
}

func ReadFile(filepath string) (arr []string){
	
	f,err := os.Open(filepath)
	
	if err != nil{
		
		panic(err)
	}
	
	
	
	return arr
	
}
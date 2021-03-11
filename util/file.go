package util

import (
	"fmt"
	"bufio"
	"os"
	"sync"
	"reporl/config"
)

// 读取文件中仓库列表返回列表数组
func ReadFiles(filepaths []string) (arr []string){
	var lock sync.Mutex
	if 0 == len(filepaths) {
		
		return nil
		
	}
	waitGroup := sync.WaitGroup{}
	
	for _,v := range filepaths{
		
		waitGroup.Add(1)
		go func(val []string){
			lock.Lock()
			arr = append(arr,ReadFile(val)...)
			lock.Unlock()
			defer waitGroup.Done()
			
		}(v)
		
	}
	
	waitGroup.Wait()
	
	return
}

func ReadFile(filepath string) (arr []string){
	
	f,err := os.Open(filepath)
	
	if err != nil{
		
		panic(err)
		
	}
	
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	
	for scanner.Scan(){
		
		repo := scanner.Text()
		if "" != repo{
			arr = append(arr,repo)
		}
	}

	
	return arr
	
}

func TestReadRepos(){
	
	
	
	fmt.Println(ReadFiles(config.FPaths.Paths))
	
	
}

package util

// go query 

import (
	"time"
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	"reporl/model"
	"net/http"
	"strings"
)

func GetRepoReleases(repoUrl string) (repos []model.Repo_rl){
	
	_s := strings.Split(repoUrl,"/")
	repoName := fmt.Sprintf("/%s/%s",_s[3],_s[4])
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	h, err := client.Get(repoUrl)
	if err != nil{
		
		panic(err)
		
	}
	
	defer h.Body.Close()
	
	if 200 != h.StatusCode{
		
		log.Fatalf("http connect failed :%d %s",h.StatusCode, h.Status)
		
	}
	
	dom,err := goquery.NewDocumentFromReader(h.Body)
	
	if err != nil{
		panic(err)
	}
	
	dom.Find("entry").Each(func(i int, s *goquery.Selection){
		reporl := model.Repo_rl{
			RepoName: repoName,
			Title: s.Find("title").Text(),
			Description: s.Find("content").Text(),
			SimpleName: _s[4],
		}
		link, _ := s.Find("link").Attr("href")
		reporl.Link = link
		
		tm,err := time.ParseInLocation("2006-01-02T15:04:05Z",s.Find("updated").Text(),time.Local)
		
		if err != nil{
			panic(err)
		}
		
		reporl.CreateTime = tm

		//fmt.Println(tm.Format("2006-01-02 15:04:05"))

		repos = append(repos, reporl)
	})

	
	return
}

func GetLatestRepoRelease(repoUrl string){
	
	
	
	
}
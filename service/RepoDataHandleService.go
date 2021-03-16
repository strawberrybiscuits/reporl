package service

import (
	"fmt"
	"strings"
	"reporl/model"
	"reporl/util"
)


func SaveRepoRssFromFile() (err error){
	
	var repos []*model.Repo_rss
	
	arr := util.ReadRepos()

	for	_,v := range arr{
		
		simpleName := strings.Split(v,"/")
		
		tmp := &model.Repo_rss{
			
			RepoName: v,
			RepoUrl: fmt.Sprintf("https://github.com%s/releases.atom", v),
			SimpleName: simpleName[1],
			
		}
		
		repos = append(repos,tmp)
		
	}
	
	err = model.AddReposRss(repos)
	
	return
}


func SaveAllRepoReleases()(err error){
	
	// get all repoUrls from db then 
	repoUrls,err := model.GetAllRepoRssAddress()
	
	if err != nil{
		return
	}
	
	err = SaveReposReleasesInformationToDb(repoUrls)
	
	if err != nil{
		
		return
	}
	
	return
	
}

func SaveReposReleasesInformationToDb(repoUrls []string) (err error){
	
	for _,v := range repoUrls{
		
		err = SaveRepoReleasesInformationToDb(v)
		
	}
	
	return
	
}


func SaveRepoReleasesInformationToDb(repoUrl string) (err error){
	
	reporls := util.GetRepoReleases(repoUrl)
	
	err = model.AddRepos(reporls)
	
	return
}

//
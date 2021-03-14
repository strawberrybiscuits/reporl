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
	
	return err
}

func SaveRepoReleasesInformationToDb(reporl model.Repo_rl) (err error){
	
	
	
	return
}

//
package model

import (
	
)

type Repo_rss struct{
	
	Id int `json:"-" gorm:"primary_key"`
	RepoUrl string `json:"repoUrl"`
	RepoName string `json:"repoName"`
	SimpleName string `json:"simpleName"`
	
}

func AddRepoRss(reporss Repo_rss) (err error){
	
	if err = db.Create(&reporss).Error; err != nil{
		return err
	}
	
	return 
	
}

func AddReposRss(repos []Repo_rss) (err error){
	
	for _,v := range repos{
		
		if err = AddRepoRss(v); err != nil{
			return err
		}
		
	}
	
	return 
	
}

func GetRepoRssAddressByRepoName(name string)(url string, err error){
	
	var reporss Repo_rss
	
	if err = db.Select("repoUrl").Where("repoName = ?", name).FirstOrInit(&reporss).Error; err != nil{
		return "",err
	}
	url = reporss.RepoUrl
	return 
	
}

func GetRepoRssAddressBySimpleName(name string)(url string, err error){
	
	var reporss Repo_rss
	
	if err = db.Select("repoUrl").Where("simpleName = ?", name).FirstOrInit(&reporss).Error; err != nil{
		return "",err
	}
	url = reporss.RepoUrl
	return 
	
}
package model

import (

	"reporl/config"
)

type Repo_rss struct{
	
	Id int `json:"-" gorm:"primary_key"`
	RepoUrl string `gorm:"column:repoUrl" json:"repoUrl"`
	RepoName string `gorm:"column:repoName" json:"repoName"`
	SimpleName string `gorm:"column:simpleName" json:"simpleName"` 
	
}


// 唯一约束，1.插入时判断。 2.设置唯一索引。

func AddRepoRss(reporss *Repo_rss) (err error){
	
	if err = config.Db.Exec("INSERT INTO repo_rss(repoUrl,repoName,simpleName) SELECT ?,?,? from dual where not exists(select * from repo_rss where repoName = ? and repoUrl = ?)",reporss.RepoUrl,reporss.RepoName,reporss.SimpleName,reporss.RepoName,reporss.RepoUrl).Error; err != nil{
		return err
	}
	
	return 
	
}

func AddReposRss(repos []*Repo_rss) (err error){
	
	for _,v := range repos{
		
		if err = AddRepoRss(v); err != nil{
			return err
		}
		
	}
	
	return 
	
}

func GetRepoRssAddressByRepoName(name string)(url string, err error){
	
	var reporss Repo_rss
	
	if err = config.Db.Select("repoUrl").Where("repoName = ?", name).FirstOrInit(&reporss).Error; err != nil{
		return "",err
	}
	url = reporss.RepoUrl
	return 
	
}

func GetRepoRssAddressBySimpleName(name string)(url string, err error){
	
	var reporss Repo_rss
	
	if err = config.Db.Select("repoUrl").Where("simpleName = ?", name).FirstOrInit(&reporss).Error; err != nil{
		return "",err
	}
	url = reporss.RepoUrl
	return 
	
}

func GetAllRepoRssAddress()(url []string, err error){
	
	var reporss []Repo_rss

	if err = config.Db.Select("repoUrl").Find(&reporss).Error; err != nil{
		return 
	}
	
	 for _,v := range reporss{

	 	url = append(url, v.RepoUrl)
	 }
	

	return
}
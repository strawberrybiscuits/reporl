package model

import (
    "reporl/config"
	"time"
	
)



type Repo_rl struct{
	
	Id int `json:"-" gorm:"primary_key"`
	RepoName string `json:"repoName"`
	Title string `json:"title"`
	Link string `json:"link"`
	Description string `json:"description"`
	SimpleName string `json:"simpleName"`
	CreateTime time.Time `json:"releaseTime"`
	
}

func AddRepo(repo Repo_rl) (err error){
	
	if err = config.Db.Create(&repo).Error; err != nil{
		return err
	}
	
	return 
	
}

func AddRepos(repos []Repo_rl)(err error){
	
	for _,v := range repos{
		
		if err = AddRepo(v); err != nil{
			return err
		}
		
	}
	
	return 
}

func GetAllRepos() (repos []Repo_rl,err error){
	
	if err = config.Db.Find(&repos).Error; err != nil{
		
		return nil,err
	}
	
	return
	
}

// 获取RepoName所有仓库
// repoName : "/abc/bcd"
func GetReposByRepoName(repoName string) (repos []Repo_rl,err error){
	
	if err = config.Db.Where("repoName = ?", repoName).Find(&repos).Error; err != nil{
		
		return nil, err
	}
	
	return
	
}

// 获取RepoName最新仓库
func GetLatestRepoByRepoName(repoName string) (repo *Repo_rl, err error){
	
	if err = config.Db.Where("repoName = ?",repoName).Order("id desc").Limit(1).Find(&repo).Error; err != nil{
		return nil,err
	}
	
	return
	
}


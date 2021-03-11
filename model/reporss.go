package model

import (
	"reporl/config"
	"github.com/jinzhu/gorm"
)

type Repo_rss struct{
	
	Id int `json:"-" gorm:"primary_key"`
	RepoUrl string `json:"repoUrl"`
	RepoName string `json:"repoName"`
	SimpleName string `json:"simpleName"`
	
}


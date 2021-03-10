package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"fmt"
)

const yamlPath string = "./config/config.yaml"
var conf = &Config{}
var Db *gorm.DB
var FPaths FilePaths

type Config struct{
	
	DbConfig Database `yaml:"mysql"`
	FilePathsConfig FilePaths `yaml:"file"`
}

type Database struct{
	
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	
} 

type FilePaths struct{
	
	Paths []string `yaml:"path"`
	
	
}




func ReadYamlConfig(path string,config *Config)  error{
	
	if f, err := os.Open(path); err != nil {
	    return err
	} else {
		yaml.NewDecoder(f).Decode(config)
	}
	return nil
}

func init(){
	
	err :=  ReadYamlConfig(yamlPath,conf)
	
	if err != nil{
		
		log.Fatalf("read yaml err: %v", err)
		
	}
	
}

func Setup(){
	
	dbSetup()
	filePathSetup()
	
}

func dbSetup(){
	
	var err error
	Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DbConfig.Username,
		conf.DbConfig.Password,
		conf.DbConfig.Host,
		conf.DbConfig.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	
	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	
	
}

func filePathSetup(){
	
	FPaths = conf.FilePathsConfig
	
}
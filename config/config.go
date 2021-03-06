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
var Db *gorm.DB

type Config struct{
	
	DbConfig Database `yaml:"mysql"`
	
}

type Database struct{
	
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Username string `yaml:username`
	Password string `yaml:password`
	
} 

var conf = &Config{}


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
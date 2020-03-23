package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Configuration struct {
	Limit         int    `yaml:"limit"`
	Interval      int    `yaml:"interval"`
	Code          string `yaml:"code"`
	UaCheck       int    `yaml:"ua_check"`
	NewFile       string `yaml:"new_file"`
	WhiteIP       string `yaml:"white_ip"`
	WhiteUa       string `yaml:"white_ua"`
	BlackUa       string `yaml:"black_ua"`
	ForceUaString string `yaml:"force_ua_string"`
	LogFile       string `yaml:"log_file"`
}

func cat(path string) *os.File {
    File, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
    }
    return File
}

func ls(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	return nil
}

func isEmptyString(data string) bool {
	if len(data) == 0 {
		return true
	}
	return false
}

func logerror(logerror error) {
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "[ERROR] [ ", log.LstdFlags)
	logger.Println("]:",logerror)
}

func ipoverflow(code,status string) {
	//filter multiple code//

	if isEmptyString(code) == false {
		for _, chk := range strings.Split(code,",") {
			if status == chk {
				tcode = 1
			}
		}
	}


}

var configuration Configuration
var tcode int
func main() {

    yamlFile := cat("C:/Users/Greypanel-kingsley/go/src/cc_defence/conf.yaml")

    //get all configuration content//
    byteValue, _ := ioutil.ReadAll(yamlFile)
	_ = yaml.Unmarshal(byteValue, &configuration)
	_ = yamlFile.Close()

	//set all configuration variables//
	code := configuration.Code
	status := "400"

	//Main Function//
	fmt.Println(ls("C:/Users/Greypanel-kingsley/go/src/"))


	ipoverflow(code,status)
	fmt.Println(tcode)
}
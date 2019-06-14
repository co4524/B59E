package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func loadJSONFile(file string) {
	filePath := "../solidity/build/contracts/" + file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
	}
	var parser map[string]interface{}
	err = json.Unmarshal(data, &parser)
	if err != nil {
		log.Println(err)
	}
	log.Println(parser["abi"])

}

func handleABIBCP(c *gin.Context) {
	loadJSONFile("BCPPoints.json")
}

func handleABILIP(c *gin.Context) {

}

func handleABIPlatform(c *gin.Context) {

}

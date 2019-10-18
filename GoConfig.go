package goconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// GoConfig loads application properties from config directory based on PROFILE Env variable
func GoConfig(config interface{}) error {
	jsonFile, err := ioutil.ReadFile("./config/global.json") // open the standard global configuration json file
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonFile, config)
	if err != nil {
		return err
	}
	profile := os.Getenv("PROFILE") // the PROFILE env variable tells us which json file to read
	if profile == "" {
		return errors.New("PROFILE environment variable not found or not set")
	}
	FileName := fmt.Sprintf("./config/%s.json", profile) // create file name based on profile
	jsonFile, err = ioutil.ReadFile(FileName)            // read the profile specific json file
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonFile, config)
	if err != nil {
		return err
	}
	return nil
}

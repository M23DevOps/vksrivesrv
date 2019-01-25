package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//Configuration setup our database
type Configuration struct {
	Database struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Dbname   string `json:"dbname"`
		Password string `json:"password"`
	} `json:"database"`
	App struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RedirectUri  string `json:"redirect_uri"`
	} `json:"app"`
	Bot struct {
		BotId     		string `json:"bot_id"`
		Token 			string `json:"token"`
		PrivateAlbumId  string `json:"private_album_id"`
		PublicAlbumId  	string `json:"public_album_id"`
	} `json:"bot"`
}

//Config is a global variable
var Config Configuration

//LoadConfiguration setup config
func LoadConfiguration(file string) Configuration {

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&Config)

	fmt.Println("Configuration loaded")

	return Config
}

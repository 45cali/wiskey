package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const serveConf = ".vindalu/wiskey"

type server struct {
	Server string `json:"server"`
}

// Server retrieves the url for the vindalu server
func Server() string {
	var b []byte
	b, err := ioutil.ReadFile(os.Getenv("HOME") + "/" + serveConf)
	if err != nil {
		fmt.Printf("could not read file at $HOME/%s\n", serveConf)
		os.Exit(0)

	}
	serv := server{}
	err = json.Unmarshal(b, &serv)
	return serv.Server
}

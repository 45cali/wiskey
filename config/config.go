package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const SERV_CONF = ".vindalu/wiskey"

type server struct {
	Server string `json:"server"`
}

func Server() string {
	var b []byte
	b, err := ioutil.ReadFile(os.Getenv("HOME") + "/" + SERV_CONF)
	if err != nil {
		fmt.Println("could not read file")
	}

	serv := server{}
	err = json.Unmarshal(b, &serv)

	return serv.Server
}

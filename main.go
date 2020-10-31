package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	authURL string = "https://hk4e-sdk-os.mihoyo.com/hk4e_global/mdk/shield/api/login?" // Default auth url. It is correct as of 31.10.2020
)

var (
	username string = "Felix_it@mail.ru"
	password string = "Blacka777"
)

func main() {
	client := &http.Client{}
	password = encryptPassword([]byte(password))
	var data = strings.NewReader(fmt.Sprintf("{\"account\":\"%s\",\"is_crypto\":true,\"password\":\"%s\"}", username, password))
	req, err := http.NewRequest("POST", authURL, data)
	if err != nil {
		log.Fatal(err)
	}
  
	req.Header.Set("x-rpc-sys_version", "14.1")
	req.Header.Set("x-rpc-channel_id", "1")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("x-rpc-device_id", "B930FA2D-6339-BCD5-E53F-6D4BB466B337") // Randomize it
	req.Header.Set("x-rpc-client_type", "1")
	req.Header.Set("x-rpc-channel_version", "1.10.1")
	req.Header.Set("x-rpc-mdk_version", "1.10.1")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("x-rpc-game_biz", "hk4e_global")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("x-rpc-language", "en")
	req.Header.Set("x-rpc-device_name", "iPhone")
	req.Header.Set("User-Agent", "Genshin%20Impact/102 CFNetwork/1197 Darwin/20.0.0")
	req.Header.Set("x-rpc-device_model", "iPhone8,4")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	bodyText, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	fmt.Scanln()
}

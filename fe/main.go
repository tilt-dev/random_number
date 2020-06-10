package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	svcNumbers = "numbers"
)

var numbersHost = flag.String("numbers-host", svcNumbers, "The host for the numbers service")
var port = flag.Int("port", 8080, "The port on which to serve http")


type Info struct {
	Number int
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handleMain)

	log.Printf("Serving the frontend on :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("> Getting a random number!")
	t, err := template.ParseFiles(templatePath("index.tpl"))
	if err != nil {
		log.Fatalf("error parsing template: %v\n", err)
	}

	num, err := getNumber()
	if err != nil {
		log.Printf("error getting number: %v\n", err)
		http.Error(w, "error getting number", http.StatusInternalServerError)
		return
	}

	info := Info{
		Number: num,
	}

	err = t.Execute(w, info)
	if err != nil {
		log.Fatalf("error executing template: %v\n", err)
	}
}

func templatePath(f string) string {
	dir := os.Getenv("TEMPLATE_DIR")
	if dir == "" {
		dir = "web/templates"
	}

	return filepath.Join(dir, f)
}

func getNumber() (int, error) {
	resp, err := pingPortForService(*numbersHost)
	if err != nil {
		return 0, err
	}

	s := strings.TrimSpace(string(resp))
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func pingPortForService(host string) ([]byte, error) {
	url := fmt.Sprintf("http://%s", host)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("request to %s failed: %v", url, err)
	}

	return ioutil.ReadAll(resp.Body)
}

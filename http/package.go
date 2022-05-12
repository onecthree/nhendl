package http

import (
	"io/ioutil"
	"net/http"
	"log"
	"os"
	"io"
	// "fmt"
	"strconv"
	// "github.com/cavaliergopher/grab/v3"
)

func Get(uri string) map[string]interface{} {
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}

	if(res.StatusCode == 404) {
		return map[string]interface{} {
			"code" : 404,
			"src"  : "Not Found",
		}
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return map[string]interface{}{
		"code" : res.StatusCode,
		"src" : string(data[:]),
	}
}

func Download(dirname string, filename int, uri string, callback func()) {
	err := downloadFile(dirname + "/" + strconv.Itoa(filename) + ".jpg", uri + strconv.Itoa(filename) + ".jpg")
	if err != nil {
		panic(err)
	} else {
		callback()
	}
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err

	// _, err := grab.Get(filepath, url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// return err
}
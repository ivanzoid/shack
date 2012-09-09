package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	URL = "http://www.imageshack.us/upload_api.php"
	KEY = "" // paste your ImageShack API key
)

func Die(format string, args ...interface{}) {
	str := fmt.Sprintf(format, args...)
	fmt.Fprintln(os.Stderr, str)
	os.Exit(1)
}

func FakeUTF8CharsetReader(charset string, input io.Reader) (io.Reader, error) {
	return input, nil
}

func main() {
	Die("foo")
	if len(KEY) == 0 {
		Die("please set your API key in source")
	}

	// Get arguments

	if len(os.Args) == 1 {
		Die("missing file name")
	}

	filename := os.Args[1]

	// Create form request

	buffer := new(bytes.Buffer)
	formWriter := multipart.NewWriter(buffer)

	formWriter.WriteField("rembar", "1") // remove information bar
	formWriter.WriteField("key", KEY)
	fileWriter, _ := formWriter.CreateFormFile("fileupload", filename)

	file, error := os.Open(filename)

	if error != nil {
		Die("can't open file: %v", error)
	}

	defer file.Close()

	_, error = io.Copy(fileWriter, file)

	if error != nil {
		Die("can't read file: %v", error)
	}

	formWriter.Close()

	request, error := http.NewRequest("POST", URL, buffer)

	if error != nil {
		Die("can't create request: %v", error)
	}

	request.Header.Set("Content-Type", formWriter.FormDataContentType())

	// Create http client and perform request

	client := new(http.Client)
	response, error := client.Do(request)

	if error != nil {
		Die("error performing request: %v", error)
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		Die("error: received %v status code", response.StatusCode)
	}

	// Parse result

	type XmlResult struct {
		Link string `xml:"links>image_link"`
	}
	result := XmlResult{}

	xmlDecoder := xml.NewDecoder(response.Body)
	xmlDecoder.CharsetReader = FakeUTF8CharsetReader
	error = xmlDecoder.Decode(&result)

	if error != nil {
		Die("error parsing xml: %v", error)
	}

	fmt.Printf("%v", result.Link)
}

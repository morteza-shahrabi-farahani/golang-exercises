package handlers

import (
	"encoding/json"
	"example.com/m/v2/pkg/models"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func DownloadFile(response http.ResponseWriter, request *http.Request) {
	jsonBody, err5 := ioutil.ReadAll(request.Body)
	if err5 != nil {
		fmt.Println("This is error")
		fmt.Println(err5)
	}

	var lastFile models.Filee

	json.Unmarshal(jsonBody, &lastFile)
	//fmt.Println(lastFile)
	address := lastFile.Id
	fmt.Println(address)

	res1 := strings.Split(address, ":")
	fmt.Println(res1[1])

	Openfile, err := os.Open(res1[1]) //Open the file to be downloaded later
	 //Close after function return

	if err != nil {
		newAdd := "./uploads/" + res1[1]
		Openfile2, err2 := os.Open(newAdd)
		if err2 != nil {
			http.Error(response, "File not found.", 404) //return 404 if file is not found
			return
		}
		defer Openfile2.Close()

		tempBuffer := make([]byte, 512) //Create a byte array to read the file later
		Openfile.Read(tempBuffer) //Read the file into  byte
		FileContentType := http.DetectContentType(tempBuffer) //Get file header

		FileStat, _ := Openfile2.Stat() //Get info from file
		FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

		Filename := "demo_download"

		//Set the headers
		response.Header().Set("Content-Type", FileContentType+";"+Filename)
		response.Header().Set("Content-Length", FileSize)
		Openfile2.Seek(0, 0) //We read 512 bytes from the file already so we reset the offset back to 0
		io.Copy(response, Openfile2) //'Copy' the file to the client

	}

	defer Openfile.Close()

	tempBuffer := make([]byte, 512) //Create a byte array to read the file later
	Openfile.Read(tempBuffer) //Read the file into  byte
	FileContentType := http.DetectContentType(tempBuffer) //Get file header

	FileStat, _ := Openfile.Stat() //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	Filename := "demo_download"

	//Set the headers
	response.Header().Set("Content-Type", FileContentType+";"+Filename)
	response.Header().Set("Content-Length", FileSize)
	Openfile.Seek(0, 0) //We read 512 bytes from the file already so we reset the offset back to 0
	io.Copy(response, Openfile) //'Copy' the file to the client
}

package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"example.com/m/v2/pkg/mocks"
	"example.com/m/v2/pkg/models"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func UploadFile(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 * 1024 * 1024)

	file, handler,err := request.FormFile("myfile")

	if err != nil {

		defer request.Body.Close()
		jsonBody, err5 := ioutil.ReadAll(request.Body)
		if err5 != nil {
			fmt.Println("This is error")
			fmt.Println(err5)
		}

		var downloadableFile models.DownloadFile

		json.Unmarshal(jsonBody, &downloadableFile)
		fmt.Println(downloadableFile)
		url := downloadableFile.FileUrl
		result, err6 := http.Get(url)

		if err6 != nil {
			fmt.Println(err6)
		}

		defer result.Body.Close()
		fmt.Println(result)
		chance := base64.StdEncoding.EncodeToString([]byte(downloadableFile.FileUrl))
		holdFile, err7 := os.Create(chance + "-sample.pdf")
		if err7 != nil {
			log.Fatal(err7)
		}
		defer holdFile.Close()

		_, err8 := io.Copy(holdFile, result.Body)
		if err8 != nil {
			log.Fatal(err8)
		}

		var newFile models.Filee

		first := sha256.New()
		first.Write([]byte(url))
		hashPart := first.Sum(nil)
		newFile.Id = string(hashPart) + ":" + holdFile.Name()
		newFile.Location = chance + "-sample.pdf"
		fmt.Println(newFile.Id)
		fmt.Println(newFile.Location)
		mocks.Files = append(mocks.Files, newFile)
		newFileJson, err4 := json.Marshal(newFile)
		if err4 != nil {
			fmt.Println(err4)
		}

		fmt.Println(newFileJson)
		response.WriteHeader(http.StatusCreated)
		response.Header().Add("Content-Type", "application/json")
		//response.Write([]byte(`{"field_id":"newFile.Id"}`))
		response.Write(newFileJson)
		fmt.Println("Success!")

		return
	}

	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File name: ", handler.Filename)
	fmt.Println("File Size: ", handler.Size)
	fmt.Println("File Type: ", handler.Header.Get("Content-Type"))

	chance2 := base64.StdEncoding.EncodeToString([]byte(handler.Filename))
	tempFile, err2 := ioutil.TempFile("uploads", chance2 + "-*.jpg")
	if err2 != nil {
		fmt.Println("This is error 2")
		fmt.Println(err2)
	}


	defer tempFile.Close()

	fileBytes, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		fmt.Println(err3)
	}

	tempFile.Write(fileBytes)

	var newFile models.Filee

	first := sha256.New()
	first.Write([]byte(handler.Filename))
	hashPart := first.Sum(nil)
	newFile.Id = string(hashPart) + ":" + tempFile.Name()
	newFile.Location = tempFile.Name()
	fmt.Println(newFile.Id)
	fmt.Println(newFile.Location)
	mocks.Files = append(mocks.Files, newFile)
	newFileJson, err4 := json.Marshal(newFile)
	if err4 != nil {
		fmt.Println(err4)
	}

	fmt.Println(newFileJson)
	response.WriteHeader(http.StatusCreated)
	response.Header().Add("Content-Type", "application/json")
	//response.Write([]byte(`{"field_id":"newFile.Id"}`))
	response.Write(newFileJson)

	fmt.Println("Done")
	
}
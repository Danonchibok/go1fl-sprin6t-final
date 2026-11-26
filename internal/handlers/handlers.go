package handlers

import (
	"bytes"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const (
	indexHtlm      = "../index.html"
	uploadFileName = "myFile"
)

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	var buf bytes.Buffer

	indexHtlm, err := os.Open(indexHtlm)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	defer indexHtlm.Close()

	_, err = buf.ReadFrom(indexHtlm)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "text/html")
	response.WriteHeader(http.StatusOK)
	response.Write(buf.Bytes())
}

func UploadHandler(response http.ResponseWriter, request *http.Request) {
	var buf bytes.Buffer

	err := request.ParseMultipartForm(1 << 20)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	file, header, err := request.FormFile(uploadFileName)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = buf.ReadFrom(file)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	text := service.ConvertText(string(buf.Bytes()))

	newFileName := "../" + time.Now().UTC().Format("2006-01-02_15_04_05") + filepath.Ext(header.Filename)
	createdFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	defer createdFile.Close()

	_, err = createdFile.WriteString(text)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Write([]byte(text))
	response.WriteHeader(http.StatusCreated)
}

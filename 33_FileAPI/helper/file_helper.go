package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func GetCurrentDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	currentFileDir := filepath.Dir(ex)
	return currentFileDir, nil
}

func UploadFile(file *multipart.FileHeader) (newFileName string, err error) {
	source, err := file.Open()
	if err != nil {
		return "", err
	}
	defer source.Close()

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	currentdir, err := GetCurrentDir()
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf("%s/files/%s", currentdir, fileName)

	destination, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func DownloadFile(fileName string) (fileByte []byte, err error) {
	currentdir, err := GetCurrentDir()
	if err != nil {
		return nil, err
	}
	fileName = strings.ReplaceAll(fileName, "/", "")
	fileName = strings.ReplaceAll(fileName, "..", "")
	filePath := fmt.Sprintf("%s/files/%s", currentdir, fileName)
	byteArray, err := os.ReadFile(filePath)

	return byteArray, err
}

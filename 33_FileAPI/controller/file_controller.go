package controller

import (
	"encoding/json"
	"fileapi/helper"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	_, fileHeader, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{ErrorDescription: err.Error()})
		return
	}
	filePath, err := helper.UploadFile(fileHeader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{ErrorDescription: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UploadFileResponse{FileName: filePath})
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	fileName := &DownloadFileRequest{}

	err := json.NewDecoder(r.Body).Decode(fileName)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{ErrorDescription: err.Error()})
		return
	}

	if fileName.FileName == "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{ErrorDescription: "{file_name} is empty!"})
		return
	}

	file, err := helper.DownloadFile(fileName.FileName)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{ErrorDescription: err.Error()})
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename="+fileName.FileName)
	w.Header().Add("Content-Type", http.DetectContentType(file))
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}

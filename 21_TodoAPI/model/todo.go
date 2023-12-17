package model

import "errors"

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var NOT_FOUND = errors.New("Todo Not Found")

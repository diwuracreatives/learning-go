package models

import "libary-management/types"

type Book struct {
	Id     int
	Title  string
	Author string
	Status types.Status
}

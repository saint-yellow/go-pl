package service

import "github.com/saint-yellow/go-pl/toys/memorandum/serialization"

type Service interface{
	Run() *serialization.Result
}
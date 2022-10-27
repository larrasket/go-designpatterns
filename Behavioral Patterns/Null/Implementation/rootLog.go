package main

import (
	"fmt"
	"os"
)

type Log interface {
	Write(message string)
}
type CosoleLog struct{}

func (C CosoleLog) Write(message string) {
	fmt.Println(message)
}

type FileLog struct {
	LogFileName string
}

func (F FileLog) Write(message string) {
	if err := os.WriteFile(F.LogFileName, []byte(message), 0644); err != nil {
		panic(err)
	}
}

type Service struct {
	Log
}

func (S Service) Result(request string) {
	if S.Log != nil {
		S.Log.Write(request + "is received")
	}

	if S.Log != nil {
		S.Log.Write(request + "is handled")
	}
	//etc..
}

type NullLog struct{}

func main() {
	Log.Write("W")
}

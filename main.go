package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Чтение файла в байтовый массив
	fileData, err := os.ReadFile("image_2.jpg")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	for _, elem := range fileData {
		fmt.Printf("%d,", elem)
	}
	// Теперь fileData — это []byte, который можно передать в gRPC запрос
	//log.Printf("File data: %v", fileData)
}

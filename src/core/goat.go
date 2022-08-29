package core

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Goat struct{}

func (g Goat) GenerateService(name string) {
	//TODO implement me
	folderName := "service"
	fileName := strings.ToLower(name) + ".go"
	filePath := "./" + folderName + "/" + fileName
	fileData := "package " + folderName + "/" + strings.ToLower(name) + "\n"

	// Check if file already existed
	files, err := os.ReadDir("./" + folderName)
	if len(files) <= 0 {
		err := os.Mkdir("./"+folderName, 0750)
		if err != nil {
			log.Fatal("Something when wrong when create folder: ", err)
		}
	}
	_, err = os.ReadFile(filePath)
	if err == nil {
		log.Fatal("File already existed")
	}

	// create the file
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Something when wrong when create file: ", err)
	}
	// close the file with defer
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Something when wrong close file: ", err)
		}
	}(f)

	//write directly into file
	_, err = f.WriteString(fileData)
	if err != nil {
		log.Fatal("Something when wrong write file: ", err)
	}

	fmt.Println("===================================")
	fmt.Println("Generated Service: ", name)
	fmt.Println("Name: ", fileName)
	fmt.Println("Path: ", filePath)
	fmt.Println("===================================")
}

func (g Goat) GenerateGrpc() {
	//TODO implement me
	folderName := "grpc"
	fileName := "grpc.go"
	filePath := "./" + folderName + "/" + fileName
	fileData := "package " + folderName + "/" + strings.ToLower("grpc") + "\n"

	// Check if file already existed
	files, err := os.ReadDir("./" + folderName)
	if len(files) <= 0 {
		err := os.Mkdir("./"+folderName, 0750)
		if err != nil {
			log.Fatal("Something when wrong when create folder: ", err)
		}
	}
	_, err = os.ReadFile(filePath)
	if err == nil {
		log.Fatal("File already existed")
	}

	// create the file
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Something when wrong when create file: ", err)
	}
	// close the file with defer
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Something when wrong close file: ", err)
		}
	}(f)

	//write directly into file
	_, err = f.WriteString(fileData)
	if err != nil {
		log.Fatal("Something when wrong write file: ", err)
	}

	fmt.Println("===================================")
	fmt.Println("Generated Grpc:")
	fmt.Println("Name: ", fileName)
	fmt.Println("Path: ", filePath)
	fmt.Println("===================================")
}

func (g Goat) GenerateHttp() {
	//TODO implement me
	folderName := "http"
	name := "http"
	fileName := strings.ToLower(name) + ".go"
	filePath := "./" + folderName + "/" + fileName
	fileData := "package " + folderName + "/" + strings.ToLower(name) + "\n"

	// Check if file already existed
	files, err := os.ReadDir("./" + folderName)
	if len(files) <= 0 {
		err := os.Mkdir("./"+folderName, 0750)
		if err != nil {
			log.Fatal("Something when wrong when create folder: ", err)
		}
	}
	_, err = os.ReadFile(filePath)
	if err == nil {
		log.Fatal("File already existed")
	}

	// create the file
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Something when wrong when create file: ", err)
	}
	// close the file with defer
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Something when wrong close file: ", err)
		}
	}(f)

	//write directly into file
	_, err = f.WriteString(fileData)
	if err != nil {
		log.Fatal("Something when wrong write file: ", err)
	}

	fmt.Println("===================================")
	fmt.Println("Generated Http: ")
	fmt.Println("Name: ", fileName)
	fmt.Println("Path: ", filePath)
	fmt.Println("===================================")
}

func (g Goat) GenerateWebSocket() {
	//TODO implement me
	folderName := "websocket"
	name := "websocket"
	fileName := strings.ToLower(name) + ".go"
	filePath := "./" + folderName + "/" + fileName
	fileData := "package " + folderName + "/" + strings.ToLower(name) + "\n"

	// Check if file already existed
	files, err := os.ReadDir("./" + folderName)
	if len(files) <= 0 {
		err := os.Mkdir("./"+folderName, 0750)
		if err != nil {
			log.Fatal("Something when wrong when create folder: ", err)
		}
	}
	_, err = os.ReadFile(filePath)
	if err == nil {
		log.Fatal("File already existed")
	}

	// create the file
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Something when wrong when create file: ", err)
	}
	// close the file with defer
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Something when wrong close file: ", err)
		}
	}(f)

	//write directly into file
	_, err = f.WriteString(fileData)
	if err != nil {
		log.Fatal("Something when wrong write file: ", err)
	}

	fmt.Println("===================================")
	fmt.Println("Generated Websocket: ")
	fmt.Println("Name: ", fileName)
	fmt.Println("Path: ", filePath)
	fmt.Println("===================================")
}

func (g Goat) GenerateEntity(name string) {
	//TODO implement me
	folderName := "entities"
	fileName := strings.ToLower(name) + ".go"
	filePath := "./" + folderName + "/" + fileName
	fileData := "package " + folderName + "/" + strings.ToLower(name) + "\n"

	// Check if file already existed
	files, err := os.ReadDir("./" + folderName)
	if len(files) <= 0 {
		err := os.Mkdir("./"+folderName, 0750)
		if err != nil {
			log.Fatal("Something when wrong when create folder: ", err)
		}
	}
	_, err = os.ReadFile(filePath)
	if err == nil {
		log.Fatal("File already existed")
	}

	// create the file
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Something when wrong when create file: ", err)
	}
	// close the file with defer
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Something when wrong close file: ", err)
		}
	}(f)

	//write directly into file
	_, err = f.WriteString(fileData)
	if err != nil {
		log.Fatal("Something when wrong write file: ", err)
	}

	fmt.Println("===================================")
	fmt.Println("Generated Entity: ", name)
	fmt.Println("Name: ", fileName)
	fmt.Println("Path: ", filePath)
	fmt.Println("===================================")
}

func (g Goat) GenerateProtobuf(name string) {
	//TODO implement me
	folderName := "protobuf"
	fileName := strings.ToLower(name) + ".go"
	filePath := "./" + folderName + "/" + fileName
	fileData := "package " + folderName + "/" + strings.ToLower(name) + "\n"

	// Check if file already existed
	files, err := os.ReadDir("./" + folderName)
	if len(files) <= 0 {
		err := os.Mkdir("./"+folderName, 0750)
		if err != nil {
			log.Fatal("Something when wrong when create folder: ", err)
		}
	}
	_, err = os.ReadFile(filePath)
	if err == nil {
		log.Fatal("File already existed")
	}

	// create the file
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Something when wrong when create file: ", err)
	}
	// close the file with defer
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Something when wrong close file: ", err)
		}
	}(f)

	//write directly into file
	_, err = f.WriteString(fileData)
	if err != nil {
		log.Fatal("Something when wrong write file: ", err)
	}

	fmt.Println("===================================")
	fmt.Println("Generated Protobuf: ", name)
	fmt.Println("Name: ", fileName)
	fmt.Println("Path: ", filePath)
	fmt.Println("===================================")
}

func (g Goat) GenerateProtocol(name string) {
	switch name {
	case "http":
		g.GenerateHttp()
	case "grpc":
		g.GenerateGrpc()
	default:
		log.Fatal("The protocol is not valid!")
	}
}

func (g Goat) GenerateRepo(name string) {
	//TODO implement me
	folderName := "repositories"
	fileName := strings.ToLower(name) + ".go"
	filePath := "./" + folderName + "/" + fileName
	fileData := "package " + folderName + "/" + strings.ToLower(name) + "\n"

	// Check if file already existed
	files, err := os.ReadDir("./" + folderName)
	if len(files) <= 0 {
		err := os.Mkdir("./"+folderName, 0750)
		if err != nil {
			log.Fatal("Something when wrong when create folder: ", err)
		}
	}
	_, err = os.ReadFile(filePath)
	if err == nil {
		log.Fatal("File already existed")
	}

	// create the file
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Something when wrong when create file: ", err)
	}
	// close the file with defer
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Something when wrong close file: ", err)
		}
	}(f)

	//write directly into file
	_, err = f.WriteString(fileData)
	if err != nil {
		log.Fatal("Something when wrong write file: ", err)
	}

	fmt.Println("===================================")
	fmt.Println("Generated Repository: ", name)
	fmt.Println("Name: ", fileName)
	fmt.Println("Path: ", filePath)
	fmt.Println("===================================")
}

func (g Goat) GenerateDocker(name string) {
	//TODO implement me
	folderName := "docker"
	fileName := strings.ToLower(name) + ".go"
	filePath := "./" + folderName + "/" + fileName
	fileData := "package " + folderName + "/" + strings.ToLower(name) + "\n"

	// Check if file already existed
	files, err := os.ReadDir("./" + folderName)
	if len(files) <= 0 {
		err := os.Mkdir("./"+folderName, 0750)
		if err != nil {
			log.Fatal("Something when wrong when create folder: ", err)
		}
	}
	_, err = os.ReadFile(filePath)
	if err == nil {
		log.Fatal("File already existed")
	}

	// create the file
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Something when wrong when create file: ", err)
	}
	// close the file with defer
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Something when wrong close file: ", err)
		}
	}(f)

	//write directly into file
	_, err = f.WriteString(fileData)
	if err != nil {
		log.Fatal("Something when wrong write file: ", err)
	}

	fmt.Println("===================================")
	fmt.Println("Generated Dockerfile: ", name)
	fmt.Println("Name: ", fileName)
	fmt.Println("Path: ", filePath)
	fmt.Println("===================================")
}

func (g Goat) GenerateEntryPoint() {
	//TODO implement me
	folderName := "entrypoint"
	name := "entrypoint"
	fileName := strings.ToLower(name) + ".go"
	filePath := "./" + folderName + "/" + fileName
	fileData := "package " + folderName + "/" + strings.ToLower(name) + "\n"

	// Check if file already existed
	files, err := os.ReadDir("./" + folderName)
	if len(files) <= 0 {
		err := os.Mkdir("./"+folderName, 0750)
		if err != nil {
			log.Fatal("Something when wrong when create folder: ", err)
		}
	}
	_, err = os.ReadFile(filePath)
	if err == nil {
		log.Fatal("File already existed")
	}

	// create the file
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Something when wrong when create file: ", err)
	}
	// close the file with defer
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Something when wrong close file: ", err)
		}
	}(f)

	//write directly into file
	_, err = f.WriteString(fileData)
	if err != nil {
		log.Fatal("Something when wrong write file: ", err)
	}

	fmt.Println("===================================")
	fmt.Println("Generated Entrypoint:")
	fmt.Println("Name: ", fileName)
	fmt.Println("Path: ", filePath)
	fmt.Println("===================================")
}

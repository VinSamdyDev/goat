package core

import "fmt"

type Goat struct{}

func (g Goat) GenerateService(name string) {
	//TODO implement me
	fmt.Println("GenerateService:", name)
}

func (g Goat) GenerateGrpc() {
	//TODO implement me
	fmt.Println("GenerateGrpc")
}

func (g Goat) GenerateHttp() {
	//TODO implement me
	fmt.Println("GenerateHttp")
}

func (g Goat) GenerateWebSocket() {
	//TODO implement me
	fmt.Println("GenerateWebSocket")
}

func (g Goat) GenerateEntity(name string) {
	//TODO implement me
	fmt.Println("GenerateEntity:", name)
}

func (g Goat) GenerateProtobuf(name string) {
	//TODO implement me
	fmt.Println("GenerateProtobuf:", name)
}

func (g Goat) GenerateProtocol(name string) {
	//TODO implement me
	fmt.Println("GenerateProtocol:", name)
}

func (g Goat) GenerateRepo(name string) {
	//TODO implement me
	fmt.Println("GenerateRepo:", name)
}

func (g Goat) GenerateDocker(name string) {
	//TODO implement me
	fmt.Println("GenerateDocker:", name)
}

func (g Goat) GenerateEntryPoint() {
	//TODO implement me
	fmt.Println("GenerateEntryPoint")
}

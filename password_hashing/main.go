package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

// Esse exemplo mostrará como gerar hashes de senhas usando o bcrypt. Para isso, precisamos obter a lib bcrypt do Golang
// go get golang.org/x/crypto/bcrypt

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite a senha: ")
	password, _ := reader.ReadString('\n')
	password = strings.Replace(password, "\r\n", "", -1)

	hash, _ := HashPassword(password)

	fmt.Println("Senha:", password)
	fmt.Println("Hash:", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:    ", match)
}

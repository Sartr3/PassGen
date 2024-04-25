package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Função para gerar a senha
func generatePassword(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+,.?/:;{}[]~"
	rand.Seed(time.Now().UnixNano())

	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}

// Manipulador para a rota "/generate-password"
func passwordHandler(w http.ResponseWriter, r *http.Request) {
	lstr := r.URL.Query().Get("length")
	length, err := strconv.Atoi(lstr)
	if err != nil {
		http.Error(w, "Erro: o comprimento deve ser um número inteiro", http.StatusBadRequest)
		return
	}

	// Gera a senha
	password := generatePassword(length)

	// Escreve a senh
	w.Write([]byte("Senha gerada: " + password))
}

func main() {
  // Mapeia o html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/face.htm")
	})

	// Mapeia a rota
	http.HandleFunc("/generate-password", passwordHandler)

	// Inicia o servidor na porta 8080
	fmt.Println("Servidor escutando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %s\n", err)
	}
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Inicializa e carrega as variáveis de ambiente
func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}
}

// Função para acessar uma variável de ambiente
func Get(key string) string {
	return os.Getenv(key)
}

package main

import (
    "carnegestao/internal/usuarios" // puxando handler e repositorio de usuarios
	"carnegestao/internal/auth"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
    "net/http"
    "os"
    "log"
    "database/sql"
    "fmt"
)


var db *sql.DB

func main() {
	// Carrega o .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	// Carregar configurações do ambiente
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// String de conexão
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	fmt.Println("Conectando com string:", dsn)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	// Testar conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Banco de dados indisponível:", err)
	}
	fmt.Println("Conectado ao banco de dados com sucesso!")

	// Inicializa repositórios
	auth.InicializarAuth(db)
	usuarios.InicializarRepositorio(db)

	// Roteamento
	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/usuarios", usuarios.CriarUsuarioHandler)


	fmt.Println("Servidor rodando na porta 8080 🚀")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Servidor online!"))
}

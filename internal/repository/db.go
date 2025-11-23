package repository

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

func Connect() *sql.DB {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    name := os.Getenv("DB_NAME")

    if port == "" {
        port = "5432"
    }

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, name,
    )

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Cannot connect to PostgreSQL: %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("Cannot ping PostgreSQL: %v", err)
    }

    log.Println("Connected to PostgreSQL")
    return db
}

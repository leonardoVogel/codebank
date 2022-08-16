package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/leonardoVogel/codebank/domain"
	"github.com/leonardoVogel/codebank/infra/repository"
	"github.com/leonardoVogel/codebank/usecase"
	_"github.com/lib/pq"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234-1234-1234"
	cc.Name = "Leonardo"
	cc.ExpirationMonth = 2025
	cc.ExpirationYear = 11
	cc.CVV = 999
	cc.Limit = 5000
	cc.Balance = 200

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	"db",
	"5432",
	"postgres",
	"root",
	"codebank",
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}
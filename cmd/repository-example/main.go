package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/saisai/todo-api/internal"
	"github.com/saisai/todo-api/internal/postgresql"
	"github.com/saisai/todo-api/internal/service"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const DATABASE_URL_PGX = "postgres://postgres:postgres@127.0.0.1:5432/todo"

func main() {
	db, err := newDB()
	if err != nil {
		log.Fatalln("couldn't instantiate db", err)
	}
	defer db.Close()

	//-

	repo := postgresql.NewTask(db) // Task Repository
	svc := service.NewTask(repo)   // Task Application Service
	task, err := svc.Create(context.Background(), "new task", internal.PriorityLow, internal.Dates{})

	fmt.Printf("NEW task %#v, err %s\n", task, err)

	//-

	if err := svc.Update(context.Background(),
		task.ID,
		"changed task",
		internal.PriorityHigh,
		internal.Dates{
			Due: time.Now().Add(2 * time.Hour),
		},
		false); err != nil {
		log.Fatalln("couldn't update task", err)
	}

	updatedTask, err := svc.Task(context.Background(), task.ID)
	if err != nil {
		log.Fatalln("couldn't find task", err)
	}

	fmt.Printf("UPDATED task %#v, err %s\n", updatedTask, err)

}

func newDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", DATABASE_URL_PGX)
	if err != nil {
		return db, fmt.Errorf("db open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return db, fmt.Errorf("db ping: %w", err)
	}

	return db, nil
}

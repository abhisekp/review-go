package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(ctx context.Context) (*sql.DB, error) {
	postgresURL := os.Getenv("POSTGRES_URL")
	fmt.Println("Postgres URL:", postgresURL)

	poolConfig, err := pgxpool.ParseConfig(postgresURL)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDBFromPool(pool)

	err = db.PingContext(ctx)
	if err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("could not ping db: %w", err)
	}

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-exitCh
		log.Printf("received signal %v, shutting down...", sig)
		err := db.Close()
		if err != nil {
			fmt.Println("Could not close the connection")
		}

		signal.Stop(exitCh)
		os.Exit(0)
	}()

	err = InitPreparedStatements(ctx, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

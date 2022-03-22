package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest/libraries/config"
	"rest/libraries/database"
	"rest/router"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if _, ok := os.LookupEnv("APP_ENV"); !ok {
		config.Setup(".env")
	}

	log := log.New(os.Stdout, "rest-skeleton : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(log); err != nil {
		log.Fatalf("error: shutting down: %s", err)
	}
}

func run(log *log.Logger) error {
	log.Printf("main : Started")
	defer log.Println("main : Completed")

	db, err := database.OpenDB()
	if err != nil {
		return fmt.Errorf("connecting to db: %v", err)
	}
	defer db.Close()

	// parameter server
	server := http.Server{
		Addr:         os.Getenv("APP_PORT"),
		Handler:      router.API(db, log),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Println("server listening on", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Membuat channel untuk mendengarkan sinyal interupsi/terminate dari OS.
	// Menggunakan channel buffered karena paket signal membutuhkannya.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Mengontrol penerimaan data dari channel,
	// jika ada error saat listenAndServe server maupun ada sinyal shutdown yang diterima
	select {
	case err := <-serverErrors:
		return fmt.Errorf("listening and serving: %s", err)

	case <-shutdown:
		log.Println("caught signal, shutting down")

		// Jika ada shutdown, meminta tambahan waktu 5 detik untuk menyelesaikan proses yang sedang berjalan.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("error: gracefully shutting down server: %s", err)
			if err := server.Close(); err != nil {
				return fmt.Errorf("error: closing server: %s", err)
			}
		}
	}

	log.Println("done")
	return nil
}

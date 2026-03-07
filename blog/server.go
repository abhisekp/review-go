package blog

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"path/filepath"
	postHandlers "review-go/blog/handlers/posts"
	dbService "review-go/blog/services/db"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/joho/godotenv"
)

var isShuttingDown atomic.Bool

var serverCtx, cancelServer = context.WithCancel(context.Background())

func init() {
	// Current file path with .env
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("Error Getting .env file from the current dir")
	}

	dir := filepath.Dir(filename)
	if err := godotenv.Load(filepath.Join(dir, ".env")); err != nil {
		log.Fatalln("error loading .env from the current directory of the server file")
	}
}

func Run() {
	Server(4560)
}

func statusHandler(w http.ResponseWriter, _ *http.Request) {
	if isShuttingDown.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, err := w.Write([]byte("Server is shutting down"))
		if err != nil {
			fmt.Println(fmt.Errorf("status error: %w", err))
		}
		return
	}

	_, err := fmt.Fprintf(w, "OK")
	if err != nil {
		fmt.Println(fmt.Errorf("status error: %w", err))
	}
}

func Sleep(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func Server(port int) {
	dbContext, cancelDB := context.WithCancel(serverCtx)
	_ = cancelDB

	db, err := dbService.InitDB(dbContext)
	if err != nil {
		log.Fatalln(fmt.Errorf("getting sql.DB failed: %w", err))
	}
	_ = db

	if createPostsTable, ok := dbService.PrepStmts.Statements["createPostsTable"]; ok {
		_, err := createPostsTable.ExecContext(dbContext)
		if err != nil {
			fmt.Println(fmt.Errorf("could not create posts table: %w", err))
			return
		}
	}

	if createUsersTable, ok := dbService.PrepStmts.Statements["createUsersTable"]; ok {
		_, err := createUsersTable.ExecContext(dbContext)
		if err != nil {
			fmt.Println(fmt.Errorf("could not create users table: %w", err))
			return
		}
	}

	postHandler := postHandlers.NewPostsHandler(postHandlers.PostsHandlerOptions{
		DB: db,
	})

	postsHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			postHandler.AddPost(w, r)
		} else if r.Method == "GET" {
			id := r.PathValue("id")
			if id != "" {
				postHandler.GetById(w, r)
				return
			} else {
				// TODO: implement pagination and filtering
				// postHandler.GetAll(w, r)
				return
			}
		} else {
			w.Header().Add("Allow", "POST")
			w.Header().Add("Allow", "GET")
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, err = w.Write([]byte("Method Not Allowed"))
			if err != nil {
				fmt.Println(fmt.Errorf("could not write response: %w", err))
				return
			}
			return
		}

	}

	http.HandleFunc("/", statusHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/posts/{id}", postsHandler)

	fmt.Printf("Running in http://localhost:%d\n", port)

	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
		BaseContext: func(_ net.Listener) context.Context {
			return serverCtx
		},
	}
	log.Fatalln(server.ListenAndServe())
}

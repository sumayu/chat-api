package main

import (
 "log"
 "net/http"
 "os"
 "time"

 "yourmodule/internal/db"
 "yourmodule/internal/httpapi"
 "yourmodule/internal/httpapi/handlers"
 "yourmodule/internal/repo"
)

func main() {
 dsn := os.Getenv("POSTGRES_DSN")
 if dsn == "" {
  log.Fatal("POSTGRES_DSN is required")
 }

 gdb, err := db.Open(dsn)
 if err != nil {
  log.Fatal(err)
 }

 chatRepo := repo.NewChatRepo(gdb)
 msgRepo := repo.NewMessageRepo(gdb)

 ch := handlers.NewChatHandler(chatRepo, msgRepo)
 mh := handlers.NewMessageHandler(chatRepo, msgRepo)

 srv := &http.Server{
  Addr:              ":8080",
  Handler:           httpapi.NewRouter(httpapi.Deps{Chats: ch, Messages: mh}),
  ReadHeaderTimeout: 5 * time.Second,
 }

 log.Println("listening on :8080")
 log.Fatal(srv.ListenAndServe())
}

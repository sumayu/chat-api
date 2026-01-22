package httpapi

import (
 "net/http"

 "yourmodule/internal/httpapi/handlers"
)

type Deps struct {
 Chats    *handlers.ChatHandler
 Messages *handlers.MessageHandler
}

func NewRouter(d Deps) http.Handler {
 mux := http.NewServeMux()

 mux.HandleFunc("POST /chats/", d.Chats.CreateChat)
 mux.HandleFunc("GET /chats/{id}", d.Chats.GetChat)
 mux.HandleFunc("DELETE /chats/{id}", d.Chats.DeleteChat)
 mux.HandleFunc("POST /chats/{id}/messages/", d.Messages.CreateMessage)

 return mux
}

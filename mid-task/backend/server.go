package server

import (
	"fmt"

	"github.com/yutaro1031/treasure-app/mid-task/backend/sample"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/justinas/alice"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	"github.com/yutaro1031/treasure-app/mid-task/backend/controller"
	db2 "github.com/yutaro1031/treasure-app/mid-task/backend/db"
	"github.com/yutaro1031/treasure-app/mid-task/backend/middleware"
)

type Server struct {
	db     *sqlx.DB
	router *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) {
	db := db2.NewDB(datasource)
	dbcon, err := db.Open()
	if err != nil {
		log.Fatalf("failed db init. %s", err)
	}
	s.db = dbcon
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/public").Handler(commonChain.Then(sample.NewPublicHandler()))

	bookController := controller.NewBook(s.db)
	r.Methods(http.MethodGet).Path("/books").Handler(commonChain.Then(AppHandler{bookController.Index}))
	r.Methods(http.MethodPost).Path("/books").Handler(commonChain.Then(AppHandler{bookController.Create}))
	r.Methods(http.MethodDelete).Path("/books/{id}").Handler(commonChain.Then(AppHandler{bookController.Destroy}))
	r.Methods(http.MethodGet).Path("/search").Handler(commonChain.Then(AppHandler{bookController.Search}))

	tagController := controller.NewTag(s.db)
	r.Methods(http.MethodGet).Path("/tags").Handler(commonChain.Then(AppHandler{tagController.Index}))
	r.Methods(http.MethodPost).Path("/tags").Handler(commonChain.Then(AppHandler{tagController.Create}))
	r.Methods(http.MethodPut).Path("/tags/{id}").Handler(commonChain.Then(AppHandler{tagController.Update}))
	r.Methods(http.MethodDelete).Path("/tags/{id}").Handler(commonChain.Then(AppHandler{tagController.Destroy}))

	r.PathPrefix("").Handler(commonChain.Then(http.StripPrefix("/img", http.FileServer(http.Dir("./img")))))
	return r
}

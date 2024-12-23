package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Rotas(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/tarefas", CreateTarefaHandle(db))
	r.GET("/tarefas", ReadTarefasHandler(db))
	r.PUT("/tarefas/:id", UpdateTarefaHandler(db))
	r.DELETE("/tarefas/:id", DeleteTarefaHandler(db))

	return r
}

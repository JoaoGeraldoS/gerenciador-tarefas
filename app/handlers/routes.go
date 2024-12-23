package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Rotas(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/tarefas", CreateTarefaHandle(db))

	return r
}

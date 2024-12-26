package handlers

import (
	"database/sql"

	cat "github.com/JoaoGeraldoS/gerenciador-tarefas/app/handlers/categoria"
	"github.com/gin-gonic/gin"
)

func Rotas(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Rotas das tarefas
	r.POST("/tarefas", CreateTarefaHandle(db))
	r.GET("/tarefas", ReadTarefasHandler(db))
	r.PUT("/tarefas/:id", UpdateTarefaHandler(db))
	r.DELETE("/tarefas/:id", DeleteTarefaHandler(db))

	// Rotas das Categorias

	r.GET("/categorias", cat.ReadCategoriaHandler(db))
	r.POST("/categorias", cat.CreateCategoriaHandler(db))
	r.PUT("/categorias", cat.UpdateCategoriaHandler(db))
	r.DELETE("/categorias", cat.DeleteCategoriaHandler(db))

	// Rotas de Filtro
	r.GET("/tarefas/filtro", FiltroTarefaHandler(db))

	return r
}

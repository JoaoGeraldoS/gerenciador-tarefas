package handlers

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/controllers"
	"github.com/gin-gonic/gin"
)

func DeleteTarefaHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		if err := controllers.DeleteTarefa(id, db); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"erro": "erro ao apagar tarefa"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "tarefa apagada!"})
	}
}

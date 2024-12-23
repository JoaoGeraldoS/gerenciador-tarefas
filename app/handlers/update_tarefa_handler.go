package handlers

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/controllers"
	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
	"github.com/gin-gonic/gin"
)

func UpdateTarefaHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var tarefa models.Tarefa

		if err := ctx.ShouldBindJSON(&tarefa); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"erro": "erro ao atualizar dados"})
			return
		}

		response, err := controllers.UpdateTarefa(id, db, &tarefa)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "tarefa atualizada!", "tarefa": response})
	}
}

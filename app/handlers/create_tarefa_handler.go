package handlers

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/controllers"
	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
	"github.com/gin-gonic/gin"
)

func CreateTarefaHandle(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tarefa models.Tarefa

		if err := ctx.ShouldBindJSON(&tarefa); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"erro": "erro ao criar tarefa"})
			return
		}

		tarefas, err := controllers.CreateTarefa(db, &tarefa)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Tarefa criada com sucesso", "tarefa": tarefas})

	}
}

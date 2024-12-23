package handlers

import (
	"database/sql"
	"net/http"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/controllers"
	"github.com/gin-gonic/gin"
)

func ReadTarefasHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tarefas, err := controllers.ReadTarefas(db)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"erro": "erro ao encotrar dados"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"tarefas": tarefas})
	}
}

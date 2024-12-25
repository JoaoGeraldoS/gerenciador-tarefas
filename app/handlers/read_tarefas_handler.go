package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/controllers"
	"github.com/gin-gonic/gin"
)

func ReadTarefasHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// tarefas, err := controllers.ReadTarefas(db)
		tarefas, err := controllers.ReadIntermediaria(db)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"erro": "erro ao encotrar dados"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"tarefas": tarefas})
	}
}

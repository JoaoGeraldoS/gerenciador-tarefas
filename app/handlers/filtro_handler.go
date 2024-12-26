package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/controllers"
	"github.com/gin-gonic/gin"
)

func FiltroTarefaHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Query("status")
		categoria := ctx.Query("categoria")

		tarefas, err := controllers.FiltroTarefa(status, categoria, db)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"erro": "Filtro não pode ser aplicado"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"Tarefas": tarefas})

		fmt.Println(status, categoria)
	}
}

package categoria

import (
	"database/sql"
	"net/http"

	cat "github.com/JoaoGeraldoS/gerenciador-tarefas/app/controllers/categorias"
	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
	"github.com/gin-gonic/gin"
)

func CreateCategoriaHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var categoria models.Categoria

		if err := ctx.ShouldBindJSON(&categoria); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"erro": "erro ao adicionar categoria!"})
			return
		}

		response, err := cat.CreateCategoria(db, &categoria)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro interno do servidor"})
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": "Categoria criada", "categoria": response})
	}
}

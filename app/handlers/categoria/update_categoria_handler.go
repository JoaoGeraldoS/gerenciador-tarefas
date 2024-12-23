package categoria

import (
	"database/sql"
	"net/http"

	cat "github.com/JoaoGeraldoS/gerenciador-tarefas/app/controllers/categorias"
	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
	"github.com/gin-gonic/gin"
)

func UpdateCategoriaHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var categoria models.Categoria

		if err := ctx.ShouldBindJSON(&categoria); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"erro": "erro ao atualizar categoria"})
		}

		response, err := cat.UpdateCategoria(id, db, &categoria)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "erro interno do servidor"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "categoria atualizada!", "categoria": response})
	}
}

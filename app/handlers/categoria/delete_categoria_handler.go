package categoria

import (
	"database/sql"
	"net/http"

	cat "github.com/JoaoGeraldoS/gerenciador-tarefas/app/controllers/categorias"
	"github.com/gin-gonic/gin"
)

func DeleteCategoriaHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if err := cat.DeleteCategoria(id, db); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"erro": "erro ao apagar dados"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "categoria apagada."})
	}
}

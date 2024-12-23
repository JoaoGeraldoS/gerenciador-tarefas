package categoria

import (
	"database/sql"
	"log"
	"net/http"

	cat "github.com/JoaoGeraldoS/gerenciador-tarefas/app/controllers/categorias"
	"github.com/gin-gonic/gin"
)

func ReadCategoriaHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := cat.ReadCategoria(db)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusNotFound, gin.H{"erro": "erro ao encotrar dados"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"categorias": response})
	}
}

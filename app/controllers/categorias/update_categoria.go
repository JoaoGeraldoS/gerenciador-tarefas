package categoria

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
)

func UpdateCategoria(id string, db *sql.DB, categoria *models.Categoria) (*models.Categoria, error) {
	update := `UPDATE categorias SET nome = ? WHERE id = ?`

	response, err := db.Exec(update, categoria.Nome, id)
	if err != nil {
		return nil, fmt.Errorf("Erro ao autualizar categoria! %v", err.Error())
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return nil, fmt.Errorf("erro: %v", err.Error())
	}

	return categoria, nil
}

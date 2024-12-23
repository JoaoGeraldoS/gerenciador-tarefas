package categoria

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
)

func CreateCategoria(db *sql.DB, categoria *models.Categoria) (*models.Categoria, error) {
	inserte := `INSERT INTO categorias (nome, descricao) VALUES (?, ?)`
	response, err := db.Exec(inserte, categoria.Nome, categoria.Descricao)
	if err != nil {
		return nil, fmt.Errorf("Erro ao salvar os dados")
	}

	id, err := response.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontrar o id")
	}

	categoria.ID = id
	return categoria, nil
}

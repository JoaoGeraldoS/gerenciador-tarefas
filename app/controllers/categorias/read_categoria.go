package categoria

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
)

func ReadCategoria(db *sql.DB) ([]models.Categoria, error) {
	query, err := db.Query(`SELECT id, nome, descricao FROM categorias`)
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar dados. %v", err)
	}
	defer query.Close()

	var categorias []models.Categoria

	for query.Next() {
		var categoria models.Categoria

		err = query.Scan(&categoria.ID, &categoria.Nome)
		if err != nil {
			return nil, fmt.Errorf("Erro no mapeamento de dados. %v", err)
		}
		categorias = append(categorias, categoria)
	}

	if err := query.Err(); err != nil {
		return nil, fmt.Errorf("Erro ao retornar dados. %v", err)
	}
	return categorias, nil

}

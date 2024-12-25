package utils

import (
	"database/sql"
	"fmt"
)

func IDsCategories(tx *sql.Tx, nomes []string) ([]int64, error) {
	ids := []int64{}
	query := `SELECT id FROM categorias WHERE nome = ?`

	for _, nome := range nomes {
		var id int64

		err := tx.QueryRow(query, nome).Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("Categoria '%s' não encontrada", nome)
			}
			return nil, fmt.Errorf("Erro ao buscar ID da categoria '%s': %v", nome, err)
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func InserirRelacionamento(tx *sql.Tx, tarefaID int64, categoriaIDs []int64) error {
	query := `INSERT INTO intermediaria(tarefa_id, categoria_id) VALUES (?, ?)`

	for _, categoriaID := range categoriaIDs {
		_, err := tx.Exec(query, tarefaID, categoriaID)
		if err != nil {
			return fmt.Errorf("Erro ao inserir relacionamento tarefa-categoria: %v", err)
		}
	}
	return nil
}

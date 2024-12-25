package validates

import (
	"database/sql"
	"fmt"
)

func ValidaCategoria(db *sql.DB, cataegoriaNome string) (int64, error) {
	var categoriaID int64
	query := `SELECT id FROM categorias WHERE nome = ?`

	err := db.QueryRow(query, cataegoriaNome).Scan(&categoriaID)
	if err != nil {
		if err == sql.ErrNoRows {
			inserte := `INSERT INTO categorias (nome) VALUES (?)`
			response, err := db.Exec(inserte, cataegoriaNome)
			if err != nil {
				return 0, fmt.Errorf("Erro ao inserir a categoria: %v", err)
			}

			categoriaID, err = response.LastInsertId()
			if err != nil {
				return 0, fmt.Errorf("Erro ao obter o ID da categoria inserida: %v", err)
			}
			return categoriaID, nil
		}
		return 0, fmt.Errorf("Erro ao verificar a categoria: %v", err)
	}
	return categoriaID, nil
}

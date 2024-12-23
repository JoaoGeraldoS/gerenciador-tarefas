package categoria

import (
	"database/sql"
	"log"
)

func DeleteCategoria(id string, db *sql.DB) error {
	delete := `DELETE FORM categorias WHERE id = ?`

	response, err := db.Exec(delete, id)
	if err != nil {
		log.Println("Erro ao apagar tarefa")
		return nil
	}

	rowAffected, err := response.RowsAffected()
	if err != nil || rowAffected == 0 {
		log.Println("Tarefa não encontrada")
		return nil
	}
	return nil
}

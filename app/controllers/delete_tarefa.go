package controllers

import (
	"database/sql"
	"log"
)

func DeleteTarefa(id string, db *sql.DB) error {
	delete := `DELETE FROM tarefas WHERE id = ?`

	response, err := db.Exec(delete, id)
	if err != nil {
		log.Println("Erro ao apagar tarefa")
		return nil
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil || rowsAffected == 0 {
		log.Println("Tarefa não encontrada")
		return nil
	}
	return nil
}

package controllers

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
)

func UpdateTarefa(id string, db *sql.DB, tarefa *models.Tarefa) (*models.Tarefa, error) {
	update := `UPDATE tarefas SET titulo = ?, descricao = ?, status = ? WHERE id = ?`

	response, err := db.Exec(update, tarefa.Titulo, tarefa.Descricao, tarefa.Status, id)
	if err != nil {
		return nil, fmt.Errorf("Erro a atualizar tarefa")
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return nil, fmt.Errorf("Erro, tarefa não encontrada")
	}
	return tarefa, nil
}

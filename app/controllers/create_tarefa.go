package controllers

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
)

func CreateTarefa(db *sql.DB, tarefa *models.Tarefa) (*models.Tarefa, error) {
	adiciona := `INSERT INTO tarefas(titulo, descricao, status) VALUES (?, ?, ?)`
	response, err := db.Exec(adiciona, tarefa.Titulo, tarefa.Descricao, tarefa.Status)
	if err != nil {
		return nil, fmt.Errorf("Erro ao inserir dados!")
	}

	id, err := response.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontar id!")
	}

	tarefa.ID = id
	return tarefa, nil

}

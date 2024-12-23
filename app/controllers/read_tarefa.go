package controllers

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
)

func ReadTarefas(db *sql.DB) ([]models.Tarefa, error) {
	query, err := db.Query(`SELECT id, titulo, descricao, status, DATE_FORMAT(data_criacao, '%d/%m/%y %H:%i:%s') AS data_criacao FROM tarefas`)
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar dados. %v", err)
	}
	defer query.Close()

	var tarefas []models.Tarefa

	for query.Next() {
		var tarefa models.Tarefa

		err = query.Scan(&tarefa.ID, &tarefa.Titulo, &tarefa.Descricao, &tarefa.Status, &tarefa.DataCriacao)
		if err != nil {
			return nil, fmt.Errorf("Erro no mapeamento de dados. %v", err)
		}
		tarefas = append(tarefas, tarefa)
	}

	if err := query.Err(); err != nil {
		return nil, fmt.Errorf("Erro ao retornar dados. %v", err)
	}
	return tarefas, nil
}

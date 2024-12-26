package controllers

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
)

func FiltroTarefa(status, categoria string, db *sql.DB) ([]models.Tarefa, error) {
	busca := `select t.id, t.titulo,t.descricao, t.status, t.data_criacao, c.id, c.nome from tarefas t join intermediaria i on i.tarefa_id = t.id
				join categorias c on i.categoria_id = c.id where t.status like ? and c.nome like ?;`

	fmt.Printf("Status: %s, NomeCategoria: %s\n", status, categoria)

	paramStatus := fmt.Sprintf("%s%%", status)
	paramCategoria := fmt.Sprintf("%s%%", categoria)

	rows, err := db.Query(busca, paramStatus, paramCategoria)
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar tarefa: %w", err)
	}

	if !rows.Next() {
		return nil, fmt.Errorf("Nenhuma tarefa encontrada")
	}

	defer rows.Close()

	var tarefas []models.Tarefa

	for rows.Next() {
		var tarefa models.Tarefa
		var categoria models.Categoria

		err := rows.Scan(&tarefa.ID, &tarefa.Titulo, &tarefa.Descricao, &tarefa.Status, &tarefa.DataCriacao, &categoria.ID, &categoria.Nome)
		if err != nil {
			return nil, fmt.Errorf("Erro ao escanear resultados: %w", err)
		}

		exists := false
		for i := range tarefas {
			if tarefas[i].ID == tarefa.ID {
				tarefas[i].Categorias = append(tarefas[i].Categorias, categoria)
				exists = true
				break
			}
		}

		if !exists {
			tarefa.Categorias = []models.Categoria{categoria}
			tarefas = append(tarefas, tarefa)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre os resultados: %w", err)
	}

	return tarefas, nil

}

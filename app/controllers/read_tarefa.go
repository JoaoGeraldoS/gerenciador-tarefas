package controllers

import (
	"database/sql"
	"log"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
)

func ReadIntermediaria(db *sql.DB) ([]models.Tarefa, error) {
	tasksMap := make(map[int64]*models.Tarefa)
	rows, err := db.Query(`select t.id, t.titulo, t.descricao, t.status, DATE_FORMAT(t.data_criacao, '%d/%m/%y %H:%i:%s')  AS data_criacao, c.id, c.nome 
		from intermediaria i 
		join tarefas t on i.tarefa_id = t.id 
		join categorias c on i.categoria_id = c.id;`)

	if err != nil {
		log.Fatalln("Erro ao buscar dados")
	}
	defer rows.Close()

	for rows.Next() {
		var (
			tarefaID, categoriaID                  int64
			titulo, descricao, status, dataCriacao string
			categoriaNome                          string
		)

		err := rows.Scan(&tarefaID, &titulo, &descricao, &status, &dataCriacao, &categoriaID, &categoriaNome)
		if err != nil {
			log.Fatalln("Erro ao mapear dados!")
		}

		if _, exists := tasksMap[tarefaID]; !exists {
			tasksMap[tarefaID] = &models.Tarefa{
				ID:          tarefaID,
				Titulo:      titulo,
				Descricao:   descricao,
				Status:      status,
				DataCriacao: dataCriacao,
				Categorias:  []models.Categoria{},
			}
		}

		tasksMap[tarefaID].Categorias = append(tasksMap[tarefaID].Categorias, models.Categoria{
			ID:   categoriaID,
			Nome: categoriaNome,
		})
	}

	var tarefas []models.Tarefa
	for _, tarefa := range tasksMap {
		tarefas = append(tarefas, *tarefa)
	}

	return tarefas, nil

}

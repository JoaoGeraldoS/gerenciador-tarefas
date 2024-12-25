package controllers

import (
	"database/sql"
	"fmt"

	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/models"
	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/utils"
	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/validates"
)

func CreateTarefa(db *sql.DB, tarefa *models.Tarefa, categorias []string) (*models.Tarefa, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("Erro ao iniciar transação: %v", err)
	}

	adiciona := `INSERT INTO tarefas(titulo, descricao, status) VALUES (?, ?, ?)`

	response, err := db.Exec(adiciona, tarefa.Titulo, tarefa.Descricao, tarefa.Status)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("Erro ao inserir dados!")
	}

	id, err := response.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("Erro ao encontar id!")
	}

	tarefa.ID = id

	for _, categoria := range tarefa.Categorias {
		_, err := validates.ValidaCategoria(db, categoria.Nome)
		if err != nil {
			return nil, fmt.Errorf("Erro ao verificar ou criar categoria: %v", err)
		}
	}

	categoriaIDs, err := utils.IDsCategories(tx, categorias)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := utils.InserirRelacionamento(tx, tarefa.ID, categoriaIDs); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("Erro ao confirmar transação: %v", err)
	}

	return tarefa, nil

}

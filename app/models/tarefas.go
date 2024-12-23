package models

type Tarefa struct {
	ID          int64     `json:"id,omitempty"`
	Titulo      string    `json:"titulo"`
	Descricao   string    `json:"descricao"`
	Status      string    `json:"status"`
	DataCriacao string    `json:"data_criacao"`
	Categoria   Categoria `json:"categorias"`
}

type Categoria struct {
	ID        int64  `json:"id,omitempty"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
}

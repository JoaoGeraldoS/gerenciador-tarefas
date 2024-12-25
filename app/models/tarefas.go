package models

type Tarefa struct {
	ID          int64       `json:"id,omitempty"`
	Titulo      string      `json:"titulo"`
	Descricao   string      `json:"descricao"`
	Status      string      `json:"status"`
	DataCriacao string      `json:"data_criacao"`
	Categorias  []Categoria `json:"categorias"`
}

type Intermediaria struct {
	ID           int64 `json:"id,omitempty"`
	Tarefa_id    int64 `json:"tarefa_id"`
	Categoria_id int64 `json:"categoria_id"`
}

type Categoria struct {
	ID   int64  `json:"id,omitempty"`
	Nome string `json:"nome"`
}

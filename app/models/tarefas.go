package models

type Tarefa struct {
	ID          int64  `json:"id,omitempty"`
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	Status      string `json:"status"`
	DataCriacao string `json:"data_criacao"`
}

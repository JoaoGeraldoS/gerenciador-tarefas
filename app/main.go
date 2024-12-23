package main

import (
	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/config"
	"github.com/JoaoGeraldoS/gerenciador-tarefas/app/handlers"
)

func main() {
	db := config.Conection()
	defer db.Close()

	r := handlers.Rotas(db)
	r.Run(":8080")
}

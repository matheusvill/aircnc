package storage

type Storage interface {
	// inserir os dados e voltar a chave do storage
	InsertUser(email, password string) int
	// pegar todos os dados do user pelo id (que é chave)
	GetUser(id int) map[string]interface{}
	// atualizar todos os dados do user
	UpdateUser(email, password string) bool

	// inserir os dados e voltar a chave do storage
	InsertHome(name, address, country, city string, user int) int
	// pegar todos os dados do home pelo id
	GetHome(id string) map[string]interface{}
	// atualizar todos os dados do home
	UpdateHome(email, password string) bool
}

package repository

type Repositoy interface {
	create(interface{})
	update(interface{})
	delete(interface{})
}

package controllers

import "FORMULARIO_WEB_MONGODB_GO/db"

type Subscription struct {
	Name  string
	Email string
}

func Create(name string, email string) error {
	s := Subscription{name, email}
	// ou Subscription{Name: name, Email: email} // dessa forma é mais segura pois valida a edição do type

	// retorna o erro da iserção definida no db.save
	// o nome da colection newsletter
	return db.Insert("newsletter", s)
}

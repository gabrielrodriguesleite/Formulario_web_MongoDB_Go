package handlers

import (
	"FORMULARIO_WEB_MONGODB_GO/controllers"
	"fmt"
	"net/http"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {

	// DEFINE A ROTA POST (em express seria o router.post(...))
	if r.Method == http.MethodPost {

		// CARREGA OS DADOS DO FORM
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "erro ao fazer parse do form: %v", err)
			return
		}

		// chama o controller passando os dados parseados do forms
		err := controllers.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))
		if err != nil {
			fmt.Fprintf(w, "erro ao salvar os dados: %v", err)
			return
		}
	}

	// DEFINE O FRONTEND COM O TEMPLATE DE FORMULARIO
	http.ServeFile(w, r, "handlers/templates/subscription_form.html")
}

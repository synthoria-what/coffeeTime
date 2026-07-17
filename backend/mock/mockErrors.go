package mock

import (
	"net/http"

	"synthori.space/coffeeTime/internal/services"
)

func SendErrorNotFound(w http.ResponseWriter, r *http.Request) {
	services.WriteError(w, http.StatusForbidden, "you dont have admin permission")
}

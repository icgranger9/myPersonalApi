package internal

import (
    "net/http"
)

type Controller struct {
    Router *http.ServeMux
    Mw     Middleware
}

type AddressController struct {
    *Controller
}

// lives inside /api
func (c Controller) Handle(route string) {
    c.Router.Handle(route+"/user/", c.Mw(GetUserHandler))
}



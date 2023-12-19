package handler

func(h *Handler) CreateUser() {
	h.serv.User.CreateUser("Hi")
}
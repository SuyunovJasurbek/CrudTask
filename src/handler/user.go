package handler

import "fmt"

func(h *Handler) CreateUser() {
	fmt.Println("Assalomu Alaykum")
	h.service.CreateUser()
}
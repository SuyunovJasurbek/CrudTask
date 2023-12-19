package handler

import "github.com/SuyunovJasurbek/CrudTask/src/service"

type Handler struct {
	serv *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{serv: service}
}

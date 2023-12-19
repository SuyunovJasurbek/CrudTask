package handler

import "github.com/SuyunovJasurbek/CrudTask/src/service"

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

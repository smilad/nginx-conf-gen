package controller

import (
	"nginx/service_contract"
)

type AdminHandler struct {
	uc service_contract.IDomainService
}

func NewAdminHandler(uc service_contract.IDomainService) AdminHandler {
	return AdminHandler{uc: uc}
}

package lb

import (
	"testappwink/internal/service"
	gen "testappwink/pkg/lb_v1"
)

type Implementation struct {
	gen.UnimplementedLBV1Server
	lbService service.LBService
}

func NewImplementation(lbService service.LBService) *Implementation {
	return &Implementation{
		lbService: lbService,
	}
}

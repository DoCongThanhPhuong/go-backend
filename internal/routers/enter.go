package routers

import (
	"github.com/DoCongThanhPhuong/go-backend/internal/routers/manager"
	"github.com/DoCongThanhPhuong/go-backend/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var AppRouterGroup = new(RouterGroup)
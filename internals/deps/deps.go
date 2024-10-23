package deps

import (
	"htmx-go/internals/repository"
)

type Dependencies struct {
	OrderRepo repository.OrderRepository
	UserRepo  repository.UserRepository
}

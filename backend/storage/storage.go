package storage

import "github.com/CrossStack-Q/SphereOne/backend/types"

type Storage interface {
	Get(int) *types.User
}

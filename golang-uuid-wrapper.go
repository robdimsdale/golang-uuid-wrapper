package golang_uuid_wrapper

import "github.com/google/uuid"

type UUID uuid.UUID

func RandomUUID() (UUID, error) {
	return uuid.NewRandom()
}

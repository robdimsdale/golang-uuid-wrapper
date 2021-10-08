package uuid

import guuid "github.com/google/uuid"

type UUID guuid.UUID

func RandomUUID() (UUID, error) {
	u, err := guuid.NewRandom()
	if err != nil {
		return UUID{}, err
	}

	return UUID(u), nil
}

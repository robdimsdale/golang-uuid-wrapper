package uuid

import guuid "github.com/google/uuid"

type UUID guuid.UUID

func NewRandom() (UUID, error) {
	u, err := guuid.NewRandom()
	if err != nil {
		return UUID{}, err
	}

	return UUID(u), nil
}

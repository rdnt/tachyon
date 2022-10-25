package uuid

import (
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v3"
)

type UUID uuid.UUID

func (id UUID) String() string {
	return shortuuid.DefaultEncoder.Encode(uuid.UUID(id))
}

func Parse(id string) (UUID, error) {
	uid, err := shortuuid.DefaultEncoder.Decode(id)
	if err != nil {
		return UUID{}, err
	}

	return UUID(uid), nil
}

func New() UUID {
	return UUID(uuid.New())
}

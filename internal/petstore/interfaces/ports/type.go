package api

type Pet struct {
	name string
}

type PetStore struct {
	Pets   map[int64]Pet
	NextId int64
}

func NewPetStore() *PetStore {

	return &PetStore{
		Pets:   make(map[int64]Pet),
		NextId: 1000,
	}
}

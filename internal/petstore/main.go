package main

import api "backend/internal/petstore/interfaces/ports"

func main() {
	petStore := api.NewPetStore()
	_ = petStore
}

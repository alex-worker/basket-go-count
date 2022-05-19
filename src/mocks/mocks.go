package mocks

import (
	"syreclabs.com/go/faker"
)

func GetConnectionStringRealDb() string {
	return "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"
}

func GetConnectionString() string {
	return faker.RandomString(10)
}

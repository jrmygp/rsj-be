package services

import (
	customerRepositories "server/repositories/customer"
	repositories "server/repositories/invoice"
)

type service struct {
	repository         repositories.Repository
	customerRepository customerRepositories.Repository
}

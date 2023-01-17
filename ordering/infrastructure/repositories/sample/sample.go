package sample

import "github.com/JiangTaoShi/eShopOnContainers/ordering/infrastructure/mysql"

type ISampleRepository interface {
	GetList()
	Add() error
}

type SampleRepository struct {
	db mysql.Repo
}

func New() ISampleRepository {
	return &SampleRepository{db: nil}
}

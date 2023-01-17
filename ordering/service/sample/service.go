package sample

import (
	"github.com/JiangTaoShi/eShopOnContainers/ordering/infrastructure/repositories/sample"
)

type SampleService struct {
	Reposity sample.ISampleRepository
}

func New() SampleService {
	return SampleService{Reposity: sample.New()}
}

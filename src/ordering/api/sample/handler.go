package sample

import (
	"github.com/JiangTaoShi/eShopOnContainers/ordering/service/sample"
)

type handler struct {
	SampleService sample.SampleService
}

func New() sample.SampleService {
	return sample.SampleService{}
}

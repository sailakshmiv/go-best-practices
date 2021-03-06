// This file was generated by counterfeiter
package apifakes

import (
	"sync"

	"github.com/tjarratt/go-best-practices/api"
	"github.com/tjarratt/go-best-practices/domain"
)

type FakePizzaRepository struct {
	MakePizzaStub        func(domain.DoughType, []domain.Ingredient) (domain.Pizza, int64, error)
	makePizzaMutex       sync.RWMutex
	makePizzaArgsForCall []struct {
		arg1 domain.DoughType
		arg2 []domain.Ingredient
	}
	makePizzaReturns struct {
		result1 domain.Pizza
		result2 int64
		result3 error
	}
}

func (fake *FakePizzaRepository) MakePizza(arg1 domain.DoughType, arg2 []domain.Ingredient) (domain.Pizza, int64, error) {
	fake.makePizzaMutex.Lock()
	fake.makePizzaArgsForCall = append(fake.makePizzaArgsForCall, struct {
		arg1 domain.DoughType
		arg2 []domain.Ingredient
	}{arg1, arg2})
	fake.makePizzaMutex.Unlock()
	if fake.MakePizzaStub != nil {
		return fake.MakePizzaStub(arg1, arg2)
	} else {
		return fake.makePizzaReturns.result1, fake.makePizzaReturns.result2, fake.makePizzaReturns.result3
	}
}

func (fake *FakePizzaRepository) MakePizzaCallCount() int {
	fake.makePizzaMutex.RLock()
	defer fake.makePizzaMutex.RUnlock()
	return len(fake.makePizzaArgsForCall)
}

func (fake *FakePizzaRepository) MakePizzaArgsForCall(i int) (domain.DoughType, []domain.Ingredient) {
	fake.makePizzaMutex.RLock()
	defer fake.makePizzaMutex.RUnlock()
	return fake.makePizzaArgsForCall[i].arg1, fake.makePizzaArgsForCall[i].arg2
}

func (fake *FakePizzaRepository) MakePizzaReturns(result1 domain.Pizza, result2 int64, result3 error) {
	fake.MakePizzaStub = nil
	fake.makePizzaReturns = struct {
		result1 domain.Pizza
		result2 int64
		result3 error
	}{result1, result2, result3}
}

var _ api.PizzaRepository = new(FakePizzaRepository)

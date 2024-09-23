// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/food_usecase.go
//
// Generated by this command:
//
//	mockgen -source ./usecase/food_usecase.go -destination usecase/mocks/food_usecase.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	model "RefrigeratorWatchdog-server/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIFoodUsecase is a mock of IFoodUsecase interface.
type MockIFoodUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIFoodUsecaseMockRecorder
}

// MockIFoodUsecaseMockRecorder is the mock recorder for MockIFoodUsecase.
type MockIFoodUsecaseMockRecorder struct {
	mock *MockIFoodUsecase
}

// NewMockIFoodUsecase creates a new mock instance.
func NewMockIFoodUsecase(ctrl *gomock.Controller) *MockIFoodUsecase {
	mock := &MockIFoodUsecase{ctrl: ctrl}
	mock.recorder = &MockIFoodUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFoodUsecase) EXPECT() *MockIFoodUsecaseMockRecorder {
	return m.recorder
}

// CreateFood mocks base method.
func (m *MockIFoodUsecase) CreateFood(food model.Food) (model.FoodResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFood", food)
	ret0, _ := ret[0].(model.FoodResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFood indicates an expected call of CreateFood.
func (mr *MockIFoodUsecaseMockRecorder) CreateFood(food any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFood", reflect.TypeOf((*MockIFoodUsecase)(nil).CreateFood), food)
}

// DeleteFood mocks base method.
func (m *MockIFoodUsecase) DeleteFood(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFood", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFood indicates an expected call of DeleteFood.
func (mr *MockIFoodUsecaseMockRecorder) DeleteFood(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFood", reflect.TypeOf((*MockIFoodUsecase)(nil).DeleteFood), id)
}

// GetFoodsByUserID mocks base method.
func (m *MockIFoodUsecase) GetFoodsByUserID(userID uint) ([]model.FoodResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFoodsByUserID", userID)
	ret0, _ := ret[0].([]model.FoodResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFoodsByUserID indicates an expected call of GetFoodsByUserID.
func (mr *MockIFoodUsecaseMockRecorder) GetFoodsByUserID(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFoodsByUserID", reflect.TypeOf((*MockIFoodUsecase)(nil).GetFoodsByUserID), userID)
}

// UpdateFood mocks base method.
func (m *MockIFoodUsecase) UpdateFood(food model.Food, id uint) (model.FoodResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFood", food, id)
	ret0, _ := ret[0].(model.FoodResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFood indicates an expected call of UpdateFood.
func (mr *MockIFoodUsecaseMockRecorder) UpdateFood(food, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFood", reflect.TypeOf((*MockIFoodUsecase)(nil).UpdateFood), food, id)
}
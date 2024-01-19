// Code generated by MockGen. DO NOT EDIT.
// Source: ./generated/swagger/configunstable/client/operations/operations_client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	operations "github.com/chronosphereio/chronoctl-core/src/generated/swagger/configunstable/client/operations"
	runtime "github.com/go-openapi/runtime"
	gomock "github.com/golang/mock/gomock"
)

// MockClientService is a mock of ClientService interface.
type MockClientService struct {
	ctrl     *gomock.Controller
	recorder *MockClientServiceMockRecorder
}

// MockClientServiceMockRecorder is the mock recorder for MockClientService.
type MockClientServiceMockRecorder struct {
	mock *MockClientService
}

// NewMockClientService creates a new mock instance.
func NewMockClientService(ctrl *gomock.Controller) *MockClientService {
	mock := &MockClientService{ctrl: ctrl}
	mock.recorder = &MockClientServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientService) EXPECT() *MockClientServiceMockRecorder {
	return m.recorder
}

// CreateDashboard mocks base method.
func (m *MockClientService) CreateDashboard(params *operations.CreateDashboardParams, opts ...operations.ClientOption) (*operations.CreateDashboardOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDashboard", varargs...)
	ret0, _ := ret[0].(*operations.CreateDashboardOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDashboard indicates an expected call of CreateDashboard.
func (mr *MockClientServiceMockRecorder) CreateDashboard(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDashboard", reflect.TypeOf((*MockClientService)(nil).CreateDashboard), varargs...)
}

// CreateLinkTemplate mocks base method.
func (m *MockClientService) CreateLinkTemplate(params *operations.CreateLinkTemplateParams, opts ...operations.ClientOption) (*operations.CreateLinkTemplateOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateLinkTemplate", varargs...)
	ret0, _ := ret[0].(*operations.CreateLinkTemplateOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLinkTemplate indicates an expected call of CreateLinkTemplate.
func (mr *MockClientServiceMockRecorder) CreateLinkTemplate(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLinkTemplate", reflect.TypeOf((*MockClientService)(nil).CreateLinkTemplate), varargs...)
}

// CreateNoopEntity mocks base method.
func (m *MockClientService) CreateNoopEntity(params *operations.CreateNoopEntityParams, opts ...operations.ClientOption) (*operations.CreateNoopEntityOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNoopEntity", varargs...)
	ret0, _ := ret[0].(*operations.CreateNoopEntityOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNoopEntity indicates an expected call of CreateNoopEntity.
func (mr *MockClientServiceMockRecorder) CreateNoopEntity(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNoopEntity", reflect.TypeOf((*MockClientService)(nil).CreateNoopEntity), varargs...)
}

// CreateSavedTraceSearch mocks base method.
func (m *MockClientService) CreateSavedTraceSearch(params *operations.CreateSavedTraceSearchParams, opts ...operations.ClientOption) (*operations.CreateSavedTraceSearchOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSavedTraceSearch", varargs...)
	ret0, _ := ret[0].(*operations.CreateSavedTraceSearchOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSavedTraceSearch indicates an expected call of CreateSavedTraceSearch.
func (mr *MockClientServiceMockRecorder) CreateSavedTraceSearch(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSavedTraceSearch", reflect.TypeOf((*MockClientService)(nil).CreateSavedTraceSearch), varargs...)
}

// CreateService mocks base method.
func (m *MockClientService) CreateService(params *operations.CreateServiceParams, opts ...operations.ClientOption) (*operations.CreateServiceOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateService", varargs...)
	ret0, _ := ret[0].(*operations.CreateServiceOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService.
func (mr *MockClientServiceMockRecorder) CreateService(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockClientService)(nil).CreateService), varargs...)
}

// CreateTraceJaegerRemoteSamplingStrategy mocks base method.
func (m *MockClientService) CreateTraceJaegerRemoteSamplingStrategy(params *operations.CreateTraceJaegerRemoteSamplingStrategyParams, opts ...operations.ClientOption) (*operations.CreateTraceJaegerRemoteSamplingStrategyOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTraceJaegerRemoteSamplingStrategy", varargs...)
	ret0, _ := ret[0].(*operations.CreateTraceJaegerRemoteSamplingStrategyOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTraceJaegerRemoteSamplingStrategy indicates an expected call of CreateTraceJaegerRemoteSamplingStrategy.
func (mr *MockClientServiceMockRecorder) CreateTraceJaegerRemoteSamplingStrategy(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTraceJaegerRemoteSamplingStrategy", reflect.TypeOf((*MockClientService)(nil).CreateTraceJaegerRemoteSamplingStrategy), varargs...)
}

// CreateTraceTailSamplingRules mocks base method.
func (m *MockClientService) CreateTraceTailSamplingRules(params *operations.CreateTraceTailSamplingRulesParams, opts ...operations.ClientOption) (*operations.CreateTraceTailSamplingRulesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTraceTailSamplingRules", varargs...)
	ret0, _ := ret[0].(*operations.CreateTraceTailSamplingRulesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTraceTailSamplingRules indicates an expected call of CreateTraceTailSamplingRules.
func (mr *MockClientServiceMockRecorder) CreateTraceTailSamplingRules(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTraceTailSamplingRules", reflect.TypeOf((*MockClientService)(nil).CreateTraceTailSamplingRules), varargs...)
}

// DeleteDashboard mocks base method.
func (m *MockClientService) DeleteDashboard(params *operations.DeleteDashboardParams, opts ...operations.ClientOption) (*operations.DeleteDashboardOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDashboard", varargs...)
	ret0, _ := ret[0].(*operations.DeleteDashboardOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDashboard indicates an expected call of DeleteDashboard.
func (mr *MockClientServiceMockRecorder) DeleteDashboard(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDashboard", reflect.TypeOf((*MockClientService)(nil).DeleteDashboard), varargs...)
}

// DeleteLinkTemplate mocks base method.
func (m *MockClientService) DeleteLinkTemplate(params *operations.DeleteLinkTemplateParams, opts ...operations.ClientOption) (*operations.DeleteLinkTemplateOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteLinkTemplate", varargs...)
	ret0, _ := ret[0].(*operations.DeleteLinkTemplateOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLinkTemplate indicates an expected call of DeleteLinkTemplate.
func (mr *MockClientServiceMockRecorder) DeleteLinkTemplate(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLinkTemplate", reflect.TypeOf((*MockClientService)(nil).DeleteLinkTemplate), varargs...)
}

// DeleteNoopEntity mocks base method.
func (m *MockClientService) DeleteNoopEntity(params *operations.DeleteNoopEntityParams, opts ...operations.ClientOption) (*operations.DeleteNoopEntityOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteNoopEntity", varargs...)
	ret0, _ := ret[0].(*operations.DeleteNoopEntityOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNoopEntity indicates an expected call of DeleteNoopEntity.
func (mr *MockClientServiceMockRecorder) DeleteNoopEntity(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNoopEntity", reflect.TypeOf((*MockClientService)(nil).DeleteNoopEntity), varargs...)
}

// DeleteSavedTraceSearch mocks base method.
func (m *MockClientService) DeleteSavedTraceSearch(params *operations.DeleteSavedTraceSearchParams, opts ...operations.ClientOption) (*operations.DeleteSavedTraceSearchOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSavedTraceSearch", varargs...)
	ret0, _ := ret[0].(*operations.DeleteSavedTraceSearchOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSavedTraceSearch indicates an expected call of DeleteSavedTraceSearch.
func (mr *MockClientServiceMockRecorder) DeleteSavedTraceSearch(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSavedTraceSearch", reflect.TypeOf((*MockClientService)(nil).DeleteSavedTraceSearch), varargs...)
}

// DeleteService mocks base method.
func (m *MockClientService) DeleteService(params *operations.DeleteServiceParams, opts ...operations.ClientOption) (*operations.DeleteServiceOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteService", varargs...)
	ret0, _ := ret[0].(*operations.DeleteServiceOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockClientServiceMockRecorder) DeleteService(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockClientService)(nil).DeleteService), varargs...)
}

// DeleteTraceJaegerRemoteSamplingStrategy mocks base method.
func (m *MockClientService) DeleteTraceJaegerRemoteSamplingStrategy(params *operations.DeleteTraceJaegerRemoteSamplingStrategyParams, opts ...operations.ClientOption) (*operations.DeleteTraceJaegerRemoteSamplingStrategyOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTraceJaegerRemoteSamplingStrategy", varargs...)
	ret0, _ := ret[0].(*operations.DeleteTraceJaegerRemoteSamplingStrategyOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTraceJaegerRemoteSamplingStrategy indicates an expected call of DeleteTraceJaegerRemoteSamplingStrategy.
func (mr *MockClientServiceMockRecorder) DeleteTraceJaegerRemoteSamplingStrategy(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTraceJaegerRemoteSamplingStrategy", reflect.TypeOf((*MockClientService)(nil).DeleteTraceJaegerRemoteSamplingStrategy), varargs...)
}

// DeleteTraceTailSamplingRules mocks base method.
func (m *MockClientService) DeleteTraceTailSamplingRules(params *operations.DeleteTraceTailSamplingRulesParams, opts ...operations.ClientOption) (*operations.DeleteTraceTailSamplingRulesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTraceTailSamplingRules", varargs...)
	ret0, _ := ret[0].(*operations.DeleteTraceTailSamplingRulesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTraceTailSamplingRules indicates an expected call of DeleteTraceTailSamplingRules.
func (mr *MockClientServiceMockRecorder) DeleteTraceTailSamplingRules(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTraceTailSamplingRules", reflect.TypeOf((*MockClientService)(nil).DeleteTraceTailSamplingRules), varargs...)
}

// ListDashboards mocks base method.
func (m *MockClientService) ListDashboards(params *operations.ListDashboardsParams, opts ...operations.ClientOption) (*operations.ListDashboardsOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDashboards", varargs...)
	ret0, _ := ret[0].(*operations.ListDashboardsOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDashboards indicates an expected call of ListDashboards.
func (mr *MockClientServiceMockRecorder) ListDashboards(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDashboards", reflect.TypeOf((*MockClientService)(nil).ListDashboards), varargs...)
}

// ListLinkTemplates mocks base method.
func (m *MockClientService) ListLinkTemplates(params *operations.ListLinkTemplatesParams, opts ...operations.ClientOption) (*operations.ListLinkTemplatesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLinkTemplates", varargs...)
	ret0, _ := ret[0].(*operations.ListLinkTemplatesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLinkTemplates indicates an expected call of ListLinkTemplates.
func (mr *MockClientServiceMockRecorder) ListLinkTemplates(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLinkTemplates", reflect.TypeOf((*MockClientService)(nil).ListLinkTemplates), varargs...)
}

// ListNoopEntities mocks base method.
func (m *MockClientService) ListNoopEntities(params *operations.ListNoopEntitiesParams, opts ...operations.ClientOption) (*operations.ListNoopEntitiesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNoopEntities", varargs...)
	ret0, _ := ret[0].(*operations.ListNoopEntitiesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNoopEntities indicates an expected call of ListNoopEntities.
func (mr *MockClientServiceMockRecorder) ListNoopEntities(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNoopEntities", reflect.TypeOf((*MockClientService)(nil).ListNoopEntities), varargs...)
}

// ListSavedTraceSearches mocks base method.
func (m *MockClientService) ListSavedTraceSearches(params *operations.ListSavedTraceSearchesParams, opts ...operations.ClientOption) (*operations.ListSavedTraceSearchesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSavedTraceSearches", varargs...)
	ret0, _ := ret[0].(*operations.ListSavedTraceSearchesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSavedTraceSearches indicates an expected call of ListSavedTraceSearches.
func (mr *MockClientServiceMockRecorder) ListSavedTraceSearches(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSavedTraceSearches", reflect.TypeOf((*MockClientService)(nil).ListSavedTraceSearches), varargs...)
}

// ListServices mocks base method.
func (m *MockClientService) ListServices(params *operations.ListServicesParams, opts ...operations.ClientOption) (*operations.ListServicesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*operations.ListServicesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockClientServiceMockRecorder) ListServices(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockClientService)(nil).ListServices), varargs...)
}

// ListTraceJaegerRemoteSamplingStrategies mocks base method.
func (m *MockClientService) ListTraceJaegerRemoteSamplingStrategies(params *operations.ListTraceJaegerRemoteSamplingStrategiesParams, opts ...operations.ClientOption) (*operations.ListTraceJaegerRemoteSamplingStrategiesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTraceJaegerRemoteSamplingStrategies", varargs...)
	ret0, _ := ret[0].(*operations.ListTraceJaegerRemoteSamplingStrategiesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTraceJaegerRemoteSamplingStrategies indicates an expected call of ListTraceJaegerRemoteSamplingStrategies.
func (mr *MockClientServiceMockRecorder) ListTraceJaegerRemoteSamplingStrategies(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTraceJaegerRemoteSamplingStrategies", reflect.TypeOf((*MockClientService)(nil).ListTraceJaegerRemoteSamplingStrategies), varargs...)
}

// ReadDashboard mocks base method.
func (m *MockClientService) ReadDashboard(params *operations.ReadDashboardParams, opts ...operations.ClientOption) (*operations.ReadDashboardOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadDashboard", varargs...)
	ret0, _ := ret[0].(*operations.ReadDashboardOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDashboard indicates an expected call of ReadDashboard.
func (mr *MockClientServiceMockRecorder) ReadDashboard(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDashboard", reflect.TypeOf((*MockClientService)(nil).ReadDashboard), varargs...)
}

// ReadLinkTemplate mocks base method.
func (m *MockClientService) ReadLinkTemplate(params *operations.ReadLinkTemplateParams, opts ...operations.ClientOption) (*operations.ReadLinkTemplateOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadLinkTemplate", varargs...)
	ret0, _ := ret[0].(*operations.ReadLinkTemplateOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadLinkTemplate indicates an expected call of ReadLinkTemplate.
func (mr *MockClientServiceMockRecorder) ReadLinkTemplate(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLinkTemplate", reflect.TypeOf((*MockClientService)(nil).ReadLinkTemplate), varargs...)
}

// ReadNoopEntity mocks base method.
func (m *MockClientService) ReadNoopEntity(params *operations.ReadNoopEntityParams, opts ...operations.ClientOption) (*operations.ReadNoopEntityOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadNoopEntity", varargs...)
	ret0, _ := ret[0].(*operations.ReadNoopEntityOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadNoopEntity indicates an expected call of ReadNoopEntity.
func (mr *MockClientServiceMockRecorder) ReadNoopEntity(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadNoopEntity", reflect.TypeOf((*MockClientService)(nil).ReadNoopEntity), varargs...)
}

// ReadSavedTraceSearch mocks base method.
func (m *MockClientService) ReadSavedTraceSearch(params *operations.ReadSavedTraceSearchParams, opts ...operations.ClientOption) (*operations.ReadSavedTraceSearchOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadSavedTraceSearch", varargs...)
	ret0, _ := ret[0].(*operations.ReadSavedTraceSearchOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadSavedTraceSearch indicates an expected call of ReadSavedTraceSearch.
func (mr *MockClientServiceMockRecorder) ReadSavedTraceSearch(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadSavedTraceSearch", reflect.TypeOf((*MockClientService)(nil).ReadSavedTraceSearch), varargs...)
}

// ReadService mocks base method.
func (m *MockClientService) ReadService(params *operations.ReadServiceParams, opts ...operations.ClientOption) (*operations.ReadServiceOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadService", varargs...)
	ret0, _ := ret[0].(*operations.ReadServiceOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadService indicates an expected call of ReadService.
func (mr *MockClientServiceMockRecorder) ReadService(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadService", reflect.TypeOf((*MockClientService)(nil).ReadService), varargs...)
}

// ReadTraceJaegerRemoteSamplingStrategy mocks base method.
func (m *MockClientService) ReadTraceJaegerRemoteSamplingStrategy(params *operations.ReadTraceJaegerRemoteSamplingStrategyParams, opts ...operations.ClientOption) (*operations.ReadTraceJaegerRemoteSamplingStrategyOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadTraceJaegerRemoteSamplingStrategy", varargs...)
	ret0, _ := ret[0].(*operations.ReadTraceJaegerRemoteSamplingStrategyOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTraceJaegerRemoteSamplingStrategy indicates an expected call of ReadTraceJaegerRemoteSamplingStrategy.
func (mr *MockClientServiceMockRecorder) ReadTraceJaegerRemoteSamplingStrategy(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTraceJaegerRemoteSamplingStrategy", reflect.TypeOf((*MockClientService)(nil).ReadTraceJaegerRemoteSamplingStrategy), varargs...)
}

// ReadTraceTailSamplingRules mocks base method.
func (m *MockClientService) ReadTraceTailSamplingRules(params *operations.ReadTraceTailSamplingRulesParams, opts ...operations.ClientOption) (*operations.ReadTraceTailSamplingRulesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadTraceTailSamplingRules", varargs...)
	ret0, _ := ret[0].(*operations.ReadTraceTailSamplingRulesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTraceTailSamplingRules indicates an expected call of ReadTraceTailSamplingRules.
func (mr *MockClientServiceMockRecorder) ReadTraceTailSamplingRules(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTraceTailSamplingRules", reflect.TypeOf((*MockClientService)(nil).ReadTraceTailSamplingRules), varargs...)
}

// SetTransport mocks base method.
func (m *MockClientService) SetTransport(transport runtime.ClientTransport) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTransport", transport)
}

// SetTransport indicates an expected call of SetTransport.
func (mr *MockClientServiceMockRecorder) SetTransport(transport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransport", reflect.TypeOf((*MockClientService)(nil).SetTransport), transport)
}

// SyncPrometheus mocks base method.
func (m *MockClientService) SyncPrometheus(params *operations.SyncPrometheusParams, opts ...operations.ClientOption) (*operations.SyncPrometheusOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncPrometheus", varargs...)
	ret0, _ := ret[0].(*operations.SyncPrometheusOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncPrometheus indicates an expected call of SyncPrometheus.
func (mr *MockClientServiceMockRecorder) SyncPrometheus(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPrometheus", reflect.TypeOf((*MockClientService)(nil).SyncPrometheus), varargs...)
}

// UpdateDashboard mocks base method.
func (m *MockClientService) UpdateDashboard(params *operations.UpdateDashboardParams, opts ...operations.ClientOption) (*operations.UpdateDashboardOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateDashboard", varargs...)
	ret0, _ := ret[0].(*operations.UpdateDashboardOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDashboard indicates an expected call of UpdateDashboard.
func (mr *MockClientServiceMockRecorder) UpdateDashboard(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDashboard", reflect.TypeOf((*MockClientService)(nil).UpdateDashboard), varargs...)
}

// UpdateLinkTemplate mocks base method.
func (m *MockClientService) UpdateLinkTemplate(params *operations.UpdateLinkTemplateParams, opts ...operations.ClientOption) (*operations.UpdateLinkTemplateOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateLinkTemplate", varargs...)
	ret0, _ := ret[0].(*operations.UpdateLinkTemplateOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLinkTemplate indicates an expected call of UpdateLinkTemplate.
func (mr *MockClientServiceMockRecorder) UpdateLinkTemplate(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLinkTemplate", reflect.TypeOf((*MockClientService)(nil).UpdateLinkTemplate), varargs...)
}

// UpdateNoopEntity mocks base method.
func (m *MockClientService) UpdateNoopEntity(params *operations.UpdateNoopEntityParams, opts ...operations.ClientOption) (*operations.UpdateNoopEntityOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateNoopEntity", varargs...)
	ret0, _ := ret[0].(*operations.UpdateNoopEntityOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNoopEntity indicates an expected call of UpdateNoopEntity.
func (mr *MockClientServiceMockRecorder) UpdateNoopEntity(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNoopEntity", reflect.TypeOf((*MockClientService)(nil).UpdateNoopEntity), varargs...)
}

// UpdateSavedTraceSearch mocks base method.
func (m *MockClientService) UpdateSavedTraceSearch(params *operations.UpdateSavedTraceSearchParams, opts ...operations.ClientOption) (*operations.UpdateSavedTraceSearchOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSavedTraceSearch", varargs...)
	ret0, _ := ret[0].(*operations.UpdateSavedTraceSearchOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSavedTraceSearch indicates an expected call of UpdateSavedTraceSearch.
func (mr *MockClientServiceMockRecorder) UpdateSavedTraceSearch(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSavedTraceSearch", reflect.TypeOf((*MockClientService)(nil).UpdateSavedTraceSearch), varargs...)
}

// UpdateService mocks base method.
func (m *MockClientService) UpdateService(params *operations.UpdateServiceParams, opts ...operations.ClientOption) (*operations.UpdateServiceOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateService", varargs...)
	ret0, _ := ret[0].(*operations.UpdateServiceOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockClientServiceMockRecorder) UpdateService(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockClientService)(nil).UpdateService), varargs...)
}

// UpdateTraceJaegerRemoteSamplingStrategy mocks base method.
func (m *MockClientService) UpdateTraceJaegerRemoteSamplingStrategy(params *operations.UpdateTraceJaegerRemoteSamplingStrategyParams, opts ...operations.ClientOption) (*operations.UpdateTraceJaegerRemoteSamplingStrategyOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTraceJaegerRemoteSamplingStrategy", varargs...)
	ret0, _ := ret[0].(*operations.UpdateTraceJaegerRemoteSamplingStrategyOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTraceJaegerRemoteSamplingStrategy indicates an expected call of UpdateTraceJaegerRemoteSamplingStrategy.
func (mr *MockClientServiceMockRecorder) UpdateTraceJaegerRemoteSamplingStrategy(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTraceJaegerRemoteSamplingStrategy", reflect.TypeOf((*MockClientService)(nil).UpdateTraceJaegerRemoteSamplingStrategy), varargs...)
}

// UpdateTraceTailSamplingRules mocks base method.
func (m *MockClientService) UpdateTraceTailSamplingRules(params *operations.UpdateTraceTailSamplingRulesParams, opts ...operations.ClientOption) (*operations.UpdateTraceTailSamplingRulesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{params}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTraceTailSamplingRules", varargs...)
	ret0, _ := ret[0].(*operations.UpdateTraceTailSamplingRulesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTraceTailSamplingRules indicates an expected call of UpdateTraceTailSamplingRules.
func (mr *MockClientServiceMockRecorder) UpdateTraceTailSamplingRules(params interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{params}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTraceTailSamplingRules", reflect.TypeOf((*MockClientService)(nil).UpdateTraceTailSamplingRules), varargs...)
}

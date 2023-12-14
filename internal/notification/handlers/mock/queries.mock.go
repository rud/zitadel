// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zitadel/zitadel/internal/notification/handlers (interfaces: Queries)
//
// Generated by this command:
//
//	mockgen -package mock -destination ./mock/queries.mock.go github.com/zitadel/zitadel/internal/notification/handlers Queries
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/zitadel/zitadel/internal/domain"
	query "github.com/zitadel/zitadel/internal/query"
	gomock "go.uber.org/mock/gomock"
	language "golang.org/x/text/language"
)

// MockQueries is a mock of Queries interface.
type MockQueries struct {
	ctrl     *gomock.Controller
	recorder *MockQueriesMockRecorder
}

// MockQueriesMockRecorder is the mock recorder for MockQueries.
type MockQueriesMockRecorder struct {
	mock *MockQueries
}

// NewMockQueries creates a new mock instance.
func NewMockQueries(ctrl *gomock.Controller) *MockQueries {
	mock := &MockQueries{ctrl: ctrl}
	mock.recorder = &MockQueriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueries) EXPECT() *MockQueriesMockRecorder {
	return m.recorder
}

// ActiveLabelPolicyByOrg mocks base method.
func (m *MockQueries) ActiveLabelPolicyByOrg(arg0 context.Context, arg1 string, arg2 bool) (*query.LabelPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveLabelPolicyByOrg", arg0, arg1, arg2)
	ret0, _ := ret[0].(*query.LabelPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveLabelPolicyByOrg indicates an expected call of ActiveLabelPolicyByOrg.
func (mr *MockQueriesMockRecorder) ActiveLabelPolicyByOrg(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveLabelPolicyByOrg", reflect.TypeOf((*MockQueries)(nil).ActiveLabelPolicyByOrg), arg0, arg1, arg2)
}

// CustomTextListByTemplate mocks base method.
func (m *MockQueries) CustomTextListByTemplate(arg0 context.Context, arg1, arg2 string, arg3 bool) (*query.CustomTexts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CustomTextListByTemplate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*query.CustomTexts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CustomTextListByTemplate indicates an expected call of CustomTextListByTemplate.
func (mr *MockQueriesMockRecorder) CustomTextListByTemplate(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomTextListByTemplate", reflect.TypeOf((*MockQueries)(nil).CustomTextListByTemplate), arg0, arg1, arg2, arg3)
}

// GetDefaultLanguage mocks base method.
func (m *MockQueries) GetDefaultLanguage(arg0 context.Context) language.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultLanguage", arg0)
	ret0, _ := ret[0].(language.Tag)
	return ret0
}

// GetDefaultLanguage indicates an expected call of GetDefaultLanguage.
func (mr *MockQueriesMockRecorder) GetDefaultLanguage(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultLanguage", reflect.TypeOf((*MockQueries)(nil).GetDefaultLanguage), arg0)
}

// GetNotifyUserByID mocks base method.
func (m *MockQueries) GetNotifyUserByID(arg0 context.Context, arg1 bool, arg2 string) (*query.NotifyUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifyUserByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*query.NotifyUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifyUserByID indicates an expected call of GetNotifyUserByID.
func (mr *MockQueriesMockRecorder) GetNotifyUserByID(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifyUserByID", reflect.TypeOf((*MockQueries)(nil).GetNotifyUserByID), arg0, arg1, arg2)
}

// MailTemplateByOrg mocks base method.
func (m *MockQueries) MailTemplateByOrg(arg0 context.Context, arg1 string, arg2 bool) (*query.MailTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MailTemplateByOrg", arg0, arg1, arg2)
	ret0, _ := ret[0].(*query.MailTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MailTemplateByOrg indicates an expected call of MailTemplateByOrg.
func (mr *MockQueriesMockRecorder) MailTemplateByOrg(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MailTemplateByOrg", reflect.TypeOf((*MockQueries)(nil).MailTemplateByOrg), arg0, arg1, arg2)
}

// NotificationPolicyByOrg mocks base method.
func (m *MockQueries) NotificationPolicyByOrg(arg0 context.Context, arg1 bool, arg2 string, arg3 bool) (*query.NotificationPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotificationPolicyByOrg", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*query.NotificationPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationPolicyByOrg indicates an expected call of NotificationPolicyByOrg.
func (mr *MockQueriesMockRecorder) NotificationPolicyByOrg(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationPolicyByOrg", reflect.TypeOf((*MockQueries)(nil).NotificationPolicyByOrg), arg0, arg1, arg2, arg3)
}

// NotificationProviderByIDAndType mocks base method.
func (m *MockQueries) NotificationProviderByIDAndType(arg0 context.Context, arg1 string, arg2 domain.NotificationProviderType) (*query.DebugNotificationProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotificationProviderByIDAndType", arg0, arg1, arg2)
	ret0, _ := ret[0].(*query.DebugNotificationProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationProviderByIDAndType indicates an expected call of NotificationProviderByIDAndType.
func (mr *MockQueriesMockRecorder) NotificationProviderByIDAndType(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationProviderByIDAndType", reflect.TypeOf((*MockQueries)(nil).NotificationProviderByIDAndType), arg0, arg1, arg2)
}

// SMSProviderConfig mocks base method.
func (m *MockQueries) SMSProviderConfig(arg0 context.Context, arg1 ...query.SearchQuery) (*query.SMSConfig, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SMSProviderConfig", varargs...)
	ret0, _ := ret[0].(*query.SMSConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SMSProviderConfig indicates an expected call of SMSProviderConfig.
func (mr *MockQueriesMockRecorder) SMSProviderConfig(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMSProviderConfig", reflect.TypeOf((*MockQueries)(nil).SMSProviderConfig), varargs...)
}

// SMTPConfigByAggregateID mocks base method.
func (m *MockQueries) SMTPConfigByAggregateID(arg0 context.Context, arg1 string) (*query.SMTPConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMTPConfigByAggregateID", arg0, arg1)
	ret0, _ := ret[0].(*query.SMTPConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SMTPConfigByAggregateID indicates an expected call of SMTPConfigByAggregateID.
func (mr *MockQueriesMockRecorder) SMTPConfigByAggregateID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMTPConfigByAggregateID", reflect.TypeOf((*MockQueries)(nil).SMTPConfigByAggregateID), arg0, arg1)
}

// SearchInstanceDomains mocks base method.
func (m *MockQueries) SearchInstanceDomains(arg0 context.Context, arg1 *query.InstanceDomainSearchQueries) (*query.InstanceDomains, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchInstanceDomains", arg0, arg1)
	ret0, _ := ret[0].(*query.InstanceDomains)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchInstanceDomains indicates an expected call of SearchInstanceDomains.
func (mr *MockQueriesMockRecorder) SearchInstanceDomains(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchInstanceDomains", reflect.TypeOf((*MockQueries)(nil).SearchInstanceDomains), arg0, arg1)
}

// SearchMilestones mocks base method.
func (m *MockQueries) SearchMilestones(arg0 context.Context, arg1 []string, arg2 *query.MilestonesSearchQueries) (*query.Milestones, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMilestones", arg0, arg1, arg2)
	ret0, _ := ret[0].(*query.Milestones)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMilestones indicates an expected call of SearchMilestones.
func (mr *MockQueriesMockRecorder) SearchMilestones(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMilestones", reflect.TypeOf((*MockQueries)(nil).SearchMilestones), arg0, arg1, arg2)
}

// SessionByID mocks base method.
func (m *MockQueries) SessionByID(arg0 context.Context, arg1 bool, arg2, arg3 string) (*query.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SessionByID", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*query.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SessionByID indicates an expected call of SessionByID.
func (mr *MockQueriesMockRecorder) SessionByID(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionByID", reflect.TypeOf((*MockQueries)(nil).SessionByID), arg0, arg1, arg2, arg3)
}

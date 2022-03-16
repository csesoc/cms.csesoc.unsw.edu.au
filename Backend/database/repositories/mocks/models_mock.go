// Code generated by MockGen. DO NOT EDIT.
// Source: models.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	repositories "cms.csesoc.unsw.edu.au/database/repositories"
	gomock "github.com/golang/mock/gomock"
)

// MockIFilesystemRepository is a mock of IFilesystemRepository interface.
type MockIFilesystemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIFilesystemRepositoryMockRecorder
}

// MockIFilesystemRepositoryMockRecorder is the mock recorder for MockIFilesystemRepository.
type MockIFilesystemRepositoryMockRecorder struct {
	mock *MockIFilesystemRepository
}

// NewMockIFilesystemRepository creates a new mock instance.
func NewMockIFilesystemRepository(ctrl *gomock.Controller) *MockIFilesystemRepository {
	mock := &MockIFilesystemRepository{ctrl: ctrl}
	mock.recorder = &MockIFilesystemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFilesystemRepository) EXPECT() *MockIFilesystemRepositoryMockRecorder {
	return m.recorder
}

// CreateEntry mocks base method.
func (m *MockIFilesystemRepository) CreateEntry(file repositories.FilesystemEntry) (repositories.FilesystemEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", file)
	ret0, _ := ret[0].(repositories.FilesystemEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockIFilesystemRepositoryMockRecorder) CreateEntry(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockIFilesystemRepository)(nil).CreateEntry), file)
}

// DeleteEntryWithID mocks base method.
func (m *MockIFilesystemRepository) DeleteEntryWithID(ID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntryWithID", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntryWithID indicates an expected call of DeleteEntryWithID.
func (mr *MockIFilesystemRepositoryMockRecorder) DeleteEntryWithID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntryWithID", reflect.TypeOf((*MockIFilesystemRepository)(nil).DeleteEntryWithID), ID)
}

// GetEntryWithID mocks base method.
func (m *MockIFilesystemRepository) GetEntryWithID(ID int) (repositories.FilesystemEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntryWithID", ID)
	ret0, _ := ret[0].(repositories.FilesystemEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntryWithID indicates an expected call of GetEntryWithID.
func (mr *MockIFilesystemRepositoryMockRecorder) GetEntryWithID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntryWithID", reflect.TypeOf((*MockIFilesystemRepository)(nil).GetEntryWithID), ID)
}

// GetEntryWithParentID mocks base method.
func (m *MockIFilesystemRepository) GetEntryWithParentID(ID int) (repositories.FilesystemEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntryWithParentID", ID)
	ret0, _ := ret[0].(repositories.FilesystemEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntryWithParentID indicates an expected call of GetEntryWithParentID.
func (mr *MockIFilesystemRepositoryMockRecorder) GetEntryWithParentID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntryWithParentID", reflect.TypeOf((*MockIFilesystemRepository)(nil).GetEntryWithParentID), ID)
}

// GetRoot mocks base method.
func (m *MockIFilesystemRepository) GetRoot() (repositories.FilesystemEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoot")
	ret0, _ := ret[0].(repositories.FilesystemEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoot indicates an expected call of GetRoot.
func (mr *MockIFilesystemRepositoryMockRecorder) GetRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoot", reflect.TypeOf((*MockIFilesystemRepository)(nil).GetRoot))
}

// RenameEntity mocks base method.
func (m *MockIFilesystemRepository) RenameEntity(ID int, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameEntity", ID, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameEntity indicates an expected call of RenameEntity.
func (mr *MockIFilesystemRepositoryMockRecorder) RenameEntity(ID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameEntity", reflect.TypeOf((*MockIFilesystemRepository)(nil).RenameEntity), ID, name)
}

// MockIPersonRepository is a mock of IPersonRepository interface.
type MockIPersonRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPersonRepositoryMockRecorder
}

// MockIPersonRepositoryMockRecorder is the mock recorder for MockIPersonRepository.
type MockIPersonRepositoryMockRecorder struct {
	mock *MockIPersonRepository
}

// NewMockIPersonRepository creates a new mock instance.
func NewMockIPersonRepository(ctrl *gomock.Controller) *MockIPersonRepository {
	mock := &MockIPersonRepository{ctrl: ctrl}
	mock.recorder = &MockIPersonRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPersonRepository) EXPECT() *MockIPersonRepositoryMockRecorder {
	return m.recorder
}

// GetPersonWithDetails mocks base method.
func (m *MockIPersonRepository) GetPersonWithDetails(arg0 repositories.Person) repositories.Person {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPersonWithDetails", arg0)
	ret0, _ := ret[0].(repositories.Person)
	return ret0
}

// GetPersonWithDetails indicates an expected call of GetPersonWithDetails.
func (mr *MockIPersonRepositoryMockRecorder) GetPersonWithDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersonWithDetails", reflect.TypeOf((*MockIPersonRepository)(nil).GetPersonWithDetails), arg0)
}

// PersonExists mocks base method.
func (m *MockIPersonRepository) PersonExists(arg0 repositories.Person) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersonExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// PersonExists indicates an expected call of PersonExists.
func (mr *MockIPersonRepositoryMockRecorder) PersonExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersonExists", reflect.TypeOf((*MockIPersonRepository)(nil).PersonExists), arg0)
}

// MockIGroupsRepository is a mock of IGroupsRepository interface.
type MockIGroupsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIGroupsRepositoryMockRecorder
}

// MockIGroupsRepositoryMockRecorder is the mock recorder for MockIGroupsRepository.
type MockIGroupsRepositoryMockRecorder struct {
	mock *MockIGroupsRepository
}

// NewMockIGroupsRepository creates a new mock instance.
func NewMockIGroupsRepository(ctrl *gomock.Controller) *MockIGroupsRepository {
	mock := &MockIGroupsRepository{ctrl: ctrl}
	mock.recorder = &MockIGroupsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGroupsRepository) EXPECT() *MockIGroupsRepositoryMockRecorder {
	return m.recorder
}

// GetGroupInfo mocks base method.
func (m *MockIGroupsRepository) GetGroupInfo(arg0 repositories.Groups) repositories.Groups {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupInfo", arg0)
	ret0, _ := ret[0].(repositories.Groups)
	return ret0
}

// GetGroupInfo indicates an expected call of GetGroupInfo.
func (mr *MockIGroupsRepositoryMockRecorder) GetGroupInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupInfo", reflect.TypeOf((*MockIGroupsRepository)(nil).GetGroupInfo), arg0)
}
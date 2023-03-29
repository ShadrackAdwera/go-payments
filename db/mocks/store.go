// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ShadrackAdwera/go-payments/db/sqlc (interfaces: TxStore)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/ShadrackAdwera/go-payments/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockTxStore is a mock of TxStore interface.
type MockTxStore struct {
	ctrl     *gomock.Controller
	recorder *MockTxStoreMockRecorder
}

// MockTxStoreMockRecorder is the mock recorder for MockTxStore.
type MockTxStoreMockRecorder struct {
	mock *MockTxStore
}

// NewMockTxStore creates a new mock instance.
func NewMockTxStore(ctrl *gomock.Controller) *MockTxStore {
	mock := &MockTxStore{ctrl: ctrl}
	mock.recorder = &MockTxStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxStore) EXPECT() *MockTxStoreMockRecorder {
	return m.recorder
}

// ApproveRequestTx mocks base method.
func (m *MockTxStore) ApproveRequestTx(arg0 context.Context, arg1 db.ApproveRequestTxRequest) (db.ApproveRequestTxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveRequestTx", arg0, arg1)
	ret0, _ := ret[0].(db.ApproveRequestTxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApproveRequestTx indicates an expected call of ApproveRequestTx.
func (mr *MockTxStoreMockRecorder) ApproveRequestTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveRequestTx", reflect.TypeOf((*MockTxStore)(nil).ApproveRequestTx), arg0, arg1)
}

// CreateClient mocks base method.
func (m *MockTxStore) CreateClient(arg0 context.Context, arg1 db.CreateClientParams) (db.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClient", arg0, arg1)
	ret0, _ := ret[0].(db.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClient indicates an expected call of CreateClient.
func (mr *MockTxStoreMockRecorder) CreateClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClient", reflect.TypeOf((*MockTxStore)(nil).CreateClient), arg0, arg1)
}

// CreatePermission mocks base method.
func (m *MockTxStore) CreatePermission(arg0 context.Context, arg1 db.CreatePermissionParams) (db.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePermission", arg0, arg1)
	ret0, _ := ret[0].(db.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePermission indicates an expected call of CreatePermission.
func (mr *MockTxStoreMockRecorder) CreatePermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePermission", reflect.TypeOf((*MockTxStore)(nil).CreatePermission), arg0, arg1)
}

// CreateRequest mocks base method.
func (m *MockTxStore) CreateRequest(arg0 context.Context, arg1 db.CreateRequestParams) (db.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRequest", arg0, arg1)
	ret0, _ := ret[0].(db.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRequest indicates an expected call of CreateRequest.
func (mr *MockTxStoreMockRecorder) CreateRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRequest", reflect.TypeOf((*MockTxStore)(nil).CreateRequest), arg0, arg1)
}

// CreateRole mocks base method.
func (m *MockTxStore) CreateRole(arg0 context.Context, arg1 db.CreateRoleParams) (db.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", arg0, arg1)
	ret0, _ := ret[0].(db.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockTxStoreMockRecorder) CreateRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockTxStore)(nil).CreateRole), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockTxStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockTxStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockTxStore)(nil).CreateUser), arg0, arg1)
}

// CreateUserPayment mocks base method.
func (m *MockTxStore) CreateUserPayment(arg0 context.Context, arg1 db.CreateUserPaymentParams) (db.UserPayment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserPayment", arg0, arg1)
	ret0, _ := ret[0].(db.UserPayment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserPayment indicates an expected call of CreateUserPayment.
func (mr *MockTxStoreMockRecorder) CreateUserPayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserPayment", reflect.TypeOf((*MockTxStore)(nil).CreateUserPayment), arg0, arg1)
}

// CreateUserRole mocks base method.
func (m *MockTxStore) CreateUserRole(arg0 context.Context, arg1 db.CreateUserRoleParams) (db.UsersRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserRole", arg0, arg1)
	ret0, _ := ret[0].(db.UsersRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserRole indicates an expected call of CreateUserRole.
func (mr *MockTxStoreMockRecorder) CreateUserRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserRole", reflect.TypeOf((*MockTxStore)(nil).CreateUserRole), arg0, arg1)
}

// DeleteClient mocks base method.
func (m *MockTxStore) DeleteClient(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClient", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClient indicates an expected call of DeleteClient.
func (mr *MockTxStoreMockRecorder) DeleteClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClient", reflect.TypeOf((*MockTxStore)(nil).DeleteClient), arg0, arg1)
}

// DeletePermission mocks base method.
func (m *MockTxStore) DeletePermission(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePermission", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePermission indicates an expected call of DeletePermission.
func (mr *MockTxStoreMockRecorder) DeletePermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePermission", reflect.TypeOf((*MockTxStore)(nil).DeletePermission), arg0, arg1)
}

// DeleteRequest mocks base method.
func (m *MockTxStore) DeleteRequest(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRequest indicates an expected call of DeleteRequest.
func (mr *MockTxStoreMockRecorder) DeleteRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRequest", reflect.TypeOf((*MockTxStore)(nil).DeleteRequest), arg0, arg1)
}

// DeleteRole mocks base method.
func (m *MockTxStore) DeleteRole(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockTxStoreMockRecorder) DeleteRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockTxStore)(nil).DeleteRole), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockTxStore) DeleteUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockTxStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockTxStore)(nil).DeleteUser), arg0, arg1)
}

// DeleteUserPayment mocks base method.
func (m *MockTxStore) DeleteUserPayment(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserPayment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserPayment indicates an expected call of DeleteUserPayment.
func (mr *MockTxStoreMockRecorder) DeleteUserPayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserPayment", reflect.TypeOf((*MockTxStore)(nil).DeleteUserPayment), arg0, arg1)
}

// DeleteUserRole mocks base method.
func (m *MockTxStore) DeleteUserRole(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserRole", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserRole indicates an expected call of DeleteUserRole.
func (mr *MockTxStoreMockRecorder) DeleteUserRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserRole", reflect.TypeOf((*MockTxStore)(nil).DeleteUserRole), arg0, arg1)
}

// GetClient mocks base method.
func (m *MockTxStore) GetClient(arg0 context.Context, arg1 int64) (db.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient", arg0, arg1)
	ret0, _ := ret[0].(db.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClient indicates an expected call of GetClient.
func (mr *MockTxStoreMockRecorder) GetClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockTxStore)(nil).GetClient), arg0, arg1)
}

// GetClients mocks base method.
func (m *MockTxStore) GetClients(arg0 context.Context, arg1 db.GetClientsParams) ([]db.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClients", arg0, arg1)
	ret0, _ := ret[0].([]db.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClients indicates an expected call of GetClients.
func (mr *MockTxStoreMockRecorder) GetClients(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClients", reflect.TypeOf((*MockTxStore)(nil).GetClients), arg0, arg1)
}

// GetPermission mocks base method.
func (m *MockTxStore) GetPermission(arg0 context.Context, arg1 int64) (db.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermission", arg0, arg1)
	ret0, _ := ret[0].(db.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermission indicates an expected call of GetPermission.
func (mr *MockTxStoreMockRecorder) GetPermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermission", reflect.TypeOf((*MockTxStore)(nil).GetPermission), arg0, arg1)
}

// GetPermissionByName mocks base method.
func (m *MockTxStore) GetPermissionByName(arg0 context.Context, arg1 string) (db.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermissionByName", arg0, arg1)
	ret0, _ := ret[0].(db.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermissionByName indicates an expected call of GetPermissionByName.
func (mr *MockTxStoreMockRecorder) GetPermissionByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissionByName", reflect.TypeOf((*MockTxStore)(nil).GetPermissionByName), arg0, arg1)
}

// GetPermissions mocks base method.
func (m *MockTxStore) GetPermissions(arg0 context.Context, arg1 db.GetPermissionsParams) ([]db.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermissions", arg0, arg1)
	ret0, _ := ret[0].([]db.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermissions indicates an expected call of GetPermissions.
func (mr *MockTxStoreMockRecorder) GetPermissions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissions", reflect.TypeOf((*MockTxStore)(nil).GetPermissions), arg0, arg1)
}

// GetPermissionsByRole mocks base method.
func (m *MockTxStore) GetPermissionsByRole(arg0 context.Context, arg1 int64) ([]db.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermissionsByRole", arg0, arg1)
	ret0, _ := ret[0].([]db.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermissionsByRole indicates an expected call of GetPermissionsByRole.
func (mr *MockTxStoreMockRecorder) GetPermissionsByRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissionsByRole", reflect.TypeOf((*MockTxStore)(nil).GetPermissionsByRole), arg0, arg1)
}

// GetRequest mocks base method.
func (m *MockTxStore) GetRequest(arg0 context.Context, arg1 int64) (db.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest", arg0, arg1)
	ret0, _ := ret[0].(db.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequest indicates an expected call of GetRequest.
func (mr *MockTxStoreMockRecorder) GetRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockTxStore)(nil).GetRequest), arg0, arg1)
}

// GetRequests mocks base method.
func (m *MockTxStore) GetRequests(arg0 context.Context, arg1 db.GetRequestsParams) ([]db.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequests", arg0, arg1)
	ret0, _ := ret[0].([]db.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequests indicates an expected call of GetRequests.
func (mr *MockTxStoreMockRecorder) GetRequests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequests", reflect.TypeOf((*MockTxStore)(nil).GetRequests), arg0, arg1)
}

// GetRole mocks base method.
func (m *MockTxStore) GetRole(arg0 context.Context, arg1 int64) (db.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0, arg1)
	ret0, _ := ret[0].(db.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockTxStoreMockRecorder) GetRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockTxStore)(nil).GetRole), arg0, arg1)
}

// GetRoles mocks base method.
func (m *MockTxStore) GetRoles(arg0 context.Context, arg1 db.GetRolesParams) ([]db.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoles", arg0, arg1)
	ret0, _ := ret[0].([]db.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoles indicates an expected call of GetRoles.
func (mr *MockTxStoreMockRecorder) GetRoles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoles", reflect.TypeOf((*MockTxStore)(nil).GetRoles), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockTxStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockTxStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockTxStore)(nil).GetUser), arg0, arg1)
}

// GetUserPayment mocks base method.
func (m *MockTxStore) GetUserPayment(arg0 context.Context, arg1 int64) (db.UserPayment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPayment", arg0, arg1)
	ret0, _ := ret[0].(db.UserPayment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPayment indicates an expected call of GetUserPayment.
func (mr *MockTxStoreMockRecorder) GetUserPayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPayment", reflect.TypeOf((*MockTxStore)(nil).GetUserPayment), arg0, arg1)
}

// GetUserPayments mocks base method.
func (m *MockTxStore) GetUserPayments(arg0 context.Context, arg1 db.GetUserPaymentsParams) ([]db.UserPayment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPayments", arg0, arg1)
	ret0, _ := ret[0].([]db.UserPayment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPayments indicates an expected call of GetUserPayments.
func (mr *MockTxStoreMockRecorder) GetUserPayments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPayments", reflect.TypeOf((*MockTxStore)(nil).GetUserPayments), arg0, arg1)
}

// GetUserRole mocks base method.
func (m *MockTxStore) GetUserRole(arg0 context.Context, arg1 int64) (db.UsersRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", arg0, arg1)
	ret0, _ := ret[0].(db.UsersRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRole indicates an expected call of GetUserRole.
func (mr *MockTxStoreMockRecorder) GetUserRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockTxStore)(nil).GetUserRole), arg0, arg1)
}

// GetUserRolesByUserIdAndRoleId mocks base method.
func (m *MockTxStore) GetUserRolesByUserIdAndRoleId(arg0 context.Context, arg1 db.GetUserRolesByUserIdAndRoleIdParams) (db.UsersRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRolesByUserIdAndRoleId", arg0, arg1)
	ret0, _ := ret[0].(db.UsersRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRolesByUserIdAndRoleId indicates an expected call of GetUserRolesByUserIdAndRoleId.
func (mr *MockTxStoreMockRecorder) GetUserRolesByUserIdAndRoleId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRolesByUserIdAndRoleId", reflect.TypeOf((*MockTxStore)(nil).GetUserRolesByUserIdAndRoleId), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockTxStore) GetUsers(arg0 context.Context, arg1 db.GetUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockTxStoreMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockTxStore)(nil).GetUsers), arg0, arg1)
}

// GetUsersRoles mocks base method.
func (m *MockTxStore) GetUsersRoles(arg0 context.Context, arg1 db.GetUsersRolesParams) ([]db.UsersRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersRoles", arg0, arg1)
	ret0, _ := ret[0].([]db.UsersRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersRoles indicates an expected call of GetUsersRoles.
func (mr *MockTxStoreMockRecorder) GetUsersRoles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersRoles", reflect.TypeOf((*MockTxStore)(nil).GetUsersRoles), arg0, arg1)
}

// UpdateClient mocks base method.
func (m *MockTxStore) UpdateClient(arg0 context.Context, arg1 db.UpdateClientParams) (db.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClient", arg0, arg1)
	ret0, _ := ret[0].(db.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateClient indicates an expected call of UpdateClient.
func (mr *MockTxStoreMockRecorder) UpdateClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClient", reflect.TypeOf((*MockTxStore)(nil).UpdateClient), arg0, arg1)
}

// UpdatePermission mocks base method.
func (m *MockTxStore) UpdatePermission(arg0 context.Context, arg1 db.UpdatePermissionParams) (db.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePermission", arg0, arg1)
	ret0, _ := ret[0].(db.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePermission indicates an expected call of UpdatePermission.
func (mr *MockTxStoreMockRecorder) UpdatePermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePermission", reflect.TypeOf((*MockTxStore)(nil).UpdatePermission), arg0, arg1)
}

// UpdateRequest mocks base method.
func (m *MockTxStore) UpdateRequest(arg0 context.Context, arg1 db.UpdateRequestParams) (db.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequest", arg0, arg1)
	ret0, _ := ret[0].(db.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRequest indicates an expected call of UpdateRequest.
func (mr *MockTxStoreMockRecorder) UpdateRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequest", reflect.TypeOf((*MockTxStore)(nil).UpdateRequest), arg0, arg1)
}

// UpdateRole mocks base method.
func (m *MockTxStore) UpdateRole(arg0 context.Context, arg1 db.UpdateRoleParams) (db.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", arg0, arg1)
	ret0, _ := ret[0].(db.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockTxStoreMockRecorder) UpdateRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockTxStore)(nil).UpdateRole), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockTxStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockTxStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockTxStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateUserPayment mocks base method.
func (m *MockTxStore) UpdateUserPayment(arg0 context.Context, arg1 db.UpdateUserPaymentParams) (db.UserPayment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPayment", arg0, arg1)
	ret0, _ := ret[0].(db.UserPayment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPayment indicates an expected call of UpdateUserPayment.
func (mr *MockTxStoreMockRecorder) UpdateUserPayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPayment", reflect.TypeOf((*MockTxStore)(nil).UpdateUserPayment), arg0, arg1)
}

// UpdateUserRole mocks base method.
func (m *MockTxStore) UpdateUserRole(arg0 context.Context, arg1 db.UpdateUserRoleParams) (db.UsersRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserRole", arg0, arg1)
	ret0, _ := ret[0].(db.UsersRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserRole indicates an expected call of UpdateUserRole.
func (mr *MockTxStoreMockRecorder) UpdateUserRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserRole", reflect.TypeOf((*MockTxStore)(nil).UpdateUserRole), arg0, arg1)
}

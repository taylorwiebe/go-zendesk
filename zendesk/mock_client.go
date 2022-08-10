// Code generated by mockery v2.14.0. DO NOT EDIT.

package zendesk

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the MockClient type
type MockClient struct {
	mock.Mock
}

// AddUserTags provides a mock function with given fields: _a0, _a1
func (_m *MockClient) AddUserTags(_a0 int64, _a1 []string) ([]string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []string
	if rf, ok := ret.Get(0).(func(int64, []string) []string); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AutocompleteOrganizations provides a mock function with given fields: _a0
func (_m *MockClient) AutocompleteOrganizations(_a0 string) ([]Organization, error) {
	ret := _m.Called(_a0)

	var r0 []Organization
	if rf, ok := ret.Get(0).(func(string) []Organization); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchUpdateManyTickets provides a mock function with given fields: _a0
func (_m *MockClient) BatchUpdateManyTickets(_a0 []Ticket) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]Ticket) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BulkUpdateManyTickets provides a mock function with given fields: _a0, _a1
func (_m *MockClient) BulkUpdateManyTickets(_a0 []int64, _a1 *Ticket) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]int64, *Ticket) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateArticle provides a mock function with given fields: _a0, _a1
func (_m *MockClient) CreateArticle(_a0 int64, _a1 *Article) (*Article, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Article
	if rf, ok := ret.Get(0).(func(int64, *Article) *Article); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *Article) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCategory provides a mock function with given fields: _a0
func (_m *MockClient) CreateCategory(_a0 *Category) (*Category, error) {
	ret := _m.Called(_a0)

	var r0 *Category
	if rf, ok := ret.Get(0).(func(*Category) *Category); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Category) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateIdentity provides a mock function with given fields: _a0, _a1
func (_m *MockClient) CreateIdentity(_a0 int64, _a1 *UserIdentity) (*UserIdentity, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *UserIdentity
	if rf, ok := ret.Get(0).(func(int64, *UserIdentity) *UserIdentity); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UserIdentity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *UserIdentity) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrUpdateOrganization provides a mock function with given fields: _a0
func (_m *MockClient) CreateOrUpdateOrganization(_a0 *Organization) (*Organization, error) {
	ret := _m.Called(_a0)

	var r0 *Organization
	if rf, ok := ret.Get(0).(func(*Organization) *Organization); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Organization) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrUpdateUser provides a mock function with given fields: _a0
func (_m *MockClient) CreateOrUpdateUser(_a0 *User) (*User, error) {
	ret := _m.Called(_a0)

	var r0 *User
	if rf, ok := ret.Get(0).(func(*User) *User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrganization provides a mock function with given fields: _a0
func (_m *MockClient) CreateOrganization(_a0 *Organization) (*Organization, error) {
	ret := _m.Called(_a0)

	var r0 *Organization
	if rf, ok := ret.Get(0).(func(*Organization) *Organization); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Organization) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrganizationMembership provides a mock function with given fields: _a0
func (_m *MockClient) CreateOrganizationMembership(_a0 *OrganizationMembership) (*OrganizationMembership, error) {
	ret := _m.Called(_a0)

	var r0 *OrganizationMembership
	if rf, ok := ret.Get(0).(func(*OrganizationMembership) *OrganizationMembership); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*OrganizationMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*OrganizationMembership) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePermissionGroup provides a mock function with given fields: _a0
func (_m *MockClient) CreatePermissionGroup(_a0 *PermissionGroup) (*PermissionGroup, error) {
	ret := _m.Called(_a0)

	var r0 *PermissionGroup
	if rf, ok := ret.Get(0).(func(*PermissionGroup) *PermissionGroup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PermissionGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*PermissionGroup) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSection provides a mock function with given fields: _a0, _a1
func (_m *MockClient) CreateSection(_a0 int64, _a1 *Section) (*Section, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Section
	if rf, ok := ret.Get(0).(func(int64, *Section) *Section); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *Section) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTicket provides a mock function with given fields: _a0
func (_m *MockClient) CreateTicket(_a0 *Ticket) (*Ticket, error) {
	ret := _m.Called(_a0)

	var r0 *Ticket
	if rf, ok := ret.Get(0).(func(*Ticket) *Ticket); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Ticket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Ticket) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: _a0
func (_m *MockClient) CreateUser(_a0 *User) (*User, error) {
	ret := _m.Called(_a0)

	var r0 *User
	if rf, ok := ret.Get(0).(func(*User) *User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteArticle provides a mock function with given fields: _a0
func (_m *MockClient) DeleteArticle(_a0 int64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCategory provides a mock function with given fields: _a0
func (_m *MockClient) DeleteCategory(_a0 int64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteIdentity provides a mock function with given fields: _a0, _a1
func (_m *MockClient) DeleteIdentity(_a0 int64, _a1 int64) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrganization provides a mock function with given fields: _a0
func (_m *MockClient) DeleteOrganization(_a0 int64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrganizationMembershipByID provides a mock function with given fields: _a0
func (_m *MockClient) DeleteOrganizationMembershipByID(_a0 int64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePermissionGroup provides a mock function with given fields: _a0
func (_m *MockClient) DeletePermissionGroup(_a0 int64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSection provides a mock function with given fields: _a0
func (_m *MockClient) DeleteSection(_a0 int64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTicket provides a mock function with given fields: _a0
func (_m *MockClient) DeleteTicket(_a0 int64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: _a0
func (_m *MockClient) DeleteUser(_a0 int64) (*User, error) {
	ret := _m.Called(_a0)

	var r0 *User
	if rf, ok := ret.Get(0).(func(int64) *User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListIdentities provides a mock function with given fields: _a0
func (_m *MockClient) ListIdentities(_a0 int64) ([]UserIdentity, error) {
	ret := _m.Called(_a0)

	var r0 []UserIdentity
	if rf, ok := ret.Get(0).(func(int64) []UserIdentity); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]UserIdentity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLocales provides a mock function with given fields:
func (_m *MockClient) ListLocales() ([]Locale, error) {
	ret := _m.Called()

	var r0 []Locale
	if rf, ok := ret.Get(0).(func() []Locale); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Locale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrganizationMembershipsByUserID provides a mock function with given fields: id
func (_m *MockClient) ListOrganizationMembershipsByUserID(id int64) ([]OrganizationMembership, error) {
	ret := _m.Called(id)

	var r0 []OrganizationMembership
	if rf, ok := ret.Get(0).(func(int64) []OrganizationMembership); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]OrganizationMembership)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrganizationTickets provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockClient) ListOrganizationTickets(_a0 int64, _a1 *ListOptions, _a2 ...SideLoad) (*ListResponse, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ListResponse
	if rf, ok := ret.Get(0).(func(int64, *ListOptions, ...SideLoad) *ListResponse); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *ListOptions, ...SideLoad) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrganizationUsers provides a mock function with given fields: _a0, _a1
func (_m *MockClient) ListOrganizationUsers(_a0 int64, _a1 *ListUsersOptions) ([]User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []User
	if rf, ok := ret.Get(0).(func(int64, *ListUsersOptions) []User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *ListUsersOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrganizations provides a mock function with given fields: _a0
func (_m *MockClient) ListOrganizations(_a0 *ListOptions) ([]Organization, error) {
	ret := _m.Called(_a0)

	var r0 []Organization
	if rf, ok := ret.Get(0).(func(*ListOptions) []Organization); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ListOptions) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRequestedTickets provides a mock function with given fields: _a0
func (_m *MockClient) ListRequestedTickets(_a0 int64) ([]Ticket, error) {
	ret := _m.Called(_a0)

	var r0 []Ticket
	if rf, ok := ret.Get(0).(func(int64) []Ticket); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Ticket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTicketComments provides a mock function with given fields: _a0
func (_m *MockClient) ListTicketComments(_a0 int64) ([]TicketComment, error) {
	ret := _m.Called(_a0)

	var r0 []TicketComment
	if rf, ok := ret.Get(0).(func(int64) []TicketComment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TicketComment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTicketFields provides a mock function with given fields:
func (_m *MockClient) ListTicketFields() ([]TicketField, error) {
	ret := _m.Called()

	var r0 []TicketField
	if rf, ok := ret.Get(0).(func() []TicketField); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TicketField)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTicketIncidents provides a mock function with given fields: _a0
func (_m *MockClient) ListTicketIncidents(_a0 int64) ([]Ticket, error) {
	ret := _m.Called(_a0)

	var r0 []Ticket
	if rf, ok := ret.Get(0).(func(int64) []Ticket); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Ticket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: _a0
func (_m *MockClient) ListUsers(_a0 *ListUsersOptions) ([]User, error) {
	ret := _m.Called(_a0)

	var r0 []User
	if rf, ok := ret.Get(0).(func(*ListUsersOptions) []User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ListUsersOptions) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentlyDeleteTicket provides a mock function with given fields: _a0
func (_m *MockClient) PermanentlyDeleteTicket(_a0 int64) (*JobStatus, error) {
	ret := _m.Called(_a0)

	var r0 *JobStatus
	if rf, ok := ret.Get(0).(func(int64) *JobStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*JobStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentlyDeleteUser provides a mock function with given fields: _a0
func (_m *MockClient) PermanentlyDeleteUser(_a0 int64) (*User, error) {
	ret := _m.Called(_a0)

	var r0 *User
	if rf, ok := ret.Get(0).(func(int64) *User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RedactCommentString provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockClient) RedactCommentString(_a0 int64, _a1 int64, _a2 string) (*TicketComment, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *TicketComment
	if rf, ok := ret.Get(0).(func(int64, int64, string) *TicketComment); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TicketComment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchOrganizationsByExternalID provides a mock function with given fields: _a0
func (_m *MockClient) SearchOrganizationsByExternalID(_a0 string) ([]Organization, error) {
	ret := _m.Called(_a0)

	var r0 []Organization
	if rf, ok := ret.Get(0).(func(string) []Organization); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchTickets provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockClient) SearchTickets(_a0 string, _a1 *ListOptions, _a2 ...Filters) (*TicketSearchResults, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *TicketSearchResults
	if rf, ok := ret.Get(0).(func(string, *ListOptions, ...Filters) *TicketSearchResults); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TicketSearchResults)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *ListOptions, ...Filters) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchUserByExternalID provides a mock function with given fields: _a0
func (_m *MockClient) SearchUserByExternalID(_a0 string) (*User, error) {
	ret := _m.Called(_a0)

	var r0 *User
	if rf, ok := ret.Get(0).(func(string) *User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchUsers provides a mock function with given fields: _a0
func (_m *MockClient) SearchUsers(_a0 string) ([]User, error) {
	ret := _m.Called(_a0)

	var r0 []User
	if rf, ok := ret.Get(0).(func(string) []User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowArticle provides a mock function with given fields: _a0
func (_m *MockClient) ShowArticle(_a0 int64) (*Article, error) {
	ret := _m.Called(_a0)

	var r0 *Article
	if rf, ok := ret.Get(0).(func(int64) *Article); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowCategory provides a mock function with given fields: _a0
func (_m *MockClient) ShowCategory(_a0 int64) (*Category, error) {
	ret := _m.Called(_a0)

	var r0 *Category
	if rf, ok := ret.Get(0).(func(int64) *Category); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowComplianceDeletionStatuses provides a mock function with given fields: _a0
func (_m *MockClient) ShowComplianceDeletionStatuses(_a0 int64) ([]ComplianceDeletionStatus, error) {
	ret := _m.Called(_a0)

	var r0 []ComplianceDeletionStatus
	if rf, ok := ret.Get(0).(func(int64) []ComplianceDeletionStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ComplianceDeletionStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowIdentity provides a mock function with given fields: _a0, _a1
func (_m *MockClient) ShowIdentity(_a0 int64, _a1 int64) (*UserIdentity, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *UserIdentity
	if rf, ok := ret.Get(0).(func(int64, int64) *UserIdentity); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UserIdentity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowJobStatus provides a mock function with given fields: _a0
func (_m *MockClient) ShowJobStatus(_a0 string) (*JobStatus, error) {
	ret := _m.Called(_a0)

	var r0 *JobStatus
	if rf, ok := ret.Get(0).(func(string) *JobStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*JobStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowLocale provides a mock function with given fields: _a0
func (_m *MockClient) ShowLocale(_a0 int64) (*Locale, error) {
	ret := _m.Called(_a0)

	var r0 *Locale
	if rf, ok := ret.Get(0).(func(int64) *Locale); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Locale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowLocaleByCode provides a mock function with given fields: _a0
func (_m *MockClient) ShowLocaleByCode(_a0 string) (*Locale, error) {
	ret := _m.Called(_a0)

	var r0 *Locale
	if rf, ok := ret.Get(0).(func(string) *Locale); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Locale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowManyUsers provides a mock function with given fields: _a0
func (_m *MockClient) ShowManyUsers(_a0 []int64) ([]User, error) {
	ret := _m.Called(_a0)

	var r0 []User
	if rf, ok := ret.Get(0).(func([]int64) []User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowOrganization provides a mock function with given fields: _a0
func (_m *MockClient) ShowOrganization(_a0 int64) (*Organization, error) {
	ret := _m.Called(_a0)

	var r0 *Organization
	if rf, ok := ret.Get(0).(func(int64) *Organization); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowPermissionGroup provides a mock function with given fields: _a0
func (_m *MockClient) ShowPermissionGroup(_a0 int64) (*PermissionGroup, error) {
	ret := _m.Called(_a0)

	var r0 *PermissionGroup
	if rf, ok := ret.Get(0).(func(int64) *PermissionGroup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PermissionGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowSection provides a mock function with given fields: _a0
func (_m *MockClient) ShowSection(_a0 int64) (*Section, error) {
	ret := _m.Called(_a0)

	var r0 *Section
	if rf, ok := ret.Get(0).(func(int64) *Section); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowTicket provides a mock function with given fields: _a0
func (_m *MockClient) ShowTicket(_a0 int64) (*Ticket, error) {
	ret := _m.Called(_a0)

	var r0 *Ticket
	if rf, ok := ret.Get(0).(func(int64) *Ticket); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Ticket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowUser provides a mock function with given fields: _a0
func (_m *MockClient) ShowUser(_a0 int64) (*User, error) {
	ret := _m.Called(_a0)

	var r0 *User
	if rf, ok := ret.Get(0).(func(int64) *User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateArticle provides a mock function with given fields: _a0, _a1
func (_m *MockClient) UpdateArticle(_a0 int64, _a1 *Article) (*Article, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Article
	if rf, ok := ret.Get(0).(func(int64, *Article) *Article); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *Article) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCategory provides a mock function with given fields: _a0, _a1
func (_m *MockClient) UpdateCategory(_a0 int64, _a1 *Category) (*Category, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Category
	if rf, ok := ret.Get(0).(func(int64, *Category) *Category); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Category)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *Category) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIdentity provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockClient) UpdateIdentity(_a0 int64, _a1 int64, _a2 *UserIdentity) (*UserIdentity, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *UserIdentity
	if rf, ok := ret.Get(0).(func(int64, int64, *UserIdentity) *UserIdentity); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UserIdentity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, *UserIdentity) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrganization provides a mock function with given fields: _a0, _a1
func (_m *MockClient) UpdateOrganization(_a0 int64, _a1 *Organization) (*Organization, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Organization
	if rf, ok := ret.Get(0).(func(int64, *Organization) *Organization); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *Organization) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePermissionGroup provides a mock function with given fields: _a0, _a1
func (_m *MockClient) UpdatePermissionGroup(_a0 int64, _a1 *PermissionGroup) (*PermissionGroup, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *PermissionGroup
	if rf, ok := ret.Get(0).(func(int64, *PermissionGroup) *PermissionGroup); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PermissionGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *PermissionGroup) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSection provides a mock function with given fields: _a0, _a1
func (_m *MockClient) UpdateSection(_a0 int64, _a1 *Section) (*Section, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Section
	if rf, ok := ret.Get(0).(func(int64, *Section) *Section); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Section)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *Section) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTicket provides a mock function with given fields: _a0, _a1
func (_m *MockClient) UpdateTicket(_a0 int64, _a1 *Ticket) (*Ticket, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *Ticket
	if rf, ok := ret.Get(0).(func(int64, *Ticket) *Ticket); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Ticket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *Ticket) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0, _a1
func (_m *MockClient) UpdateUser(_a0 int64, _a1 *User) (*User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *User
	if rf, ok := ret.Get(0).(func(int64, *User) *User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadFile provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockClient) UploadFile(_a0 string, _a1 *string, _a2 io.Reader) (*Upload, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *Upload
	if rf, ok := ret.Get(0).(func(string, *string, io.Reader) *Upload); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Upload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *string, io.Reader) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithHeader provides a mock function with given fields: name, value
func (_m *MockClient) WithHeader(name string, value string) Client {
	ret := _m.Called(name, value)

	var r0 Client
	if rf, ok := ret.Get(0).(func(string, string) Client); ok {
		r0 = rf(name, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Client)
		}
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockClient(t mockConstructorTestingTNewClient) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

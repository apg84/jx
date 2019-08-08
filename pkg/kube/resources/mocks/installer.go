// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/pkg/kube/resources (interfaces: Installer)

package resources_test

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockInstaller struct {
	fail func(message string, callerSkip ...int)
}

func NewMockInstaller(options ...pegomock.Option) *MockInstaller {
	mock := &MockInstaller{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockInstaller) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockInstaller) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockInstaller) Install(_param0 string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockInstaller().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Install", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockInstaller) InstallDir(_param0 string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockInstaller().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("InstallDir", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockInstaller) VerifyWasCalledOnce() *VerifierMockInstaller {
	return &VerifierMockInstaller{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockInstaller) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockInstaller {
	return &VerifierMockInstaller{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockInstaller) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockInstaller {
	return &VerifierMockInstaller{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockInstaller) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockInstaller {
	return &VerifierMockInstaller{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockInstaller struct {
	mock                   *MockInstaller
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockInstaller) Install(_param0 string) *MockInstaller_Install_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Install", params, verifier.timeout)
	return &MockInstaller_Install_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockInstaller_Install_OngoingVerification struct {
	mock              *MockInstaller
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockInstaller_Install_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockInstaller_Install_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockInstaller) InstallDir(_param0 string) *MockInstaller_InstallDir_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "InstallDir", params, verifier.timeout)
	return &MockInstaller_InstallDir_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockInstaller_InstallDir_OngoingVerification struct {
	mock              *MockInstaller
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockInstaller_InstallDir_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockInstaller_InstallDir_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}
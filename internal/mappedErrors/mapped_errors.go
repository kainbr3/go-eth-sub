package mappederrors

import (
	"fmt"
	"strings"
)

type MappedError struct {
	module      int
	code        int
	parentError string
	text        string
	message     string
}

// Format formats the message with the suplied values
func (m *MappedError) Format(values ...interface{}) *MappedError {
	return &MappedError{
		module:  m.module,
		code:    m.code,
		message: fmt.Sprintf(m.message, values...),
	}
}

func New(code int, text string) *MappedError {
	return &MappedError{
		code: code,
		text: text,
	}
}

func (m *MappedError) BuildErrorMessage(module int, err error) *MappedError {
	m.module = module
	if err != nil {
		m.parentError = err.Error()
		if strings.Contains(m.message, "ETH-SUB") {
			removeMessageFromString(&m.message, m.message)
		}
	}
	m.message = fmt.Sprintf("%s %s", m.Code(), m.text)
	return m
}

func removeMessageFromString(messages *string, messageToRemove string) {
	*messages = strings.ReplaceAll(*messages, messageToRemove, "")
}

func (k *MappedError) BuildLogMessage(module int, err error) error {
	k.module = module
	return err
}

func (m *MappedError) Message() string {
	return m.message
}

func (m *MappedError) ParentError() string {
	return m.parentError
}

func (m *MappedError) Code() string {
	return fmt.Sprintf("ETH-SUB-%03d%03d", m.module, m.code)
}

func (m *MappedError) Module() int {
	return m.module
}

func (m *MappedError) Text() string {
	return m.text
}

func (m *MappedError) Error() string {
	return fmt.Sprintf("ETH-SUB-%03d%03d %s", m.module, m.code, m.text)
}

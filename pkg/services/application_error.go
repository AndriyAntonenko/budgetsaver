package service

type ServiceError struct {
	Id      string
	Message string
}

func NewServiceError(id string, message string) *ServiceError {
	return &ServiceError{
		Id:      id,
		Message: message,
	}
}

func (se *ServiceError) Error() string {
	return se.Message
}

const (
	UnexpectedError                = "UNEXPECTED"
	UnknownFinanceGroupMemberError = "UNKNOWN_GROUP_MEMBER"
	ActionForbiddenError           = "ACTION_FORBIDDEN"
)

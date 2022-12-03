package domain

type DomainError struct {
	msg string
}

func (err *DomainError) Error() string {
	return "DomainError: " + err.msg
}

func IsRequired(field string) error {
	return &DomainError{field + " is required"}
}

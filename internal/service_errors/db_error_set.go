package service_errors

import (
	"errors"

	"google.golang.org/grpc/codes"
)

// ErrDbConnFailure outlines connection failures against databases of interest
var ErrDbConnFailure *DatabaseError = &DatabaseError{
	err: &ServiceError{
		Code:           ErrorCodeDatabaseConnectionFailure,
		Err:            errors.New("failed to connect to a database"),
		GRPCStatusCode: codes.Internal,
		Type:           ErrorTypeDependencyConnection,
	},
}

// ErrInvalidGormDbOject outlines nil or invalid gorm database connection object
var ErrInvalidGormDbOject *DatabaseError = &DatabaseError{
	err: &ServiceError{
		Code:                  ErrorCodeDatabaseConnectionFailure,
		Err:                   errors.New("invalid gorm database connection object"),
		GRPCStatusCode:        codes.Internal,
		Msg:                   "",
		Param:                 "",
		RequestID:             "",
		Type:                  ErrorTypeDependencyConnection,
		OAuthError:            "",
		OAuthErrorDescription: "",
	},
}

package service_errors

import (
	"errors"

	"google.golang.org/grpc/codes"
)

var ErrInvalidAcctParam *FieldConfigurationError = &FieldConfigurationError{
	err: &ServiceError{
		Code:                  ErrorCodeInvalidConfigurationParameters,
		Err:                   errors.New("invalid account parameters"),
		GRPCStatusCode:        codes.Internal,
		Msg:                   "",
		Param:                 "",
		RequestID:             "",
		Type:                  ErrorTypeParameters,
		OAuthError:            "",
		OAuthErrorDescription: "",
	},
}

var ErrInvalidInputParam *FieldConfigurationError = &FieldConfigurationError{
	err: &ServiceError{
		Code:                  ErrorCodeInvalidConfigurationParameters,
		Err:                   errors.New("invalid parameters. please check input parameters and assert they are of expected value"),
		GRPCStatusCode:        codes.Internal,
		Msg:                   "",
		Param:                 "",
		RequestID:             "",
		Type:                  ErrorTypeParameters,
		OAuthError:            "",
		OAuthErrorDescription: "",
	},
}

var ErrInvalidDbObject *FieldConfigurationError = &FieldConfigurationError{
	err: &ServiceError{
		Code:                  ErrorCodeInvalidConfigurationParameters,
		Err:                   errors.New("invalid parameters. please check database object and assert its parameters are of expected value"),
		GRPCStatusCode:        codes.Internal,
		Msg:                   "",
		Param:                 "",
		RequestID:             "",
		Type:                  ErrorTypeParameters,
		OAuthError:            "",
		OAuthErrorDescription: "",
	},
}

package mappederrors

const (
	ErrUnexpectedCode = iota + 1
	ErrJsonUnmarshalCode
	ErrJsonMarshalCode
	ErrParseQueryParamCode
	ErrParseURLPathParamCode
	ErrParseBodyParamCode
	ErrStructValidationCode
	ErrStorageGetCode
	ErrStorageSetCode
	ErrStorageSetDuplicateCode
	ErrStorageDeleteCode
	ErrInvalidEthereumAddressCode
	ErrRequestCode
	ErrNotFoundCode
)

var (
	ErrUnexpected             = New(ErrUnexpectedCode, "unexpected error")
	ErrJsonUnmarshal          = New(ErrJsonUnmarshalCode, "json: error unmarshalling json")
	ErrJsonMarshal            = New(ErrJsonMarshalCode, "json: error marshalling json")
	ErrParseQueryParam        = New(ErrParseQueryParamCode, "error parsing query parameter")
	ErrParseURLPathParam      = New(ErrParseURLPathParamCode, "error parsing url path parameter")
	ErrParseBodyParam         = New(ErrParseBodyParamCode, "error parsing body parameter")
	ErrStructValidation       = New(ErrStructValidationCode, "error validating struct")
	ErrStorageGet             = New(ErrStorageGetCode, "error retrieving data from storage")
	ErrStorageSet             = New(ErrStorageSetCode, "error persisting data from storage")
	ErrStorageSetDuplicate    = New(ErrStorageSetDuplicateCode, "already exists")
	ErrStorageDelete          = New(ErrStorageDeleteCode, "error deleting data from storage")
	ErrInvalidEthereumAddress = New(ErrInvalidEthereumAddressCode, "error address not exists")
	ErrRequest                = New(ErrRequestCode, "error while executing a request")
	ErrNotFound               = New(ErrNotFoundCode, "not found")
)

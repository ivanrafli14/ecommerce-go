package response

var ErrorMapping = map[string]Error{
	ErrGeneral.Error():    ErrorGeneral,
	ErrRepository.Error(): ErrorRepository,
	ErrBadRequest.Error(): ErrorBadRequest,

	ErrEmailNotFound.Error():         ErrorEmailNotFound,
	ErrEmailInvalid.Error():          ErrorEmailInvalid,
	ErrEmailRequired.Error():         ErrorEmailRequired,
	ErrPasswordRequired.Error():      ErrorPasswordRequired,
	ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
	ErrEmailAlreadyUsed.Error():      ErrorEmailAlreadyUsed,
	ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,

	ErrGenderRequired.Error():           ErrorGenderRequired,
	ErrGenderInvalid.Error():            ErrorGenderInvalid,
	ErrPhoneNumberRequired.Error():      ErrorPhoneNumberRequired,
	ErrPhoneNumberInvalidLength.Error(): ErrorPhoneNumberInvalidLength,
	ErrNameRequired.Error():             ErrorNameRequired,
	ErrAddressRequired.Error():          ErrorAddressRequired,
	ErrDateOfBirthRequired.Error():      ErrorDateOfBirthRequired,
	ErrDateOfBirthInvalid.Error():       ErrorDateOfBirthInvalid,
	ErrImageUrlRequired.Error():         ErrorImageUrlRequired,
	ErrRoleInvalid.Error():              ErrorRoleInvalid,
	ErrUserNotFound.Error():             ErrorUserNotFound,
	ErrJWTExpired.Error():               ErrorJWTExpired,
	ErrJWTEmpty.Error():                 ErrorJWTEmpty,
	ErrJWTInvalid.Error():               ErrorJWTInvalid,
	ErrUserAlreadyExists.Error():        ErrorUserAlreadyExists,

	ErrNameMerchantRequired.Error():          ErrorNameMerchantRequired,
	ErrAddressMerchantRequired.Error():       ErrorAddressMerchantRequired,
	ErrPhoneNumberMerchantRequired.Error():   ErrorPhoneMerchantRequired,
	ErrPhoneNumberMerchantInvalidLen.Error(): ErrorPhoneMerchantLength,
	ErrImageUrlMerchantRequired.Error():      ErrorImageUrlMerchantRequired,
	ErrCityMerchantRequired.Error():          ErrorCityMerchantRequired,
	ErrMerchantNotFound.Error():              ErrorMerchantNotFound,
	ErrMerchantAlreadyExits.Error():          ErrorMerchantAlreadyExits,

	ErrPriceProductRequired.Error():       ErrorPriceProductInvalid,
	ErrPriceProductInvalid.Error():        ErrorPriceProductRequired,
	ErrStockProductRequired.Error():       ErrorStockProductRequired,
	ErrStockProductInvalid.Error():        ErrorStockProductInvalid,
	ErrNameProductRequired.Error():        ErrorNameProductRequired,
	ErrDescriptionProductRequired.Error(): ErrorDescriptionProductRequired,
	ErrImageUrlProductRequired.Error():    ErrorImageUrlProductRequired,
	ErrCategoryIDProductRequired.Error():  ErrorCategoryIDProductRequired,
	ErrCategoryIDProductNotFound.Error():  ErrorCategoryIDProductNotFound,
	ErrProductNotFound.Error():            ErrorProductNotFound,

	ErrInvalidQuantity.Error():        ErrorInvalidQuantity,
	ErrMaxReachQuantity.Error():       ErrorMaxReachQuantity,
	ErrProductOrderNotFound.Error():   ErrorProductOrderNotFound,
	ErrProductIDOrderRequired.Error(): ErrorProductIDOrderRequired,
	ErrQuantityOrderRequired.Error():  ErrorQuantityOrderRequired,
}

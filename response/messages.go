package response

const (
	USER_EXISTS           = "A user with this email already exists."
	USER_NOEXIST          = "A user with this email or id doesn't exist."
	INCORRECT_CREDENTIALS = "The specified credentials are incorrect."
	UNAUTHORIZED          = "You have failed the authorization."
	INVALID_TOKEN         = "The authorization token is invalid."
	INVALID_RESET_TOKEN   = "The rese token is invalid."
)

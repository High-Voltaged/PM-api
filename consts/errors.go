package consts

const (
	USER_EXISTS            = "A user with this email already exists."
	USER_NOEXIST           = "A user with this email or id doesn't exist."
	INCORRECT_CREDENTIALS  = "The specified credentials are incorrect."
	UNAUTHORIZED           = "You have failed the authorization."
	INVALID_TOKEN          = "The authorization token is invalid."
	INVALID_RESET_TOKEN    = "The reset token is invalid."
	PROJECTS_NOT_FOUND     = "The projects for the currently logged in user weren't found."
	INVALID_PROJECT_ID     = "Please provide a valid id for the project to update."
	INVALID_PROJECT_ACTION = "Either you can't edit this project, or it doesn't exist."
)

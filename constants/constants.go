// constants package contains error messages and success messages used in various projects.
package constants

// Error messages
const (
	NilUser                     = "User pointer not initialized."
	JSONError                   = "Error occurred while creating JSON."
	NilTpl                      = "Template pointer is not initialized."
	NilMongoClient              = "DB Client is not initialized."
	DatabaseNotConnected        = "Unable to connect to database."
	CollectionNotFound          = "Collection not found."
	NilFireStoreHelper          = "FileStore helper is not initialized"
	NilMongoHelper              = "Mongo helper is not initialized."
	InvalidString               = "Enter a string value."
	InvalidInteger              = "Enter an integer value."
	InvalidName                 = "Name must be at least 5 characters."
	InvalidPassword             = "Password must be at least 8 characters."
	InvalidEmail                = "Invalid email address."
	EmailIDIsAvailable          = "Email ID is available."
	AccountVerificationLinkSent = "An account verification e-mail has been sent to your account."
	LoginSuccess                = "Login successful."
	AccountAlreadyExists        = "Email ID is already in use."
	AccountNotExists            = "No record found with this email ID."
	RegistrationError           = "Unable to register account."
	LoginError                  = "Username/password combination is incorrect, please retry or click on forgot password."
	InactiveAccountError        = "Verify your email ID to proceed login"
	InvalidDataFormat           = "The data is in incorrect format."
	ClientAuthenticationError   = "Invalid client ID"
	InvalidUserId               = "Invalid user ID"
	InvalidCategoryId           = "Invalid category ID"
	NoRecordFound               = "No record found."
	InvalidArguments            = "Invalid arguments"
	InvalidPageNumber           = "Invalid page number"
	InvalidRecordsCount         = "Invalid record count"
	UnableToUpdate              = "Unable to update record."
	UnableToDelete              = "Unable to delete record."
	UnableToAdd                 = "Unable to add record."
)

// Success messages
const (
	RecordUpdated  = "Record updated successfully."
	RecordDeleted  = "Record deleted successfully."
	RecordAdded    = "Record added successfully."
	ListingSuccess = "The records fetched successfully."
)

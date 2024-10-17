package response

const (
	ErrorCodeSuccess      = 20001 //Success
	ErrorCodeParamInvalid = 20002 //Param invalid
	ErrorCodeEmailInvalid = 20003 //Param invalid

	ErrorCodeTokenInvalid = 30000 //invalid auth token
)

var mgs = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeEmailInvalid: "Email is invalid",
	ErrorCodeTokenInvalid: "token is invalid",
}

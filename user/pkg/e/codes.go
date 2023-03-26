package e

var codes = map[uint32]string{}

var (
	SUCCESS				=		newErrCode(0,	  "Success")	
	UserNotExist		=		newErrCode(10001, "User does not exist")
	UserExist			=		newErrCode(10002, "User already exist")
	UserNameInvalid		=		newErrCode(10003, "User name is invalid")
	PasswordInvalid		=		newErrCode(10004, "Password is invalid")
	UserInternalErr		=		newErrCode(10005, "Operate user occurs an internal error")
	UserPasswdNotMatch	=		newErrCode(10006, "User name and password are incorrect")
	ServiceInternalErr	=		newErrCode(11001, "Service internal error")
)


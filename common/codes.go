package common

const (
	SUCCESS         int = 0
	PARAM_ERROR     int = 101
	NEED_LOGIN      int = 201
	NOT_CATCH_ERROR int = 999
	RPC_CALL_FAILED int = 998

	DB_QUERY_FAILED = 206

	ACCESS_TOKEN_CREATE_FAILED   = 301
	ACCESS_TOKEN_TRANSFER_FAILED = 302
	CONFIG_FILE_READ_FAILED      = 303
)

var AppCodes = map[int]string{
	0:   "ok",
	101: "param error",
	201: "need login",
	301: "access token create failed",
	302: "access token transfer failed",
	303: "config read failed",
	999: "unkown exception",
}

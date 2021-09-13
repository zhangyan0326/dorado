package bizerror

type BizError struct {
	Errors  error
	Code    string
	Message string
}

func (e *BizError) Error() string {
	return e.Message
}

func NewBizerror(code string, message string) BizError {
	return BizError{
		Code:    code,
		Message: message,
	}
}

var (
	IsFailed                  = NewBizerror("9999", "failed")
	IsSuccess                 = NewBizerror("0000", "success")
	ParamError                = NewBizerror("80001", "参数错误")
	SignError                 = NewBizerror("80002", "签名错误")
	LimitError                = NewBizerror("80003", "系统限流，请稍后重试")
	HandlerNotFoundError      = NewBizerror("80004", "处理器未找到")
	HandlerPointNilError      = NewBizerror("80005", "处理器指针为空")
	WhiteListError            = NewBizerror("80006", "IP地址非法")
	MerchantError             = NewBizerror("80007", "商户信息错误")
	RsaDecryptError           = NewBizerror("80008", "RSA解密错误")
	PrivateKeyError           = NewBizerror("80009", "获取商户密钥信息失败")
	Base64DecodeError         = NewBizerror("80010", "Base64解析失败")
	X509ParseError            = NewBizerror("80011", "X509解析失败")
	JsonUnmarshalError        = NewBizerror("80012", "json解析失败")
	RsaTextLengthError        = NewBizerror("80013", "ras加密长度错误")
	AesDecryptTextLengthError = NewBizerror("80014", "aes解密长度错误")
	AesEncryptError           = NewBizerror("80015", "加密错误")
	AesDecryptError           = NewBizerror("80015", "aes解密错误")

	JwtSetError         = NewBizerror("80016", "token生成错误")
	JwtCheckError       = NewBizerror("80017", "未传token")
	JwtTokenError       = NewBizerror("80018", "该用户token不存在")
	JwtParameterError   = NewBizerror("80019", "token格式错误")
	JwtTokenIsNullError = NewBizerror("80020", "token不正确")

	AttestationFailure    = NewBizerror("80021", "验签失败")
	JsonUnmarshalMapError = NewBizerror("80022", "json解析Map失败")
	HttpError             = NewBizerror("80023", "http连接错误")
	ParesIntError         = NewBizerror("80024", "字符串转换数字错误")
	ParesUintError        = NewBizerror("80025", "字符串转换数字错误")

	RedisConnError       = NewBizerror("81001", "redis连接失败")
	RpcRequestError      = NewBizerror("81002", "rpc调用异常")
	RpcConnectError      = NewBizerror("81003", "rpc连接异常")
	HttpRemoteError      = NewBizerror("81004", "http请求异常")
	RedisGetKeyError     = NewBizerror("81005", "redis获取key失败")
	D404Error            = NewBizerror("81006", "请求失败,请重新登录")
	RpcServerError       = NewBizerror("81007", "rpc服务失败")
	RpcServerListenError = NewBizerror("81009", "rpc服务监听失败")
	JwtMethodError       = NewBizerror("81010", "token加密方法错误")

	DatabaseInsertError = NewBizerror("82001", "数据库插入失败")

	ActivityRankingError      = NewBizerror("82002", "获取拉新排名失败")
	ActivityRankingListError  = NewBizerror("82003", "获取拉新排名列表失败")
	PosterConfigListError     = NewBizerror("82004", "查询海报配置信息报错")
	PosterTagListError        = NewBizerror("82005", "查询海报标签报错")
	PosterGenerateLogError    = NewBizerror("82006", "查询海报生成记录报错")
	PosterGenerateIdError     = NewBizerror("82007", "海报id不能为空")
	PosterStyleNotFoundError  = NewBizerror("82008", "未查询到海报样式")
	PosterQrcodeNotFoundError = NewBizerror("82009", "未查询到海报链接")
)

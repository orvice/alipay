// @authors     ascoders

package alipay

import (
	"net/url"
	"strconv"
)

type Client struct {
	Partner   string // 合作者ID
	Key       string // 合作者私钥
	ReturnUrl string // 同步返回地址
	NotifyUrl string // 网站异步返回地址
	Email     string // 网站卖家邮箱地址
}

type Result struct {
	// 状态
	Status int
	// 本网站订单号
	OrderNo string
	// 支付宝交易号
	TradeNo string
	// 买家支付宝账号
	BuyerEmail string
	// 错误提示
	Message string
}

// 生成订单的参数
type Options struct {
	OrderId  string  // 订单唯一id
	Fee      float32 // 价格
	NickName string  // 充值账户名称
	Subject  string  // 充值描述
}

/* 生成支付宝即时到帐提交表单html代码 */
func (this *Client) Form(opts Options) string {
	//实例化参数
	param := &AlipayParameters{}
	param.InputCharset = "utf-8"
	param.Body = "为" + opts.NickName + "充值" + strconv.FormatFloat(float64(opts.Fee), 'f', 2, 32) + "元"
	param.NotifyUrl = this.NotifyUrl
	param.OutTradeNo = opts.OrderId
	param.Partner = this.Partner
	param.PaymentType = 1
	param.ReturnUrl = this.ReturnUrl
	param.SellerEmail = this.Email
	param.Service = "create_direct_pay_by_user"
	param.Subject = opts.Subject
	param.TotalFee = opts.Fee

	//生成签名
	sign := sign(param)

	//追加参数
	param.Sign = sign
	param.SignType = "MD5"

	//生成自动提交form
	return `
		<form id="alipaysubmit" name="alipaysubmit" action="https://mapi.alipay.com/gateway.do?_input_charset=utf-8" method="get" style='display:none;'>
			<input type="hidden" name="_input_charset" value="` + param.InputCharset + `">
			<input type="hidden" name="body" value="` + param.Body + `">
			<input type="hidden" name="notify_url" value="` + param.NotifyUrl + `">
			<input type="hidden" name="out_trade_no" value="` + param.OutTradeNo + `">
			<input type="hidden" name="partner" value="` + param.Partner + `">
			<input type="hidden" name="payment_type" value="` + strconv.Itoa(int(param.PaymentType)) + `">
			<input type="hidden" name="return_url" value="` + param.ReturnUrl + `">
			<input type="hidden" name="seller_email" value="` + param.SellerEmail + `">
			<input type="hidden" name="service" value="` + param.Service + `">
			<input type="hidden" name="subject" value="` + param.Subject + `">
			<input type="hidden" name="total_fee" value="` + strconv.FormatFloat(float64(param.TotalFee), 'f', 2, 32) + `">
			<input type="hidden" name="sign" value="` + param.Sign + `">
			<input type="hidden" name="sign_type" value="` + param.SignType + `">
		</form>
		<script>
			document.forms['alipaysubmit'].submit();
		</script>
	`
}

/* 生成支付宝即时到帐提交表单html代码 */
func (this *Client) WebUrl(opts Options) string {
	//实例化参数
	param := &AlipayParameters{}
	param.InputCharset = "utf-8"
	param.Body = "为" + opts.NickName + "充值" + strconv.FormatFloat(float64(opts.Fee), 'f', 2, 32) + "元"
	param.NotifyUrl = this.NotifyUrl
	param.OutTradeNo = opts.OrderId
	param.Partner = this.Partner
	param.PaymentType = 1
	param.ReturnUrl = this.ReturnUrl
	param.SellerEmail = this.Email
	param.Service = "create_direct_pay_by_user"
	param.Subject = opts.Subject
	param.TotalFee = opts.Fee

	//生成签名
	sign := sign(param)

	//追加参数
	param.Sign = sign
	param.SignType = "MD5"

	var baseUrl = "https://mapi.alipay.com/gateway.do"

	var params = map[string]string{
		"_input_charset": param.InputCharset,
		"body":           param.Body,
		"notify_url":     param.NotifyUrl,
		"service":        param.Service,
		"return_url":     param.ReturnUrl,
		"seller_email":   param.SellerEmail,
		"out_trade_no":   param.OutTradeNo,
		"partner":        param.Partner,
		"payment_type":   strconv.Itoa(int(param.PaymentType)),
		"subject":        param.Subject,
		"total_fee":      strconv.FormatFloat(float64(param.TotalFee), 'f', 2, 32),
		"sign":           param.Sign,
		"sign_type":      param.SignType,
	}

	vls := url.Values{}

	for k,v := range params{
		vls.Add(k,v)
	}

	finalUrl := baseUrl + vls.Encode()
	return finalUrl
}

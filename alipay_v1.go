// @authors     ascoders
// 保持旧api可用
// 依赖beego

package alipay

import (
	"strconv"
)

var (
	AlipayPartner  string //合作者ID
	AlipayKey      string //合作者私钥
	WebReturnUrl   string //网站同步返回地址
	WebNotifyUrl   string //网站异步返回地址
	WebSellerEmail string //网站卖家邮箱地址
)

/* 生成支付宝即时到帐提交表单html代码
 * @params string 订单唯一id
 * @params int 价格
 * @params int 获得代金券的数量
 * @params string 充值账户的名称
 * @params string 充值描述
 */
func CreateAlipaySign(orderId string, fee float32, nickname string, subject string) string {
	//实例化参数
	param := &AlipayParameters{}
	param.InputCharset = "utf-8"
	param.Body = "为" + nickname + "充值" + strconv.FormatFloat(float64(fee), 'f', 2, 32) + "元"
	param.NotifyUrl = WebNotifyUrl
	param.OutTradeNo = orderId
	param.Partner = AlipayPartner
	param.PaymentType = 1
	param.ReturnUrl = WebReturnUrl
	param.SellerEmail = WebSellerEmail
	param.Service = "create_direct_pay_by_user"
	param.Subject = subject
	param.TotalFee = fee

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


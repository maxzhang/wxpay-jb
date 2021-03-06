package wxpay

import (
	"encoding/xml"
	"net/http"
	"net/url"
)

//xml 解析
func (this *Client) Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

// GetTradeNotification https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_7&index=3
//微信支付异步验证
func (this *Client) GetTradeNotification(data []byte) (result *TradeNotification, err error) {
	return result, this.verifyResponse(data, &result)
}

// GetRefundNotification 退款结果通知
// docs: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10
func (this *Client) GetRefundNotification(data []byte) (result *RefundNotification, err error) {
	return result, this.verifyResponse(data, &result)
}

//签约解约异步通知
func (this *Client) GetContractNotification(data []byte) (result *ContractNotification, err error) {
	return result, this.verifyResponse(data, &result)
}

//签约扣款异步通知
func (this *Client) GetPayApplyNotification(data []byte) (result *PayApplyNotification, err error) {
	return result, this.verifyResponse(data, &result)
}

//验签
func (this *Client) verifyResponse(data []byte, result interface{}) (err error) {
	key, err := this.getKey()
	if err != nil {
		return err
	}
	if _, err := VerifyResponseData(data, key); err != nil {
		return err
	}
	return this.Unmarshal(data, result)
}

func (this *Client) AckNotification(w http.ResponseWriter) {
	AckNotification(w)
}

func AckNotification(w http.ResponseWriter) {
	var v = url.Values{}
	v.Set("return_code", "SUCCESS")
	v.Set("return_msg", "OK")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(URLValueToXML(v)))
}

package handler

import (
	"bytes"
	"context"
	pb "emailservice/proto"
	"log"
	"net/smtp"
	"strconv"

	"github.com/jordan-wright/email"
)

type DummyEmailService struct{}

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func (c *DummyEmailService) SendOrderConfirmation(ctx context.Context, in *pb.SendOrderConfirmationRequest) (*pb.Empty, error) {

	// 创建邮件
	emailContent := "<h1>您的购物订单确认</h1>"
	emailContent += "<p>订单号：" + in.Order.OrderId + "</p>"
	emailContent += "<p>配送跟踪号：" + in.Order.ShippingTrackingId + "</p>"
	emailContent += "<p>配送地址：" + in.Order.ShippingAddress.StreetAddress + ", " + in.Order.ShippingAddress.City + ", " + in.Order.ShippingAddress.State + ", " + in.Order.ShippingAddress.Country + ", " + string(in.Order.ShippingAddress.ZipCode) + "</p>"
	emailContent += "<p>商品列表：</p>"
	var totalCost float64
	for _, item := range in.Order.Items {
		// 获取商品数量
		quantity := strconv.Itoa(int(item.Item.Quantity))
		// 计算总价
		totalUnits := item.Cost.Units
		totalNanos := item.Cost.Nanos
		totalPrice := float64(totalUnits) + float64(totalNanos)/1000000000.0
		totalPrice *= float64(item.Item.Quantity)

		totalCost += totalPrice

		emailContent += "<p>" + item.Item.ProductId + " - 数量：" + quantity + " - 总价：" + strconv.FormatFloat(totalPrice, 'f', 2, 64) + " " + item.Cost.CurrencyCode + "</p>"
	}
	// 计算总配送费用
	shippingCost := float64(in.Order.ShippingCost.Units) + float64(in.Order.ShippingCost.Nanos)/1000000000.0

	// 将配送费用添加到邮件内容中
	emailContent += "<p>配送费用：" + strconv.FormatFloat(shippingCost, 'f', 2, 64) + " " + in.Order.ShippingCost.CurrencyCode + "</p>"

	// 计算总付款金额
	totalPayment := totalCost + shippingCost

	emailContent += "<p>总付款金额：" + strconv.FormatFloat(totalPayment, 'f', 2, 64) + " " + in.Order.ShippingCost.CurrencyCode + "</p>"
	e := email.NewEmail()
	e.From = "glengail <1370025136@qq.com>"
	e.To = []string{in.Email}
	e.Bcc = []string{"1370025136@qq.com"} // 密送人，BCC无法看到其他收件人，也无法被其他收件人看到
	e.Cc = []string{"1370025136@qq.com"}  // 抄送人，可看见其他CC收件人
	e.Subject = "您的购物订单确认"
	e.HTML = []byte(emailContent)

	// 发送邮件
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1370025136@qq.com", "XXXXXXXXX", "smtp.qq.com"))
	if err != nil {
		log.Printf("邮件发送失败: %v", err)
		return nil, err
	}

	logger.Printf("邮件已发送到：%s.", in.Email)
	out := new(pb.Empty)
	return out, nil
}

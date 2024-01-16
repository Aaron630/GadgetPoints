package gadgetPoints

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type GadgetPointsIn interface {
	GetRedeemCode(orderId int) (string, error) // 获取兑换码
	UseRedeemCode(redeemCode string) error     // 使用兑换码
	Subscribe(userId int) error                // 订阅
	UnSubscribe(userId int) error              // 取消订阅
	SubscribePush() (string, error)            // 订阅推送
}

type GadgetPointsUseCase struct {
}

func (d *GadgetPointsUseCase) GetRedeemCode(orderId int) (string, error) {
	// 1.验证订单ID
	// 2.验证配件数量 下发对应数量兑换码
	// 3.记录兑换码与订单号（标记已兑换）（事务）

	return randomString(20), nil
}

func (d *GadgetPointsUseCase) UseRedeemCode(redeemCode string) error {
	// 1.验证兑换码
	// 2.记录兑换信息（用户 店铺 时间 地点······）
	// 3.兑换成功 减去兑换商品库存（事务）
	return nil
}

func (d *GadgetPointsUseCase) Subscribe(userId int) error {
	// // 1.记录订阅用户 存入数据库
	return nil
}

func (d *GadgetPointsUseCase) UnSubscribe(userId int) error {
	// 1.取消用户订阅
	return nil
}

func (d *GadgetPointsUseCase) SubscribePush() (string, error) {
	// 根据规则 给订阅用户推送相关消息， 如（邮件、短信）
	return "您订阅的产品又打折了：xxx", nil
}

func randomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[r.Intn(len(charset))])
	}
	return sb.String()
}

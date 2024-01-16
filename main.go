// 根据下文业务描述，使用Golang建模（无需完成 Repo层和 API 层，无需运行, 但是需要包含业务逻辑)。完成后提交代码到Github Public Repo.
// TechSolutions is a national electronics retailer. They experienced rapid growth in the last year and have opened 50 new outlets. Each outlet sells electronic devices and related accessories, as well as store-specific gadgets. Outlets often have individual offers, but national marketing campaigns are often run, which influence the price of an item too.
// TechSolutions recently launched a loyalty program called GadgetPoints, which allows customers to get 1 free accessory for every 10 they purchase. It doesn’t matter which outlet they purchase an item at or which they redeem it at.
// TechSolutions has been thinking of launching an online store. They are also considering a monthly subscription that allows purchasers to get unlimited gadget servicing every month, as well as a discount on other devices. Now that we understand the business domain, we can start to explore how we can build systems to help TechSolutions achieve its goals!
package main

import (
	"GadgetPoints/gadgetPoints"
	"fmt"
)

func main() {
	var gp gadgetPoints.GadgetPointsIn
	gpUseCase := gadgetPoints.GadgetPointsUseCase{}
	gp = &gpUseCase

	code, _ := gp.GetRedeemCode(1)
	fmt.Println("获取兑换码：", code)

	err := gp.UseRedeemCode(code)
	if err == nil {
		fmt.Println("兑换成功")
	}

	err = gp.Subscribe(1)
	if err == nil {
		fmt.Println("订阅成功")
	}

	pushString, err := gp.SubscribePush()
	fmt.Println(pushString)
}

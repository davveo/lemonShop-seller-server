package consts

type ContextKey string

const (
	DefaultDisabled = 0 // 是否被删除0 删除   1  没有删除
	DefaultEnabled  = 1 // 是否被删除0 删除   1  没有删除
	DefaultSellerId = 0 // 默认卖家
)

const (
	Buyer  = Role("Buyer")
	Seller = Role("Seller")
	Clerk  = Role("Clerk")
)

const (
	BuyerPermission  = Permission("Buyer")
	SellerPermission = Permission("Seller")
	AdminPermission  = Permission("Admin")
	ClientPermission = Permission("Client")
)

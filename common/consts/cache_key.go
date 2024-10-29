package consts

const (
	CacheApiKey                  = "c:sys:api:"
	CompanyStatisticsCompanyUser = "companyStatistics:companyUser"
	CompanyStatisticsAccount     = "companyStatistics:account"
	CompanyStatisticsDevice      = "companyStatistics:device"
	CompanyStatisticsSelf        = "companyStatistics:self"
	CompanyStatisticsThird       = "companyStatistics:third"
	CompanyStatisticsApi         = "companyStatistics:api"
	CompanyStatisticsMeal        = "companyStatistics:meal"
	CompanyStatisticsOrder       = "companyStatistics:order"
	CompanyUserAuthentication    = "companyStatistics:authentication" // 统计企业认证信息key
	CompanyStatisticsUserType    = "companyStatistics:userType"       //更新企业用户类型

	// 同步企业住相关信息
	CompanyOnwerStatisticsUserType  = "companyStatisticsOnwer:userType"  // 更新企业主用户类型
	CompanyOnwerStatisticsAccount   = "companyStatisticsOnwer:account"   // 同步企业住环境信息
	CompanyOnwerStatisticsSelf      = "companyStatisticsOnwer:self"      // 同步企业主自有代理信息
	CompanyOnwerStatisticsThird     = "companyStatisticsOnwer:third"     // 同步企业主三方代理信息
	CompanyOnwerStatisticsApi       = "companyStatisticsOnwer:api"       // 同步企业主api代理信息
	CompanyOnwerStatisticsLastLogin = "companyStatisticsOnwer:lastLogin" // 同步用户最后登录时间

	//分片上传存储Key
	PackageSliceKey = "packageSliceKey:"

	// 退出登录黑名单
	LogoutJwtKey = "logoutJwtKey:"

	//测试企业用户缓存
	TestCompanyIdSliceKey = "testCompanyIdSliceKey"

	//测试用户缓存
	TestUserIdSliceKey = "testUserIdSliceKey"

	//后台登录用户数据
	AdminLoginUserKey = "adminLoginUserKey:"

	//同步所有订单相关键值
	SyncAllOrderToCompanyOrder    = "syncAllOrder:to:company_order" // 同步客户创建订单
	SyncAllOrderToTripartiteOrder = "syncAllOrder:tripartite_order" // 同步crm销售创建订单
	SyncAllOrderToOrderRefund     = "syncAllOrder:order_refund"     // 同步退款订单
	SyncAllOrderToCPerformance    = "syncAllOrder:c_performance"    // 同步销售促单
)

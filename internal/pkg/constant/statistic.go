// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/16$ 14:37$

package constant

//密码服务类型
const (
	TypeCipher      = iota + 1 //密码机服务器
	TypeSign                   //签章服务器
	TypeTimeStamp              //时间戳服务器
	TypeSignVer                //签名验签服务器
	TypeCA                     //CA数字证书认证系统
	TypeCloudCipher            //云密码机
)

var CipherType = [7]string{}

//接口类型
const (
	ApiTypeSm4Encrypt      = iota + 1 //加密				1
	ApiTypeSm4Decrypt                 //解密				2
	ApiTypeSm3HMAC                    //HMAC计算			3
	ApiTypeSm3HMACVerify              //HMAC校验			4
	ApiTypeEncUpdate                  //密文更新			5
	ApiTypeTimeStampCreate            //生成时间戳		6
	ApiTypeTimeStampVerify            //验证时间戳		7
	ApiTypeTimeStampParse             //获取时间戳信息	8
	ApiTypeSm2Encrypt                 //非对称加密		9
	ApiTypeSm2Decrypt                 //非对称解密		10
	ApiTypeSm2Sign                    //签名				11
	ApiTypeSm2SignVerify              //验签				12
)

var ApiType = [13]string{}

// TimeLen 交易统计坐标时间点统计个数
const TimeLen = 7

//TimeStatic 当前时间需要减掉的时间为统计时间段
type TimeStatic [TimeLen]int64

//TimeRange 对应时间范围下统计时间段
var TimeRange map[int]TimeStatic

//业务交易-交易统计-时间范围
const (
	TimeOneHour    = iota + 1 //1小时	1
	TimeTwelveHour            //12小时 	2
	TimeOneDay                //1天		3
	TimeOneWeek               //1周		4
	TimeOneMonth              //1个月	5
	TimeThreeMonth            //3个月	6
	TimeOneYear               //1年		7
)

func init() {
	//密码资源类型
	FillCipherType()

	//时间范围
	FillTimeRange()

	//接口对应名称
	FillApiType()
}

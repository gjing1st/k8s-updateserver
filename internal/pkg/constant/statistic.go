// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/16$ 14:37$

package constant

import "time"

//密码服务类型
const (
	TypeCipher      = iota + 1 //密码机服务器
	TypeSign                   //签章服务器
	TypeTimeStamp              //时间戳服务器
	TypeSignVer                //签名验签服务器
	TypeCA                     //CA数字证书认证系统
	TypeCloudCipher            //云密码机
)

var CipherType = [7]string{"", "密码机服务器", "签章服务器", "时间戳服务器", "签名验签服务器", "CA数字证书认证系统", "云密码机"}

// TimeLen 交易统计坐标时间点统计个数
const TimeLen = 7

//TimeStatic 当前时间需要减掉的时间为统计时间段
type TimeStatic [TimeLen]int64

//TimeRange 对应时间范围下统计时间段
var TimeRange map[int]TimeStatic

func init() {
	//密码资源类型
	CipherType[TypeCipher] = "密码机服务器"

	//时间范围
	FillTimeRange()
}

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

// FillTimeRange
// @description: 填充时间范围解析
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/10 18:18
// @success:
func FillTimeRange() {
	TimeRange = make(map[int]TimeStatic)
	//1小时要查询的统计区间
	//1小时：当前时间-60min，当前时间-50min，当前时间-40min，当前时间-30min，当前时间-20min，当前时间-10min，当前时间（21:00），统计的数据为10分钟内的总数，第一个值计算（当前时间-70min~当前时间-60min的，后同）
	oneHour := TimeStatic{
		int64(time.Minute * -70),
		int64(time.Minute * -60),
		int64(time.Minute * -50),
		int64(time.Minute * -40),
		int64(time.Minute * -30),
		int64(time.Minute * -20),
		int64(time.Minute * -10),
	}
	TimeRange[TimeOneHour] = oneHour
	//12小时要查询的统计区间
	//12小时：当前时间-12h，当前时间-10h，当前时间-8h，当前时间-6h，当前时间-4h，当前时间-2h（19:00），当前时间（21:00），统计数据为2小时内的总数
	twelveHour := TimeStatic{
		int64(time.Hour * -14),
		int64(time.Hour * -12),
		int64(time.Hour * -10),
		int64(time.Hour * -8),
		int64(time.Hour * -6),
		int64(time.Hour * -4),
		int64(time.Hour * -2),
	}
	TimeRange[TimeTwelveHour] = twelveHour
	//1天要查询的统计区间
	//1天：当前时间-24h，当前时间-20h，当前时间-16h，当前时间-12h，当前时间-8h，当前时间-4h，当前时间（21:00），统计数据为4小时内的总数
	oneDay := TimeStatic{
		int64(time.Hour * -28),
		int64(time.Hour * -24),
		int64(time.Hour * -20),
		int64(time.Hour * -16),
		int64(time.Hour * -12),
		int64(time.Hour * -8),
		int64(time.Hour * -4),
	}
	TimeRange[TimeOneDay] = oneDay
	//一周：当前时间-6d，当前时间-5d，当前时间-4d，当前时间-3d，当前时间-2d，当前时间-1d，当前时间（10-09），统计数据为1天内的总数
	//单位天，需要当前时间减掉的天数
	oneWeek := TimeStatic{-7, -6, -5, -4, -3, -2, -1}
	TimeRange[TimeOneWeek] = oneWeek
	//一个月：当前时间-30d，当前时间-25d，当前时间-20d，当前时间-15d，当前时间-10d，当前时间-5d，当前时间（月-日），统计数据为5天内的总数
	//单位天，需要当前时间减掉的天数
	oneMonth := TimeStatic{-35, -30, -25, -20, -15, -10, -5}
	TimeRange[TimeOneMonth] = oneMonth
	//三个月：当前时间-3m（07-09），当前时间-15d-2m（07-24），当前时间-2m（08-09），当前时间-15d-1m（08-24），当前时间-1m（09-09），当前时间-15d（09-24），当前时间（10-09），统计数据为15天内的总数
	//单位天，需要当前时间减掉的天数
	threeMonth := TimeStatic{-105, -90, -75, -60, -45, -30, -15}
	TimeRange[TimeThreeMonth] = threeMonth
	//一年：当前时间-12m（10-09），当前时间-10m（12-09），当前时间-8m（02-09），当前时间-6m（04-09），当前时间-4m（06-09），当前时间-2m（08-09），当前时间（10-09），统计数据为60天内的总数
	//单位月，需要减掉的月数
	oneYear := TimeStatic{-14, -12, -10, -8, -6, -4, -2}
	TimeRange[TimeOneYear] = oneYear
}

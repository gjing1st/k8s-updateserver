// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/16$ 14:37$

package constant

//密码服务类型
const (
	TypeCipher    = iota + 1 //密码机服务器
	TypeSign                 //签章服务器
	TypeTimeStamp            //时间戳服务器
	TypeSignVer              //签名验签服务器
	TypeCA                   //CA数字证书认证系统
)

var CipherType = [6]string{"", "密码机服务器", "签章服务器", "时间戳服务器", "签名验签服务器", "CA数字证书认证系统"}

// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/10/18$ 14:45$

package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

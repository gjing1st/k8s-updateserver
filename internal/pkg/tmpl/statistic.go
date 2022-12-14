// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/7$ 10:07$

package tmpl

//RealTime 实时调度
var RealTime = `{
    "query": {
        "bool": {
            "must": [
                {
                    "match_phrase": {
                        "msg": "{{msgField}}"
                    }
                },
                {
                    "match_phrase": {
                        "event_tid": {{eventTidField}}
                    }
                },
                {
                    "match_phrase": {
                        "kubernetes.namespace_name": "{{nameSpaceField}}"
                    }
                }
            ]
        }
    },
    "from": 0,
    "size": 50,
    "sort": {
    }
}`

//Ranking 排名
var Ranking = `{
	"query": {
		"bool": {
			"must": [
				{
					"match_phrase": {
						"msg": "{{msgField}}"
					}
				},
				{
					"match_phrase": {
						"event_tid": {{eventTidField}}
					}
				},
				{
					"range": {
						"event_time": {
							"gt": "{{gtField}}",
							"lt": "{{ltField}}"
						}
					}
				},
                {
                    "match_phrase": {
                        "kubernetes.namespace_name": "{{nameSpaceField}}"
                    }
                }
			]
		}
	},
	"aggs": {
		"event_appname": {
			"terms": {
				"field": "event_appname.keyword"
			}
		}
	},
	"size": 0
}`

//Flow 业务流量
var Flow = `{
    "query": {
        "bool": {
            "must": [
                {
                    "match_phrase": {
                        "msg": "{{msgField}}"
                    }
                },
                {
                    "range": {
                        "event_time": {
                            "gt": "{{gtField}}",
                            "lt": "{{ltField}}"
                        }
                    }
                },
                {
                    "match_phrase": {
                        "event_tid": {{eventTidField}}
                    }
                },
                {
                    "match_phrase": {
                        "kubernetes.namespace_name": "{{nameSpaceField}}"
                    }
                }
            ]
        }
    },
    "aggs": {
        "event_appid": {
            "terms": {
                "field": "event_appname.keyword",
                "order": {
                    "flow": "DESC"
                }
            },
            "aggs": {
                "flow": {
                    "sum": {
                        "field": "event_data"
                    }
                }
            }
        }
    },
    "size": 0
}`

//Statistics 统计
var Statistics = `{
    "query": {
        "bool": {
            "must": [
                {
                    "match_phrase": {
                        "msg": "{{msgField}}"
                    }
                },
                {
                    "match_phrase": {
                        "event_tid": {{eventTidField}}
                    }
                },
                {
                    "match_phrase": {
                        "kubernetes.namespace_name": "{{nameSpaceField}}"
                    }
                }
            ],
            "filter": [
                {
                    "range": {
                        "event_time": {
                            "gt": "{{gtField}}",
							"lt": "{{ltField}}"
                        }
                    }
                }
            ]
        }
    },
    "aggs": {
        "event_serial,event_dtype": {
            "terms": {
                "script": "doc['event_serial.keyword'].value + ',' + doc['event_dtype'].value"
            }
        }
    },
    "size": 0
}`

var Latest = `{
    "query": {
        "bool": {
            "must": [
                {
                    "match_phrase": {
                        "msg": "{{msgField}}"
                    }
                },
                {
                    "match_phrase": {
                        "kubernetes.namespace_name": "{{nameSpaceField}}"
                    }
                },{
                    "range": {
                        "event_time": {
                            "gte": {{gtField}},
                            "lt": {{ltField}}
                        }
                    }
                },{
                    "match_phrase": {
                        "event_tid": {{eventTidField}}
                    }
                }
            ]
        }
    },
    "from": {{fromField}},
    "size": {{sizeField}}
}`

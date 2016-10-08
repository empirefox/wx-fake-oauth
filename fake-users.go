package main

import (
	"encoding/json"

	"github.com/RangelReale/osin"
	mpoauth2 "github.com/chanxuehong/wechat.v2/mp/oauth2"
)

type Users map[string]osin.ResponseData

func getFakeUsers() (r Users) {
	var _users []*mpoauth2.UserInfo
	err := json.Unmarshal([]byte(rawFakeUsers), &_users)
	if err != nil {
		panic(err)
	}

	var users []osin.ResponseData
	err = json.Unmarshal([]byte(rawFakeUsers), &users)
	if err != nil {
		panic(err)
	}

	r = make(Users)
	for _, user := range users {
		r[user["nickname"].(string)] = user
	}

	return
}

const rawFakeUsers = `[{
    "openid": " OPENID1",
    "nickname": "NICKNAME1",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL01"
}, {
    "openid": " OPENID2",
    "nickname": "NICKNAME2",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL02"
}, {
    "openid": " OPENID3",
    "nickname": "NICKNAME3",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL03"
}, {
    "openid": " OPENID4",
    "nickname": "NICKNAME4",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL04"
}, {
    "openid": " OPENID5",
    "nickname": "NICKNAME5",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL05"
}, {
    "openid": " OPENID6",
    "nickname": "NICKNAME6",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL06"
}, {
    "openid": " OPENID7",
    "nickname": "NICKNAME7",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL07"
}, {
    "openid": " OPENID8",
    "nickname": "NICKNAME8",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL08"
}, {
    "openid": " OPENID9",
    "nickname": "NICKNAME9", 
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL09"
}, {
    "openid": " OPENID10",
    "nickname": "NICKNAME10",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL10"
}, {
    "openid": " OPENID11",
    "nickname": "NICKNAME11",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL11"
}, {
    "openid": " OPENID12",
    "nickname": "NICKNAME12",
    "sex": 1,
    "province": "PROVINCE",
    "city": "CITY",
    "country": "COUNTRY",
    "headimgurl": "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
    "privilege": ["PRIVILEGE1", "PRIVILEGE2"],
    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL12"
}]
`

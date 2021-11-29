package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/gogap/logrus"
	"github.com/iiinsomnia/goadmin/consts"
	"github.com/iiinsomnia/goadmin/helpers"
	"github.com/iiinsomnia/goadmin/service"
	"github.com/shenghui0779/yiigo"
	"github.com/thedevsaddam/gojsonq"
)

type event struct {
	Events []string
	Count  int
}

type WeiboUserListItem struct {
	UID         string  `json:"uid"`
	Name        string  `json:"name"`
	NickName    string  `json:"nickname"`
	Watch       string  `json:"watch"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	Fans        float64 `json:"fans"`
	Follow      float64 `json:"follow"`
	WeiboConut  float64 `json:"weibo_conut"`
}
type WeiboUserListReply struct {
	Count int64                `json:"count"`
	List  []*WeiboUserListItem `json:"list"`
}

type weiboUserAdd struct {
	Name string `json:"name" valid:"required"`
	Uid  string `json:"uid" valid:"required"`
}

type weiboUserUpdate struct {
	Name string `json:"name" valid:"required"`
	Uid  string `json:"uid" valid:"required"`
}

type watchUserDelete struct {
	Uid string `json:"uid"`
}

type weiboUserInfo struct {
	Name        string
	Uid         string
	Watch       string
	Fans        float64
	Follow      float64
	WeiboConut  float64
	Location    string
	Description string
	Nickname    string
}

// Events ...
func WeiboEvents(c *gin.Context) {
	count := 0
	events, err := etcdcli.GetAll("/events/")
	if err != nil {
		if err.Error() == "Key Not Found" {
			events = []string{"暂时没有事件产生。"}
		} else {
			log.Error(err)
			events = []string{"获取 events 失败: " + err.Error()}
		}
	} else {
		count = len(events)
	}

	Render(c, "weibo", gin.H{
		"events": events,
		"count":  count,
	})
}

func WeiboUsers(c *gin.Context) {
	Render(c, "weibo_users", gin.H{"menu": "9"})
}

func WeiboUsersQuery(c *gin.Context) {
	s := new(service.UserList)
	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))
		return
	}

	db := yiigo.DB()
	query := "select name,nickname,uid,fans,follows,watch,weibo_count,description,location from weibo"

	querys := make([]string, 0)
	if s.Name != "" {
		querys = append(querys, query+" where name like '%"+s.Name+"%'")
		querys = append(querys, query+" where uid like '%"+s.Name+"%'")
		querys = append(querys, query+" where nickname like '%"+s.Name+"%'")
	} else {
		querys = append(querys, query)
	}

	datas := make([]*WeiboUserListItem, 0)

	for _, q := range querys {
		rows, err := db.Query(q)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var uid, name, nickname, watch, description, location string
			var fans, follows, weibo_count float64
			err := rows.Scan(&name, &nickname, &uid, &fans, &follows, &watch, &weibo_count, &description, &location)
			if err != nil {
				fmt.Println("err: ", err)
			}

			data := new(WeiboUserListItem)
			data.Name = name
			data.NickName = nickname
			data.UID = uid
			data.Fans = fans
			data.Follow = follows
			data.Watch = watch
			data.WeiboConut = weibo_count
			data.Description = description
			data.Location = location
			datas = append(datas, data)
		}
	}
	obj := gin.H{
		"err":  false,
		"code": 1000,
		"msg":  "success",
	}

	resp := &WeiboUserListReply{
		Count: int64(len(datas)),
		List:  make([]*WeiboUserListItem, 0),
	}

	resp.List = datas

	obj["data"] = resp

	c.Set("response", obj)
	c.JSON(http.StatusOK, obj)
}

func WeiboUsersAdd(c *gin.Context) {
	s := new(weiboUserAdd)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))
		return
	}

	u, err := queryWeiboUserInfoByUid(s)
	if err != nil {
		log.Error(err)
	}

	db := yiigo.DB()
	sql := fmt.Sprintf(`insert into weibo(name,nickname,uid,fans,follows,weibo_count,description,location) values("%v","%v","%v","%v","%v","%v","%v","%v")`,
		u.Name, u.Nickname, u.Uid, u.Fans, u.Follow, u.WeiboConut, u.Description, u.Location)
	_, err = db.Exec(sql)
	if err != nil {
		Err(c, helpers.Error(helpers.ErrParams), "添加失败:"+err.Error())
		log.Error(err)
		return
	}

	OK(c)
}

func WeiboUsersUpdate(c *gin.Context) {
	s := new(weiboUserUpdate)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))
		log.Error("err: ", err)
		return
	}
	db := yiigo.DB()
	sql := fmt.Sprintf(`update weibo set name="%v" where uid="%v"`, s.Name, s.Uid)
	_, err := db.Exec(sql)
	if err != nil {
		Err(c, err, "更新失败")
		return
	}

	OK(c)
}

func WeiboUsersDelete(c *gin.Context) {
	identity, err := Identity(c)

	if err != nil || identity.Role != consts.SuperManager {
		Err(c, helpers.Error(helpers.ErrForbid, err))
		return
	}

	s := new(watchUserDelete)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))
		return
	}

	db := yiigo.DB()
	sql := fmt.Sprintf(`delete from weibo where uid="%v"`, s.Uid)
	_, err = db.Exec(sql)
	if err != nil {
		Err(c, helpers.Error(helpers.ErrParams), "删除失败")
		return
	}

	OK(c)
}

func queryWeiboUserInfoByUid(s *weiboUserAdd) (*weiboUserInfo, error) {
	url := "https://weibo.com/ajax/profile/info?uid=" + s.Uid
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("cookie", "SUB=_2AkMW_Msrf8NxqwJRmPoXzG3lb491yArEieKgoDrwJRMxHRl-yT8XqkUrtRB6PXzlxDF03pdbGMvtYocdCLMK6GGaGO_O; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WWfVYMDd.cejLMndXZjw9Bo; XSRF-TOKEN=S0R2thXnTMpJYVrxnn2AIR1F; _s_tentry=weibo.com; Apache=7886961433352.875.1637893164987; SINAGLOBAL=7886961433352.875.1637893164987; ULV=1637893164993:1:1:1:7886961433352.875.1637893164987:; WBPSESS=Jx_XaCleItbWmjWmltuZpPHAWWfc0ARg_GHEyi04aIEYP3kVyoef4sQe3SvKUkGIshIrMGCqiU76Uc85lAwWWKsDEd7DGh0eVfqCoHojEDnr_rGNmX0wxqSdHcSNDXnmueDwkDAIrueAu9cwffr01ohexVllJNrl21-bf30v344=")

	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if gojsonq.New().FromString(string(body)).Find("ok").(float64) != 1 {
		return nil, fmt.Errorf("请求出错")
	}

	user := &weiboUserInfo{
		Name:        s.Name,
		Uid:         s.Uid,
		Fans:        gojsonq.New().FromString(string(body)).Find("data.user.followers_count").(float64),
		Follow:      gojsonq.New().FromString(string(body)).Find("data.user.friends_count").(float64),
		WeiboConut:  gojsonq.New().FromString(string(body)).Find("data.user.statuses_count").(float64),
		Location:    gojsonq.New().FromString(string(body)).Find("data.user.location").(string),
		Description: gojsonq.New().FromString(string(body)).Find("data.user.description").(string),
		Nickname:    gojsonq.New().FromString(string(body)).Find("data.user.screen_name").(string),
	}

	if user.Location == "" {
		user.Location = "未知"
	}
	if user.Description == "" {
		user.Description = "未知"
	}
	if user.Nickname == "" {
		user.Nickname = "未知"
	}

	return user, nil
}

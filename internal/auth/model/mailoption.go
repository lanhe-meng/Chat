package model

type EmailOptions struct {
	MailHost string //smtp.163.com
	MailPort int    //465
	MailUser string // 发件人
	MailPass string // 发件人密码
	MailTo   string // 收件人 多个用,分割
	Subject  string // 邮件主题
	Body     string // 邮件内容
}

// redis中verify_code:data:*@*.com对应的值
type EmailRedisValue struct {
	Code    int
	Attemps int
}

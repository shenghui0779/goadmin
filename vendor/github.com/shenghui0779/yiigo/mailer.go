package yiigo

import (
	"sync"

	"github.com/pelletier/go-toml"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

type emailConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

// EMail email
type EMail struct {
	Title   string
	Subject string
	From    string
	To      []string
	Cc      []string
	Content string
	Attach  []string
}

// emailOptions email options
type emailOptions struct {
	charset     string
	encoding    gomail.Encoding
	contentType string
}

// EMailOption configures how we set up the email
type EMailOption interface {
	apply(*emailOptions)
}

// funcEMailOption implements email option
type funcEMailOption struct {
	f func(*emailOptions)
}

func (fo *funcEMailOption) apply(o *emailOptions) {
	fo.f(o)
}

func newFuncEMailOption(f func(*emailOptions)) *funcEMailOption {
	return &funcEMailOption{f: f}
}

// WithEMailCharset specifies the `Charset` to email.
func WithEMailCharset(s string) EMailOption {
	return newFuncEMailOption(func(o *emailOptions) {
		o.charset = s
	})
}

// WithEMailEncoding specifies the `Encoding` to email.
func WithEMailEncoding(e gomail.Encoding) EMailOption {
	return newFuncEMailOption(func(o *emailOptions) {
		o.encoding = e
	})
}

// WithEMailContentType specifies the `ContentType` to email.
func WithEMailContentType(s string) EMailOption {
	return newFuncEMailOption(func(o *emailOptions) {
		o.contentType = s
	})
}

// EMailDialer email dialer
type EMailDialer struct {
	dialer *gomail.Dialer
}

// Send send an email.
func (m *EMailDialer) Send(e *EMail, options ...EMailOption) error {
	o := &emailOptions{contentType: "text/html"}

	if len(options) > 0 {
		for _, option := range options {
			option.apply(o)
		}
	}

	settings := make([]gomail.MessageSetting, 0, 2)

	if o.charset != "" {
		settings = append(settings, gomail.SetCharset(o.charset))
	}

	if o.encoding != "" {
		settings = append(settings, gomail.SetEncoding(o.encoding))
	}

	msg := gomail.NewMessage(settings...)

	msg.SetHeader("Subject", e.Subject)
	msg.SetAddressHeader("From", e.From, e.Title)
	msg.SetHeader("To", e.To...)

	if len(e.Cc) > 0 {
		msg.SetHeader("Cc", e.Cc...)
	}

	if len(e.Attach) > 0 {
		for _, v := range e.Attach {
			msg.Attach(v)
		}
	}

	msg.SetBody(o.contentType, e.Content)

	// Send the email
	err := m.dialer.DialAndSend(msg)

	return err
}

var (
	defaultMailer *EMailDialer
	mailerMap     sync.Map
)

func initMailer() {
	tree, ok := env.get("email").(*toml.Tree)

	if !ok {
		return
	}

	keys := tree.Keys()

	if len(keys) == 0 {
		return
	}

	for _, v := range keys {
		node, ok := tree.Get(v).(*toml.Tree)

		if !ok {
			continue
		}

		cfg := new(emailConfig)

		if err := node.Unmarshal(cfg); err != nil {
			logger.Error("yiigo: email dialer init error", zap.String("name", v), zap.Error(err))
		}

		dialer := &EMailDialer{dialer: gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)}

		if v == AsDefault {
			defaultMailer = dialer
		}

		mailerMap.Store(v, dialer)
	}
}

// Mailer returns an email dialer.
func Mailer(name ...string) *EMailDialer {
	if len(name) == 0 {
		if defaultMailer == nil {
			logger.Panic("yiigo: invalid email dialer", zap.String("name", AsDefault))
		}

		return defaultMailer
	}

	v, ok := mailerMap.Load(name[0])

	if !ok {
		logger.Panic("yiigo: invalid email dialer", zap.String("name", name[0]))
	}

	return v.(*EMailDialer)
}

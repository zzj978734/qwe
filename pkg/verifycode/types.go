//go:generate  mcube enum -m

package verifycode

const (
	// NotifyTypeMail (mail) 邮件通知
	NotifyTypeMail NotifyType = iota
	// NotifyTypeSMS (sms) 短信通知
	NotifyTypeSMS
)

// NotifyType 通知方式
type NotifyType uint

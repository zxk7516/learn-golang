package common

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Addr                string `yaml:"addr"`
	PgHost              string `yaml:"pg_host"`
	PgPort              int    `yaml:"pg_port"`
	PgUser              string `yaml:"pg_user"`
	PgPass              string `yaml:"pg_pass"`
	PgDb                string `yaml:"pg_db"`
	UserJwtSecret       string `yaml:"user_jwt_secret"`
	AdminJwtSecret      string `yaml:"admin_jwt_secret"`
	MaxIdleConn         int    `yaml:"max_idle_conn"`
	TokenExpiredDays    int    `yaml:"token_expired_days"`
	TokenRefreshDays    int    `yaml:"token_refresh_days"`
	CookieUser          string `yaml:"cookie_user"`
	CookieToken         string `yaml:"cookie_token"`
	BlogAnnouncementURL string `yaml:"blog_announcement_url"`
	BlogLinkCount       int    `yaml:"blog_link_count"`
	BlogCommentCount    int    `yaml:"blog_comment_count"`
	DebugMode           bool   `yaml:"debug_mode"`
	MediaFolder         string `yaml:"media_folder"`
	SendEmailReply      bool   `yaml:"send_email_reply"`
	EmailSMTPHost       string `yaml:"email_smtp_host"`
	EmailSMTPPort       int    `yaml:"email_smtp_port"`
	EmailSMTPUsername   string `yaml:"email_smtp_username"`
	EmailSMTPPassword   string `yaml:"email_smtp_password"`
	EmailSMTPEmail      string `yaml:"email_admin_email"`
	SiteName            string `yaml:"site_name"`
	SiteURL             string `yaml:"site_url"`
	SiteDescription     string `yaml:"site_description"`
	SiteAuthor          string `yaml:"site_author"`
}

const (
	SessionName    = "gofiberdemo"
	SessionUserKey = "g.id"
)

var config *Config

func LoadConfig(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(result, &config)
}

func GetConfig() *Config {
	return config
}

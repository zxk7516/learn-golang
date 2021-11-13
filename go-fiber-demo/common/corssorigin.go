package common

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type CorssOriginSettings struct {
	Next             func(c *fiber.Ctx) bool // nil
	AllowOrigins     string                  // "*"
	AllowMethods     string                  // "GET,POST,HEAD,PUT,DELETE,PATCH"
	AllowHeaders     string                  // ""
	AllowCredentials bool                    // false
	ExposeHeaders    string                  // ""
	MaxAge           int                     // 0
}

var SettingsDefault = CorssOriginSettings{
	Next:         nil,
	AllowOrigins: "*",
	AllowMethods: strings.Join([]string{
		fiber.MethodGet,
		fiber.MethodPost,
		fiber.MethodHead,
		fiber.MethodPut,
		fiber.MethodDelete,
		fiber.MethodPatch,
	}, ","),
	AllowHeaders:     "",
	AllowCredentials: false,
	ExposeHeaders:    "",
	MaxAge:           0,
}

func CrossOrigin(configs ...CorssOriginSettings) fiber.Handler {
	var cfg CorssOriginSettings
	if len(configs) > 0 {
		cfg = configs[0]
		if cfg.AllowMethods == "" {
			cfg.AllowMethods = SettingsDefault.AllowMethods
		}
		if cfg.AllowOrigins == "" {
			cfg.AllowOrigins = SettingsDefault.AllowOrigins
		}
	} else {
		cfg = SettingsDefault
	}
	allowOrigins := strings.Split(strings.Replace(cfg.AllowOrigins, " ", "", -1), ",")
	allowMethods := strings.Replace(cfg.AllowMethods, " ", "", -1)
	allowHeaders := strings.Replace(cfg.AllowHeaders, " ", "", -1)
	exposeHeaders := strings.Replace(cfg.ExposeHeaders, " ", "", -1)
	maxAge := strconv.Itoa(cfg.MaxAge)

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}
		origin := c.Get(fiber.HeaderOrigin)
		allowOrigin := ""
		for _, o := range allowOrigins {
			if o == "*" && cfg.AllowCredentials {
				allowOrigin = origin
			}
			if o == "*" || o == origin {
				allowOrigin = o
				break
			}
			if matchSubdomain(origin, o) {
				allowOrigin = origin
				break
			}
		}
		if c.Method() != http.MethodOptions {
			c.Vary(fiber.HeaderOrigin)
			c.Set(fiber.HeaderAccessControlAllowOrigin, allowOrigin)

			if cfg.AllowCredentials {
				c.Set(fiber.HeaderAccessControlAllowOrigin, allowOrigin)
			}
			if exposeHeaders != "" {
				c.Set(fiber.HeaderAccessControlExposeHeaders, exposeHeaders)
			}
			return c.Next()
		}

		// preflight request
		c.Vary(fiber.HeaderOrigin)
		c.Vary(fiber.HeaderAccessControlRequestMethod)
		c.Vary(fiber.HeaderAccessControlRequestHeaders)
		c.Set(fiber.HeaderAccessControlAllowOrigin, allowOrigin)
		c.Set(fiber.HeaderAccessControlAllowMethods, allowMethods)
		if cfg.AllowCredentials {
			c.Set(fiber.HeaderAccessControlAllowCredentials, "true")
		}
		if allowHeaders != "" {
			c.Set(fiber.HeaderAccessControlAllowHeaders, allowHeaders)
		} else {
			h := c.Get(fiber.HeaderAccessControlRequestHeaders)
			if h != "" {
				c.Set(fiber.HeaderAccessControlAllowHeaders, h)
			}
		}
		if cfg.MaxAge > 0 {
			c.Set(fiber.HeaderAccessControlMaxAge, maxAge)
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func matchScheme(domain, pattern string) bool {
	didx := strings.Index(domain, ":")
	pidx := strings.Index(pattern, ":")
	return didx != -1 && pidx != -1 && domain[:didx] == pattern[:pidx]
}

func matchSubdomain(domain, pattern string) bool {
	if !matchScheme(domain, pattern) {
		return false
	}
	didx := strings.Index(domain, "://")
	pidx := strings.Index(pattern, "://")
	if didx == -1 || pidx == -1 {
		return false
	}
	domAuth := domain[didx+3:]
	// to avoid long loop by invalid long domain
	if len(domAuth) > 253 {
		return false
	}
	patAuth := pattern[pidx+3:]

	domComp := strings.Split(domAuth, ".")
	patComp := strings.Split(patAuth, ".")
	for i := len(domComp)/2 - 1; i >= 0; i-- {
		opp := len(domComp) - 1 - i
		domComp[i], domComp[opp] = domComp[opp], domComp[i]
	}
	for i := len(patComp)/2 - 1; i >= 0; i-- {
		opp := len(patComp) - 1 - i
		patComp[i], patComp[opp] = patComp[opp], patComp[i]
	}

	for i, v := range domComp {
		if len(patComp) <= i {
			return false
		}
		p := patComp[i]
		if p == "*" {
			return true
		}
		if p != v {
			return false
		}
	}
	return false
}

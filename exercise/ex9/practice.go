package ex9

func DomainForLocale(domain, locale string) string {
	if locale != "" {
		return locale + "." + domain
	} else {
		return "en." + domain
	}
}

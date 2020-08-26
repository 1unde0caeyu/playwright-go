// Package playwright is a library to automate Chromium, Firefox and WebKit with
// a single API. Playwright is built to enable cross-browser web automation that
// is ever-green, capable, reliable and fast.
package playwright

type Playwright struct {
	ChannelOwner
	Chromium *BrowserType
	Firefox  *BrowserType
	WebKit   *BrowserType
}

func (p *Playwright) Stop() error {
	return p.connection.Stop()
}

func newPlaywright(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Playwright {
	// TODO: add devices and selectors
	pw := &Playwright{
		Chromium: (initializer["chromium"]).(*Channel).object.(*BrowserType),
		Firefox:  (initializer["firefox"]).(*Channel).object.(*BrowserType),
		WebKit:   (initializer["webkit"]).(*Channel).object.(*BrowserType),
	}
	pw.createChannelOwner(pw, parent, objectType, guid, initializer)
	return pw
}

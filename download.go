package crx3lite

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var chromeExtURL = "https://clients2.google.com/service/update2/crx?response=redirect&prodversion=80.0&acceptformat=crx3&x=id%3D{id}%26installsource%3Dondemand%26uc"

// SetWebStoreURL sets the web store url to download extensions.
// web store URL is expected to contain {id} which will be replaced with the ExtensionID
func SetWebStoreURL(u string) {
	if len(u) == 0 {
		return
	}
	if !strings.HasPrefix(u, "http") {
		u = "https://" + u
	}
	chromeExtURL = u
}

// DownloadFromWebStore downloads a chrome extension from the web store.
// ExtensionID can be an identifier or an url.
func DownloadFromWebStore(extensionID string) (rawPkg []byte, err error) {
	if len(extensionID) == 0 {
		return nil, ErrExtensionNotSpecified
	}

	extensionURL := makeChromeURL(extensionID)

	resp, err := http.Get(extensionURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("crx3lite: bad status: %s", resp.Status)
	}

	rawPkg, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("crx3lite: %s: %e", extensionID, err)
	}

	return rawPkg, nil
}

func makeChromeURL(chromeID string) string {
	return strings.Replace(chromeExtURL, "{id}", chromeID, 1)
}

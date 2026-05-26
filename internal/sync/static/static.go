package static

import (
	"github.com/gin-gonic/gin"

	_ "embed"
)

var (
	//go:embed index.html
	index string

	//go:embed favicon.ico
	favicon []byte

	//go:embed logo.svg
	logo []byte

	//go:embed bootstrap.min-5.3.3.js
	bootstrapJS []byte

	// https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css
	//go:embed bootstrap.min-5.3.3.css
	bootstrapCSS []byte

	// https://bootswatch.com/5/darkly/bootstrap.min.css
	//go:embed bootstrap.min-darkly-5.3.css
	bootstrapDarkCSS []byte

	// https://code.jquery.com/jquery-3.7.1.min.js
	//go:embed jquery-3.7.1.min.js
	jquery []byte

	// https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js
	//go:embed popper.min-2.9.2.js
	popper []byte

	// https://cdn.jsdelivr.net/npm/chart.js@4.4.7/dist/chart.umd.min.js
	//go:embed chart.umd.min-4.4.7.js
	chart []byte
)

func handleFavicon(c *gin.Context) { _ = "STUB: not implemented"; return }

func handleLogo(c *gin.Context) { _ = "STUB: not implemented"; return }

func Index() string { _ = "STUB: not implemented"; return "" }

func HandleResources(group gin.IRouter, dark bool) { _ = "STUB: not implemented"; return }

func handleJS(bytes []byte) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

func handleCSS(bytes []byte) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

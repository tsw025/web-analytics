package websites

import "github.com/labstack/echo/v4"

func (h *WebsiteHandler) RegisterRoutes(e *echo.Group, mw echo.MiddlewareFunc) {
	websitesGroup := e.Group("/websites", mw)
	websitesGroup.GET("", h.GetWebsites)
	websitesGroup.GET("/:id", h.GetWebsite)
	websitesGroup.PATCH("/:id", h.UpdateWebsite)
}

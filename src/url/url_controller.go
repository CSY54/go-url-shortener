package url

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UrlController struct {
	UrlService UrlService
}

func ProvideUrlController(u UrlService) UrlController {
	return UrlController{UrlService: u}
}

func (u *UrlController) UploadUrl(c *gin.Context) {
	var uploadUrlDTO UploadUrlDTO
	err := c.BindJSON(&uploadUrlDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key missing or value invalid"})
		return
	}

	if uploadUrlDTO.ExpireAt.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Time already expired"})
		return
	}

	res := u.UrlService.Create(ToUrl(uploadUrlDTO))
	c.JSON(http.StatusOK, ToResponse(res))
}

func (u *UrlController) RedirectUrl(c *gin.Context) {
	var redirectUrlDTO RedirectUrlDTO
	err := c.BindUri(&redirectUrlDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key missing or value invalid"})
		return
	}

	id, err := B64ToUint32(redirectUrlDTO.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	res := u.UrlService.FindByID(int(id))
	if res == (Url{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	if res.ExpireAt.Before(time.Now()) {
		// c.status(http.StatusGone)
		c.JSON(http.StatusNotFound, gin.H{"error": "Link expired"})
		return
	}

	c.Redirect(http.StatusFound, res.Url)
}

package router

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leonardogcsoares/phonebook-api/router/repo"
)

var (
	errUnmarshEntry = "Unable to unmarshal entry"
	errCreateEntry  = "error creating entry in repo"
	errGetEntry     = "error retrieving entry from repo"
	errUpdateEntry  = "error updating entry to repo"
	errDeleteEntry  = "error deleting entry  repo"
	errInvalidEntry = "invalid entry field"
	errInvalidID    = "invalid id param"
)

const letterBytes = "0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	idSize        = 128
)

var src = rand.NewSource(time.Now().UnixNano())

func randID() string {
	n := idSize
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// CreateEntry TODO
func (r Router) CreateEntry(c *gin.Context) {

	var entry repo.Entry
	err := c.BindJSON(&entry)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResp{
			Msg:   errUnmarshEntry,
			Error: err.Error(),
		})
		return
	}

	// do validations here
	if errValid := r.validator.IsValidEntry(entry); errValid != nil {
		c.JSON(http.StatusBadRequest, errorResp{
			Msg:   errInvalidEntry,
			Error: errValid.Error(),
		})
		return
	}

	entry.ID = randID()

	createdEntry, err := r.Repo.CreateEntry(entry)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResp{
			Msg:   errCreateEntry,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, createdEntry)
}

// GetEntry TODO
func (r Router) GetEntry(c *gin.Context) {
	id := c.Param("id")

	entry, err := r.Repo.GetEntry(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResp{
			Msg:   errGetEntry,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entry)
}

// UpdateEntry TODO
func (r Router) UpdateEntry(c *gin.Context) {
	id := c.Param("id")

	var entry repo.Entry
	err := c.BindJSON(&entry)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResp{
			Msg:   errUnmarshEntry,
			Error: err.Error(),
		})
		return
	}

	updatedEntry, err := r.Repo.UpdateEntry(id, entry)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResp{
			Msg:   errUpdateEntry,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedEntry)
}

// DeleteEntry TODO
func (r Router) DeleteEntry(c *gin.Context) {
	id := c.Param("id")

	err := r.Repo.DeleteEntry(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResp{
			Msg:   errDeleteEntry,
			Error: err.Error(),
		})
		return
	}

}

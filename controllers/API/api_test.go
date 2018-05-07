package API

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	// "strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	// ---  init ---
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// --- exec ---
	Index(c)

	// --- check ---
	if rec.Body.String() != "{\"status\":200,\"message\":\"ヽ（　＾ω＾）ﾉｻｸｾｽ！\"}" {
		t.Errorf("expected response, got %s", rec.Body.String())
	}
}

func TestConnection(t *testing.T) {
	// ---  init ---
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/connection/:lang")
	c.SetParamNames("lang")
	c.SetParamValues("rb")

	if assert.NoError(t, Connection(c)) {
		assert.Equal(t, http.StatusOK, rec.Code, "Two http status codes are different")
		hoge := strconv.Itoa(123)
		assert.Exactly(t, hoge, "123")
	}
}

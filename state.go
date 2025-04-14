package components

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type State interface {
	// Echo context functions
	Request() *http.Request
	ResponseWriter() http.ResponseWriter
	SetRequest(r *http.Request)
	IsTLS() bool
	IsWebSocket() bool
	Scheme() string
	RealIP() string
	Path() string
	SetPath(p string)
	Param(name string) string
	ParamNames() []string
	SetParamNames(names ...string)
	ParamValues() []string
	SetParamValues(values ...string)
	QueryParam(name string) string
	QueryParams() url.Values
	QueryString() string
	FormValue(name string) string
	FormParams() (url.Values, error)
	FormFile(name string) (*multipart.FileHeader, error)
	MultipartForm() (*multipart.Form, error)
	Cookie(name string) (*http.Cookie, error)
	SetCookie(cookie *http.Cookie)
	Cookies() []*http.Cookie
	Get(key string) interface{}
	Set(key string, val interface{})
	Bind(i interface{}) error
	Validate(i interface{}) error
	Render(code int, name string, data interface{}) error
	HTML(code int, html string) error
	HTMLBlob(code int, b []byte) error
	String(code int, s string) error
	JSON(code int, i interface{}) error
	JSONPretty(code int, i interface{}, indent string) error
	JSONBlob(code int, b []byte) error
	JSONP(code int, callback string, i interface{}) error
	JSONPBlob(code int, callback string, b []byte) error
	XML(code int, i interface{}) error
	XMLPretty(code int, i interface{}, indent string) error
	XMLBlob(code int, b []byte) error
	Blob(code int, contentType string, b []byte) error
	Stream(code int, contentType string, r io.Reader) error
	File(file string) error
	Attachment(file string, name string) error
	Inline(file string, name string) error
	NoContent(code int) error
	Redirect(code int, url string) error
	Error(err error)
	Reset(r *http.Request, w http.ResponseWriter)

	// Raptor context functions
	JSONResponse(data interface{}, status ...int) error

	Controller() string
	Action() string
}

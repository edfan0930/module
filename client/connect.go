package client

import (
	"context"
	"net/http"
	"io"
	"time"
	"github.com/pkg/errors"
)

var timeout = 30 * time.Second
var port = ":1324"
var host = "http://localhost/" + port

type config struct {
	Method string
	Path   string
	Body  io.Reader
}

func NewConnect(method,path string,body io.Reader)*config  {
	return &config{
		Method:method,
		Path:path,
		Body:body
	}
}

func (c *config)Do() (resp []byte,err error) {
	client:= &http.Client{
		Timeout:time.Duration(timeout)
	}
	req,err :=http.NewRequest(c.Method,host+c.Path,c.Body)
	if err!=nil {
		return nil,errors.Wrap(err,"")
	}
	ctx,cancel:=context.WithTimeout(context.Background,20*time.Second)
	defer cancel()

	req=req.WithContext(ctx)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)	
	response,err:=client.Do(req)
	defer response.Body.Close()
	if err!=nil {
		return nil ,errors.Wrap(err,"")
	}

	if response.StatusCode!=200 {
		return nil,errors.New("http status error")
	}
	resp, err = ioutil.ReadAll(response.Body)
	return 
}
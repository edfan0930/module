package client

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type ex struct {
	Name     string `json:"name"`
	NickName string `json:"nickname"`
}

func (e *ex) example() error {

	//	body := strings.NewReader("ownerid=" + ownerID + "&category=" + category + "&allowed=" + allowed)
	//or form := make(url.Values) form.Set(`mtcode`, t.MtCode)	form.Set(`amount`, t.Amount) form.Set(`eventTime`, t.EventTime)
	// body := strings.NewReader(form)
	con := NewConnect(
		"GET",
		"person/1",
		nil,
	)
	response, err := con.Do()
	if err != nil {
		return errors.Wrap(err, "")
	}
	if err = jsoniter.Unmarshal(response, e); err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

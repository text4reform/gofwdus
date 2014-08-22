package fwdus

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	StagingHost    = "https://staging.fwd.us/api/v1/"
	ProductionHost = "https://app.fwd.us/api/v1/"
)

type FWDus struct {
	Key  string
	Host string
}

func NewFWDusClient(key string) *FWDus {
	return &FWDus{Key: key, Host: StagingHost}
}

func (c *FWDus) CallLegislator(args CallLegislatorArgs) (*CallRequest, error) {
	body := struct {
		Call CallLegislatorArgs `json:"call"`
	}{Call: args}
	buf, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}
	url := c.Host + "legislators/call.json?key=" + c.Key
	res, err := http.Post(url, "application/json", strings.NewReader(string(buf)))
	if err != nil {
		return nil, err
	}
	var bodyRes Response
	if err := unmarshal(res, &bodyRes); err != nil {
		return nil, err
	}
	return &bodyRes.CallRequest, nil
}

func (c *FWDus) SearchLegislators(args SearchLegislatorArgs) (*[]Legislator, error) {
	v := url.Values{}
	v.Set("key", c.Key)
	if len(args.Zip) > 0 {
		v.Set("zip", args.Zip)
	}
	if args.District > 0 {
		v.Set("district", strconv.Itoa(args.District))
	}
	if len(args.State) > 0 {
		v.Set("state", args.State)
	}
	if len(args.Party) > 0 {
		v.Set("party", args.Party)
	}
	url := c.Host + "legislators/search.json?" + v.Encode()
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var bodyRes Response
	if err := unmarshal(res, &bodyRes); err != nil {
		return nil, err
	}
	return &bodyRes.Legislators, nil
}

func (c *FWDus) CreateLetter(args CreateLetterArgs) (*Letter, error) {
	body := struct {
		Letter CreateLetterArgs `json:"letter"`
	}{Letter: args}
	buf, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}
	url := c.Host + "letters.json?key=" + c.Key
	res, err := http.Post(url, "application/json", strings.NewReader(string(buf)))
	if err != nil {
		return nil, err
	}
	var bodyRes Response
	if err := unmarshal(res, &bodyRes); err != nil {
		return nil, err
	}
	return &bodyRes.Letter, nil
}

func unmarshal(res *http.Response, bodyRes *Response) error {
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &bodyRes); err != nil {
		return err
	}
	if len(bodyRes.Error) > 0 {
		return errors.New(bodyRes.Error)
	}
	if len(bodyRes.Errors) > 0 {
		msg := ""
		for key, reasons := range bodyRes.Errors {
			msg += key + ": " + strings.Join(reasons, ", ") + "."
		}
		return errors.New(msg)
	}
	return nil
}

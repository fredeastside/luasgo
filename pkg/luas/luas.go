package luas

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const (
	luasforecastsURL = "http://luasforecasts.rpa.ie/xml/get.ashx?encrypt=false"
)

// Stops struct
type Stops struct {
	XMLName xml.Name `xml:"stops" json:"-"`
	Lines   []Line   `xml:"line" json:"lines"`
}

// Line struct
type Line struct {
	XMLName xml.Name `xml:"line" json:"-"`
	Name    string   `xml:"name,attr" json:"name"`
	Stops   []Stop   `xml:"stop" json:"stops"`
}

// Stop struct
type Stop struct {
	XMLName       xml.Name `xml:"stop" json:"-"`
	Name          string   `xml:"abrev,attr" json:"abrev"`
	IsParkRide    bool     `xml:"isParkRide,attr" json:"isParkRide"`
	IsCycleRide   bool     `xml:"isCycleRide,attr" json:"isCycleRide"`
	Pronunciation string   `xml:"pronunciation,attr" json:"pronunciation"`
	Lat           float32  `xml:"lat,attr" json:"lat"`
	Long          float32  `xml:"long,attr" json:"long"`
}

// StopInfo struct
type StopInfo struct {
	XMLName   xml.Name    `xml:"stopInfo" json:"-"`
	Created   string      `xml:"created,attr" json:"created"`
	Stop      string      `xml:"stop,attr" json:"stop"`
	StopAbv   string      `xml:"stopAbv,attr" json:"stopAbv"`
	Message   string      `xml:"message" json:"message"`
	Direction []Direction `xml:"direction" json:"directions"`
}

// Direction struct
type Direction struct {
	XMLName xml.Name `xml:"direction" json:"-"`
	Name    string   `xml:"name,attr" json:"name"`
	Trams   []Tram   `xml:"tram" json:"trams"`
}

// Tram struct
type Tram struct {
	XMLName     xml.Name `xml:"tram" json:"-"`
	DueMins     string   `xml:"dueMins,attr" json:"dueMins"`
	Destination string   `xml:"destination,attr" json:"destination"`
}

// Farecalc struct
type Farecalc struct {
	XMLName xml.Name `xml:"farecalc" json:"-"`
	Created string   `xml:"created,attr" json:"created"`
	Params  Params   `xml:"params" json:"params"`
	Result  Result   `xml:"result" json:"result"`
}

// Params struct
type Params struct {
	From     string `xml:"from,attr" json:"from"`
	To       string `xml:"to,attr" json:"to"`
	Adults   string `xml:"adults,attr" json:"adults"`
	Children string `xml:"children,attr" json:"children"`
}

// Result struct
type Result struct {
	Peek           string `xml:"peak,attr" json:"peak"`
	Offpeak        string `xml:"offpeak,attr" json:"offpeak"`
	ZonesTravelled string `xml:"zonesTravelled,attr" json:"zonesTravelled"`
}

// GetStops - returns stops
func GetStops() (Stops, error) {
	url := fmt.Sprintf("%s&action=stops", luasforecastsURL)

	var stops Stops
	err := decodeXML(url, &stops)
	if err != nil {
		return stops, err
	}

	return stops, nil
}

// GetStop - returns shedule by stop
func GetStop(stop string) (StopInfo, error) {
	url := fmt.Sprintf("%s&action=forecast&stop=%s", luasforecastsURL, strings.ToUpper(stop))

	var stopInfo StopInfo
	err := decodeXML(url, &stopInfo)
	if err != nil {
		return stopInfo, err
	}

	return stopInfo, nil
}

// GetFares - returns fares
func GetFares(from, to string, isChildren bool) (Farecalc, error) {
	urlChildsPart := "adults=1&children=0"
	if isChildren {
		urlChildsPart = "adults=0&children=1"
	}
	url := fmt.Sprintf("%s&action=farecalc&from=%s&to=%s&%s", luasforecastsURL, strings.ToUpper(from), strings.ToUpper(to), urlChildsPart)

	var farecalc Farecalc
	err := decodeXML(url, &farecalc)
	if err != nil {
		return farecalc, err
	}

	return farecalc, nil
}

func decodeXML(url string, st interface{}) error {
	body, err := getResponseBody(url)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, st)
	if err != nil {
		return errors.Wrap(err, "can not unmarshal xml data.")
	}

	return nil
}

func getResponseBody(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "can not get remote xml data.")
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "body read failed.")
	}

	return body, nil
}

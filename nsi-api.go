package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// NsiProperties is a reflection of the JSON feature property attributes from the NSI-API
type NsiProperties struct {
	Name             int     `json:"fd_id"`
	X                float64 `json:"x"`
	Y                float64 `json:"y"`
	Occtype          string  `json:"occtype"`
	FoundHt          float64 `json:"found_ht"`
	FoundType        string  `json:"found_type"`
	DamCat           string  `json:"st_damcat"`
	StructVal        float64 `json:"val_struct"`
	ContVal          float64 `json:"val_cont"`
	CB               string  `json:"cbfips"`
	Pop2amu65        int32   `json:"pop2amu65"`
	Pop2amo65        int32   `json:"pop2amo65"`
	Pop2pmu65        int32   `json:"pop2pmu65"`
	Pop2pmo65        int32   `json:"pop2pmo65"`
	NumStories       int32   `json:"num_story"`
	FirmZone         string  `json:"firmzone"`
	GroundElevation  float64 `json:"ground_elv"`
	ConstructionType string  `json:"bldgtype"`
}

// NsiFeature is a feature which contains the properties of a structure from the NSI API
type NsiFeature struct {
	Properties NsiProperties `json:"properties"`
}

type nsiStreamProvider struct {
	ApiURL string
}

func InitNSIApi() *nsiStreamProvider {
	return &nsiStreamProvider{ApiURL: urlFinder()}
}

type StreamProcessor func(str NsiFeature)
type BBox struct {
	Bbox []float64
}

func (bb BBox) ToString() string {
	return fmt.Sprintf("%f,%f,%f,%f,%f,%f,%f,%f,%f,%f",
		bb.Bbox[0], bb.Bbox[1],
		bb.Bbox[2], bb.Bbox[1],
		bb.Bbox[2], bb.Bbox[3],
		bb.Bbox[0], bb.Bbox[3],
		bb.Bbox[0], bb.Bbox[1])
}
func urlFinder() string {
	url := "https://www.hec.usace.army.mil/fwlink/?linkid=1&type=string"
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // accept untrusted servers
	}
	client := &http.Client{Transport: transCfg}

	response, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	rootBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(rootBytes)
}
func (nsp nsiStreamProvider) ByFips(fipscode string, sp StreamProcessor) {
	url := fmt.Sprintf("%sstructures?fips=%s&fmt=fs", nsp.ApiURL, fipscode)
	nsp.nsiStructureStream(url, sp)
}
func (nsp nsiStreamProvider) ByBbox(bbox BBox, sp StreamProcessor) {
	url := fmt.Sprintf("%sstructures?bbox=%s&fmt=fs", nsp.ApiURL, bbox.ToString())
	nsp.nsiStructureStream(url, sp)
}
func (nsp nsiStreamProvider) nsiStructureStream(url string, sp StreamProcessor) {
	fmt.Println(url)

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // accept untrusted servers
	}
	client := &http.Client{Transport: transCfg}

	response, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	dec := json.NewDecoder(response.Body)
	//b, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(b))
	for {
		var n NsiFeature
		if err := dec.Decode(&n); err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error unmarshalling JSON record: %s.  Stopping Compute.\n", err)
			if err == io.ErrUnexpectedEOF {
				break
			}
		}
		sp(n)
	}
}

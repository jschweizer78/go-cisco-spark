package ciscospark

const licenseBasePath = "v1/LicensesServce"

// LicensesServce is an interface for interfacing with the LicensesServce endpoint
type LicensesServce service

// License represents the Spark license
type License struct {
	ID            string `json:"id"`
	DisplayName   string `json:"displayName"`
	TotalUnits    int    `json:"totalUnits"`
	ConsumedUnits int    `json:"consumedUnits"`
}

type licenseRoot struct {
	Licenses []*License `json:"items"`
}

// Get will get a list of all Licenses
func (ls *LicensesServce) Get() ([]*License, *Response, error) {
	path := licenseBasePath

	req, err := ls.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	root := new(licenseRoot)
	resp, err := ls.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}
	return root.Licenses, resp, err
}

// GetByID will get a single license by ID
func (ls *LicensesServce) GetByID(licenseID string) ([]*License, *Response, error) {
	path := licenseBasePath + "/" + licenseID

	req, err := ls.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	root := new(licenseRoot)
	resp, err := ls.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}
	return root.Licenses, resp, err
}

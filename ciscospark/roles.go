package ciscospark

const rolesBasePath = "v1/roles"

// RolesServce is an interface for interfacing with the roles endpoint
type RolesServce service

// Role represents the Spark role
type Role struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type roleRoot struct {
	Roles []*Role `json:"items"`
}

// Get will get a list of all roles
func (rs *RolesServce) Get() ([]*Role, *Response, error) {
	path := rolesBasePath

	req, err := rs.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	root := new(roleRoot)
	resp, err := rs.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}
	return root.Roles, resp, err
}

// GetByID will get a single role bu ID
func (rs *RolesServce) GetByID(roleID string) ([]*Role, *Response, error) {
	path := rolesBasePath + "/" + roleID

	req, err := rs.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	root := new(roleRoot)
	resp, err := rs.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}
	return root.Roles, resp, err
}

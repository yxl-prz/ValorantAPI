package ValorantAPI

import (
	"encoding/json"
	"io/ioutil"
)

func (dt *API_Response) Save(path string) error {
	dat, err := json.MarshalIndent(dt, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, dat, 0777)
}

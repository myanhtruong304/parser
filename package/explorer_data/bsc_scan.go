package explorerData

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/myanhtruong304/parser/package/model"
)

func (e *Explorer) GetBlockAtTimestamp(timestamp string) (*model.BlockByTimeStampResponse, error) {
	const (
		module  = "block"
		action  = "getblocknobytime"
		closest = "before"
	)

	Chain := e.ChainSelect("bsc")

	uri := fmt.Sprintf("%s?module=%s&action=%s&timestamp=%s&closest=%s&apikey=%s", Chain.ExplorerUri, module, action, timestamp, closest, Chain.APIKey)

	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var blockAtTimestamp model.BlockByTimeStampResponse

	err = json.Unmarshal(body, &blockAtTimestamp)
	if err != nil {
		return nil, err
	}

	return &blockAtTimestamp, nil
}

package explorerData

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/myanhtruong304/parser/package/model"
	"github.com/myanhtruong304/parser/utils"
)

func (e *Explorer) GetBlockAtTimestamp(timestamp string) (*model.BlockByTimeStampResponse, error) {
	const (
		module  = "block"
		action  = "getblocknobytime"
		closest = "before"
	)

	chain := utils.ChainSelect(e.chain, e.cfg)
	uri := fmt.Sprintf("%s?module=%s&action=%s&timestamp=%s&closest=%s&apikey=%s", chain.ExplorerUri, module, action, timestamp, closest, chain.ExplorerApiKey)

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

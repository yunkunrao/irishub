package record

//-----------------------------------------------------------
// Record which include metadata
type Record struct {
	ownerAddress string `json:"ownerAddress"`
	submitTime   string `json:"submitTime"`
	dataHash     string `json:"datahash"`
	dataSize     string `json:"datasize"`
	pinedNode    string `json:"pineNode"`
}

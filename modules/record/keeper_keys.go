package record

import "fmt"

var (
	fileHashKey = "ipfs/%s"
)

func GetFileHashKey(fileHash string) []byte {
	return []byte(fmt.Sprintf(fileHashKey, fileHash))
}

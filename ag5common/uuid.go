package ag5common

import "fmt"

type UUID [16]byte

func (uuid UUID) String() string {
	return EncodeUUID(uuid)
}

func EncodeUUID(uuid [16]byte) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

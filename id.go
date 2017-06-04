package randutil

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
)

func RandomID() (string, error) {
	var buf [26]byte
	_, err := rand.Read(buf[:16])
	if err != nil {
		return "", fmt.Errorf("failed to generate RandomID, err=%v", err)
	}

	u1 := binary.LittleEndian.Uint64(buf[:8])
	u2 := binary.LittleEndian.Uint64(buf[8:])
	p := strconv.AppendUint(buf[:0], u1, 36)
	p = strconv.AppendUint(p, u2, 36)
	return string(p), nil
}

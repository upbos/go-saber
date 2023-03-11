package uid

import (
	"github.com/speps/go-hashids/v2"
	"strconv"
)

const (
	base      = 16                 // 16进制
	bitSize   = 64                 // int64
	alphabet  = "0123456789abcdef" // 16进制字符表
	minLength = 8                  // 最小长度
)

var hashID *hashids.HashID

func Init(salt string) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	hd.Alphabet = alphabet
	hashID, _ = hashids.NewWithData(hd)
}

func Uid(id int64) int64 {
	hash, _ := hashID.EncodeInt64([]int64{id})
	uid, _ := strconv.ParseInt(hash, base, bitSize)
	return uid
}

func Id(uid int64) int64 {
	hex := strconv.FormatInt(uid, base)
	id, _ := hashID.DecodeInt64WithError(hex)
	return id[0]
}

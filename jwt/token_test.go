package jwt

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	Setup("acc@2022")
}

func TestGenerateToken(t *testing.T) {
	duration := 1 * time.Hour
	token, err := GenerateToken(1, "200.22.11.31", duration)
	if err != nil {
		t.Errorf("生成令牌错误, %v", err)
		return
	}

	fmt.Printf("jwt: %s \n", token)
}

func TestParseToken(t *testing.T) {
	duration := 1 * time.Hour
	token, err := GenerateToken(1, "200.22.11.31", duration)
	if err != nil {
		t.Errorf("生成令牌错误, %v", err)
		return
	}
	claims, err := ParseToken(token)
	if err != nil {
		t.Errorf("解析token错误, %v", err)
		return
	}

	if claims.Uid != 1 {
		t.Error("get uid error")
	}
	if claims.IP != "200.22.11.31" {
		t.Error("get ip error")
	}

	fmt.Printf("uid: %v\n", claims.Uid)
	fmt.Printf("ip: %v\n", claims.IP)
	fmt.Printf("id: %v\n", claims.RegisteredClaims.ID)
	fmt.Printf("expire time: %v\n", claims.RegisteredClaims.ExpiresAt)
}

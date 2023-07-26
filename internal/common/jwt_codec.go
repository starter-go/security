package common

import (
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/starter-go/base/lang"
	sejwt "github.com/starter-go/security/jwt"
)

// JWTCodec ...
type JWTCodec struct {

	//starter:component
	_as func(sejwt.Registry) //starter:as(".")

	Secret string //starter:inject("${security.jwt.secret}")

}

func (inst *JWTCodec) _impl() (sejwt.CODEC, sejwt.Registry) {
	return inst, inst
}

// ListRegistrations ...
func (inst *JWTCodec) ListRegistrations() []*sejwt.Registration {
	r1 := &sejwt.Registration{
		Priority: 0,
		Enabled:  true,
		CODEC:    inst,
	}
	return []*sejwt.Registration{r1}
}

func (inst *JWTCodec) getSecretKey() []byte {
	hex := lang.Hex(inst.Secret)
	return hex.Bytes()
}

// Encode ...
func (inst *JWTCodec) Encode(t1 *sejwt.DTO) (sejwt.Text, error) {

	key := inst.getSecretKey()
	t := jwt.New(jwt.SigningMethodHS256)

	if t1 == nil {
		t1 = &sejwt.DTO{}
	}

	tokenJSON, err := json.Marshal(t1)
	if err != nil {
		return "", err
	}

	claims := make(jwt.MapClaims)
	claims["token"] = string(tokenJSON)
	t.Claims = claims

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return sejwt.Text(s), err
}

// Decode ...
func (inst *JWTCodec) Decode(t1 sejwt.Text) (*sejwt.DTO, error) {

	tokenString := t1.String()
	hmacSampleSecret := inst.getSecretKey()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("bad jwt.claims")
	}

	if !token.Valid {
		return nil, fmt.Errorf("bad jwt.valid")
	}

	tokenJSON := fmt.Sprint(claims["token"])
	t2 := &sejwt.DTO{}
	err = json.Unmarshal([]byte(tokenJSON), t2)
	if err != nil {
		return nil, err
	}

	err = inst.verify(t2)
	if err != nil {
		return nil, err
	}

	return t2, nil
}

func (inst *JWTCodec) verify(o *sejwt.DTO) error {
	now := lang.Now()
	t1 := o.CreatedAt
	t2 := o.ExpiredAt
	if now < t1 || t2 < now {
		return fmt.Errorf("bad token")
	}
	return nil
}

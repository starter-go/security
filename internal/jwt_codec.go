package internal

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/jwt"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// JWTCODEC ... 实现 JWT 的编解码
type JWTCODEC struct {

	//starter:component
	_as func(jwt.Registry) //starter:as(".")

	secret lang.Base64
}

func (inst *JWTCODEC) _impl() (jwt.CODEC, jwt.Registry, application.Lifecycle) {
	return inst, inst, inst
}

// Life ...
func (inst *JWTCODEC) Life() *application.Life {
	return &application.Life{OnCreate: inst.init}
}

func (inst *JWTCODEC) init() error {
	data := make([]byte, 36)
	rand.Reader.Read(data)
	inst.secret = lang.Base64FromBytes(data)
	return nil
}

// ListRegistrations ...
func (inst *JWTCODEC) ListRegistrations() []*jwt.Registration {
	r1 := &jwt.Registration{
		Enabled:  true,
		Priority: 0,
		CODEC:    inst,
	}
	return []*jwt.Registration{r1}
}

// Encode ...
func (inst *JWTCODEC) Encode(t1 *jwt.Token) (jwt.Text, error) {

	key := inst.secret.Bytes()
	t := jwtv5.New(jwtv5.SigningMethodHS256)

	if t1 == nil {
		t1 = &jwt.Token{}
	}

	tokenJSON, err := json.Marshal(t1)
	if err != nil {
		return "", err
	}

	claims := make(jwtv5.MapClaims)
	claims["token"] = string(tokenJSON)
	t.Claims = claims

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return jwt.Text(s), err
}

// Decode ...
func (inst *JWTCODEC) Decode(t1 jwt.Text) (*jwt.Token, error) {

	tokenString := t1.String()
	hmacSampleSecret := inst.secret.Bytes()

	token, err := jwtv5.Parse(tokenString, func(token *jwtv5.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtv5.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwtv5.MapClaims)
	if !ok {
		return nil, fmt.Errorf("bad jwt.claims")
	}

	if !token.Valid {
		return nil, fmt.Errorf("bad jwt.valid")
	}

	tokenJSON := fmt.Sprint(claims["token"])
	t2 := &jwt.Token{}
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

func (inst *JWTCODEC) verify(o *jwt.Token) error {
	now := lang.Now()
	t1 := o.StartedAt
	t2 := o.ExpiredAt
	if t1 <= now && now <= t2 {
		return nil
	}
	return fmt.Errorf("bad token")
}

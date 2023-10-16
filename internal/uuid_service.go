package internal

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"strconv"
	"sync"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
)

// UUIDServiceImpl 实现 UUID 创建服务
type UUIDServiceImpl struct {

	//starter:component
	_as func(random.UUIDService) //starter:as("#")

	Host string //starter:inject("${security.uuid.service.hostname}")

	lock      sync.Mutex
	sn        int64
	startedAt lang.Time
	chain     string
}

func (inst *UUIDServiceImpl) _impl() {
	inst._as(inst)
}

// Life ...
func (inst *UUIDServiceImpl) Life() *application.Life {
	l := &application.Life{}
	l.OnCreate = func() error {
		inst.startedAt = lang.Now()
		inst.chain = inst.nextNonce()
		return nil
	}
	return l
}

// Build ...
func (inst *UUIDServiceImpl) Build() random.UUIDBuilder {
	return &uuidBuilder{service: inst}
}

func (inst *UUIDServiceImpl) prepareGenerateSafe(b *uuidBuilder) {
	inst.lock.Lock()
	defer inst.lock.Unlock()

	inst.sn++

	sum := sha1.Sum([]byte(inst.chain))
	hex := lang.HexFromBytes(sum[:])
	next := hex.String()
	inst.chain = next

	b.sn = inst.sn
	b.chain = next
}

func (inst *UUIDServiceImpl) updateChain(data []byte) {
	hex := lang.HexFromBytes(data)
	inst.lock.Lock()
	defer inst.lock.Unlock()
	inst.chain = hex.String()
}

func (inst *UUIDServiceImpl) nextNonce() string {
	var data [20]byte
	buffer := data[:]
	rand.Reader.Read(buffer)
	hex := lang.HexFromBytes(buffer)
	return hex.String()
}

func (inst *UUIDServiceImpl) prepareGenerate(b *uuidBuilder) []byte {

	b.createdAt = lang.Now()
	b.startedAt = inst.startedAt
	b.host = inst.Host
	b.nonce = inst.nextNonce()

	// set: sn,chain
	inst.prepareGenerateSafe(b)

	buffer := bytes.Buffer{}
	buffer.WriteString("sn:" + strconv.FormatInt(b.sn, 10))
	buffer.WriteString("    t1:" + b.createdAt.String())
	buffer.WriteString("    t0:" + b.startedAt.String())
	buffer.WriteString(" nonce:" + b.nonce)
	buffer.WriteString("  host:" + b.host)
	buffer.WriteString(" chain:" + b.chain)
	buffer.WriteString(" class:" + b.class)
	buffer.WriteString("    id:" + b.id)
	for _, o := range b.others {
		buffer.WriteString(" other:" + o)
	}
	return buffer.Bytes()
}

////////////////////////////////////////////////////////////////////////////////

type uuidBuilder struct {
	service *UUIDServiceImpl

	sn        int64
	nonce     string
	chain     string
	host      string
	createdAt lang.Time
	startedAt lang.Time
	class     string
	id        string
	others    []string
}

func (inst *uuidBuilder) _impl() random.UUIDBuilder {
	return inst
}

func (inst *uuidBuilder) Generate() lang.UUID {
	data := inst.service.prepareGenerate(inst)
	sum := sha1.Sum(data)
	uuid := sum[:]
	inst.service.updateChain(uuid)
	return lang.NewUUID(uuid)
}

func (inst *uuidBuilder) Others(value ...string) random.UUIDBuilder {
	inst.others = value
	return inst
}

func (inst *uuidBuilder) Class(value string) random.UUIDBuilder {
	inst.class = value
	return inst
}

func (inst *uuidBuilder) ID(value string) random.UUIDBuilder {
	inst.id = value
	return inst
}

////////////////////////////////////////////////////////////////////////////////

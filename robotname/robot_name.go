package robotname

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var reg nameregistry

type nameregistry map[string]struct{}

type Robot struct{ name string }

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func New() *Robot {
	if reg == nil {
		reg = make(map[string]struct{})
	}
	s := make([]string, 5)
	for {
		s[0] = string(charset[random.Intn(len(charset))])
		s[1] = string(charset[random.Intn(len(charset))])
		s[2] = strconv.Itoa(random.Intn(10))
		s[3] = strconv.Itoa(random.Intn(10))
		s[4] = strconv.Itoa(random.Intn(10))
		name := strings.Join(s, "")
		if _, ok := reg[name]; !ok {
			reg[name] = struct{}{}
			return &Robot{name}
		}
	}
}

func (r *Robot) Name() string {
	return r.name
}

func (r *Robot) Reset() {
	temp := New()
	r.name = temp.Name()
}

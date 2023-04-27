package docs

import (
	"testing"

	"github.com/kr/pretty"

	"github.com/stretchr/testify/assert"
)

func TestPlace(t *testing.T) {

	p := NewPlace("UK")
	q := NewPlace("")

	println(p.ID)
	println(q.ID)

	if p.ID == q.ID {
		t.Fail()
	}

	pretty.Println(p)
	pretty.Println(p.ParentHashes())

}

func TestPlaceParenthashes(t *testing.T) {
	assert := assert.New(t)

	input1 := "43, Reading Street, Broadstoars, Kend, UK, ENgladk, World, Universe, Stuff"

	input2 := "43, Reading Street, Broadstoars, Kend, UN, ENgladk, World, Universe, Stuff"

	p := NewPlace(input1)
	q := NewPlace(input2)
	pretty.Println(p)
	pretty.Println(q)

	ph := p.ParentHashes()
	qh := q.ParentHashes()

	for x := 0; x < 4; x++ {
		assert.Equal(ph[x], qh[x])
	}
	for x := 4; x < len(ph); x++ {
		assert.NotEqual(ph[x], qh[x])
	}
	pretty.Println(ph)
	pretty.Println(qh)

}

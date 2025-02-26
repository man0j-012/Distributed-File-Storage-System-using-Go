// run package tests | run file tests
package p2p

import(
	"testing"
	"github.com/stretchr/testify/assert"
)


// run test | debug test

func TestTCPTransport(t *testing.T){
	listenAddr := ":4000"
	tr:= NewTCPTransport(listenAddr)

	assert.Equal(t, tr.listenAddress, listenAddr)

tr.ListenAndAccept()

}
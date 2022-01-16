package goods

import (
	"testing"

	"github.com/gorilla/websocket"
)

func Test_websocket(t *testing.T) {
	d := websocket.Dialer{}
	conn, resp, err := d.Dial("wss://ws-test.yime.app", nil)
	t.Logf("%v", resp)
	AssertNoError(t, err)
	conn.Close()
}

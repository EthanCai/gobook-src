package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"success", "ECHO:" + method + " " + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("method1", "params1")
	resp2, _ := client2.Call("method2", "params2")

	if resp1.Body != "ECHO:method1 params1" || resp2.Body != "ECHO:method2 params2" {
		t.Error("IpcClient.Call failed. resp1:", resp1.Body, "resp2:", resp2.Body)
	}

	client1.Close()
	client2.Close()
}

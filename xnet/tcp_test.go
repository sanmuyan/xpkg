package xnet

import (
	"bufio"
	"encoding/json"
	"io"
	"net"
	"testing"
)

func TestTCPEncodeAndDecode(t *testing.T) {
	type Data struct {
		Name string
	}
	name := "hello"
	d := Data{Name: name}
	encode := func(conn net.Conn) error {
		pkg, err := json.Marshal(d)
		if err != nil {
			return err
		}
		pkg, err = Encode(pkg)
		if err != nil {
			return err
		}
		_, err = conn.Write(pkg)
		if err != nil {
			return err
		}
		return nil
	}

	decode := func(reader *bufio.Reader) (*Data, error) {
		pkg, err := Decode(reader)
		if err != nil {
			return nil, err
		}
		d := &Data{}
		err = json.Unmarshal(pkg, d)
		if err != nil {
			return nil, err
		}
		return d, nil
	}

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Error(err)
		return
	}
	go func() {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			t.Error(err)
			return
		}
		defer func() {
			_ = conn.Close()
		}()
		reader := bufio.NewReader(conn)
		for i := 0; i < 100; i++ {
			err := encode(conn)
			if err != nil {
				t.Error(err)
				return
			}
			d, err := decode(reader)
			if err != nil {
				t.Error(err)
				return
			}
			if d.Name != name {
				t.Errorf("name: %s != %s", d.Name, name)
				return
			}
		}
	}()
	conn, err := listener.Accept()
	if err != nil {
		t.Error(err)
		return
	}
	count := 1
	reader := bufio.NewReader(conn)
	for {
		t.Logf("count: %d", count)
		d, err := decode(reader)
		if err != nil {
			if err == io.EOF {
				return
			}
			t.Error(err)
			return
		}
		if d.Name != name {
			t.Errorf("name: %s != %s", d.Name, name)
			return
		}
		err = encode(conn)
		if err != nil {
			t.Error(err)
			return
		}
		count++
	}
}

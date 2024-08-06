package benchmark_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpc_server "github.com/plutov/packagemain/benchmark-http-grpc/grpc"
	"github.com/plutov/packagemain/benchmark-http-grpc/grpc/gen"
	http_server "github.com/plutov/packagemain/benchmark-http-grpc/http"
)

func init() {
	go grpc_server.Start()
	go http_server.StartHTTP1()
	go http_server.StartHTTP2()

	// just to make sure the servers are running
	time.Sleep(2 * time.Second)
}

func BenchmarkGRPCProtobuf(b *testing.B) {
	conn, err := grpc.NewClient("localhost:60000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		b.Fatalf("grpc connection failed: %v", err)
	}

	client := gen.NewUsersClient(conn)

	for n := 0; n < b.N; n++ {
		doGRPC(client, b)
	}
}

func doGRPC(client gen.UsersClient, b *testing.B) {
	resp, err := client.CreateUser(context.Background(), &gen.User{
		Id:    "1001",
		Email: "foo@bar.com",
		Name:  "Bench",
	})

	if err != nil {
		b.Fatalf("grpc request failed: %v", err)
	}

	if resp.Code != 201 || resp.User.Id != "1001" {
		b.Fatalf("grpc response is wrong: %v", resp)
	}
}

func newClientTLS() *http.Client {
	caCert, err := os.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Failed to load CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 配置TLS
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}
	// tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	return &http.Client{
		Transport: transport,
	}
}

func BenchmarkHTTP1JSON(b *testing.B) {
	client := &http.Client{}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		doPost(client, "http://localhost:60001", b)
	}
}

func BenchmarkHTTPS1JSON(b *testing.B) {
	client := newClientTLS()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		doPost(client, "https://localhost:60002", b)
	}
}

func BenchmarkHTTP2JSON(b *testing.B) {
	client := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			// Pretend we are dialing a TLS endpoint. (Note, we ignore the passed tls.Config)
			DialTLSContext: func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
				var d net.Dialer
				return d.DialContext(ctx, network, addr)
			},
		},
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		doPost(client, "http://localhost:60003", b)
	}
}

func doPost(client *http.Client, url string, b *testing.B) {
	u := &http_server.User{
		ID:    "1001",
		Email: "foo@bar.com",
		Name:  "Bench",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(u)

	req, _ := http.NewRequest("POST", url, buf)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		b.Fatalf("http request failed: %v", err)
	}

	defer resp.Body.Close()

	var r http_server.Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		b.Fatalf("unable to decode json: %v", err)
	}

	if r.Code != 201 || r.User.ID != "1001" {
		b.Fatalf("http response is wrong: %v", resp)
	}
}

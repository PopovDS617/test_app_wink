package lb

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"testappwink/internal/service/mocks"
	"testing"

	"testappwink/pkg/customerr"
	gen "testappwink/pkg/lb_v1"

	"github.com/gojuno/minimock/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestLBImplementation_Get_Success(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	mc := minimock.NewController(t)
	srvcMockWithoutErr := mocks.NewLBServiceMock(mc).GetURLMock.Return("http://cdn.com/video/123", nil)

	i := NewImplementation(srvcMockWithoutErr)

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() { lis.Close() })

	server := grpc.NewServer()
	t.Cleanup(func() { server.Stop() })

	gen.RegisterLBV1Server(server, i)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("server.Serve: %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	t.Cleanup(func() { cancel() })

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	t.Cleanup(func() { conn.Close() })

	if err != nil {
		t.Fatalf("grpc.DialContext: %v", err)
	}

	client := gen.NewLBV1Client(conn)

	_, err = client.GetVideoURL(context.Background(), &gen.GetVideoURLRequest{Video: "http://s1.test.com/video/123"})

	if err != nil {
		t.Fatalf("client.GetVideoURL: %v", err)
	}

}
func TestLBImplementation_Get_Error(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	mc := minimock.NewController(t)
	srvcMockWithErr := mocks.NewLBServiceMock(mc).GetURLMock.Return("", customerr.ErrInvalidURL)

	i := NewImplementation(srvcMockWithErr)

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() { lis.Close() })

	server := grpc.NewServer()
	t.Cleanup(func() { server.Stop() })

	gen.RegisterLBV1Server(server, i)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("server.Serve: %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	t.Cleanup(func() { cancel() })

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	t.Cleanup(func() { conn.Close() })

	if err != nil {
		t.Fatalf("grpc.DialContext: %v", err)
	}

	client := gen.NewLBV1Client(conn)

	res, err := client.GetVideoURL(context.Background(), &gen.GetVideoURLRequest{Video: "http://s1.test.com/video/123"})

	if err == nil {
		t.Fatalf("expected err to be returned, but received: %v", err)
	}

	fmt.Println(err)

	if res != nil {
		t.Fatalf("expected res to be nil, but received: %v", res.Video)
	}

}

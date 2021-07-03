package dbsdk

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/r-coffee/db-append-only-sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	connectionTimeout = 5 * time.Second
	requestTimeout    = 3 * time.Second
)

type AppendDbSDKClient struct {
	stub proto.DBServiceClient
}

// CreateAppendDBClient creates a new sdk client
// host is the hostname of the server
// pathToCert is the path to the server's public certificate
// port is the port number the server service is running on
func CreateAppendDBClient(host, pathToCert string, port int) *AppendDbSDKClient {
	var sdk AppendDbSDKClient
	creds, err := credentials.NewClientTLSFromFile(pathToCert, host)
	if err != nil {
		log.Fatal(err)
	}

	// connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	sdk.stub = proto.NewDBServiceClient(conn)
	return &sdk
}

// Append will write a new row to the table
func (s *AppendDbSDKClient) Append(table string, ts time.Time, dat []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var tup proto.DBTuple
	tup.Ts = ts.UnixNano()
	tup.Data = dat

	_, err := s.stub.Append(ctx, &proto.AppendRequest{Table: table, Data: &tup})
	return err
}

// Query will return all the rows for a table that are between start and stop inclusive
func (s *AppendDbSDKClient) Query(table string, start, stop time.Time) ([]*proto.DBTuple, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	resp, err := s.stub.Query(ctx, &proto.QueryRequest{Table: table, Start: start.UnixNano(), Stop: stop.UnixNano()})

	// handle nil response
	if resp == nil {
		return nil, err
	}

	return resp.Data, err
}

// Stats returns some statistics about the table
func (s *AppendDbSDKClient) Stats(table string) (*proto.TableStatTuple, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	return s.stub.Stats(ctx, &proto.TableRequest{Table: table})
}

// ListTables returns a list of all the tables in the server
func (s *AppendDbSDKClient) ListTables() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	resp, err := s.stub.ListTables(ctx, &proto.Empty{})
	return resp.Tables, err
}

// Purge removes a table and all of it's data from the server
func (s *AppendDbSDKClient) Purge(table string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	_, err := s.stub.Purge(ctx, &proto.TableRequest{Table: table})
	return err
}

package dbsdk

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type AppendDbSDKClient struct {
	stub DBServiceClient
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

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(creds), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	sdk.stub = NewDBServiceClient(conn)
	return &sdk
}

// Append will write a new row to the table
func (s *AppendDbSDKClient) Append(table string, ts time.Time, dat []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tup DBTuple
	tup.Ts = ts.UnixNano()
	tup.Data = dat

	_, err := s.stub.Append(ctx, &AppendRequest{Table: table, Data: &tup})
	return err
}

// Query will return all the rows for a table that are between start and stop inclusive
func (s *AppendDbSDKClient) Query(table string, start, stop time.Time) ([]*DBTuple, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := s.stub.Query(ctx, &QueryRequest{Table: table, Start: start.UnixNano(), Stop: stop.UnixNano()})

	return resp.Data, err
}

package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"docker-grpc-service-tutorial/configreader"
	"docker-grpc-service-tutorial/poetrydb"
	poetry "docker-grpc-service-tutorial/proto"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

// These global variables makes it easy
// to mock these dependencies
// in unit tests.
var (
	netListen           = net.Listen
	configreaderReadEnv = configreader.ReadEnv
	jsonMarshal         = json.Marshal
	protojsonUnmarshal  = protojson.Unmarshal
)

// Server defines the available operations for gRPC server.
type Server interface {
	// Serve is called for serving requests.
	Serve() error
	// GracefulStop is called for stopping the server.
	GracefulStop()
	// RandomPoetries returns a random list of poetries.
	RandomPoetries(ctx context.Context, in *poetry.RandomPoetriesRequest) (*poetry.PoetryList, error)
}

// server implements Server.
type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	poetryDb   poetrydb.PoetryDb
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) GracefulStop() {
	s.grpcServer.GracefulStop()
}

// NewServer creates a new gRPC server.
func NewServer(port int) (Server, error) {
	server := new(server)
	listener, err := netListen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return server, errors.Wrap(err, "tcp listening")
	}
	server.listener = listener
	config, err := configreaderReadEnv()
	if err != nil {
		return server, errors.Wrap(err, "reading env vars")
	}
	server.poetryDb = poetrydb.NewPoetryDb(config.PoetrydbBaseUrl, config.PoetrydbHttpTimeout)
	server.grpcServer = grpc.NewServer()
	poetry.RegisterProtobufServiceServer(server.grpcServer, server)
	reflection.Register(server.grpcServer)
	return server, nil
}

func (s *server) RandomPoetries(ctx context.Context, in *poetry.RandomPoetriesRequest) (*poetry.PoetryList, error) {
	pbPoetryList := new(poetry.PoetryList)
	poetryList, err := s.poetryDb.Random(int(in.NumberOfPoetries))
	if err != nil {
		return pbPoetryList, errors.Wrap(err, "requesting random poetry")
	}
	json, err := jsonMarshal(poetryList)
	if err != nil {
		return pbPoetryList, errors.Wrap(err, "marshalling json")
	}
	err = protojsonUnmarshal(json, pbPoetryList)
	if err != nil {
		return pbPoetryList, errors.Wrap(err, "unmarshalling proto")
	}
	return pbPoetryList, nil
}

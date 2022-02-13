package streaminginterceptor

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/manofthelionarmy/learngRPC/ch5/interceptors/streaming"
	"google.golang.org/grpc"
)

var _ pb.PokemonServiceServer = (*PokemonStreamServer)(nil)

// PokemonStreamServer server implements pokemon grpc service
type PokemonStreamServer struct {
	Pokemon map[string][]*pb.Pokemon
}

var _ grpc.ServerStream = (*wrappedPokemonStream)(nil)

// TODO: study interface wrapper pattern
type wrappedPokemonStream struct {
	grpc.ServerStream
}

// AddPokemon implements the remote funciton exposed by our Pokemen grpc service
func (s PokemonStreamServer) AddPokemon(
	ctx context.Context,
	pokemon *pb.Pokemon,
) (*pb.PokemonId, error) {
	tag := pokemon.Tag.Value
	if _, ok := s.Pokemon[tag]; !ok {
		s.Pokemon[tag] = make([]*pb.Pokemon, 0)
	}

	s.Pokemon[tag] = append(s.Pokemon[tag], pokemon)
	for _, p := range s.Pokemon {
		pokemon.Id += int64(len(p))
	}
	return &pb.PokemonId{Id: pokemon.Id}, nil
}

// GetPokemonByTag implements the remote function exposed by our Pokemon grpc service
func (s PokemonStreamServer) GetPokemonByTag(
	tag *pb.PokemonTag,
	serverStream pb.PokemonService_GetPokemonByTagServer,
) error {
	pokemon, ok := s.Pokemon[tag.Value]
	if !ok {
		return fmt.Errorf("no pokemon found for tag %s", tag.Value)
	}

	// I'm confused... we can send messages to the serverStream too, what is the use case?
	// var msg interface{}
	// err := serverStream.RecvMsg(msg) // blocks until a message is received... what are the use cases?

	for _, p := range pokemon {
		err := serverStream.Send(p)
		return err
	}
	return nil
}

// StreamServerInterceptor is a funciton type defined by grpc.StreamServerInterceptor
// This will intercept the message before it reaches the service
func pokmeonStreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	// Pre-process
	log.Println("=== [Server Stream Interceptor] ===", info.FullMethod)
	// Invoke
	// wrap our stream server
	err := handler(srv, newWrappedPokemonStream(ss))
	// Post-process
	if err != nil {
		return fmt.Errorf("RPC failed with error %+v", err)
	}
	return nil
}

// pass in our server stream
func newWrappedPokemonStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedPokemonStream{ServerStream: s}
}

// These are invoked when the service sends a message to the server stream
func (w *wrappedPokemonStream) SendMsg(m interface{}) error {
	log.Printf("=== [Server Stream Interceptor] === Received a message (Type: %T) at %s\n",
		m,
		time.Now().Format(time.RFC3339),
	)
	return w.ServerStream.SendMsg(m)
}

// This is invoked when the service recieves a message on the server stream
func (w *wrappedPokemonStream) RecvMsg(m interface{}) error {
	log.Printf("=== [Server Stream Interceptor] === Sent a message (Type: %T) at %s\n",
		m,
		time.Now().Format(time.RFC3339),
	)
	return w.ServerStream.RecvMsg(m)
}

type foo interface {
	implementFunction(string)
}

// defining a function type
type containerFunction func(string)

// containerFucntion has a method called embeddedFunciton, and it satisfies foo interface
func (cf containerFunction) implementFunction(s string) {
	// once we invoke containerFunction's method, we can call the containerFunction beause,
	// well it's a function
	s += " world"
	cf(s)
}

func result() {
	cf := containerFunction(func(greeting string) {
		fmt.Println(greeting + "!!!!!")
	})

	fooer := newFooWrapper(cf)

	fooer.implementFunction("hello")
}

type fooWrapper struct {
	foo // we're embedding the interface which promotes implementfunction to fooWrapper
}

var _ foo = (*fooWrapper)(nil)

// in terms of the call stack, the container/wrapper will be called first
func (w *fooWrapper) implementFunction(greeting string) {
	greeting += " I'm excited to code!"
	// we can call implementFunction by using our embedded interface since it is a field
	w.foo.implementFunction(greeting)
}

func newFooWrapper(fooer foo) foo {
	return &fooWrapper{fooer}
}

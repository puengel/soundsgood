package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/faiface/beep/wav"
	pb "github.com/puengel/soundsgood/soundService"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedAudioStreamServer
}

func (s *server) GetFormat(context.Context, *emptypb.Empty) (*pb.AudioFormat, error) {
	fmt.Println("GetFormat handler called")

	f, err := os.Open("../CantinaBand60.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.AudioFormat{
		SampleRate:  int32(format.SampleRate),
		NumChannels: int32(format.NumChannels),
		Precision:   int32(format.Precision),
	}, nil
}

func (s *server) GetStream(stream pb.AudioStream_GetStreamServer) error {
	fmt.Println("GetStream handler called")

	f, err := os.Open("../CantinaBand60.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	streamer, _, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// Create StreamService
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Printf("what's that?: %s\n", err)
			return err
		}

		chunk := make([][2]float64, int(r.Amount))

		n, ok := streamer.Stream(chunk)
		if !ok || n == 0 {
			fmt.Println("Finished sending file")
			return nil
		}

		ch1 := make([]float64, int(r.Amount))
		ch2 := make([]float64, int(r.Amount))

		for i := 0; i < int(r.Amount); i++ {
			ch1[i] = chunk[i][0]
			ch2[i] = chunk[i][0]
		}

		err = stream.Send(&pb.AudioSample{
			Timestamp: "hello",
			Channel1:  ch1,
			Channel2:  ch2,
		})
		if err != nil {
			fmt.Printf("come on: %s\n", err)
			return err
		}

	}
}

func main() {

	RunServer()

}

func RunServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAudioStreamServer(s, &server{})

	fmt.Println("listening...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

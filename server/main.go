package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/puengel/soundsgood/soundService"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

const (
	port    = ":50051"
	address = "localhost:50051"
)

type server2 struct {
	pb.UnimplementedAudioStreamServer
}

func (s *server2) SendFormat(ctx context.Context, in *pb.AudioFormat) (*emptypb.Empty, error) {
	log.Printf("AudioFormat: %+v\n", in)
	return &emptypb.Empty{}, nil
}

func main() {
	GetClient()

	select {}
}

func GetClient() {
	ctx := context.Background()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewAudioStreamClient(conn)

	format, err := client.GetFormat(ctx, &emptypb.Empty{})
	if err != nil {
		fmt.Printf("oh no: %s\n", err)
		return
	}

	fmt.Println(format)

	bufferSize := int(time.Second / 10 * time.Duration(format.SampleRate) / time.Second)

	err = speaker.Init(beep.SampleRate(format.SampleRate), bufferSize)

	if err != nil {
		fmt.Printf("Init Speaker: %s\n", err)
		return
	}

	stream, err := client.GetStream(ctx)
	if err != nil {
		fmt.Printf("nonononono: %s\n", err)
		return
	}

	var playing bool

	aStream := audioStream{
		err:      nil,
		channel1: make(chan []float64, 512),
		channel2: make(chan []float64, 512),
	}

	var done bool
	for !done {

		err := stream.Send(&pb.SampleRequest{
			Amount: int32(bufferSize),
		})
		if err != nil {
			fmt.Printf("Send sample request: %s\n", err)
			return
		}

		sample, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Recieve sample: %s\n", err)
			return
		}
		aStream.channel1 <- sample.Channel1
		aStream.channel2 <- sample.Channel2

		if !playing {
			playing = true
			speaker.Play(&aStream, beep.Callback(func() {
				// done = true
			}))
		}
	}

}

type audioStream struct {
	err      error
	channel1 chan []float64
	channel2 chan []float64

	pos int

	ch1 []float64
	ch2 []float64
}

func (a *audioStream) Stream(samples [][2]float64) (n int, ok bool) {

	for i := range samples {

		if (len(a.ch1) == 0 || a.pos >= len(a.ch1)) ||
			(len(a.ch2) == 0 || a.pos >= len(a.ch2)) {
			select {
			case a.ch1 = <-a.channel1:
			default:
				fmt.Println("no stream data")
				return n, false
			}
			select {
			case a.ch2 = <-a.channel2:
			default:
				fmt.Println("no stream data")
				return n, false
			}
			a.pos = 0
		}

		samples[i][0] = a.ch1[a.pos]
		samples[i][1] = a.ch2[a.pos]

		a.pos++
		n++
	}

	return n, true
}

func (a *audioStream) Err() error {
	return a.err
}

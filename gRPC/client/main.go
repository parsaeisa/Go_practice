package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "gitlab.snapp.ir/backend/proto/golang/sanjeh"
)

func main() {
	address := "0.0.0.0:50052"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect %", err)
	}
	defer conn.Close()

	client := pb.NewSanjehClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	HelloReply, err := client.GetDriverRatingsByDriverID(ctx, &pb.GetDriverRatingsByDriverIDRequest{
		RatingsCount:     2,
		DriverId:         110,
		WithReasons:      true,
		OrderByCreatedAt: pb.Sorted_DESC,
	})
	if err != nil {
		log.Fatalf("error while greeting %v", err)
	}

	println(HelloReply.Ratings[0].DriverId)
	println(HelloReply.Ratings[0].RideId)

	println(HelloReply.Ratings[1].DriverId)
	println(HelloReply.Ratings[1].RideId)

	//println(HelloReply.Ratings[0].Reasons[0])
	//println(HelloReply.Ratings[0].Reasons[1])
	//println(HelloReply.Ratings[0].Reasons[2])
	println(len(HelloReply.Ratings[0].Reasons))
	println(len(HelloReply.Ratings[1].Reasons))
}

package handler

import (
	"github.com/micro/go-log"

	example "mytest/proto/example"

	"golang.org/x/net/context"
	"mytest/diffcode"
	

)

type Example struct{}

func (e *Example) Diff(ctx context.Context, req *example.DiffRequest, rsp *example.DiffResponse) error {
	log.Log("Received Example.Diff request")
	diff := diffcode.New(req.Original, req.Revised, req.Bincardoriginal, req.Bincardrevised) 
	diff.Compose()
	diff.PrintSes()
	
//	log.Log(diffcode.CharacterCount(req.Original, req.Bincardoriginal))
//	log.Log(diffcode.TagCount(req.Original, req.Bincardoriginal))
	log.Log(diffcode.WordCount(req.Original))

	rsp.Symbols = diff.SprintSes()
	addedLength := diff.AddedLenght()
	deletedLength := diff.DeletedLenght()
	
	
	deletedPositions, addedPositions := diff.AddDelPositions()
	for i, al := range addedLength {
		rsp.Added = append(rsp.Added, &example.DiffResponse_Position{
			Pos: int64(addedPositions[i]),
			Len: int64(al),
		})
	}
	
	for i, dl := range deletedLength {
		rsp.Deleted = append(rsp.Deleted, &example.DiffResponse_Position{
			Pos: int64(deletedPositions[i]), 
			Len: int64(dl),
		})
	}


	return nil
}



// Call is a single request handler called via client.Call or the generated client code
func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	log.Log("Received Example.Call request")
	
	rsp.Msg = "hello" + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Example) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
	log.Logf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&example.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

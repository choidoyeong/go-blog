package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/choidoyeong/go-blog/blog"
)

const (
	port = ":50051"
)

var nextPostID = 1

type server struct {
	pb.UnimplementedBlogServer
	savedPosts map[string]*pb.Post
	comments   map[string][]*pb.Comment
}

func (s *server) ListPosts(empty *pb.EmptyRequest, stream pb.Blog_ListPostsServer) error {
	for _, post := range s.savedPosts {
		if err := stream.Send(post); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) CreatePost(ctx context.Context, post *pb.Post) (*pb.Post, error) {
	post.PostId = strconv.Itoa(nextPostID)
	nextPostID++
	s.savedPosts[post.PostId] = post
	s.comments[post.PostId] = make([]*pb.Comment, 0, 10)
	return s.savedPosts[post.PostId], nil
}

func (s *server) GetPost(ctx context.Context, postID *pb.PostId) (*pb.Post, error) {
	if val, ok := s.savedPosts[postID.PostId]; ok {
		return val, nil
	}

	return nil, nil
}

func (s *server) CreateComment(ctx context.Context, comment *pb.Comment) (*pb.Comment, error) {
	if val, ok := s.comments[comment.PostId]; ok {
		val = append(val, comment)
		return comment, nil
	}

	return nil, nil
}

func (s *server) ListComments(postID *pb.PostId, stream pb.Blog_ListCommentsServer) error {
	if val, ok := s.comments[postID.PostId]; ok {
		for _, comment := range val {
			if err := stream.Send(comment); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBlogServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

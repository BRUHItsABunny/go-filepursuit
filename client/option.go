package client

type Option interface {
	execute(fClient *FilePursuitClient) error
}

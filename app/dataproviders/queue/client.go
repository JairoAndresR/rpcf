package queue

type Client interface {
	Push(key, content string) error
}

package pubsub

type PubSub interface {
	Publish(topic string, payload []byte) error
	Subscribe(topic string, callback func(topic string, payload []byte))
}

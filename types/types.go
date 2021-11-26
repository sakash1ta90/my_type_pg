package types

// Type向けinterface
type MyTypes interface {
	New(value interface{})
	Get()
}

// Any型
type MyAny struct {
	value interface{}
}

func (mss MyAny) New(value interface{}) MyAny {
	mss.value = value
	return mss
}

func (mss MyAny) Get() interface{} {
	return mss.value
}

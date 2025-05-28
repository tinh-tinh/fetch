package fetch

import "context"

type CancelToken struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewCancelToken() *CancelToken {
	ctx, cancel := context.WithCancel(context.Background())
	return &CancelToken{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (t *CancelToken) Cancel() {
	t.cancel()
}

func (t *CancelToken) Context() context.Context {
	return t.ctx
}

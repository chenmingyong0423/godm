// Generated by [optioner] command-line tool; DO NOT EDIT
// If you have any questions, please create issues and submit contributions at:
// https://github.com/chenmingyong0423/go-optioner

package finder

type AfterOpContextOption[T any] func(*AfterOpContext[T])

func NewAfterOpContext[T any](opContext *OpContext, opts ...AfterOpContextOption[T]) *AfterOpContext[T] {
	afterOpContext := &AfterOpContext[T]{
		OpContext: opContext,
	}

	for _, opt := range opts {
		opt(afterOpContext)
	}

	return afterOpContext
}

func WithDoc[T any](doc *T) AfterOpContextOption[T] {
	return func(afterOpContext *AfterOpContext[T]) {
		afterOpContext.Doc = doc
	}
}

func WithDocs[T any](docs []*T) AfterOpContextOption[T] {
	return func(afterOpContext *AfterOpContext[T]) {
		afterOpContext.Docs = docs
	}
}

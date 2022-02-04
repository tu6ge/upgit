package main

import "os"

type Result[T any] struct {
	value T
	err   error
}

func (r Result[T]) Ok() bool {
	return r.err == nil
}

func FromGoRet[T any](in ...interface{}) Result[T] {
	if nil != in[1] {
		return Result[T]{
			err: in[1].(error),
		}
	}
	return Result[T]{
		value: in[0].(T),
	}
}

func (r Result[T]) ValueOrDefault(default_ T) T {
	if r.err == nil {
		return r.value
	}
	return default_
}

func (r Result[T]) ValueOrPanic() T {
	if r.err == nil {
		return r.value
	}
	panic(r.err)
}

func (r Result[T]) ValueOrExit() T {
	if r.err == nil {
		return r.value
	}
	abortErr(r.err)
	panic(r.err)
}

func abortErr(err error) {
	if err != nil {
		GVerbose.Error("abort: " + err.Error())
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func panicErr(err error) {
	if err != nil {
		GVerbose.Error("panic: " + err.Error())
		panic(err)
	}
}

func panicErrWithoutLog(err error) {
	if err != nil {
		panic(err)
	}
}

package logger

import "go.uber.org/zap"

var (
	Binary     = zap.Binary
	Bool       = zap.Bool
	ByteString = zap.ByteString
	Float64    = zap.Float64
	Float32    = zap.Float32
	Int        = zap.Int
	Int64      = zap.Int64
	Int32      = zap.Int32
	String     = zap.String
	Stringp    = zap.Stringp
	Stack      = zap.Stack
	StackSkip  = zap.StackSkip
	Durationp  = zap.Durationp
	Any        = zap.Any
)

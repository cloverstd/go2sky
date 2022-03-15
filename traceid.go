package go2sky

import "github.com/SkyAPM/go2sky/internal/idgen"

func SetTraceIDFunc(f func() (string, error)) {
	idgen.Set(f)
}

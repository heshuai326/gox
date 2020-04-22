package runtime

import "github.com/gopub/gox"

type Namer interface {
	Name(srcName string) (dstName string)
}

type NameFunc func(srcName string) (dstName string)

func (f NameFunc) Name(srcName string) (dstName string) {
	return f(srcName)
}

func MapNamer(srcToDst map[string]string) Namer {
	return NameFunc(func(srcName string) (dstName string) {
		return srcToDst[srcName]
	})
}

var SnakeToCamelNamer NameFunc = func(snakeSrcName string) (camelDstName string) {
	return gox.SnakeToCamel(snakeSrcName)
}

var CamelToSnakeNamer NameFunc = func(camelSrcName string) (snakeDstName string) {
	return gox.CamelToSnake(camelSrcName)
}

var EqualNamer NameFunc = func(srcName string) (dstName string) {
	return srcName
}

var DefaultNamer Namer = EqualNamer

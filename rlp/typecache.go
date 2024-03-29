package rlp

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var (
	typeCacheMutex sync.RWMutex
	typeCache      = make(map[typekey]*typeinfo)
)

type typeinfo struct {
	decoder    decoder
	decoderErr error // error from makeDecoder
	writer     writer
	writerErr  error // error from makeWriter
}

// represents struct tags
type tags struct {
	// rlp:"nil" controls whether empty input results in a nil pointer.
	nilOK bool
	// rlp:"tail" controls whether this field swallows additional list
	// elements. It can only be set for the last field, which must be
	// of slice types.
	tail bool
	// rlp:"-" ignores fields.
	ignored bool
}

type typekey struct {
	reflect.Type
	// the key must include the struct tags because they
	// might generate a different decoder.
	tags
}

type decoder func(*Stream, reflect.Value) error

type writer func(reflect.Value, *encbuf) error

func cachedDecoder(typ reflect.Type) (decoder, error) {
	info := cachedTypeInfo(typ, tags{})
	return info.decoder, info.decoderErr
}

func cachedWriter(typ reflect.Type) (writer, error) {
	info := cachedTypeInfo(typ, tags{})
	return info.writer, info.writerErr
}

func cachedTypeInfo(typ reflect.Type, tags tags) *typeinfo {
	typeCacheMutex.RLock()
	info := typeCache[typekey{typ, tags}]
	typeCacheMutex.RUnlock()
	if info != nil {
		return info
	}
	// not in the cache, need to generate info for this types.
	typeCacheMutex.Lock()
	defer typeCacheMutex.Unlock()
	return cachedTypeInfo1(typ, tags)
}

func cachedTypeInfo1(typ reflect.Type, tags tags) *typeinfo {
	key := typekey{typ, tags}
	info := typeCache[key]
	if info != nil {
		// another goroutine got the write lock first
		return info
	}
	// put a dummy value into the cache before generating.
	// if the generator tries to lookup itself, it will get
	// the dummy value and won't call itself recursively.
	info = new(typeinfo)
	typeCache[key] = info
	info.generate(typ, tags)
	return info
}

type field struct {
	index int
	info  *typeinfo
}

func structFields(typ reflect.Type) (fields []field, err error) {
	lastPublic := lastPublicField(typ)
	for i := 0; i < typ.NumField(); i++ {
		if f := typ.Field(i); f.PkgPath == "" { // exported
			tags, err := parseStructTag(typ, i, lastPublic)
			if err != nil {
				return nil, err
			}
			if tags.ignored {
				continue
			}
			info := cachedTypeInfo1(f.Type, tags)
			fields = append(fields, field{i, info})
		}
	}
	return fields, nil
}

func parseStructTag(typ reflect.Type, fi, lastPublic int) (tags, error) {
	f := typ.Field(fi)
	var ts tags
	for _, t := range strings.Split(f.Tag.Get("rlp"), ",") {
		switch t = strings.TrimSpace(t); t {
		case "":
		case "-":
			ts.ignored = true
		case "nil":
			ts.nilOK = true
		case "tail":
			ts.tail = true
			if fi != lastPublic {
				return ts, fmt.Errorf(`rlp: invalid struct tag "tail" for %v.%s (must be on last field)`, typ, f.Name)
			}
			if f.Type.Kind() != reflect.Slice {
				return ts, fmt.Errorf(`rlp: invalid struct tag "tail" for %v.%s (field types is not slice)`, typ, f.Name)
			}
		default:
			return ts, fmt.Errorf("rlp: unknown struct tag %q on %v.%s", t, typ, f.Name)
		}
	}
	return ts, nil
}

func lastPublicField(typ reflect.Type) int {
	last := 0
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).PkgPath == "" {
			last = i
		}
	}
	return last
}

func (i *typeinfo) generate(typ reflect.Type, tags tags) {
	i.decoder, i.decoderErr = makeDecoder(typ, tags)
	i.writer, i.writerErr = makeWriter(typ, tags)
}

func isUint(k reflect.Kind) bool {
	return k >= reflect.Uint && k <= reflect.Uintptr
}

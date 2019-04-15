package geos

import (
	"runtime"
	"unsafe"
)

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lgeos_c
#include <geos_c.h>
#include <stdlib.h>
*/
import "C"

type wktReader struct {
	c *C.GEOSWKTReader
}

type wktWriter struct {
	c *C.GEOSWKTWriter
}

func (r *wktReader) read(wkt string) *Geom {
	cs := C.CString(wkt)
	defer C.free(unsafe.Pointer(cs))

	c := C.GEOSWKTReader_read_r(ctxHandler, r.c, cs)
	return GenerateGEOM(c)
}

func (w *wktWriter) write(g *Geom) string {
	dims := C.GEOSGeom_getCoordinateDimension_r(ctxHandler, g.cGeom)
	C.GEOSWKTWriter_setOutputDimension_r(ctxHandler, w.c, dims)
	return C.GoString(C.GEOSWKTWriter_write_r(ctxHandler, w.c, g.cGeom))
}

func createWktReader() *wktReader {
	c := C.GEOSWKTReader_create_r(ctxHandler)
	if c == nil {
		return nil
	}

	r := &wktReader{c: c}

	// Instruct Go garbage collector to clean C memory blocks
	runtime.SetFinalizer(r, func(r *wktReader) {
		C.GEOSWKTReader_destroy_r(ctxHandler, r.c)
	})

	return r
}

func (r *wktReader) Destroy() {
	C.GEOSWKTReader_destroy_r(ctxHandler, r.c)
}

func createWktWriter() *wktWriter {
	c := C.GEOSWKTWriter_create_r(ctxHandler)
	if c == nil {
		return nil
	}

	w := &wktWriter{c: c}

	// Instruct Go garbage collector to clean C memory blocks
	runtime.SetFinalizer(w, func(w *wktWriter) {
		C.GEOSWKTWriter_destroy_r(ctxHandler, w.c)
	})

	return w
}

func (w *wktWriter) Destroy() {
	C.GEOSWKTWriter_destroy_r(ctxHandler, w.c)
}

//go:build !nomimalloc
// +build !nomimalloc

package dl

/*
#cgo pkg-config: mimalloc

#include <string.h>
#include <mimalloc.h>
#include <mimalloc-override.h>
*/
import "C"

func CMiStatsPrint() { C.mi_stats_print(nil) }

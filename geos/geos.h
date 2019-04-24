#include <stdlib.h>
#include <stdarg.h>
#include <stdio.h>

// Avoid using by accident non _r functions
#ifndef GEOS_USE_ONLY_R_API
#define GEOS_USE_ONLY_R_API
#endif

#include <geos_c.h>

void notice(const char *fmt, ...);

void log_and_exit(const char *fmt, ...);

GEOSContextHandle_t ctx;

GEOSContextHandle_t init_geos();

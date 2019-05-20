#include <liblwgeom.h>
#include <geos_c.h>

#ifndef GEOM_H
#define GEOM_H

GEOSGeometry *from_wkt(char *wkt);
char *to_wkt(GEOSGeometry *g);
GEOSGeometry *simplified_buffer(GEOSGeometry *g, double width, double tolerance);
char *simplified_buffer_from_wkt(char *inwkt, double width, double tolerance);

LWGEOM* lwgeom_line_substring(LWGEOM *ingeom, double from, double to);
LWGEOM* buffer(LWGEOM* lwg, double width, int quadsegs);
LWGEOM* buffer_with_params(LWGEOM* lwg, double width, GEOSBufferParams *buffer_params);
LWGEOM* closest_point(LWGEOM *lwg1, LWGEOM *lwg2);
LWGEOM* split(LWGEOM *lwg_in, LWGEOM *blade);
LWGEOM* get_subgeom(LWGEOM *lwg, int index);
double line_locate_point(LWGEOM *linestring, LWGEOM *point);
LWGEOM * geos_union(LWGEOM *lwg1, LWGEOM *lwg2);

#endif

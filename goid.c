#include <runtime.h>

void ·ID(int32 id) {
	id = g->goid;
	USED(&id);
}

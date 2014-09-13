#include <runtime.h>

void ·GoID(int64 id) {
	id = g->goid;
	USED(&id);
}

#/bin/sh

set -ex

if [ ! -z "$1" ]; then GOROOT=$1; fi

cp -f $GOROOT/src/runtime/go_tls.h go_tls.h
WORK=$(go build -a -x -work runtime 2>&1|sed -e 's/^WORK=//g; t; d;')
cp $WORK/runtime/_obj/go_asm.h ./
rm -rf $WORK


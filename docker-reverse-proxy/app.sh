#!/bin/bash
echo "pwd: `pwd`"
echo "\$0: $0"
echo "basename: `basename $0`"
echo "dirname: `dirname $0`"
echo "dirname/readlink: $(dirname $(readlink -f $0))"
/usr/sbin/nginx
./bin/server1&
# ./bin/server2
node server2/app.js
#!/bin/bash

TAPES="add edit move toggle delete clean"
for TAPE in $TAPES
do
    echo "processing $TAPE ..."
    vhs ${TAPE}.tape -q
done

echo "finished"

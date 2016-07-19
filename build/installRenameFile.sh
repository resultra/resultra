#!/bin/bash
# This is a helper script to copy a file while creating any parent directories as necessary
cp $1 /tmp/$3 && mkdir -p $2 && cp /tmp/$3 $2 && rm -f /tmp/$3
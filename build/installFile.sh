#!/bin/bash
# This is a helper script to copy a file while creating any parent directories as necessary
mkdir -p $2 && cp $1 $2
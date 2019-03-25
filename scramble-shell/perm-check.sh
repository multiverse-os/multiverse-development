#!/bin/sh

echo "$1 permissions are 0$(stat -c %a $1)"

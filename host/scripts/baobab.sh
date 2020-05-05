#!/bin/sh
###############################################################################

echo "Multiverse OS: Analyzing file tree for large files..."

# A very simple 1 link implementation of baobab like functionality 
# (i.e. finding the biggest files in the fs) using only tree+grep
tree / -lah | grep 'G]'

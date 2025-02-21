#! /bin/bash

subdir="$1"

mkdir "${subdir}"
while read x; do
  mkdir "${subdir}/${x}"
done <dirs
while read x; do
  touch "${subdir}/${x}"
done <fils

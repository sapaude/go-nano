#!/bin/bash
BASE=${1:-HEAD~1}

git diff --name-only "$BASE" HEAD \
  | grep -v '^\.github' \
  | while read -r f; do
      dir=$(dirname "$f")
      while [ "$dir" != "." ]; do
        if [ -f "$dir/go.mod" ]; then echo "$dir"; break; fi
        dir=$(dirname "$dir")
      done
    done | sort -u

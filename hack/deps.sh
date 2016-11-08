#!/bin/bash

set -o errexit

echo "Updating deps"
rm -f glide.lock
rm -f glide.yaml
glide create --non-interactive
glide up --update-vendored

echo "Add vendored files to git"
find ./vendor/ -type f -not -path "*/.git*" -exec git add {} \;

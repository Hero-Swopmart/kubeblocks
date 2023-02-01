#!/bin/bash
# Copyright ApeCloud, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Easy & Dumb header check for CI jobs, currently checks ".go" files only.
#
# This will be called by the CI system (with no args) to perform checking and
# fail the job if headers are not correctly set. It can also be called with the
# 'fix' argument to automatically add headers to the missing files.
#
# Check if headers are fine:
#   $ ./hack/header-check.sh
# Check and fix headers:
#   $ ./hack/header-check.sh fix

set -e -o pipefail

# Initialize vars
ERR=false
FAIL=false

for file in $(git ls-files | grep '\.cue\|\.go$' | grep -v vendor/); do
  echo -n "Header check: $file... "
  if [[ -z $(cat ${file} | grep "Copyright ApeCloud, Inc.") ]]; then
      ERR=true
  fi
  if [ $ERR == true ]; then
    if [[ $# -gt 0 && $1 =~ [[:upper:]fix] ]]; then
      ext="${file##*.}"
      cat ./hack/boilerplate."${ext}".txt "${file}" > "${file}".new
      mv "${file}".new "${file}"
      echo "$(tput -T xterm setaf 3)FIXING$(tput -T xterm sgr0)"
      ERR=false
    else
      echo "$(tput -T xterm setaf 1)FAIL$(tput -T xterm sgr0)"
      ERR=false
      FAIL=true
    fi
  else
    echo "$(tput -T xterm setaf 2)OK$(tput -T xterm sgr0)"
  fi
done

# If we failed one check, return 1
[ $FAIL == true ] && exit 1 || exit 0
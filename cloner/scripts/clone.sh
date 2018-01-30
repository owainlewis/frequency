#!/bin/sh

set -euo pipefail
set -x

git clone "${GIT_REPO}" "${GIT_WORKSPACE}"

cd ${GIT_WORKSPACE}

git fetch origin "${GIT_REVISION}"
git checkout -qf FETCH_HEAD
git reset --hard -q "${GIT_REVISION}"

#!/usr/bin/env bash
# Copyright 2019 Cisco Systems
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Strips comments etc. from protobufs to compare the raw content.
# ./compare.sh xr/telemetry.proto nx/telemetry_bis.proto
function strip_comments() {
    grep -vE '^\s*(/|#|$|\*)' $1
}

diff -s --ignore-all-space <(strip_comments $1)  <(strip_comments $2)

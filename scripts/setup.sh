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
COLOR_RED='\033[0;91m'
COLOR_GREEN='\033[0;92m'
COLOR_CYAN='\033[0;96m'
COLOR_DEFAULT='\033[0m'
echo -e "${COLOR_CYAN}Ensuring Python is installed...${COLOR_DEFAULT}"
if ! python --version; then
    echo -e "${COLOR_RED}Please install Python before usage!${COLOR_DEFAULT}"
    exit 1
fi
echo -e "${COLOR_CYAN}Ensuring pipenv is installed...${COLOR_DEFAULT}"
if ! pipenv --version &> /dev/null; then
    python -m ensurepip
    pip install pipenv
fi
if ! pipenv --version; then
    echo -e "${COLOR_RED}Please ensure Python packages are reachable in your PATH!${COLOR_DEFAULT}"
    exit 2
fi
echo -e "${COLOR_CYAN}Installing scripts/ dependencies...${COLOR_DEFAULT}"
pipenv --three install
echo -e "${COLOR_GREEN}Ready to go!${COLOR_DEFAULT}"

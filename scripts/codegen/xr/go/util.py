"""Copyright 2019 Cisco Systems

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

import os
import re
import json
import glob
import logging
import subprocess

def camel_case(snake_str):
    """Modelled on protoc-gen-go/generator/generator.go Remember to
    catch the first one too. Logical OR of regexp might have been a
    little neater.
    """
    capital_snake = re.sub("^([a-z])", lambda m: m.group(1).upper(), snake_str)
    camel_onepass = re.sub("_([a-z])", lambda m: m.group(1).upper(), capital_snake)
    camel_twopass = re.sub("([0-9][a-z])", lambda m: m.group(1).upper(), camel_onepass)
    return camel_twopass

def parse_yang2proto_files(yang2proto_file_glob):
    """Parses a glob specifying yang2proto files to load as JSON."""
    yang2proto = []
    for yang2proto_filename in list(glob.glob(yang2proto_file_glob)):
        with open(yang2proto_filename) as yang2proto_fd:
            yang2proto.extend(json.load(yang2proto_fd))
    return yang2proto

def compile_proto(root_proto, root_codegen, proto,
        command_template = 'protoc --proto_path={root_proto} --go_out=plugins=grpc:{root_codegen} {root_proto}/{proto}'
    ):
    command = command_template.format(
        root_proto=root_proto,
        root_codegen=root_codegen,
        proto=proto
    )
    logging.debug('Executing: %s', command)
    return subprocess.check_output(
        command,
        shell=True,
        stderr=subprocess.STDOUT
    )

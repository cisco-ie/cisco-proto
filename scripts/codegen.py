#!/usr/bin/env python
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

"""Entrypoint for codegen related tasks.
Accepts an OS and/or language designation.
Not specifying an OS or language will default to all
permutations of OS/languages.
python codegen.py -h
"""
import logging
import argparse
import json
import time

from codegen import start_codegen

def load_config(config_file='codegen.json'):
    config = None
    with open(config_file, 'r') as config_fd:
        config = json.load(config_fd)
    logging.basicConfig(
        level=logging.DEBUG if config['debug'] else logging.INFO,
        format='%(levelname)s ~ %(asctime)s | %(message)s',
        datefmt='%H:%M:%S'
    )
    return config

def parse_args():
    parser = argparse.ArgumentParser(
        description='Codegen for Cisco protobufs'
    )
    parser.add_argument('--os',
        help='OS to codegen for',
        choices=['NX', 'XE', 'XR'],
        nargs='*',
        dest='os'
    )
    parser.add_argument('--language',
        help='target codegen language',
        choices=['Go'],
        nargs='*',
        dest='language'
    )
    return parser.parse_args()

def main():
    config = load_config()
    args = parse_args()
    start_time = time.time()
    start_codegen(args.os, args.language, config)
    logging.info('Took %i seconds.', time.time() - start_time)

if __name__ == '__main__':
    main()

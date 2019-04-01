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

"""Ensures the license header is applied to codegen files."""

import os
import argparse
import logging

__copyright = "Copyright 2019 Cisco Systems"

__license_str = """
{copyright}

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
""".format(copyright=__copyright).strip()

license_python = """\"\"\"
{license_str}
\"\"\"
""".format(license_str=__license_str).strip()

license_go = """/*
{license_str}
*/
""".format(license_str=__license_str).strip()

extension_license_map = {
    '.py': license_python,
    '.go': license_go
}

def check_license(file_path):
    license_present = False
    file_license = get_file_license(file_path)
    head_content = ''
    with open(file_path, 'r') as file_fd:
        head_content = file_fd.read(len(file_license))
    if file_license != head_content:
        if check_copyright(head_content):
            logging.error('License does not match, but copyright is present in %s', file_path)
            license_present = True
    else:
        license_present = True
    return license_present

def check_copyright(content):
    copyright_in_content = False
    if __copyright in content:
        logging.error('Expected copyright in content!')
        copyright_in_content = True
    elif 'copyright' in content.lower():
        logging.error('A copyright in content, uncertain which one!')
        copyright_in_content = True
    return copyright_in_content

def prepend_license(file_path):
    if not os.path.isfile(file_path):
        raise FileNotFoundError('%s does not exist!' % file_path)
    og_content = ''
    with open(file_path, 'r') as file_fd:
        og_content = file_fd.read()
    licensed_content = '{license}\n\n{content}'.format(
        license=get_file_license(file_path),
        content=og_content
    )
    with open(file_path, 'w') as file_fd:
        file_fd.write(licensed_content)

def get_file_license(file_path):
    _, extension = os.path.splitext(file_path)
    if extension not in extension_license_map.keys():
        raise Exception('%s is not supported!' % extension)
    return extension_license_map[extension]

def ensure_license(file_path):
    if not check_license(file_path):
        logging.debug('Adding license to %s', file_path)
        prepend_license(file_path)
        return True
    else:
        logging.debug('License already present in %s', file_path)
        return False

def parse_args():
    parser = argparse.ArgumentParser(
        description='License checker for codegen'
    )
    parser.add_argument('--base_path',
        help='Base directory path to recurse checking from.',
        dest='base_path',
        default='../codegen/'
    )
    return parser.parse_args()

def main():
    logging.basicConfig(level=logging.INFO)
    args = parse_args()
    if not os.path.isdir(args.base_path):
        logging.error('--base_path must be a directory!')
        exit(1)
    logging.info('Enumerating files with %s...', str(extension_license_map.keys()))
    file_paths = (
       os.path.join(parent, name)
       for (parent, subdirs, files) in os.walk(args.base_path)
       for name in files + subdirs
       if os.path.splitext(name)[1] in extension_license_map.keys()
    )
    logging.info('Checking/adding licenses...')
    licenses_added = 0
    total_files = 0
    for file_path in file_paths:
        total_files += 1
        if ensure_license(file_path):
            licenses_added += 1
    logging.info('Added license to %i/%i of files.', licenses_added, total_files)


if __name__ == '__main__':
    main()

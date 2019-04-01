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

import logging

from .codegen import Codegen
from .xr import Go as XRGo
from .xe import Go as XEGo
from .nx import Go as NXGo


os_lang_map = {
    'XR': {
        'Go': XRGo
    },
    'XE': {
        'Go': XEGo
    },
    'NX': {
        'Go': NXGo
    }
}

def get_codegen_class(os_name, language):
    if os_name not in os_lang_map.keys():
        logging.error('%s is not a supported OS!', os_name)
        return None
    if language not in os_lang_map[os_name].keys():
        logging.error('%s is not a supported language for %s!', language, os_name)
        return None
    return os_lang_map[os_name][language]

def start_codegen(os_list, language_list, config):
    if os_list:
        logging.info('Preparing codegen for %s.', ', '.join(os_list))
        if not set(os_list).issubset(set(os_lang_map.keys())):
            logging.error('OS list contains invalid entries!')
            return
    else:
        logging.info('Preparing codegen for all supported OSes.')
        os_list = os_lang_map.keys()
    if not language_list:
        logging.info('All supported languages will be generated.')
    for _os in os_list:
        languages = language_list if language_list else os_lang_map[_os].keys()
        for language in languages:
            gen_target = get_codegen_class(_os, language)
            if gen_target:
                logging.info('Starting %s generation for %s.', language, _os)
                gen_target(config).generate()

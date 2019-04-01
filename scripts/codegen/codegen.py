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


class Codegen:

    def __init__(self, config):
        self.config = config
        self.parse_config()
    
    def parse_config(self):
        raise NotImplementedError('Must be implemented by concrete class!')

    def generate(self):
        raise NotImplementedError('Must be implemented by concrete class!')

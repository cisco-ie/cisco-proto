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

"""Uses protoc to compile IOS XE proto files into usable Go code to codegen/go/xe"""

import os
import subprocess
import logging
from codegen import Codegen

class Go(Codegen):

    def parse_config(self):
        self.root = os.path.abspath(self.config['root'])
        self.root_proto = os.path.join(self.root, 'proto/nx')
        self.root_codegen = os.path.join(self.root, 'codegen/go/nx/')
        self.baseURL = self.config['baseURL']

    def generate(self):
        if not os.path.isdir(self.root_proto):
            logging.error('Proto root %s does not exist!', self.root_proto)
            return
        if not os.path.isdir(self.root_codegen):
            logging.info('Codegen root %s does not exist! Creating...', self.root_codegen)
            try:
                os.makedirs(self.root_codegen, exist_ok=True)
            except:
                logging.exception('Could not create codegen root %s', self.root_codegen)
                return
        logging.info('Compiling NX-OS stubs...')
        self.generate_stubs(self.root_proto, self.root_codegen)

    def generate_stubs(self, root_proto, root_codegen):
        self.run_protoc(root_proto, root_codegen, 'nx_grpc.proto')
        logging.info('Stub %s generation finished!', 'nx_grpc.proto')
        nx_telemetry_proto_root = os.path.join(root_proto, 'nx-telemetry-proto/')
        self.run_protoc(nx_telemetry_proto_root, root_codegen, 'telemetry_bis.proto')
        logging.info('Stub %s generation finished!', 'telemetry_bis.proto')
        self.run_protoc(nx_telemetry_proto_root, root_codegen, 'mdt_dialout.proto')
        logging.info('Stub %s generation finished!', 'mdt_dialout.proto')
    
    def run_protoc(self, root_proto, root_codegen, proto):
        command_template = 'protoc --proto_path={root_proto} --go_out=plugins=grpc:{root_codegen} {root_proto}/{proto}'
        command = command_template.format(
            root_proto=root_proto,
            root_codegen=root_codegen,
            proto=proto
        )
        try:
            subprocess.check_output(
                command,
                shell=True,
                stderr=subprocess.STDOUT
            )
        except subprocess.CalledProcessError as e:
            logging.error('Error compiling %s with protoc! %s', proto, e.output.decode('utf-8'))

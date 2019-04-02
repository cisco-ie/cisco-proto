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

"""Derived from https://github.com/cisco/bigmuddy-network-telemetry-proto/blob/master/prep_golang.py
Thank you Chris Cassar :)
Parses/compiles IOS XR proto files and generates usable Go code into codegen/go/xr
"""
import os
import subprocess
import logging
import glob
import json
import jinja2
from codegen import Codegen
from . import util

class Go(Codegen):

    def parse_config(self):
        self.root = os.path.abspath(self.config['root'])
        self.root_proto = os.path.join(self.root, 'proto/xr/protos/')
        self.root_codegen = os.path.join(self.root, 'codegen/go/xr/')
        self.base_url = self.config['baseURL']
        self.package = self.config['codegen']['languages']['go']['package']
        self.main = self.config['codegen']['languages']['go']['main']
        self.map_file_glob = self.config['codegen']['languages']['go']['mapFileGlob']
        root_template = os.path.join(
            os.path.dirname(
                os.path.abspath(__file__)
            ),
            'template/'
        )
        self.template_engine = jinja2.Environment(
            loader=jinja2.FileSystemLoader(root_template)
        )

    def generate(self):
        """Iterates through each xr/<release> directory and compiles the associated
        stubs and messages from the proto files.
        """
        for release_dir in sorted(os.listdir(self.root_proto)):
            root_proto_release = os.path.join(self.root_proto, release_dir)
            if not os.path.isdir(root_proto_release):
                logging.error('Proto root %s does not exist!', root_proto_release)
                continue
            root_codegen_release = os.path.join(self.root_codegen, release_dir)
            if not os.path.isdir(root_codegen_release):
                logging.info('Codegen root %s does not exist! Creating...', root_codegen_release)
                try:
                    os.makedirs(root_codegen_release, exist_ok=True)
                except:
                    logging.exception('Could not create codegen root %s', root_codegen_release)
                    continue
            logging.info('Compiling IOS XR - %s stubs...', release_dir)
            self.generate_stubs(root_proto_release, root_codegen_release)
            logging.info('Compiling IOS XR - %s messages...', release_dir)
            self.generate_messages(root_proto_release, root_codegen_release)

    def generate_stubs(self, root_proto_release, root_codegen_release,
        protos=['mdt_grpc_dialin/mdt_grpc_dialin.proto', 'mdt_grpc_dialout/mdt_grpc_dialout.proto', 'telemetry.proto']
    ):
        """Compiles the stubs minimally required for using gRPC/MDT."""
        for proto in protos:
            try:
                util.compile_proto(root_proto_release, root_codegen_release, proto)
            except subprocess.CalledProcessError as e:
                error_msg = e.output.decode('utf-8').strip()
                logging.error('Error compiling %s with protoc!', proto)
                logging.error(error_msg)
    
    def generate_messages(self, root_proto_release, root_codegen_release,
        ignore_dirs=['mdt_grpc_dialin/', 'mdt_grpc_dialout/']
    ):
        """Parses the IOS XR supplied yang2proto metadata files and compiles all
        of the protos with protoc into codegen as well as generates code to make usage easier.
        """
        yang2proto = util.parse_yang2proto_files(
            os.path.join(root_proto_release, self.map_file_glob)
        )
        paths = []
        errors = {}
        def add_error(encoding_path, proto, msg):
            errors[encoding_path] = {'proto': proto, 'msg': msg}
        for entry in yang2proto:
            proto_file_path = os.path.join(root_proto_release, entry['file'])
            proto_filename = os.path.basename(proto_file_path)
            proto_root = os.path.dirname(proto_file_path)
            if not os.path.isfile(proto_file_path):
                logging.error('Proto file %s is missing!', proto_file_path)
                add_error(entry['encoding_path'], entry['file'], 'Proto file is missing!')
                continue
            codegen_file_path = os.path.join(root_codegen_release, entry['file'])
            codegen_root = os.path.dirname(codegen_file_path)
            if not os.path.isdir(codegen_root):
                try:
                    logging.debug('Codegen directory %s does not exist - creating.', codegen_root)
                    os.makedirs(os.path.dirname(codegen_root), exist_ok=True)
                except Exception as e:
                    logging.error('Error creating codegen directory %s!', codegen_root)
                    add_error(entry['encoding_path'], entry['file'], 'Codegen directory creation error!')
                    continue
            logging.debug('Compiling %s', proto_file_path)
            try:
                util.compile_proto(proto_root, codegen_root, proto_filename)
            except subprocess.CalledProcessError as e:
                error_msg = e.output.decode('utf-8').strip()
                logging.error('Error compiling %s with protoc!', entry['file'])
                logging.error(error_msg)
                add_error(entry['encoding_path'], entry['file'], error_msg)
                continue
            paths.append({
                'gather_path': entry['encoding_path'],
                'package': entry['package'].replace('.', '_'),
                'path': os.path.dirname(entry['file']),
                'content': util.camel_case(entry['message']),
                'keys': util.camel_case(entry['message_keys'])
            })
        if errors:
            errors_filename = os.path.join(root_codegen_release, 'protoc_errors.json')
            logging.info('Writing errors to %s', errors_filename)
            with open(errors_filename, 'w') as errors_fd:
                json.dump(errors, errors_fd, indent=4, sort_keys=True)
        import_prefix = '{}/{}'.format(
            self.base_url,
            os.path.relpath(root_codegen_release, self.root)
        )
        self.generate_go_code(root_codegen_release, import_prefix, paths)
        self.generate_go_test_code(root_codegen_release, import_prefix, paths)
        self.test_go_code(root_codegen_release, import_prefix)

    def generate_go_code(self, root_codegen_release, import_prefix, paths):
        """Generates Go code which helps translates encoded XPaths from MDT
        to the associated compiled protobufs. This can be vendored in to projects
        to ease usage.
        Jinja is used to template the code.
        """
        go_code_template = self.template_engine.get_template('mdt_telemetry.go.jinja2')
        go_code = go_code_template.render(package=self.package, import_prefix=import_prefix, paths=paths)
        main_filename = os.path.join(root_codegen_release, self.main)
        logging.info('Writing Go bindings/utility to %s', main_filename)
        with open(main_filename, 'w') as main_fd:
            main_fd.write(go_code)

    def generate_go_test_code(self, root_codegen_release, import_prefix, paths):
        """Creates a test file used for testing and benchmarking the generated code.
        Jinja is used to template the code.
        """
        go_test_code_template = self.template_engine.get_template('mdt_telemetry_test.go.jinja2')
        go_test_code = go_test_code_template.render(package=self.package, import_prefix=import_prefix, paths=paths)
        # Append _test. yielding filename_test.go
        main_components = self.main.split('.')
        main_components.insert(1, '_test.')
        main_test_filename = os.path.join(
            root_codegen_release,
            ''.join(main_components)
        )
        logging.info('Writing Go test to %s', main_test_filename)
        with open(main_test_filename, 'w') as main_test_fd:
            main_test_fd.write(go_test_code)
    
    def test_go_code(self, root_codegen_release, import_prefix):
        """Tests the generated code to deserialize messages."""
        # -gcflags="-e" outputs all errors instead of stopping with "too may errors"
        command = 'go test -gcflags="-e" -run=. -bench=. {import_prefix}'.format(
            import_prefix=import_prefix
        )
        logging.info('Testing generated Go code: %s', command)
        test_err_filename = os.path.join(root_codegen_release, 'test_failure.log')
        try:
            results = subprocess.check_output(
                command,
                shell=True,
                stderr=subprocess.STDOUT
            )
            logging.info('Go test passed!')
            logging.info(results.decode('utf-8').strip())
            if os.path.isfile(test_err_filename):
                os.remove(test_err_filename)
        except subprocess.CalledProcessError as e:
            error_msg = e.output.decode('utf-8').strip()
            logging.error('Go test failed! Writing output to %s', test_err_filename)
            with open(test_err_filename, 'w') as test_err_fd:
                test_err_fd.write(error_msg)

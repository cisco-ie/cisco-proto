# cisco-proto
> [Protocol Buffer (protobuf)](https://developers.google.com/protocol-buffers/docs/overview) files for Cisco networking operating systems.

This repository attempts to consolidate sources of protobufs for Cisco networking operating systems and provide generated code which may be vendored in to projects. If any issues or inconsistencies are found, please open an issue! The directory structure is evolving and open to suggestion.

## Operating Systems
* [IOS XE](proto/xe/)
* [IOS XR](proto/xr/)
* [NX-OS](proto/nx/)

## Languages
If a desired language is not present, please file an issue for support.
* [Go](codegen/go/)

## Usage
This repository utilizes submodules, thus will require some special consideration when cloning.

```bash
# For only generated code
git clone https://github.com/cisco-ie/cisco-proto

# To include all protos and sources
git clone --recursive https://github.com/cisco-ie/cisco-proto

cd cisco-proto
```

This repository comes with scripts to aid with code generation - namely [scripts/codegen.py](scripts/). This script does utilize some Python libraries thus it does require some setup. Python is being utilized as the scripting language given the relatively complex scripting operations for the codegen process which would make `bash`, etc. unwieldly.

There is a [Makefile](Makefile) provided for convience which can ensure dependencies are installed and perform basic code generation tasks on a per-OS basis.

__Makefile__
```bash
# First run only, setup!
make setup
# To build all OSes
make all
# Build per OS
make nx
make xe
make xr
# Ensure license is applied to codegen files
make license
```

For more control, it is possible to use `scripts/codegen.py` directly.

__Manual__
```bash
cd scripts/
pipenv shell
python codegen.py -h
exit
```

Compilation of protobufs is accomplished with `protoc`. If unfamiliar with this process, it is highly recommended to follow the [tutorials](https://developers.google.com/protocol-buffers/docs/tutorials) on the subject.

## Codegen
This repository contains generated code in `codegen/<language>` which may be vendored in to existing projects, as well as the scripts to go about the generation process in case modifications or updates need to be made.

This repository is open to pull requests and discussion around code generation and organization.

### Go
In order to compile for Go, this repository must either be published on GitHub with `scripts/codegen.json` field `baseURL` set appropriately, or within the `GOPATH` at the expected location e.g. `$GOPATH/src/github.com/cisco-ie/cisco-proto`.

Go, `protoc`, and `protoc-gen-go` must all be installed. If these are not installed, please refer to the [Protocol Buffer tutorials](https://developers.google.com/protocol-buffers/docs/tutorials).

## Licensing
This repository is licensed as [Apache License, Version 2.0](LICENSE).

## Related Projects
* [ios-xr/model-driven-telemetry](https://github.com/ios-xr/model-driven-telemetry)
* [cisco/bigmuddy-network-telemetry-proto](https://github.com/cisco/bigmuddy-network-telemetry-proto)
* [CiscoDevNet/nx-telemetry-proto](https://github.com/CiscoDevNet/nx-telemetry-proto)
* [openconfig/gnmi](https://github.com/openconfig/gnmi)
* [ios-xr/telegraf-plugin](https://github.com/ios-xr/telegraf-plugin)
* [cisco/bigmuddy-network-telemetry-pipeline](https://github.com/cisco/bigmuddy-network-telemetry-pipeline)

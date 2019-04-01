# Scripts
This directory contains utility scripts used to perform various operations against the protocol buffers in `/proto`.

## codegen.py
`codegen.py` is utilized to compile/generate code from `/proto` into `/codegen`.

```bash
$ python codegen.py -h
usage: codegen.py [-h] [--os [{NX,XE,XR} [{NX,XE,XR} ...]]]
                  [--language [{Go} [{Go} ...]]]

Codegen for Cisco protobufs

optional arguments:
  -h, --help            show this help message and exit
  --os [{NX,XE,XR} [{NX,XE,XR} ...]]
                        OS to codegen for
  --language [{Go} [{Go} ...]]
                        target codegen language
```

`codegen.py` utilizes the `codegen/` module which contains all of the logic for compiling the protobufs/generating code. This logic is implemented on a per-OS, per-language basis.

The `codegen` module is organized with OS modules such as `nx`, `xe`, and `xr`. Each OS module exposes supported language modules like `go` with a concrete language generation class `Go`. The code for implementing the codegen is left largely as an exercise for the developer aside from a basic `Codegen` class specified in `codegen/codegen.py` which concrete language implementations are expected to inherit from purely for common entrypoint.

The lack of shared code per OS/language is largely due to the potential for the sources of information to change which may require reconsideration on a per OS basis for how to generate the code as it is exposed today.

## compare-proto.sh
`compare-proto.sh` simply compares proto files, stripping comments in order to compare the actual proto content itself instead of potentially varying documentation.

```bash
$ ./compare-proto.sh ../proto/nx/nx-telemetry-proto/telemetry_bis.proto ../proto/xr/protos/62x/telemetry.proto
2c2
< option go_package = "telemetry_bis";
---
> package telemetry;
```

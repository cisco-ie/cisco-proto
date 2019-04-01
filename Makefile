.PHONY: all nx xe xr setup license

setup :
	pushd scripts && ./setup.sh && popd

nx :
	pushd scripts && pipenv run python codegen.py --os NX && popd

xe :
	pushd scripts && pipenv run python codegen.py --os XE && popd

xr :
	pushd scripts && pipenv run python codegen.py --os XR && popd

license :
	python scripts/ensure_license.py --base_path codegen/

all : nx xe xr license

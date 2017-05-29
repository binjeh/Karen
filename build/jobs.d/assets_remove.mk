.PHONY: assets_remove
assets_remove:
	test -d src/assets && rm -r src/assets || exit 0

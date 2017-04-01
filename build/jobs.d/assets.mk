assets:
	test -d src/assets || mkdir src/assets
	go-bindata -nomemcopy -nocompress -pkg assets -o src/assets/assets.go assets/

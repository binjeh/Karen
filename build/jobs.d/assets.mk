assets:
	test -d src/assets || mkdir src/assets
	go-bindata -pkg assets -o src/assets/assets.go assets/

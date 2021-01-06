.PHONY: all build run gotool clean help

BINARY="go-backend-bootstrap"
BUILDPATH="./build"
RPMBUILDPATH="${HOME}/rpmbuild"

rpmbuild:
# mkdir -pv ./build/rpmbuild/{BUILD,BUILDROOT,RPMS,SOURCES,SPECS,SRPMS}
# echo "%_topdir	./build/rpm_build" > ~/.rpmmacros && rpmdev-setuptree && echo "rpm building preparatory work --- Done."
# echo "%_topdir	%{getenv:HOME}/rpmbuild" > ~/.rpmmacros && rpmdev-setuptree && echo "rpm building preparatory work --- Done."
	echo "%_topdir	%{getenv:HOME}/rpmbuild" > ~/.rpmmacros && echo "%debug_package %{nil}" >> ~/.rpmmacros && rpmdev-setuptree && echo "rpm building preparatory work --- Done."
	rpmdev-setuptree && echo "rpm building preparatory work --- Done."
	cp ./resource/build/rpmbuild/hello-world-1.0.0.tar.gz ${RPMBUILDPATH}/SOURCES
	cp ./resource/build/rpmbuild/hello-world.spec ${RPMBUILDPATH}/SPECS
	rpmbuild -ba ${RPMBUILDPATH}/SPECS/hello-world.spec

	# Start to copy build target files to project build directory
	mkdir -pv ./build/rpmbuild/{RPMS,SRPMS}
	cp -r ${RPMBUILDPATH}/RPMS/* ./build/rpmbuild/RPMS
	cp -r ${RPMBUILDPATH}/SRPMS/* ./build/rpmbuild/SRPMS

	# Clean rpmbuild folder
	rm -rf ${RPMBUILDPATH}
# rpmbuild --define '_topdir ~/rpmbuild' -ba $HOME/rpmbuild/SPECS/hello-world.spec

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run main.go

check:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
# rm -rf ~/.rpmmacros
	rm -rf ${BUILDPATH}

help:
	@echo "make build - generate bin file" 
	@echo "make run - run main.go"
	@echo "make clean - clean bin file"
	@echo "make check - run Go 'fmt' and 'vet'"

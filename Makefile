

ALL_DIST:=
DIST_NAME:=snow-cmd
# define_release_pack: GOOS GOARCH ExecuteEndfix
define define_release_pack

ALL_DIST+=dist/${DIST_NAME}-${1}-${2}${3}

dist/${DIST_NAME}-${1}-${2}${3}:
	GO_ARCH=${2} GOOS=${1} go build -v -o dist/${DIST_NAME}-${1}-${2}${3} ./cmd

endef

$(eval $(call define_release_pack,linux,amd64))
$(eval $(call define_release_pack,linux,arm64))
$(eval $(call define_release_pack,linux,arm))

$(eval $(call define_release_pack,windows,arm64,.exe))
$(eval $(call define_release_pack,windows,amd64,.exe))

$(eval $(call define_release_pack,darwin,arm64))
$(eval $(call define_release_pack,darwin,amd64))

dist/config.yml:
	cp config/config.yml dist/config.yml

all: ${ALL_DIST} dist/config.yml
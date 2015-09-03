#!/bin/bash

set -e

[[ ! -z ${WORKSPACE} ]]    || WORKSPACE=$(cd $(dirname $0)/../..;pwd)

rm -Rf ${WORKSPACE}/client/iOSXCodeProject/

/Applications/Unity/Unity.app/Contents/MacOS/Unity \
	-quit \
	-batchmode \
	-executeMethod BuildScript.ExecuteViaCommandLine \
	-projectPath ${WORKSPACE}/client \
	-logFile ${WORKSPACE}/client/build.log \
	"buildTarget=${BUILD_TARGET}" \
	"bundleVersion=${BUNDLE_VERSION}" \
	"bundleIdentifier=${BUNDLE_IDENTIFIER}" \
	"productName=${PRODUCT_NAME}"

if [ ${BUILD_TARGET} = ios ]
then
	cd ${WORKSPACE}/client/iOSXCodeProject/

	xcodebuild -version
	xcodebuild clean
	xcodebuild -sdk iphoneos CODE_SIGN_IDENTITY="iPhone Distribution: ${CODE_SIGN_IDENTITY}"

	if [ ! -f build/Release-iphones/${PRODUCT_NAME}.app/ResourceRules.plist ]
	then
		cp "${WORKSPACE}/tools/build/ResourceRules.plist" build/Release-iphoneos/${PRODUCT_NAME}.app/
	fi

	xcrun -sdk iphoneos "PackageApplication" "build/Release-iphoneos/${PRODUCT_NAME}.app" \
		-o "${WORKSPACE}/client/build_${BUILD_NUMBER}.ipa" \
		--sign "${CODE_SIGN_IDENTITY}" \
		--embed "${WORKSPACE}/${PROVISIONING_PROFILE}"
fi

if [ ${BUILD_TARGET} = android ]
then
	mv ${WORKSPACE}/client/build.apk ${WORKSPACE}/client/build_${BUILD_NUMBER}.apk
fi

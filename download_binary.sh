#!/bin/bash

############################################################################################
#
# This script downloads terraform provider ec (elastic search) binary for current OS
# scripts/download_binary.sh                    - to download binary from release 0.1.0
# scripts/download_binary.sh {release_version}  - to download binary from provided release version
#
############################################################################################

# Will be used if no release version is provided
DEFAULT_RELEASE_VERSION="0.1.0"

RELEASE_VERSION=${1:-$DEFAULT_RELEASE_VERSION}
PROVIDER_NAME="terraform-provider-ec"
TERRAFORM_PLUGIN_DIRECTORY="$HOME/.terraform.d/plugins"

OS=$(uname -s | awk '{print tolower($0)}')
echo "Operation system: $OS"

BINARY_NAME="${PROVIDER_NAME}_${OS}_amd64"
BINARY_PATH="https://github.com/everonhq/terraform-provider-ec/releases/download/$RELEASE_VERSION/$BINARY_NAME"

echo "Downloading binary plugin from $BINARY_PATH to $TERRAFORM_PLUGIN_DIRECTORY ..."

mkdir -p $TERRAFORM_PLUGIN_DIRECTORY
wget -cO - $BINARY_PATH > $TERRAFORM_PLUGIN_DIRECTORY/$PROVIDER_NAME
chmod a+x $TERRAFORM_PLUGIN_DIRECTORY/$PROVIDER_NAME

echo "$PROVIDER_NAME downloaded successfully"
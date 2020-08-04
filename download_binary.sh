#!/bin/bash

RELEASE_VERSION="0.1.0"
PROVIDER_NAME="terraform-provider-ec"
TERRAFORM_PLUGIN_DIRECTORY="$HOME/.terraform.d/plugins"

OS=$(uname -a | cut -d " " -f1 | awk '{print tolower($0)}')
echo "Operation system: $OS"

BINARY_NAME="${PROVIDER_NAME}_${OS}_amd64"
BINARY_PATH="https://github.com/everonhq/terraform-provider-ec/releases/download/$RELEASE_VERSION/$BINARY_NAME"

echo "Downloading binary plugin from $BINARY_PATH to $TERRAFORM_PLUGIN_DIRECTORY ..."

mkdir $TERRAFORM_PLUGIN_DIRECTORY
wget -cO - $BINARY_PATH > $TERRAFORM_PLUGIN_DIRECTORY/$PROVIDER_NAME
chmod a+x $TERRAFORM_PLUGIN_DIRECTORY/$PROVIDER_NAME

echo "$PROVIDER_NAME downloaded successfully"
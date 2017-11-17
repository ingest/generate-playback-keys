# Playback Key Generator

This is a simple utility that helps users on multiple platforms generate the playback key pair to create and sign playback keys for the Ingest service.

## Download

Fetch the approriate binary from the GitHub Release or build it using Go.

## Usage

Once you've downloaded the tool, execute it and after a second or two, a key.pem and key.pub will be output. You should keep the key.pem file to generate your playback keys on your servers, and give the Ingest service the public key to validate that these tokens are coming from you.

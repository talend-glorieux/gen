# gen

[![Build Status](https://travis-ci.org/talend-glorieux/gen.svg?branch=master)](https://travis-ci.org/talend-glorieux/gen) [![Go Report Card](https://goreportcard.com/badge/github.com/talend-glorieux/gen)](https://goreportcard.com/report/github.com/talend-glorieux/gen)

## Install

Get the latest binaries.

## Usage

`gen` takes any number of flags and a file path as it's last argument.

`gen -foo --bar= template.txt`

For example this README was generated by using: `gen -title=gen -install -answer=42 example.md > README.md`

#!/bin/sh

export $(grep -v '^#' ../local.env | xargs)
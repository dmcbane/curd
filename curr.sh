#! /usr/bin/env bash

function curr() {
  D=$(curd "$@")
  cd "${D}"
}

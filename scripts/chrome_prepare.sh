#!/usr/bin/env bash
killall chrome
google-chrome --remote-debugging-port=9222 &
google-chrome --remote-debugging-port=9223 &
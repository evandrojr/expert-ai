#!/usr/bin/env bash
killall chromium
chromium --remote-debugging-port=9222 &
chromium --remote-debugging-port=9223 &
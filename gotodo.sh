#!/bin/bash
nginx -g "daemon on;"  # Start nginx as daemon
/gotodo/gotodo # Start our Go app

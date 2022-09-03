#!/bin/bash
./publish.sh
./deploy.sh
./configure-ingress.sh
./configure-prometheus.sh
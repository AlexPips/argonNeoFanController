#!/bin/bash

# Define the service file name
service_file="customFanController.service"

# Define the systemd service directory
service_dir="/etc/systemd/system"

# Stop the service
sudo systemctl stop $service_file

# Disable the service from starting on boot
sudo systemctl disable $service_file

# Remove the service file from the systemd service directory
sudo rm $service_dir/$service_file

# Reload the systemd daemon to recognize the service removal
sudo systemctl daemon-reload

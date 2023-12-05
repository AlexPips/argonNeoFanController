#!/bin/bash

# Edit serivce file
execScriptPath=$(pwd)/argonNeoFanController

# Replace the file line
sed -i "s|^ExecStart=.*|ExecStart=$execScriptPath|" customFanController.service

# Define the service file name
serviceFile="customFanController.service"

# Define the systemd service directory
serviceDir="/etc/systemd/system"

# Copy the service file to the systemd service directory
sudo cp $serviceFile $serviceDir

# Reload the systemd daemon to recognize the new service
sudo systemctl daemon-reload

# Enable the service to start on boot
sudo systemctl enable $serviceFile

# Start the service
sudo systemctl start $serviceFile
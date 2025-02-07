# Raspberry Pi 4 Fan Control Script

This project provides a custom fan control script for the Raspberry Pi 4, written in Go. It uses GPIO and I2C to manage fan speed based on system temperature. Follow the steps below to install and configure the script.

## Prerequisites

- Raspberry Pi 4 running Raspberry Pi OS (32-bit or 64-bit)
- Internet connection
- Basic familiarity with the terminal

## Installation Steps

### 1. Install Go 1.23

The script is written in Go, so you need to install Go 1.23 on your Raspberry Pi.

Download and install Go:

```bash
wget https://go.dev/dl/go1.23.linux-armv6l.tar.gz  # For 32-bit OS
# OR
wget https://go.dev/dl/go1.23.linux-arm64.tar.gz   # For 64-bit OS
```

Extract the tarball:

```bash
sudo tar -C /usr/local -xzf go1.23.linux-armv6l.tar.gz  # For 32-bit OS
# OR
sudo tar -C /usr/local -xzf go1.23.linux-arm64.tar.gz   # For 64-bit OS
```

Add Go to your PATH:

```bash
export PATH=$PATH:/usr/local/go/bin
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

Verify the installation:

```bash
go version
```

You should see `go version go1.23 linux/arm` or `go version go1.23 linux/arm64`.

### 2. Build the Script

Clone or download this repository to your Raspberry Pi.

Navigate to the project directory:

```bash
cd /path/to/your/project
```

Build the Go application:

```bash
go build
```

This will generate an executable file in the current directory.

### 3. Enable GPIO and I2C

The script requires GPIO and I2C to be enabled on your Raspberry Pi.

Open the Raspberry Pi configuration tool:

```bash
sudo raspi-config
```

Enable GPIO:
- Navigate to `Interfacing Options > GPIO` and enable it.

Enable I2C:
- Navigate to `Interfacing Options > I2C` and enable it.

Reboot your Raspberry Pi:

```bash
sudo reboot
```

### 4. Install the Service

The project includes an `installService.sh` script to install the fan control application as a systemd service.

Make the script executable:

```bash
chmod +x installService.sh
```

Run the installation script:

```bash
sudo ./installService.sh
```

This will:
- Copy the executable to `/usr/local/bin/`.
- Create a systemd service file.
- Enable and start the service.

Verify the service is running:

```bash
sudo systemctl status fan-control.service
```

## Usage

Once installed, the fan control script will run automatically at startup. You can manually control the service using the following commands:

Start the service:

```bash
sudo systemctl start fan-control.service
```

Stop the service:

```bash
sudo systemctl stop fan-control.service
```

Restart the service:

```bash
sudo systemctl restart fan-control.service
```

Check the service status:

```bash
sudo systemctl status fan-control.service
```

## Troubleshooting

### Missing `/dev/i2c-1`

If the script fails with an error like `Error: open /dev/i2c-1: no such file or directory`, ensure I2C is enabled in `raspi-config` and the `i2c-dev` kernel module is loaded:

```bash
sudo modprobe i2c-dev
```

### Permission Issues

If you encounter permission errors, ensure your user is in the `gpio` and `i2c` groups:

```bash
sudo usermod -aG gpio,i2c $USER
```

Log out and log back in for the changes to take effect.

## Uninstallation

To uninstall the fan control service:

Stop and disable the service:

```bash
sudo systemctl stop fan-control.service
sudo systemctl disable fan-control.service
```

Remove the service file:

```bash
sudo rm /etc/systemd/system/fan-control.service
```

Remove the executable:

```bash
sudo rm /usr/local/bin/fan-control
```

Reload systemd:

```bash
sudo systemctl daemon-reload
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Support

For issues or questions, please open an issue on the GitHub repository or contact the maintainer.


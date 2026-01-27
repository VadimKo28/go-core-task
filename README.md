# Update Cursor on Ubuntu 24.04

This guide covers the common install methods on Ubuntu. Use the
section that matches how you installed Cursor.

## 1) Identify the install method

Check where the `cursor` binary comes from:

```
which cursor
readlink -f "$(which cursor)"
```

Common locations:
- `/usr/bin/cursor` -> `.deb` package
- `/snap/bin/cursor` -> Snap
- an `.AppImage` path -> AppImage

You can also confirm with package tools:

```
dpkg -l | grep -i cursor
snap list | grep -i cursor
flatpak list | grep -i cursor
```

## 2) Update by method

### A) Installed from a `.deb` file (downloaded installer)

1. Download the latest `.deb` from the official Cursor site.
2. Install it (this upgrades the existing package):

```
sudo apt install ./cursor_*.deb
```

### B) Installed as an AppImage

1. Download the latest AppImage.
2. Replace the old file and ensure it is executable:

```
mkdir -p ~/Applications
mv ~/Downloads/Cursor-*.AppImage ~/Applications/Cursor.AppImage
chmod +x ~/Applications/Cursor.AppImage
```

### C) Installed via Snap

Find the snap name, then refresh:

```
snap list | grep -i cursor
sudo snap refresh <snap-name>
```

### D) Installed via Flatpak

Find the app id, then update:

```
flatpak list | grep -i cursor
flatpak update <app-id>
```

### E) Installed from an APT repository

If you added an APT repo for Cursor, update normally:

```
sudo apt update
sudo apt upgrade cursor
```

## 3) Verify the version

If the CLI is on your PATH:

```
cursor --version
```

Or check in the app: Help -> About.

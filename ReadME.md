# Portscanner

A simple, terminal-based portscanner.

## ⚙️ Dependencies to build

```
Go - https://go.dev
```

## 🔧 Installation for Linux-based OS

```bash
git clone https://github.com/raymiamis/portscanner.git
cd portscanner
go build -o portscanner
sudo cp portscanner /usr/local/bin/
```
or
```bash
go install https://github.com/raymiamis/portscanner
```

## ⌨️ Usage

```
portscanner -host [target host to be scanned] -start [starting port] -end [ending port] -timeout [timeout in ms]
```
or
```
portscanner -host [target host to be scanned]
```
for starting & ending ports defaulting in 1 - 1024.

```
portscanner
```
returns the main menu.
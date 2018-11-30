# consul-template-plugin-system-information
Plugin for consul-template to obtain any kind of system information like hostname, kernel version, etc.

## Installation

### Manual installation
```sh
git clone https://github.com/tears-of-noobs/consul-template-plugin-system-information.git
cd consul-template-plugin-system-information
make install
```

### Installation from package (for ArchLinux users)
```sh
git clone https://github.com/tears-of-noobs/consul-template-plugin-system-information.git
cd consul-template-plugin-system-information
git checkout pkgbuild
makepkg -Cod; PKGVER=$(cd $(pwd)/src/consul-template-plugin-system-information/ && make ver) makepkg -esd
pacman -U *.tar.xz
```

## Usage

### Get current hostname
```sh
{{ $system := "consul-template-plugin-system-information" }}
{{ $hostname := $system "hostname" }}
```

### Get first IP address from network interface
get IPv4 address
```sh
{{ $system := "consul-template-plugin-system-information" }}
{{ $ipv4 := $system "ipv4 eth0" }}
```
get IPv6 address
```sh
{{ $system := "consul-template-plugin-system-information" }}
{{ $ipv6 := $system "ipv6 eth0" }}
```

### Get current time in preferred format

```sh
{{ $system := "consul-template-plugin-system-information" }}
{{ $time := $system "time" "rfc3339"}}
{{ $time }}
```
will be rendered as 
```sh
2006-01-02T15:04:05Z07:00
```

Supported formats
- ansic       "Mon Jan _2 15:04:05 2006"
- unix_date   "Mon Jan _2 15:04:05 MST 2006"
- ruby_date   "Mon Jan 02 15:04:05 -0700 2006"
- rfc822      "02 Jan 06 15:04 MST"
- rfc822Z     "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
- rfc850      "Monday, 02-Jan-06 15:04:05 MST"
- rfc1123     "Mon, 02 Jan 2006 15:04:05 MST"
- rfc1123Z    "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
- rfc3339     "2006-01-02T15:04:05Z07:00"
- rfc3339Nano "2006-01-02T15:04:05.999999999Z07:00"

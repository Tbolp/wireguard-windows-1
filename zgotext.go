// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"ja": &dictionary{index: jaIndex, data: jaData},
		"sl": &dictionary{index: slIndex, data: slData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%.2f\u00a0GiB":                         21,
	"%.2f\u00a0KiB":                         19,
	"%.2f\u00a0MiB":                         20,
	"%.2f\u00a0TiB":                         22,
	"%d day(s)":                             13,
	"%d hour(s)":                            14,
	"%d minute(s)":                          15,
	"%d second(s)":                          16,
	"%d tunnels were unable to be removed.": 158,
	"%d year(s)":                            12,
	"%d\u00a0B":                             18,
	"%s\n\nPlease consult the log for more information.": 111,
	"%s (out of date)":                        112,
	"%s (unsigned build, no updates)":         163,
	"%s You cannot undo this action.":         154,
	"%s ago":                                  17,
	"%s received, %s sent":                    71,
	"%s: %q":                                  23,
	"&About WireGuard…":                       109,
	"&Activate":                               58,
	"&Block untunneled traffic (kill-switch)": 82,
	"&Configuration:":                         86,
	"&Copy":                                   102,
	"&Deactivate":                             57,
	"&Edit":                                   133,
	"&Import tunnel(s) from file…":            119,
	"&Manage tunnels…":                        118,
	"&Name:":                                  79,
	"&Public key:":                            80,
	"&Remove selected tunnel(s)":              141,
	"&Save":                                   84,
	"&Save to file…":                          104,
	"&Toggle":                                 138,
	"(no argument): elevate and install manager service": 1,
	"(unknown)":                             81,
	"A name is required.":                   88,
	"A tunnel was unable to be removed: %s": 156,
	"About WireGuard":                       51,
	"Activating":                            97,
	"Active":                                96,
	"Add &empty tunnel…":                    134,
	"Add Tunnel":                            135,
	"Addresses:":                            62,
	"Addresses: %s":                         124,
	"Addresses: None":                       117,
	"All peers must have public keys":       44,
	"Allowed IPs:":                          65,
	"An Update is Available!":               129,
	"An interface must have a private key":  42,
	"An update to WireGuard is available. It is highly advisable to update without delay.":            166,
	"An update to WireGuard is now available. You are advised to update as soon as possible.":         131,
	"Another tunnel already exists with the name ‘%s’":                                                144,
	"Another tunnel already exists with the name ‘%s’.":                                               92,
	"App version: %s\nGo backend version: %s\nGo version: %s\nOperating system: %s\nArchitecture: %s": 53,
	"Are you sure you would like to delete %d tunnels?":                                               151,
	"Are you sure you would like to delete tunnel ‘%s’?":                                              153,
	"Brackets must contain an IPv6 address":                                                           28,
	"Cancel":                                                                                          85,
	"Close":                                                                                           54,
	"Command Line Options":                                                                            3,
	"Configuration Files (*.zip, *.conf)|*.zip;*.conf|All Files (*.*)|*.*":                            159,
	"Configuration ZIP Files (*.zip)|*.zip":                                                           161,
	"Could not enumerate existing tunnels: %v":                                                        143,
	"Could not import selected configuration: %v":                                                     142,
	"Create new tunnel":                                                                               77,
	"DNS servers:":                                                                                    63,
	"Deactivating":                                                                                    99,
	"Delete %d tunnels":                                                                               150,
	"Delete tunnel ‘%s’":                                                                              152,
	"E&xit":                                                                                           120,
	"Edit &selected tunnel…":                                                                          140,
	"Edit tunnel":                                                                                     78,
	"Endpoint:":                                                                                       66,
	"Error":                                                                                           0,
	"Error Exiting WireGuard":                                                                         164,
	"Error in getting configuration":                                                                  45,
	"Error: %v. Please try again.":                                                                    170,
	"Export all tunnels to &zip…":                                                                     139,
	"Export all tunnels to zip":                                                                       137,
	"Export log to file":                                                                              108,
	"Export tunnels to zip":                                                                           162,
	"Failed to activate tunnel":                                                                       73,
	"Failed to deactivate tunnel":                                                                     74,
	"Failed to determine tunnel state":                                                                72,
	"File ‘%s’ already exists.\n\nDo you want to overwrite it?":                                       95,
	"Import tunnel(s) from file":                                                                      160,
	"Imported %d of %d tunnels":                                                                       148,
	"Imported %d tunnels":                                                                             147,
	"Imported tunnels":                                                                                146,
	"Inactive":                                                                                        98,
	"Interface: %s":                                                                                   75,
	"Invalid IP address":                                                                              24,
	"Invalid MTU":                                                                                     29,
	"Invalid config key is missing an equals separator":                                               38,
	"Invalid endpoint host":                                                                           27,
	"Invalid key for [Interface] section":                                                             40,
	"Invalid key for [Peer] section":                                                                  41,
	"Invalid key for interface section":                                                               46,
	"Invalid key for peer section":                                                                    48,
	"Invalid key: %v":                                                                                 32,
	"Invalid name":                                                                                    87,
	"Invalid network prefix length":                                                                   25,
	"Invalid persistent keepalive":                                                                    31,
	"Invalid port":                                                                                    30,
	"Key must have a value":                                                                           39,
	"Keys must decode to exactly 32 bytes":                                                            33,
	"Latest handshake:":                                                                               68,
	"Line must occur in a section":                                                                    37,
	"Listen port:":                                                                                    60,
	"Log":                                                                                             101,
	"Log message":                                                                                     106,
	"MTU:":                                                                                            61,
	"Missing port from endpoint":                                                                      26,
	"Now":                                                                                             10,
	"Number must be a number between 0 and 2^64-1: %v":                                                34,
	"Peer":                                76,
	"Persistent keepalive:":               67,
	"Preshared key:":                      64,
	"Protocol version must be 1":          47,
	"Public key:":                         59,
	"Remove selected tunnel(s)":           136,
	"Select &all":                         103,
	"Status:":                             56,
	"Status: %s":                          123,
	"Status: Complete!":                   171,
	"Status: Unknown":                     116,
	"Status: Waiting for updater service": 169,
	"Status: Waiting for user":            167,
	"System clock wound backward!":        11,
	"Text Files (*.txt)|*.txt|All Files (*.*)|*.*": 107,
	"The %s tunnel has been activated.":            126,
	"The %s tunnel has been deactivated.":          128,
	"Time":                                         105,
	"Transfer:":                                    69,
	"Tunnel Error":                                 110,
	"Tunnel already exists":                        91,
	"Tunnel name is not valid":                     36,
	"Tunnel name ‘%s’ is invalid.":                 89,
	"Tunnels":                                      132,
	"Two commas in a row":                          35,
	"Unable to create new configuration":           93,
	"Unable to create tunnel":                      149,
	"Unable to delete tunnel":                      155,
	"Unable to delete tunnels":                     157,
	"Unable to determine whether the process is running under WOW64: %v":                          4,
	"Unable to exit service due to: %v. You may want to stop WireGuard from the service manager.": 165,
	"Unable to import configuration: %v":                                                          145,
	"Unable to list existing tunnels":                                                             90,
	"Unable to open current process token: %v":                                                    6,
	"Unable to wait for WireGuard window to appear: %v":                                           114,
	"Unknown state":    100,
	"Update Now":       168,
	"Usage: %s [\n%s]": 2,
	"When a configuration has exactly one peer, and that peer has an allowed IPs containing at least one of 0.0.0.0/0 or ::/0, then the tunnel service engages a firewall ruleset to block all traffic that is neither to nor from the tunnel interface, with special exceptions for DHCP and NDP.": 83,
	"WireGuard Activated":        125,
	"WireGuard Deactivated":      127,
	"WireGuard Detection Error":  113,
	"WireGuard Tunnel Error":     121,
	"WireGuard Update Available": 130,
	"WireGuard is running, but the UI is only accessible from desktops of the Builtin %s group.": 8,
	"WireGuard logo image": 52,
	"WireGuard may only be used by users who are a member of the Builtin %s group.": 7,
	"WireGuard system tray icon did not appear after 30 seconds.":                   9,
	"WireGuard: %s":          122,
	"WireGuard: Deactivated": 115,
	"Writing file failed":    94,
	"You must use the 64-bit version of WireGuard on this computer.": 5,
	"[EnumerationSeparator]": 49,
	"[UnitSeparator]":        50,
	"[none specified]":       43,
	"enabled":                70,
	"http2: Framer %p: failed to decode just-written frame": 172,
	"http2: Framer %p: read %v":                             174,
	"http2: Framer %p: wrote %v":                            173,
	"http2: decoded hpack field %+v":                        175,
	"♥ &Donate!":                                            55,
}

var enIndex = []uint32{ // 177 elements
	// Entry 0 - 1F
	0x00000000, 0x00000006, 0x00000039, 0x0000004f,
	0x00000064, 0x000000aa, 0x000000e9, 0x00000115,
	0x00000166, 0x000001c4, 0x00000200, 0x00000204,
	0x00000221, 0x00000241, 0x0000025f, 0x0000027f,
	0x000002a3, 0x000002c7, 0x000002d1, 0x000002da,
	0x000002e7, 0x000002f4, 0x00000301, 0x0000030e,
	0x0000031b, 0x0000032e, 0x0000034c, 0x00000367,
	0x0000037d, 0x000003a3, 0x000003af, 0x000003bc,
	// Entry 20 - 3F
	0x000003d9, 0x000003ec, 0x00000411, 0x00000445,
	0x00000459, 0x00000472, 0x0000048f, 0x000004c1,
	0x000004d7, 0x000004fb, 0x0000051a, 0x0000053f,
	0x00000550, 0x00000570, 0x0000058f, 0x000005b1,
	0x000005cc, 0x000005e9, 0x000005ec, 0x000005ef,
	0x000005ff, 0x00000614, 0x0000067f, 0x00000685,
	0x00000692, 0x0000069a, 0x000006a6, 0x000006b0,
	0x000006bc, 0x000006c9, 0x000006ce, 0x000006d9,
	// Entry 40 - 5F
	0x000006e6, 0x000006f5, 0x00000702, 0x0000070c,
	0x00000722, 0x00000734, 0x0000073e, 0x00000746,
	0x00000761, 0x00000782, 0x0000079c, 0x000007b8,
	0x000007c9, 0x000007ce, 0x000007e0, 0x000007ec,
	0x000007f3, 0x00000800, 0x0000080a, 0x00000832,
	0x00000950, 0x00000956, 0x0000095d, 0x0000096d,
	0x0000097a, 0x0000098e, 0x000009b2, 0x000009d2,
	0x000009e8, 0x00000a21, 0x00000a44, 0x00000a58,
	// Entry 60 - 7F
	0x00000a97, 0x00000a9e, 0x00000aa9, 0x00000ab2,
	0x00000abf, 0x00000acd, 0x00000ad1, 0x00000ad7,
	0x00000ae3, 0x00000af4, 0x00000af9, 0x00000b05,
	0x00000b32, 0x00000b45, 0x00000b59, 0x00000b66,
	0x00000b9a, 0x00000bae, 0x00000bc8, 0x00000bfd,
	0x00000c14, 0x00000c24, 0x00000c34, 0x00000c47,
	0x00000c66, 0x00000c6c, 0x00000c83, 0x00000c94,
	0x00000ca2, 0x00000cb3, 0x00000cc7, 0x00000cec,
	// Entry 80 - 9F
	0x00000d02, 0x00000d29, 0x00000d41, 0x00000d5c,
	0x00000db4, 0x00000dbc, 0x00000dc2, 0x00000dd7,
	0x00000de2, 0x00000dfc, 0x00000e16, 0x00000e1e,
	0x00000e3c, 0x00000e55, 0x00000e70, 0x00000e9f,
	0x00000ecb, 0x00000f03, 0x00000f29, 0x00000f3a,
	0x00000f70, 0x00000fb7, 0x00000fcf, 0x00001001,
	0x00001073, 0x0000108d, 0x000010c7, 0x000010ea,
	0x00001102, 0x0000112b, 0x00001144, 0x0000119d,
	// Entry A0 - BF
	0x000011e2, 0x000011fd, 0x00001223, 0x00001239,
	0x0000125c, 0x00001274, 0x000012d3, 0x00001328,
	0x00001341, 0x0000134c, 0x00001370, 0x00001390,
	0x000013a2, 0x000013db, 0x000013fc, 0x0000141c,
	0x0000143e,
} // Size: 720 bytes

const enData string = "" + // Size: 5182 bytes
	"\x02Error\x02(no argument): elevate and install manager service\x02Usage" +
	": %[1]s [\x0a%[2]s]\x02Command Line Options\x02Unable to determine wheth" +
	"er the process is running under WOW64: %[1]v\x02You must use the 64-bit " +
	"version of WireGuard on this computer.\x02Unable to open current process" +
	" token: %[1]v\x02WireGuard may only be used by users who are a member of" +
	" the Builtin %[1]s group.\x02WireGuard is running, but the UI is only ac" +
	"cessible from desktops of the Builtin %[1]s group.\x02WireGuard system t" +
	"ray icon did not appear after 30 seconds.\x02Now\x02System clock wound b" +
	"ackward!\x14\x01\x81\x01\x00\x02\x0b\x02%[1]d year\x00\x0c\x02%[1]d year" +
	"s\x14\x01\x81\x01\x00\x02\x0a\x02%[1]d day\x00\x0b\x02%[1]d days\x14\x01" +
	"\x81\x01\x00\x02\x0b\x02%[1]d hour\x00\x0c\x02%[1]d hours\x14\x01\x81" +
	"\x01\x00\x02\x0d\x02%[1]d minute\x00\x0e\x02%[1]d minutes\x14\x01\x81" +
	"\x01\x00\x02\x0d\x02%[1]d second\x00\x0e\x02%[1]d seconds\x02%[1]s ago" +
	"\x02%[1]d\u00a0B\x02%.2[1]f\u00a0KiB\x02%.2[1]f\u00a0MiB\x02%.2[1]f" +
	"\u00a0GiB\x02%.2[1]f\u00a0TiB\x02%[1]s: %[2]q\x02Invalid IP address\x02I" +
	"nvalid network prefix length\x02Missing port from endpoint\x02Invalid en" +
	"dpoint host\x02Brackets must contain an IPv6 address\x02Invalid MTU\x02I" +
	"nvalid port\x02Invalid persistent keepalive\x02Invalid key: %[1]v\x02Key" +
	"s must decode to exactly 32 bytes\x02Number must be a number between 0 a" +
	"nd 2^64-1: %[1]v\x02Two commas in a row\x02Tunnel name is not valid\x02L" +
	"ine must occur in a section\x02Invalid config key is missing an equals s" +
	"eparator\x02Key must have a value\x02Invalid key for [Interface] section" +
	"\x02Invalid key for [Peer] section\x02An interface must have a private k" +
	"ey\x02[none specified]\x02All peers must have public keys\x02Error in ge" +
	"tting configuration\x02Invalid key for interface section\x02Protocol ver" +
	"sion must be 1\x02Invalid key for peer section\x02, \x02, \x02About Wire" +
	"Guard\x02WireGuard logo image\x02App version: %[1]s\x0aGo backend versio" +
	"n: %[2]s\x0aGo version: %[3]s\x0aOperating system: %[4]s\x0aArchitecture" +
	": %[5]s\x02Close\x02♥ &Donate!\x02Status:\x02&Deactivate\x02&Activate" +
	"\x02Public key:\x02Listen port:\x02MTU:\x02Addresses:\x02DNS servers:" +
	"\x02Preshared key:\x02Allowed IPs:\x02Endpoint:\x02Persistent keepalive:" +
	"\x02Latest handshake:\x02Transfer:\x02enabled\x02%[1]s received, %[2]s s" +
	"ent\x02Failed to determine tunnel state\x02Failed to activate tunnel\x02" +
	"Failed to deactivate tunnel\x02Interface: %[1]s\x02Peer\x02Create new tu" +
	"nnel\x02Edit tunnel\x02&Name:\x02&Public key:\x02(unknown)\x02&Block unt" +
	"unneled traffic (kill-switch)\x02When a configuration has exactly one pe" +
	"er, and that peer has an allowed IPs containing at least one of 0.0.0.0/" +
	"0 or ::/0, then the tunnel service engages a firewall ruleset to block a" +
	"ll traffic that is neither to nor from the tunnel interface, with specia" +
	"l exceptions for DHCP and NDP.\x02&Save\x02Cancel\x02&Configuration:\x02" +
	"Invalid name\x02A name is required.\x02Tunnel name ‘%[1]s’ is invalid." +
	"\x02Unable to list existing tunnels\x02Tunnel already exists\x02Another " +
	"tunnel already exists with the name ‘%[1]s’.\x02Unable to create new con" +
	"figuration\x02Writing file failed\x02File ‘%[1]s’ already exists.\x0a" +
	"\x0aDo you want to overwrite it?\x02Active\x02Activating\x02Inactive\x02" +
	"Deactivating\x02Unknown state\x02Log\x02&Copy\x02Select &all\x02&Save to" +
	" file…\x02Time\x02Log message\x02Text Files (*.txt)|*.txt|All Files (*.*" +
	")|*.*\x02Export log to file\x02&About WireGuard…\x02Tunnel Error\x02%[1]" +
	"s\x0a\x0aPlease consult the log for more information.\x02%[1]s (out of d" +
	"ate)\x02WireGuard Detection Error\x02Unable to wait for WireGuard window" +
	" to appear: %[1]v\x02WireGuard: Deactivated\x02Status: Unknown\x02Addres" +
	"ses: None\x02&Manage tunnels…\x02&Import tunnel(s) from file…\x02E&xit" +
	"\x02WireGuard Tunnel Error\x02WireGuard: %[1]s\x02Status: %[1]s\x02Addre" +
	"sses: %[1]s\x02WireGuard Activated\x02The %[1]s tunnel has been activate" +
	"d.\x02WireGuard Deactivated\x02The %[1]s tunnel has been deactivated." +
	"\x02An Update is Available!\x02WireGuard Update Available\x02An update t" +
	"o WireGuard is now available. You are advised to update as soon as possi" +
	"ble.\x02Tunnels\x02&Edit\x02Add &empty tunnel…\x02Add Tunnel\x02Remove s" +
	"elected tunnel(s)\x02Export all tunnels to zip\x02&Toggle\x02Export all " +
	"tunnels to &zip…\x02Edit &selected tunnel…\x02&Remove selected tunnel(s)" +
	"\x02Could not import selected configuration: %[1]v\x02Could not enumerat" +
	"e existing tunnels: %[1]v\x02Another tunnel already exists with the name" +
	" ‘%[1]s’\x02Unable to import configuration: %[1]v\x02Imported tunnels" +
	"\x14\x01\x81\x01\x00\x02\x16\x02Imported %[1]d tunnel\x00\x17\x02Importe" +
	"d %[1]d tunnels\x14\x02\x80\x01\x02\x1f\x02Imported %[1]d of %[2]d tunne" +
	"l\x00 \x02Imported %[1]d of %[2]d tunnels\x02Unable to create tunnel\x14" +
	"\x01\x81\x01\x00\x02\x14\x02Delete %[1]d tunnel\x00\x15\x02Delete %[1]d " +
	"tunnels\x14\x01\x81\x01\x00\x024\x02Are you sure you would like to delet" +
	"e %[1]d tunnel?\x005\x02Are you sure you would like to delete %[1]d tunn" +
	"els?\x02Delete tunnel ‘%[1]s’\x02Are you sure you would like to delete t" +
	"unnel ‘%[1]s’?\x02%[1]s You cannot undo this action.\x02Unable to delete" +
	" tunnel\x02A tunnel was unable to be removed: %[1]s\x02Unable to delete " +
	"tunnels\x14\x01\x81\x01\x00\x02'\x02%[1]d tunnel was unable to be remove" +
	"d.\x00)\x02%[1]d tunnels were unable to be removed.\x02Configuration Fil" +
	"es (*.zip, *.conf)|*.zip;*.conf|All Files (*.*)|*.*\x02Import tunnel(s) " +
	"from file\x02Configuration ZIP Files (*.zip)|*.zip\x02Export tunnels to " +
	"zip\x02%[1]s (unsigned build, no updates)\x02Error Exiting WireGuard\x02" +
	"Unable to exit service due to: %[1]v. You may want to stop WireGuard fro" +
	"m the service manager.\x02An update to WireGuard is available. It is hig" +
	"hly advisable to update without delay.\x02Status: Waiting for user\x02Up" +
	"date Now\x02Status: Waiting for updater service\x02Error: %[1]v. Please " +
	"try again.\x02Status: Complete!\x02http2: Framer %[1]p: failed to decode" +
	" just-written frame\x02http2: Framer %[1]p: wrote %[2]v\x02http2: Framer" +
	" %[1]p: read %[2]v\x02http2: decoded hpack field %+[1]v"

var jaIndex = []uint32{ // 177 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000a, 0x00000067, 0x00000081,
	0x000000a6, 0x000000ef, 0x00000145, 0x00000180,
	0x000001d7, 0x0000025a, 0x000002af, 0x000002b6,
	0x000002de, 0x000002e8, 0x000002f2, 0x000002ff,
	0x00000309, 0x00000313, 0x0000031d, 0x00000325,
	0x00000332, 0x0000033f, 0x0000034c, 0x00000359,
	0x00000366, 0x00000380, 0x000003ba, 0x000003e8,
	0x00000410, 0x00000450, 0x00000465, 0x0000047e,
	// Entry 20 - 3F
	0x000004a9, 0x000004bd, 0x000004ef, 0x00000542,
	0x00000566, 0x00000585, 0x000005b0, 0x000005f4,
	0x00000625, 0x00000665, 0x000006a0, 0x000006d7,
	0x000006e6, 0x00000717, 0x00000736, 0x0000077c,
	0x000007bf, 0x000007f3, 0x000007f7, 0x000007f9,
	0x00000810, 0x00000827, 0x00000892, 0x0000089c,
	0x000008b8, 0x000008c0, 0x000008ce, 0x000008dc,
	0x000008e7, 0x000008fe, 0x00000903, 0x00000911,
	// Entry 40 - 5F
	0x00000920, 0x00000931, 0x0000093e, 0x00000955,
	0x00000975, 0x00000995, 0x0000099d, 0x000009a4,
	0x000009cc, 0x000009fd, 0x00000a2b, 0x00000a59,
	0x00000a79, 0x00000a80, 0x00000a9c, 0x00000ab2,
	0x00000abe, 0x00000acd, 0x00000ad6, 0x00000b2f,
	0x00000c63, 0x00000c6e, 0x00000c7e, 0x00000c8a,
	0x00000c9a, 0x00000cb3, 0x00000ce2, 0x00000d10,
	0x00000d38, 0x00000d87, 0x00000daf, 0x00000dd4,
	// Entry 60 - 7F
	0x00000e26, 0x00000e2d, 0x00000e3a, 0x00000e41,
	0x00000e4e, 0x00000e5e, 0x00000e65, 0x00000e73,
	0x00000e87, 0x00000ea4, 0x00000eab, 0x00000ec1,
	0x00000f0b, 0x00000f36, 0x00000f53, 0x00000f69,
	0x00000f9e, 0x00000fb1, 0x00000fcb, 0x0000100d,
	0x00001028, 0x00001037, 0x0000104c, 0x00001069,
	0x0000109e, 0x000010a9, 0x000010c9, 0x000010da,
	0x000010e8, 0x000010fc, 0x00001116, 0x00001148,
	// Entry 80 - 9F
	0x00001162, 0x00001194, 0x000011b3, 0x000011dc,
	0x00001244, 0x00001251, 0x0000125c, 0x0000127f,
	0x00001295, 0x000012b7, 0x000012eb, 0x000012fc,
	0x00001337, 0x00001360, 0x00001386, 0x000013ca,
	0x00001408, 0x00001454, 0x00001483, 0x000014a2,
	0x000014d3, 0x00001514, 0x00001539, 0x00001555,
	0x00001590, 0x000015b3, 0x000015ee, 0x0000161f,
	0x00001644, 0x00001679, 0x0000169e, 0x000016d2,
	// Entry A0 - BF
	0x00001725, 0x00001756, 0x00001780, 0x000017a8,
	0x000017de, 0x0000180e, 0x00001889, 0x000018ee,
	0x00001918, 0x00001928, 0x0000195b, 0x00001990,
	0x000019a2, 0x000019e8, 0x00001a1c, 0x00001a50,
	0x00001a90,
} // Size: 720 bytes

const jaData string = "" + // Size: 6800 bytes
	"\x02エラー\x02(引数なし): 管理者権限でmanagerサービスをインストールしてください\x02使い方: %[1]s [\x0a%[2" +
	"]s]\x02コマンドラインオプション\x02プロセスがWOW64で動作しているか確認できません: %[1]v\x02このコンピュータでは64ビ" +
	"ット版の WireGuard を使ってください。\x02現在のプロセストークンを開けません: %[1]v\x02WireGuard は組み込" +
	"みの %[1]s グループのメンバーだけが使えます。\x02WireGuard は実行中ですが、UI画面は組み込みの %[1]s グループの" +
	"デスクトップからしか開けません。\x02WireGuard システムトレイアイコンは30秒後に非表示になります。\x02現在\x02システム" +
	"時刻が遅れています\x02%[1]d 年\x02%[1]d 日\x02%[1]d 時間\x02%[1]d 分\x02%[1]d 秒\x02%" +
	"[1]s 前\x02%[1]d B\x02%.2[1]f\u00a0KiB\x02%.2[1]f\u00a0MiB\x02%.2[1]f" +
	"\u00a0GiB\x02%.2[1]f\u00a0TiB\x02%[1]s: %[2]q\x02不正な IP アドレス\x02不正なネットワー" +
	"クプレフィックス長指定\x02エンドポイントのポート指定なし\x02不正なエンドポイントホスト\x02カッコ内は IPv6 アドレスでなけれ" +
	"ばなりません\x02不正な MTU 指定\x02不正なポート番号\x02不正な持続的キープアライブ値\x02不正な鍵: %[1]v\x02鍵" +
	"は 32 バイトでなければなりません\x02数値は0から2の64乗-1の範囲内の値でなければなりません: %[1]v\x021行にカンマが2" +
	"つあります\x02トンネル名が不正です\x02行がセクション内にありません\x02イコール（=）が無いためキー項目として不正です\x02キー" +
	"項目に対応する値がありません\x02[Interface] セクションのキー項目として不正です\x02[Peer] セクションのキー項目とし" +
	"て不正です\x02インターフェースには秘密鍵が必須です\x02[指定なし]\x02すべてのピアには公開鍵が必須です\x02設定の取得時にエラ" +
	"ー\x02インターフェースセクションのキー項目が不正です\x02プロトコルバージョンは 1 でなければなりません\x02ピアセクションのキー" +
	"項目が不正です\x02、\x02 \x02WireGuard について\x02WireGuard ロゴ画像\x02App version: " +
	"%[1]s\x0aGo backend version: %[2]s\x0aGo version: %[3]s\x0aOperating sys" +
	"tem: %[4]s\x0aArchitecture: %[5]s\x02閉じる\x02♥ 寄付のお願い!(&D)\x02状態:\x02無効化(" +
	"&D)\x02有効化(&A)\x02公開鍵:\x02待受ポート番号:\x02MTU:\x02アドレス:\x02DNS サーバ:\x02事前共有鍵" +
	":\x02Allowed IPs:\x02エンドポイント:\x02持続的キープアライブ:\x02直近のハンドシェイク:\x02転送:\x02有効" +
	"\x02%[1]s 受信済み、%[1]s 送信済み\x02トンネルの状態取得に失敗しました\x02トンネルの有効化に失敗しました\x02トンネル" +
	"の無効化に失敗しました\x02インターフェース: %[1]s\x02ピア\x02トンネルの新規作成\x02トンネルの編集\x02名前(&N)" +
	":\x02公開鍵(&P):\x02(不明)\x02トンネルを通らないトラフィックのブロック（キルスイッチ）(&B)\x02ピアが1つだけ設定され" +
	"ていて、さらに Allowed IPs に 0.0.0.0/0 または ::/0 が含まれている場合、トンネルサービスはトンネルインターフェ" +
	"ースを通らないすべてのトラフィックをブロックするファイアウォールルールを追加します。\x02保存(&S)\x02キャンセル\x02設定(&C" +
	"):\x02不正な名前\x02名前は必須です。\x02トンネル名 ‘%[1]s’ は不正です。\x02既存のトンネルを表示できません\x02トン" +
	"ネルはすでに存在します\x02‘%[1]s’ という名前の別のトンネルがすでに存在します。\x02新しい設定が作成できません\x02ファイル" +
	"の書き込みに失敗\x02ファイル ‘%[1]s’ はすでに存在します。\x0a\x0a上書きしますか？\x02有効\x02有効化中\x02無" +
	"効\x02無効化中\x02不明な状態\x02ログ\x02コピー(&C)\x02すべて選択(&a)\x02ファイルに保存…(&S)\x02時刻" +
	"\x02ログメッセージ\x02テキストファイル (*.txt)|*.txt|すべてのファイル (*.*)|*.*\x02ログをファイルにエクスポ" +
	"ート\x02WireGuardについて…(&A)\x02トンネルエラー\x02%[1]s\x0a\x0a詳細はログを参照してください。" +
	"\x02%[1]s 期限切れ\x02WireGuard 検出エラー\x02WireGuard ウィンドウが表示できませんでした: %[1]v" +
	"\x02WireGuard: 無効化済み\x02状態: 不明\x02アドレス: なし\x02トンネルの管理…(&M)\x02トンネルをファイルか" +
	"らインポート(&I)\x02終了(&X)\x02WireGuard トンネルエラー\x02WireGuard: %[1]s\x02状態: %" +
	"[1]s\x02アドレス: %[1]s\x02WireGuard 有効化済み\x02トンネル %[1]s は有効化されました。\x02WireG" +
	"uard 無効化済み\x02トンネル %[1]s は無効化されました。\x02更新が利用できます！\x02WireGuard の更新が利用可能で" +
	"す\x02WireGuard の更新が利用可能になりました。できるだけ早く更新してください。\x02トンネル\x02編集(&E)\x02空の" +
	"トンネルを追加…(&e)\x02トンネルの追加\x02選択したトンネルの削除\x02すべてのトンネルをzipにエクスポート\x02切り替え(" +
	"&T)\x02すべてのトンネルをzipにエクスポート…(&z)\x02選択したトンネルの編集…(&s)\x02選択したトンネルの削除(&R)" +
	"\x02選択した設定をインポートできませんでした: %[1]v\x02既存のトンネルを表示できませんでした: %[1]v\x02‘%[1]s’ " +
	"という名前の別のトンネルがすでに存在します\x02設定をインポートできません: %[1]v\x02トンネルをインポート\x02%[1]d ト" +
	"ンネルをインポートしました\x02%[2]d つ中の %[1]d トンネルをインポートしました\x02トンネルを作成できません\x02%[1" +
	"]d トンネルを削除\x02本当に %[1]d つのトンネルを削除しますか？\x02トンネル ‘%[1]s’ を削除\x02本当にトンネル ‘%" +
	"[1]s’ を削除しますか？\x02%[1]s この操作はもとに戻せません。\x02トンネルを削除できません\x02トンネルを削除できませんでし" +
	"た: %[1]s\x02トンネルを削除できません\x02%[1]d トンネルは削除できませんでした\x02設定ファイル (*.zip, *." +
	"conf)|*.zip;*.conf|すべてのファイル (*.*)|*.*\x02ファイルからトンネルをインポート\x02ZIP形式設定ファイル" +
	" (*.zip)|*.zip\x02トンネルをZIPにエクスポート\x02%[1]s (未署名のビルド、更新の提供なし)\x02エラーのため W" +
	"ireGuard を終了します\x02%[1]v のためサービスを終了できません。サービスマネージャから WireGuard を停止できます。" +
	"\x02WireGuard の更新が利用可能です。速やかに更新することを強く推奨します。\x02状態: ユーザーからの応答待ち\x02今すぐ更新" +
	"\x02状態: アップデータサービスを待機中\x02エラー: %[1]v。再度実行してください。\x02状態: 完了！\x02http2: Fr" +
	"amer %[1]p: just-writtenフレームのデコードに失敗\x02http2: Framer %[1]p: %[2]v を書き込み" +
	"ました\x02http2: Framer %[1]p: %[2]v を読み込みました\x02http2: hpack フィールド %+[1]" +
	"v をデコードしました"

var slIndex = []uint32{ // 177 elements
	// Entry 0 - 1F
	0x00000000, 0x00000007, 0x00000058, 0x00000070,
	0x00000089, 0x000000c1, 0x00000107, 0x0000013e,
	0x00000190, 0x000001f6, 0x0000023a, 0x0000023f,
	0x0000025e, 0x00000296, 0x000002cd, 0x00000301,
	0x00000341, 0x00000385, 0x00000391, 0x0000039a,
	0x000003a7, 0x000003b4, 0x000003c1, 0x000003ce,
	0x000003db, 0x000003ee, 0x00000412, 0x00000434,
	0x0000045d, 0x00000483, 0x00000490, 0x0000049f,
	// Entry 20 - 3F
	0x000004c3, 0x000004da, 0x0000050b, 0x0000053f,
	0x00000554, 0x0000056b, 0x00000586, 0x000005be,
	0x000005d9, 0x000005fe, 0x0000061e, 0x00000640,
	0x0000064e, 0x00000675, 0x00000695, 0x000006b7,
	0x000006d5, 0x000006f7, 0x000006fa, 0x000006fc,
	0x00000709, 0x00000727, 0x0000079c, 0x000007a2,
	0x000007b0, 0x000007b8, 0x000007c5, 0x000007d0,
	0x000007de, 0x000007f1, 0x000007f6, 0x000007ff,
	// Entry 40 - 5F
	0x0000080f, 0x00000825, 0x00000836, 0x00000846,
	0x00000862, 0x00000874, 0x0000087c, 0x00000887,
	0x000008a4, 0x000008c8, 0x000008e6, 0x00000906,
	0x00000915, 0x0000091d, 0x0000092f, 0x0000093b,
	0x00000941, 0x00000950, 0x0000095a, 0x00000984,
	0x00000aa5, 0x00000aad, 0x00000ab7, 0x00000ac7,
	0x00000ad4, 0x00000ae4, 0x00000b06, 0x00000b36,
	0x00000b48, 0x00000b73, 0x00000b9a, 0x00000bb8,
	// Entry 60 - 7F
	0x00000bf3, 0x00000bfb, 0x00000c07, 0x00000c11,
	0x00000c1f, 0x00000c2e, 0x00000c36, 0x00000c3f,
	0x00000c4b, 0x00000c63, 0x00000c68, 0x00000c7e,
	0x00000cb6, 0x00000cd0, 0x00000ce3, 0x00000cf1,
	0x00000d20, 0x00000d36, 0x00000d53, 0x00000d8e,
	0x00000da5, 0x00000db4, 0x00000dc2, 0x00000dd9,
	0x00000df8, 0x00000dff, 0x00000e17, 0x00000e28,
	0x00000e36, 0x00000e45, 0x00000e59, 0x00000e77,
	// Entry 80 - 9F
	0x00000e8d, 0x00000ead, 0x00000ec6, 0x00000ee6,
	0x00000f2b, 0x00000f32, 0x00000f39, 0x00000f52,
	0x00000f5e, 0x00000f76, 0x00000f8e, 0x00000f98,
	0x00000fb6, 0x00000fcf, 0x00000fe8, 0x00001016,
	0x00001049, 0x0000106e, 0x00001094, 0x000010a4,
	0x00001108, 0x00001194, 0x000011b0, 0x00001215,
	0x00001302, 0x0000131d, 0x00001358, 0x00001383,
	0x0000139d, 0x000013c5, 0x000013e0, 0x00001494,
	// Entry A0 - BF
	0x000014e1, 0x000014fa, 0x00001525, 0x00001539,
	0x00001568, 0x00001588, 0x000015f5, 0x00001648,
	0x00001664, 0x00001672, 0x00001699, 0x000016bb,
	0x000016cd, 0x00001715, 0x00001739, 0x0000175d,
	0x00001782,
} // Size: 720 bytes

const slData string = "" + // Size: 6018 bytes
	"\x02Napaka\x02(brez argumenta): povzdigni na skrbniške pravice in namest" +
	"i skrbniško storitev\x02Uporaba: %[1]s [\x0a%[2]s]\x02Možnosti ukazne vr" +
	"stice\x02Napaka pri določanju ali proces teče kot WOW64: %[1]v\x02Na tem" +
	"u računalniku morate uporabiti 64-bitno različico WireGuarda.\x02Napaka " +
	"pri odpiranju žetona trenutnega procesa: %[1]v\x02WireGuard lahko uporab" +
	"ljajo samo uporabniki, ki so člani vgrajene skupine %[1]s.\x02WireGuard " +
	"je zagnan, vendar je up. vmesnik dostopen samo z namizij uporabnikov čla" +
	"nov skupine %[1]s.\x02Ikona WireGuarda se po 30 sekundah ni pojavila v s" +
	"istemski vrstici.\x02Zdaj\x02Sistemska ura prevrtena nazaj!\x14\x01\x81" +
	"\x01\x00\x04\x0b\x02%[1]d leta\x02\x0b\x02%[1]d leto\x03\x0b\x02%[1]d le" +
	"ti\x00\x0a\x02%[1]d let\x14\x01\x81\x01\x00\x04\x0a\x02%[1]d dni\x02\x0a" +
	"\x02%[1]d dan\x03\x0c\x02%[1]d dneva\x00\x0a\x02%[1]d dni\x14\x01\x81" +
	"\x01\x00\x04\x0a\x02%[1]d ure\x02\x0a\x02%[1]d uro\x03\x0a\x02%[1]d uri" +
	"\x00\x09\x02%[1]d ur\x14\x01\x81\x01\x00\x04\x0d\x02%[1]d minute\x02\x0d" +
	"\x02%[1]d minuto\x03\x0d\x02%[1]d minuti\x00\x0c\x02%[1]d minut\x14\x01" +
	"\x81\x01\x00\x04\x0e\x02%[1]d sekunde\x02\x0e\x02%[1]d sekundo\x03\x0e" +
	"\x02%[1]d sekundi\x00\x0d\x02%[1]d sekund\x02%[1]s nazaj\x02%[1]d\u00a0B" +
	"\x02%.2[1]f\u00a0KiB\x02%.2[1]f\u00a0MiB\x02%.2[1]f\u00a0GiB\x02%.2[1]f" +
	"\u00a0TiB\x02%[1]s: %[2]q\x02Napačen naslov IP\x02Napačna dolžina predpo" +
	"ne omrežja\x02Pri končni točki manjkajo vrata\x02Pri končni točki je gos" +
	"titelj napačen\x02Oklepaji morajo vsebovati naslov IPv6\x02Napačen MTU" +
	"\x02Napačna vrata\x02Napačno trajno ohranjanje povezave\x02Napačen ključ" +
	": %[1]v\x02Dekodirani ključi morajo biti natanko 32 bajtov\x02Številka m" +
	"ora biti število med 0 in 2^64-1: %[1]v\x02Dve zaporedni vejici\x02Ime t" +
	"unela ni veljavno\x02Vrstica mora biti v odseku\x02Napačnemu ključu konf" +
	"iguracije manjka ločilo-enačaj\x02Ključ mora imeti vrednost\x02Napačen k" +
	"ljuč za odsek [Interface]\x02Napačen ključ za odsek [Peer]\x02Vmesnik mo" +
	"ra imeti zasebni ključ\x02[ni navedeno]\x02Vsi vrstniki morajo imeti jav" +
	"ni ključ\x02Napaka pri branju konfiguracije\x02Napačen ključ za odsek vm" +
	"esnika\x02Verzija protokola mora biti 1\x02Napačen ključ za odsek vrstni" +
	"ka\x02, \x02 \x02O WireGuardu\x02Slika WireGuardovega logotipa\x02Verzij" +
	"a aplikacije: %[1]s\x0aVerzija wireguard-go: %[2]s\x0aVerzija Go: %[3]s" +
	"\x0aOperacijski sistem: %[4]s\x0aArhitektura: %[5]s\x02Zapri\x02♥ &Donir" +
	"aj!\x02Status:\x02&Deaktiviraj\x02&Aktiviraj\x02Javni ključ:\x02Vrata po" +
	"slušanja:\x02MTU:\x02Naslovi:\x02Strežniki DNS:\x02Ključ v skupni rabi:" +
	"\x02Dovoljeni IP-ji:\x02Končna točka:\x02Trajno ohranjanje povezave:\x02" +
	"Zadnje rokovanje:\x02Prenos:\x02omogočeno\x02%[1]s prejeto, %[2]s poslan" +
	"o\x02Napaka pri določanju stanja tunela\x02Napaka pri aktiviranju tunela" +
	"\x02Napaka pri deaktiviranju tunela\x02Vmesnik: %[1]s\x02Vrstnik\x02Ustv" +
	"ari nov tunel\x02Uredi tunel\x02&Ime:\x02&Javni ključ:\x02(neznano)\x02&" +
	"Blokiraj promet izven tunela (varovalka)\x02Kadar ima konfiguracija nata" +
	"nko enega vrstnika in njegov spisek dovoljenih IP-jev vsebuje vsaj enega" +
	" izmed 0.0.0.0/0 ali ::/0, bo storitev tunela vzpostavila pravila požarn" +
	"ega zidu, ki bodo blokirala ves promet, ki ni niti za niti iz vmesnika t" +
	"unela s posebnimi izjemami za DHCP and NDP.\x02&Shrani\x02Prekliči\x02&K" +
	"onfiguracija:\x02Napačno ime\x02Ime je obvezno.\x02Ime tunela »%[1]s« ni" +
	" veljavno.\x02Napaka pri pripravi seznama obstoječih tunelov\x02Tunel že" +
	" obstaja\x02Drug tunel z imenom »%[1]s« že obstaja.\x02Napaka pri izdela" +
	"vi nove konfiguracije\x02Napaka pri pisanju v datoteko\x02Datoteka »%[1]" +
	"s« že obstaja.\x0a\x0aAli jo želite prepisati?\x02Aktivno\x02Se aktivira" +
	"\x02Neaktivno\x02Se deaktivira\x02Neznano stanje\x02Dnevnik\x02&Kopiraj" +
	"\x02&Izberi vse\x02&Shrani v datoteko\u00a0…\x02Čas\x02Sporočilo v dnevn" +
	"iku\x02Tekstovne datoteke (*.txt)|*.txt|Vse datoteke (*.*)|*.*\x02Izvozi" +
	" dnevnik v datoteko\x02O WireGu&ardu\u00a0…\x02Napaka tunela\x02%[1]s" +
	"\x0a\x0aDodatne informacije najdete v dnevniku.\x02%[1]s (neposodobljen)" +
	"\x02Napaka zaznavanja WireGuarda\x02Čakanje, da se pojavi WireGuardovo o" +
	"kno, ni možno: %[1]v\x02WireGuard: Deaktiviran\x02Status: Neznan\x02Nasl" +
	"ovi: Brez\x02&Upravljaj tunele\u00a0…\x02Uvoz&i tunele iz datoteke\u00a0" +
	"…\x02I&zhod\x02Napaka tunela WireGuard\x02WireGuard: %[1]s\x02Status: " +
	"%[1]s\x02Naslovi: %[1]s\x02WireGuard aktiviran\x02Tunel %[1]s je bil akt" +
	"iviran.\x02WireGuard deaktiviran\x02Tunel %[1]s je bil deaktiviran.\x02N" +
	"a voljo je posodobitev!\x02Posodobitev WireGuarda na voljo\x02Posodobite" +
	"v WireGuarda je na voljo. Svetujemo posodobitev čim prej.\x02Tuneli\x02U" +
	"r&edi\x02Dodaj praz&en tunel\u00a0…\x02Dodaj tunel\x02Odstrani izbrane t" +
	"unele\x02Izvozi vse tunele v zip\x02&Preklopi\x02Izvozi vse tunele v &zi" +
	"p\u00a0…\x02Uredi i&zbran tunel\u00a0…\x02Odst&rani izbrane tunele\x02Na" +
	"paka pri uvozu izbrane konfiguracije: %[1]v\x02Napaka pri preštevanju ob" +
	"stoječih tunelov: %[1]v\x02Tunel z imenom »%[1]s« že obstaja\x02Napaka p" +
	"ri uvozu konfiguracije: %[1]v\x02Uvoženi tuneli\x14\x01\x81\x01\x00\x04" +
	"\x16\x02Uvoženi %[1]d tuneli\x02\x14\x02Uvožen %[1]d tunel\x03\x16\x02Uv" +
	"ožena %[1]d tunela\x00\x17\x02Uvoženo %[1]d tunelov\x14\x01\x81\x01\x00" +
	"\x04 \x02Uvoženi %[1]d od %[2]d tunelov\x02\x1f\x02Uvožen %[1]d od %[2]d" +
	" tunelov\x03 \x02Uvožena %[1]d od %[2]d tunelov\x00 \x02Uvoženo %[1]d od" +
	" %[2]d tunelov\x02Napaka pri stvaritvi tunela\x14\x01\x81\x01\x00\x04" +
	"\x16\x02Izbriši %[1]d tunele\x02\x15\x02Izbriši %[1]d tunel\x03\x16\x02I" +
	"zbriši %[1]d tunela\x00\x17\x02Izbriši %[1]d tunelov\x14\x01\x81\x01\x00" +
	"\x048\x02Ali ste prepričani, da želite izbrisati %[1]d tunele?\x027\x02A" +
	"li ste prepričani, da želite izbrisati %[1]d tunel?\x038\x02Ali ste prep" +
	"ričani, da želite izbrisati %[1]d tunela?\x009\x02Ali ste prepričani, da" +
	" želite izbrisati %[1]d tunelov?\x02Izbriši tunel ‘%[1]s’\x02Ali ste pre" +
	"pričani, da želite izbrisati tunel »%[1]s«?\x02%[1]s Tega dejanja ne mor" +
	"ete razveljaviti.\x02Napaka pri izbrisu tunela\x02Napaka pri odstranjeva" +
	"nju tunela: %[1]s\x02Napaka pri izbrisu tunelov\x14\x01\x81\x01\x00\x04*" +
	"\x02%[1]d tunelov ni bilo mogoče odstraniti.\x02)\x02%[1]d tunela ni bil" +
	"o mogoče odstraniti.\x03*\x02%[1]d tunelov ni bilo mogoče odstraniti." +
	"\x00*\x02%[1]d tunelov ni bilo mogoče odstraniti.\x02Konfiguracijske dat" +
	"oteke (*.zip, *.conf)|*.zip;*.conf|Vse datoteke (*.*)|*.*\x02Uvozi tunel" +
	"e iz datoteke\x02Konfiguracijske datoteke ZIP (*.zip)|*.zip\x02Izvozi tu" +
	"nele v zip\x02%[1]s (nepodpisane izdelave, brez posodobitev)\x02Napaka p" +
	"ri izhodu iz WireGuarda\x02Storitve ni bilo mogoče zaustaviti, ker: %[1]" +
	"v. Poskusite zaustaviti WireGuard z uporabo programa Storitve.\x02Posodo" +
	"bitev WireGuarda je na voljo. Zelo priporočamo posodobitev brez odlašanj" +
	"a.\x02Status: Čaka na uporabnika\x02Posodobi zdaj\x02Status: Čaka na ser" +
	"vis za posodobitev\x02Napaka: %[1]v. Poskusite ponovno.\x02Status: Konča" +
	"no!\x02http2: Framer %[1]p: napaka pri dekodiranju ravnokar zapisanega o" +
	"kvirja\x02http2: Framer %[1]p: zapisano %[2]v\x02http2: Framer %[1]p: pr" +
	"ebrano %[2]v\x02http2: dekodirano polje hpack %+[1]v"

	// Total table size 20160 bytes (19KiB); checksum: 352FA2FD
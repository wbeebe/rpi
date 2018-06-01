90-i2c.rules and 99-gpio.rules should be copied to /etc/udev/rules.d.
Create group gpio and add group gpio to any users needed access to GPIO.

If you have an Apple USB keyboard copy hid_apple.conf to /etc/modprobe.d.

To expand the number of virtual text logins from 6 to 9 copy logind.conf
to /etc/systemd. If you already have a modified logind.conf, then set the
following parameter in that file:

NAutoVTs=9

making sure to uncomment the entry.

Reboot after these changes.


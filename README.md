# bgw-210-connection-tracking

Using https://github.com/chromedp/chromedp

BGW-210 has no SNMP.

BGW-210 has no true "bridge mode" and will always maintain its own internal connection-tracking table.

BGW-210 becomes unstable when the connection-tracking table nears full.

Extract the total and current connecting tracking stats and post them to an influx db.


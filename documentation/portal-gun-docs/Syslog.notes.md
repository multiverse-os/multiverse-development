
http://square.github.io/cube/
Time Series Data Collection & Analysis
Cube works great with Cubism, our JavaScript library for visualizing time series.

Cube is a system for collecting timestamped events and deriving metrics. By collecting events rather than metrics, Cube lets you compute aggregate statistics post hoc. It also enables richer analysis, such as quantiles and histograms of arbitrary event sets. Cube is built on MongoDB and available under the Apache License on GitHub. 

https://github.com/ekanite/ekanite
Ekanite is a high-performance syslog server with built-in text search. Its goal is to do a couple of things, and do them well -- accept log messages over the network, and make it easy to search the messages. What it lacks in feature, it makes up for in focus. Built in Go, it has no external dependencies, which makes deployment easy.
				Supports reception of log messages over UDP, TCP, and TCP with TLS.
				Full text search of all received log messages.
				Full parsing of RFC5424 headers.
				Log messages are indexed by parsed timestamp, if one is available. This means search results are presented in the order the messages occurred, not in the order they were received, ensuring sensible display even with delayed senders.
				Automatic data-retention management. Ekanite deletes indexed log data older than a configurable time period.
				Not a JVM in sight.


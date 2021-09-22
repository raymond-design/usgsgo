UsgsGo
-------
UsgsGo is a lightweight client library for working with USGS Water APIs in Go.

Install
-------
   `go get github.com/raymond-design/usgsgo`

Notes
-------
- USGS occasionally changes the API. This library aims to maintain compatibility with the current API. No gaurantees, however.
- This library is also not opinionated on query efficiency. It is up to the developer.
- The public API has rate limits. This library does not attempt to handle rate limits. According to the USGS FAQ: `For small requests (e.g., up to 7 or 14 days for a single site) limit requests to a maximum of 5-10 per second at a steady rate. A maximum burst rate of to 40 or 50 request per second is OK, but not for extended periods of time.`

Usage
-------
This client library is not complete yet. Currently, none of the web services are supported.
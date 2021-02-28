# gratisdns-dynamic-dns

DynamicDNS Client written in [Go](https://golang.org/) for usage on [GratisDNS](https://web.gratisdns.dk/)

## Installation

* Download the latest version from the [Releases](https://github.com/kasperstad/gratisdns-dynamic-dns/releases) tab
* Unpack the downloaded version on the Computer/Server which should do the DynamicDNS Update
  * When you have unpacked the executeable, run it one time to create the config file.
  * Fill out the information, this will create a **config.json** file with the info needed to run the Dynamic DNS update
* Run the program by hand or schedule it using Task Scheduler or Crontab

### Supported Platforms

| OS                     | Supported             |
| ---------------------- | --------------------- |
| Ubuntu 20.04 LTS       | Yes                   |
| Windows 10             | Yes                   |
| Windows Server 2016    | Yes                   |
| Windows Server 2019    | Yes                   |

## License

MIT License

Copyright (c) 2021 Kasper Stad

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

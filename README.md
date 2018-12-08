# Linux System Definition Parser and SPDX Generator

Open Invention Network's website includes a series of tables that list the
packages comprising their Linux System Definition.

To make this information more interchangeable between tools and services, this
repository provides the Linux System Definition tables' data in SPDX tag-value
format, together with the code used to parse and generate the SPDX documents.

## Caveats

Note that the generated SPDX files do **not** include any information about the
licenses or copyright notices associated with these packages. 

If you want to re-run the parser and generate the SPDX files themselves, note
that the code currently assumes that you have already downloaded the Tables'
HTML code to your local system, and have placed them in the specific locations
listed at the top of the `main()` function. I will be updating this code to
enable automatically retrieving and saving the Tables separately.

## Requirements

This repo uses the following dependencies:

* [`github.com/anaskhan96/soup`](https://github.com/anaskhan96/soup): to parse the Tables' HTML files
* [`github.com/swinslow/spdx-go`](https://github.com/swinslow/spdx-go): to generate and save the SPDX output

## License

The SPDX documents in `/spdxdocs/` are licensed under Creative Commons Zero, CC0-1.0.

The spdx-oin source code and documentation is provided and may be used, at your
option, under *either*:
* Apache License, version 2.0 (**Apache-2.0**), **OR**
* GNU General Public License, version 2.0 or later (**GPL-2.0-or-later**).

SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

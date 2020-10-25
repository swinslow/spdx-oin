# Linux System Definition Parser and SPDX Generator

Open Invention Network's website includes [a series of tables](https://www.openinventionnetwork.com/joining-oin/linux-system/) that list the
packages comprising the Linux System Definition.

To make this information more interchangeable between tools and services, this
repository provides the Linux System Definition tables' data in SPDX tag-value
format, together with the code used to parse and generate the SPDX documents.

## Caveats

Note that the generated SPDX files do **not** include any information about the
licenses or copyright notices associated with these packages. These are just
intended to be a structured, interchangeable formatted version of the data from
the tables.

If you want to re-run the parser and generate the SPDX files themselves, note
that the code currently assumes that you have already downloaded the Tables'
HTML code to your local system, and have placed them in the specific locations
listed at the top of the `main()` function. I plan to update this code to enable
automatically retrieving and saving the Tables separately.

## Requirements

If you just want to use the SPDX files corresponding to the Linux System Tables,
you can find them in the `/spdxdocs/` folder.

If you're looking to use the code that created them, it uses the following dependencies:

* [`github.com/anaskhan96/soup`](https://github.com/anaskhan96/soup): to parse the Tables' HTML files
* [`github.com/spdx/tools-golang`](https://github.com/spdx/tools-golang): to generate and save the SPDX output

## License

The SPDX documents in `/spdxdocs/` are licensed under [Creative Commons Zero](https://creativecommons.org/publicdomain/zero/1.0/) ([CC0-1.0](https://spdx.org/licenses/CC0-1.0.html)).

The spdx-oin source code and documentation is provided and may be used, at your
option, under *either*:
* [Apache License, version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([Apache-2.0](https://spdx.org/licenses/Apache-2.0.html)), **OR**
* [GNU General Public License, version 2.0 or later](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html) ([GPL-2.0-or-later](https://spdx.org/licenses/GPL-2.0-or-later.html)).

SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

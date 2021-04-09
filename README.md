# pngextract

This codebase was written to help someone. It is effectively similar to binwalk but only extracts PNG files. Unlike binwalk, however, it extracts entire PNGs exactly, ending on their IEND chunk. (Note that it is possible to have PNGs without IEND chunks; though not well-formed, most PNG parsers support it. This extractor will fail if there is no IEND chunk found.)

It might be interesting to extend it to be aware of more formats and essentially build a binwalk-like tool. It might be a good excuse to try implementing the Commentz-Walter algorithm.

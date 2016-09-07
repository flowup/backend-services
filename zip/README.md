## Zip

Zipping service core function is to compress and decompress given data in io.Reader and io.Writer. In addition compressing and decompressing of files was added.
This Zip service uses Gzip algorithm for zipping.

## Code Example

The code below shows operations you can do with events:

    import "flowdock.eu/flowup/services/zip"

    zip := zip.NewTracker()

    zip.Compress(source, destination) // source is io.Reader and destination io.Writer these can be simulated by buffers
    
    zip.Decompress(source, destination)

In addition there are added functions for zipping and decompressing files:

    zip.service.CompressFile("path_to_the_file", "destination_folder")
    
    zip.DecompressFile("path_to_the_file", "destination_folder")
    

## Tests

All tests can be executed by running: 

    goconvey 

in the folder with test file, alternatively you can use command: 

    go test

## License
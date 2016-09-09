## File

File service is service that is providing basic methods for storing files to the server and on the contrary downloading stored data. Data from the server can be deleted. Also there is implemented method for accessing meta data. 

## Code Example

Example how to initialize this service and `UPLOAD FILE`:

    import "flowdock.eu/flowup/services/file"

    fileGrid := file.NewGrid()

    file, _ := os.Open("path/to/the/file.ext")
    
    m := service.Upload(file, file.Name())

This way you can `DOWNLOAD FILE` by ID or possibly by hash:

    file = fileGrid.Download(ID) // ID is ID of the file that is requested  
    
    file = fileGrid.DownloadByHash("hash_string")

Same for the `DELETING FILE`:

    fileGrid.Delete(0)

For getting `META DATA` of the file there are two functions similiar to download functions:
    
    meta := fileGrid.GetMeta(ID) // ID of the file that is requested 
    
    meta = fileGrid.GetMetaByHash(hash.Hash)

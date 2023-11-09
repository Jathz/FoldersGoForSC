package folders

// Type for this response
type FetchFolderResponsePagination struct {
	Folders []*Folder
	Token   string
}

// GetAllFoldersWithPagination retrieves folders based on the provided FetchFolderRequest and token for pagination.
func GetAllFoldersWithPagination(req *FetchFolderRequest, token string) (*FetchFolderResponsePagination, error) {
	const chunkSize = 2 // The chunk size, this is the amount of folders you want returned with each request

	r, _ := FetchAllFoldersByOrgID(req.OrgID) // Retrieve all folders for the given organization ID using the fuction in folder.go

	var fp []*Folder
	start := 0
	end := chunkSize

	// If the token is provided adjust the start and end of the chunk
	if token != "" {
		for i, folder := range r {
			if folder.Id.String() == token {
				start = i + 1
				end = start + chunkSize
				break
			}
		}
	}

	// If the end if the chunk is past the last folder set the exact length of the end of the chunk needed
	if end > len(r) {
		end = len(r)
	}

	// Iterate over the data to extract the required chunk
	for i := start; i < end; i++ {
		fp = append(fp, r[i])
	}

	newToken := ""
	if end < len(r) {
		newToken = r[end-1].Id.String() // Convert UUID to string for the token to be returned
	}

	ffr := &FetchFolderResponsePagination{Folders: fp, Token: newToken} // Include the new token in the response

	return ffr, nil
}

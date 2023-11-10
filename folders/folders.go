package folders

import "github.com/gofrs/uuid"

// GetAllFolders retrieves all folders based on the provided FetchFolderRequest.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	r, _ := FetchAllFoldersByOrgID(req.OrgID) // Stores the first returned value in r and ignores the second returned variable

	ffr := &FetchFolderResponse{Folders: r} // Create a pointer to a new FetchFolderResponse struct, where the Folders field of the struct is initialized with the fp slice
	return ffr, nil                         // Return the pointer to the newly created FetchFolderResponse struct and error as nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData() // Gets a list of pointers to the folders in sample.json using this function defined in static.go

	resFolder := []*Folder{} // Creates an Empty list of ponters to folders

	// Iterates through the list of folders from the sample data
	for _, folder := range folders {

		// Checks if the folders organisation ID is the same as the one given as input to this fuction
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder) // Appends the current folder in the list to the end of the reFolder list
		}
	}

	return resFolder, nil // Returns a list of pointer to folders of the given orgID
}

// GetAllFolders retrieves all folders based on the given FetchFolderRequest.
/*func GetAllFoldersOriginal(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	var (
		err error
		_   = err
		f1  Folder
		_   = f1
		fs  []*Folder
		_   = fs
	) // Unsused unessisary variables
	f := []Folder{}                           // Create a new empty slice of type "Folder"
	r, _ := FetchAllFoldersByOrgID(req.OrgID) // Stores the first returned value in r and ignores the second returned variable (r is a list of pointers to Folders)

	// Loop to iterate over the range of all the folders
	for _, v := range r { // k is the index of the current folder in the loop, v is the value pointed to by that index
		f = append(f, *v) // This appends the value that is pointed to by v to f and sets that as the new f
	}

	var fp []*Folder

	for _, v1 := range f {
		//temp := v1 // This line is not necessary in Go 1.20, but it is needed in Go 1.21
		fp = append(fp, &v1) // Append the address of v1 to fp
	}
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: r} // Creates a new FetchFolderResponse struct with the fetched folders
	return ffr, nil
}*/

/* The Original GetAllFolders function retrieves the folders based on the provided requests organzation Id and
then performs unnecessary variable initializations and conversions to ponter and back before finally
returning the response. A more efficient way would be directly returning a response
containing the returned folders. I have renamed the Original GetAllFolders "GetAllFoldersOriginal" to show the
comments I wrote to understand what it was doing and removed all the fmt.prints I used help myself understand.
I have written a new imporoved "GetAllFolders" function above */

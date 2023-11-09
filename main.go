package main

import (
	"fmt"

	"encoding/json"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	res, err := folders.GetAllFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	folders.PrettyPrint(res)

	// Pagination
	/* This runs the pagination in folders_pagination.go by first giving the token "" then
	iterates on the returned token value till it returns a empty string as a token. */
	// Comment out the below section to test improvements for Component 1.

	req1 := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	token := ""

	for {
		response, err := folders.GetAllFoldersWithPagination(req1, token)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		s, _ := json.MarshalIndent(response.Folders, "", "\t")
		fmt.Printf("request(\"%s\") -> { data : %s, token: \"%s\" }\n", token, s, response.Token)

		token = response.Token

		if token == "" {
			break
		}

	}
}

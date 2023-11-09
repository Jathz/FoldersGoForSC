package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("Test fetch all folders with specific organization ID", func(t *testing.T) {

		orgID, err := uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
		if err != nil {
			t.Errorf("Error parsing organization ID: %v", err)
		}

		req := &folders.FetchFolderRequest{
			OrgID: orgID,
		}

		res, err := folders.GetAllFolders(req)
		if err != nil {
			t.Errorf("Error getting folders: %v", err)
		}

		assert.NotNil(t, res, "Returned folders should not be nil")
		assert.Equal(t, 666, len(res.Folders), "Expected number of folders not returned for the organization ID")
	})

	t.Run("Test getting all folders with a different organization ID", func(t *testing.T) {

		orgID, err := uuid.FromString("6591e16c-c257-4366-bf6d-650c68f71284")
		if err != nil {
			t.Errorf("Error parsing organization ID: %v", err)
		}

		req := &folders.FetchFolderRequest{
			OrgID: orgID,
		}

		res, err := folders.GetAllFolders(req)
		if err != nil {
			t.Errorf("Error getting folders: %v", err)
		}

		assert.NotNil(t, res, "Returned folders should not be nil")
		assert.Equal(t, 1, len(res.Folders), "Expected number of folders not returned for the organization ID")
	})

	t.Run("Test fetch all folders with a non existant organization ID", func(t *testing.T) {

		orgID, err := uuid.FromString("3212d618-66ff-468a-862d-ea49fef5e183")
		if err != nil {
			t.Errorf("Error parsing organization ID: %v", err)
		}

		req := &folders.FetchFolderRequest{
			OrgID: orgID,
		}

		res, err := folders.GetAllFolders(req)
		if err != nil {
			t.Errorf("Error getting folders: %v", err)
		}

		assert.NotNil(t, res, "Returned folders should not be nil")
		assert.Equal(t, 0, len(res.Folders), "Expected no folders to be retrieved for the provided organization ID")
		assert.NoError(t, err, "Expected no error for the given organization ID")
	})
}

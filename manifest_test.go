package vite

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

type testSuite struct {
	arg                 string
	expectedData        AssetsData
	expectedChuncksList []AssetsData
	expectedError       error
}

var testSuits = []testSuite{
	{arg: validManifest, expectedData: expectedMap, expectedChuncksList: expectedChuncks, expectedError: nil},
	{arg: invalidManifest, expectedData: nil, expectedChuncksList: nil, expectedError: errors.New(INVALID_MANIFEST_STRUCT)},
	{arg: multipleEntryManifest, expectedData: nil, expectedChuncksList: nil, expectedError: errors.New(MULTIPLE_ENTRY_ERR)},
}

func TestMapManifest(t *testing.T) {
	for _, ts := range testSuits {
		var jsonData any

		json.Unmarshal([]byte(ts.arg), &jsonData)

		target, chuncks, err := mapManifest(jsonData)

		targetJson, _ := json.Marshal(target)
		expectedMapJson, _ := json.Marshal(ts.expectedData)

		chuncksJson, _ := json.Marshal(chuncks)
		expectedChuncksJson, _ := json.Marshal(ts.expectedChuncksList)

		if !reflect.DeepEqual(targetJson, expectedMapJson) {
			t.Errorf("Fail: get incorrect manifest map: \n output:   %v,\n expected: %v", target, ts.expectedData)
		}

		if !reflect.DeepEqual(chuncksJson, expectedChuncksJson) {
			t.Errorf("Fail: get incorrect manifest map: \n output:   %v,\n expected: %v", chuncks, ts.expectedChuncksList)
		}

		if err != nil {
			if err.Error() != ts.expectedError.Error() {
				t.Errorf("Expected error: %v, output error: %v", err.Error(), ts.expectedError.Error())
			}
		}
	}
}

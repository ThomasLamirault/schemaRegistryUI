package pageProcessors

import (
	"testing"
)

func TestGetVersionLinkList(t *testing.T) {
	result := GetVersionLinkList("subjectTestName", []int{0,1})
	if result[0].Link != "/subjects/subjectTestName/versions/0" && result[0].Version != 0 {
		t.Errorf("incorrect output from GetVersionLinkList, having version %d, link %s", result[0].Version, result[0].Link)
	}
	if result[1].Link != "/subjects/subjectTestName/versions/1" && result[1].Version != 1 {
		t.Errorf("incorrect output from GetVersionLinkList, having version %d, link %s", result[1].Version, result[1].Link)
	}
}
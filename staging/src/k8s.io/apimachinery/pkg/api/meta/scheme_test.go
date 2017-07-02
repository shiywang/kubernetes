package meta

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kubernetes/pkg/api/testapi"
)

// These types do not follow the list convention as documented in
// docs/devel/api-convention.md
var listTypeExceptions = sets.NewString("APIGroupList", "APIResourceList")

func validateListType(target reflect.Type) error {
	// exceptions
	if listTypeExceptions.Has(target.Name()) {
		return nil
	}
	hasListSuffix := strings.HasSuffix(target.Name(), "List")
	hasMetadata := false
	hasItems := false
	for i := 0; i < target.NumField(); i++ {
		field := target.Field(i)
		tag := field.Tag.Get("json")
		switch {
		case strings.HasPrefix(tag, "metadata"):
			hasMetadata = true
		case tag == "items":
			hasItems = true
			if field.Type.Kind() != reflect.Slice {
				return fmt.Errorf("Expected items to be slice, got %s", field.Type.Kind())
			}
		}
	}
	if hasListSuffix && !hasMetadata {
		return fmt.Errorf("Expected type %s to contain \"metadata\"", target.Name())
	}
	if hasListSuffix && !hasItems {
		return fmt.Errorf("Expected type %s to contain \"items\"", target.Name())
	}
	// if a type contains field Items with JSON tag "items", its name should end with List.
	if !hasListSuffix && hasItems {
		return fmt.Errorf("Type %s has Items, its name is expected to end with \"List\"", target.Name())
	}
	return nil
}

// TestListTypes verifies that no external type violates the api convention of
// list types.
func TestListTypes(t *testing.T) {
	for groupKey, group := range testapi.Groups {
		for kind, target := range group.ExternalTypes() {
			t.Logf("working on %v in %v", kind, groupKey)
			err := validateListType(target)
			if err != nil {
				t.Error(err)
			}
		}
	}
}

type WithoutMetaDataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta
	Items []interface{} `json:"items"`
}

type WithoutItemsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
}

type WrongItemsJSONTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []interface{} `json:"items,omitempty"`
}

// If a type has Items, its name should end with "List"
type ListWithWrongName struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []interface{} `json:"items"`
}

// TestValidateListType verifies the validateListType function reports error on
// types that violate the api convention.
func TestValidateListType(t *testing.T) {
	var testTypes = []interface{}{
		WithoutMetaDataList{},
		WithoutItemsList{},
		WrongItemsJSONTagList{},
		ListWithWrongName{},
	}
	for _, testType := range testTypes {
		err := validateListType(reflect.TypeOf(testType))
		if err == nil {
			t.Errorf("Expected error")
		}
	}
}


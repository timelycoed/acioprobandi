package stemma

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestIsProtoSubjectGroup(t *testing.T) {
	// Create an instance of ProtoGroup
	p := ProtoGroup{}

	// Call the IsProtoSubject method
	result := p.IsProtoSubject()

	// Check if the result is true, as expected
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestIsProtoSubjectEvent(t *testing.T) {
	// Create an instance of ProtoEvent
	p := ProtoEvent{}

	// Call the IsProtoSubject method
	result := p.IsProtoSubject()

	// Check if the result is true, as expected
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestIsProtoSubjectPlace(t *testing.T) {
	// Create an instance of ProtoPlace
	p := ProtoPlace{}

	// Call the IsProtoSubject method
	result := p.IsProtoSubject()

	// Check if the result is true, as expected
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestIsProtoSubjectPerson(t *testing.T) {
	// Create an instance of ProtoPerson
	p := ProtoPerson{}

	// Call the IsProtoSubject method
	result := p.IsProtoSubject()

	// Check if the result is true, as expected
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestIsProtoSubjectAnimal(t *testing.T) {
	// Create an instance of ProtoPerson
	p := ProtoAnimal{}

	// Call the IsProtoSubject method
	result := p.IsProtoSubject()

	// Check if the result is true, as expected
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestProductBuilder(t *testing.T) {
	const expectedID = "id"
	const expectedName = "name"
	const expectedVersion = "version"

	builder := &ProductBuilder{}
	product := builder.SetID(expectedID).SetName(expectedName).SetVersion(expectedVersion).Build()

	if product.ID != expectedID {
		t.Errorf("Expected ID to be %s, got %s instead", expectedID, product.ID)
	}
	if product.Name != expectedName {
		t.Errorf("Expected Name to be %s, got %s instead", expectedName, product.Name)
	}
	if product.Version != expectedVersion {
		t.Errorf("Expected Version to be %s, got %s instead", expectedVersion, product.Version)
	}
}

func TestJoinFromBuilder(t *testing.T) {
	const expectedKey = "pExampleKey123"
	builder := &JoinFromBuilder{}
	joinFrom := builder.SetKey(expectedKey).Build()

	if joinFrom.Key != expectedKey {
		t.Errorf("Expected Key to be %s, got %s instead", expectedKey, joinFrom.Key)
	}
}

// TestProductXMLMarshalling validates marshalling Product to XML.
func TestProductXMLMarshalling(t *testing.T) {
	const expectedID = "id"
	product := &Product{ID: expectedID}
	xmlBytes, err := xml.Marshal(product)
	if err != nil {
		t.Fatalf("Marshalling failed: %v", err)
	}

	expectedXML := `<Product><ID>id</ID></Product>`
	if string(xmlBytes) != expectedXML {
		t.Errorf("Expected XML to be %s, got %s instead", expectedXML, xmlBytes)
	}
}

// TestProductXMLUnmarshalling validates unmarshalling XML to Product structure.
func TestProductXMLUnmarshalling(t *testing.T) {
	const xmlData = `<Product><ID>id</ID></Product>`
	var product Product
	err := xml.Unmarshal([]byte(xmlData), &product)
	if err != nil {
		t.Fatalf("Unmarshalling failed: %v", err)
	}

	expectedID := "id"
	if product.ID != expectedID {
		t.Errorf("Expected ID to be %s, got %s instead", expectedID, product.ID)
	}
}

// TestJoinFromXMLMarshalling validates marshalling JoinFrom to XML.
func TestJoinFromXMLMarshalling(t *testing.T) {
	const expectedKey = "pExampleKey123"
	joinFrom := &JoinFrom{Key: expectedKey}
	xmlBytes, err := xml.Marshal(joinFrom)
	if err != nil {
		t.Fatalf("Marshalling failed: %v", err)
	}

	expectedXML := `<JoinFrom Key="pExampleKey123"></JoinFrom>`
	if string(xmlBytes) != expectedXML {
		t.Errorf("Expected XML to be %s, got %s instead", expectedXML, xmlBytes)
	}
}

// TestJoinFromXMLUnmarshalling validates unmarshalling XML to JoinFrom structure.
func TestJoinFromXMLUnmarshalling(t *testing.T) {
	const xmlData = `<JoinFrom Key="pExampleKey123"></JoinFrom>`
	var joinFrom JoinFrom
	err := xml.Unmarshal([]byte(xmlData), &joinFrom)
	if err != nil {
		t.Fatalf("Unmarshalling failed: %v", err)
	}

	expectedKey := "pExampleKey123"
	if joinFrom.Key != expectedKey {
		t.Errorf("Expected Key to be %s, got %s instead", expectedKey, joinFrom.Key)
	}
}

// TestJoinFromKeyValidation validates if the Key is according to the prescribed format.
func TestJoinFromKeyValidation(t *testing.T) {
	tests := []struct {
		Key      string
		Expected bool
	}{
		{"pA1ValidKey123", true},
		{"cBAnotherValidKey", true},
		{"sValidWithLowercaseSecond", true},
		{"p2InvalidMissingCharAfterLowercase", true},
		{"InvalidKey", false},
		{"2InvalidKeyStart", false},
		{"cainvalidAllLowercase", false},
		{"rg1MixedInvalidStart", true},
	}

	for _, test := range tests {
		joinFrom := &JoinFrom{Key: test.Key}
		valid := isValidKey(joinFrom.Key)
		if valid != test.Expected {
			t.Errorf("Expected validity of key '%s' to be %v, got %v instead", test.Key, test.Expected, valid)
		}
	}
}

// isValidKey checks if the Key is valid according to the specified rules.
func isValidKey(key string) bool {
	if len(key) < 2 {
		return false
	}

	// First character check: must be within "acdeglmnprstw"
	if !strings.ContainsRune("acdeglmnprstw", rune(key[0])) {
		return false
	}

	// Second character check
	if 'a' <= key[1] && key[1] <= 'z' {
		// If second character is lowercase, it must be within "acdegpsw"
		if !strings.ContainsRune("acdegpsw", rune(key[1])) {
			return false
		}
		if !('A' <= key[2] && key[2] <= 'Z') && !('0' <= key[2] && key[2] <= '9') {
			return false
		}

	} else if !('A' <= key[1] && key[1] <= 'Z') && !('0' <= key[1] && key[1] <= '9') {
		// If second character is not lowercase, it must be uppercase or a digit
		return false
	}

	// Remaining characters: must be alphanumeric
	for _, ch := range key[3:] {
		if !isAlphanumeric(ch) {
			return false
		}
	}

	return true
}

// isAlphanumeric checks if the rune is a letter or a digit.
func isAlphanumeric(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9')
}

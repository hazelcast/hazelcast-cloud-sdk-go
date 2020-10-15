package hazelcastcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHazelcastVersionServiceOp_List(t *testing.T) {
	//given
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	defer server.Close()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		var request GraphQLQuery
		json.NewDecoder(r.Body).Decode(&request)

		if strings.Contains(request.Query, "hazelcastVersions") {
			fmt.Fprint(w, `{"data":{"response":[{"version":"4.0"},{"version":"3.12.6"},{"version":"3.12.5"},{"version":"3.12.4"},{"version":"3.12.3"},{"version":"3.12.2"},{"version":"3.12.1"},{"version":"3.12"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	hazelcastVersions, _, _ := NewHazelcastVersionService(client).List(context.TODO())

	//then
	assert.Len(t, *hazelcastVersions, 8)
}

func ExampleHazelcastVersionService_list() {
	client, _, _ := New()
	hazelcastVersions, _, _ := client.HazelcastVersion.List(context.Background())
	fmt.Printf("Result: %#v", hazelcastVersions)
	//Output:Result: &[]models.EnterpriseHazelcastVersion{models.EnterpriseHazelcastVersion{Version:"4.0", UpgradeableVersions:[]string{}}, models.EnterpriseHazelcastVersion{Version:"3.12.6", UpgradeableVersions:[]string{}}, models.EnterpriseHazelcastVersion{Version:"3.12.5", UpgradeableVersions:[]string{"3.12.6"}}, models.EnterpriseHazelcastVersion{Version:"3.12.4", UpgradeableVersions:[]string{"3.12.6", "3.12.5"}}, models.EnterpriseHazelcastVersion{Version:"3.12.3", UpgradeableVersions:[]string{"3.12.6", "3.12.5", "3.12.4"}}, models.EnterpriseHazelcastVersion{Version:"3.12.2", UpgradeableVersions:[]string{"3.12.6", "3.12.5", "3.12.4", "3.12.3"}}, models.EnterpriseHazelcastVersion{Version:"3.12.1", UpgradeableVersions:[]string{"3.12.6", "3.12.5", "3.12.4", "3.12.3", "3.12.2"}}, models.EnterpriseHazelcastVersion{Version:"3.12", UpgradeableVersions:[]string{"3.12.6", "3.12.5", "3.12.4", "3.12.3", "3.12.2", "3.12.1"}}}
}

package internal

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"testing"
)

func TestValidateCRDS(t *testing.T) {
	file := "./testdata/valid/etcdbackups.etcd.database.coreos.com.crd.yaml"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	crd := &v1beta1.CustomResourceDefinition{}
	err = yaml.Unmarshal(data, crd)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", crd)

	results := validateCRDs(crd)
	if len(results) > 0 {
		t.Fatalf("expected 0 errors: got %v", results)
	}
}

func TestValidateV1Beta1CRD(t *testing.T) {
	file := "./testdata/valid/etcdbackups.etcd.database.coreos.com.crd.yaml"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	crd := &v1beta1.CustomResourceDefinition{}
	err = yaml.Unmarshal(data, crd)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", crd)

	results := validateV1Beta1CRD(crd)
	if len(results.Errors) > 0 {
		t.Fatalf("expected 0 errors: got %v", results)
	}
}

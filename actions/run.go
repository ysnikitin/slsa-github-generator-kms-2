package main

import (
	"context"
	"fmt"
	intoto "github.com/in-toto/in-toto-golang/in_toto"
	"github.com/slsa-framework/slsa-github-generator/signing/sigstore"

	_ "github.com/sigstore/sigstore/pkg/signature/kms/gcp"
)

const (
	keyRef = "gcpkms://projects/etsy-builder-kms-prod/locations/global/keyRings/testkeys-services-keyring/cryptoKeys/test_a-sigining-key"
)

func main() {
	ctx := context.Background()

	statement := &intoto.Statement{
		StatementHeader: intoto.StatementHeader{
			PredicateType: "https://slsa.dev/provenance/v1",
			Subject: []intoto.Subject{
				{
					Name: "hi",
				},
			},
		},
	}

	//fulcio := sigstore.NewDefaultFulcio()
	fulcio := sigstore.NewWithKeyFulcio(keyRef)
	attest, err := fulcio.Sign(ctx, statement)
	if err != nil {
		fmt.Printf("error exiting: %v", err)
		return
	}

	fmt.Printf("attest: %v", attest.Bytes())
	fmt.Printf("attest: %v", string(attest.Cert()))
}

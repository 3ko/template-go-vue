package utils

import (
    "time"

    "github.com/MicahParks/keyfunc/v3"
)

var jwksCache *keyfunc.JWKS

func LoadJWKS(jwksURL string) (*keyfunc.JWKS, error) {
    if jwksCache != nil {
        return jwksCache, nil
    }

    options := keyfunc.Options{
        RefreshInterval: time.Hour,
        RefreshErrorHandler: func(err error) {
            // could log error
        },
    }

    var err error
    jwksCache, err = keyfunc.Get(jwksURL, options)
    if err != nil {
        return nil, err
    }

    return jwksCache, nil
}

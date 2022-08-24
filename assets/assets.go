package assets

import (
    _ "embed"
)

var (
    //go:embed swagger.redoc.html
    SwaggerRedocUI string

    //go:embed docs/swagger.json
    SwaggerSpec []byte

    //go:embed docs/swagger.yaml
    SwaggerSpecYaml []byte

    //go:embed octicons.css
    OcticonsCss []byte

    //go:embed config.yaml
    DefaultConfig []byte
)

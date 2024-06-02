package model

import (
    "vitess.io/vitess/go/vt/vtgate/vtgateconn"
    "vitess.io/vitess/go/vt/vtgate/vtgateconn/driver"
    "vitess.io/vitess/go/vt/proto/query"
)

func Create() {
    // Connect to the Vitess vtgate service
    conn, err := vtgateconn.Dial("localhost:15991")
    if err != nil {
        log.Fatalf("failed to connect to vtgate: %v", err)
    }
    defer conn.Close()
}
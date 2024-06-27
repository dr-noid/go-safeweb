package safehttp

import (
    "net/http"
    "testing"
)

func TestFlightAddCookie(t *testing.T) {
    f := flight{
        header: NewHeader(http.Header{}),
    }

    f.AddCookie(&Cookie{
        wrapped: &http.Cookie{
            Name:  "name",
            Value: "value",
        },
    })

    cookie := f.header.wrapped.Get("Set-Cookie")

    if cookie != "name=value" {
        t.Errorf("got %s, want 'name=value'", cookie)
    }
}
Little utility to read properties from json when the content is not known at compilation time.

Usage example:

    var data = `
    {
      "hey": {
        "what": {
          "a": {
            "number": 3,
            "string": "here",
          }
        }
      }
    }
    `

    myDigger := NewDigger([]byte(data), ".")
    s := myDigger.GetString("hey.what.a.string")
    n := myDigger.GetNumber("hey.what.a.number")

And there we should have 's' containing 'here', and 'n' being 3.

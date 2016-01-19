package digger

import "testing"

const mdYAMLData = `
guestbook:
  redis-master:
    rc:
      num: 1
  redis-slave:
    rc:
      num: 3
  frontend:
    rc:
      images:
        image: "google.io/guestbook"
  somebool: true
`

const mdJSONData = `
{
  "guestbook": {
    "redis-slave": {
      "rc": {
        "num": 7
      }
    },
    "frontend": {
      "rc": {
        "num": 3
      }
    },
    "someotherbool": true
  },
  "id": "ABCD1234"
}
`

func multi() Digger {
	jsonDigger, _ := NewJSONDigger([]byte(mdJSONData))
	yamlDigger, _ := NewYAMLDigger([]byte(mdYAMLData))
	multiDigger, _ := NewMultiDigger(
		jsonDigger,
		yamlDigger,
	)
	return multiDigger
}

func Test_Multi_GetString_TopLayer_Success(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetString("id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if val != "ABCD1234" {
		t.Errorf("Unexpected value: '%v', was expecting '%v'", val, "ABCD1234")
	}
}

// Test retrieving a property that's not in the first layer
func Test_Multi_GetString_BottomLayer_Success(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetString("guestbook/frontend/rc/images/image")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if val != "google.io/guestbook" {
		t.Errorf("Unexpected value: '%v', was expecting '%v'", val, "google.io/guestbook")
	}
}

func Test_Multi_GetString_IncorrectPath(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetString("some/invalid/path")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != "" {
		t.Errorf("Unexpected value: '%v', was expecting an empty string", val)
	}
}

func Test_Multi_GetString_IncorrectType(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetString("guestbook/redis-master/rc/num")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != "" {
		t.Errorf("Unexpected value: '%v', was expecting an empty string", val)
	}
}

func Test_Multi_GetNumber_TopLayer_Success(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetNumber("guestbook/redis-slave/rc/num")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if val != 7 {
		t.Errorf("Unexpected value: '%v', was expecting '%v'", val, 7)
	}
}

func Test_Multi_GetNumber_BottomLayer_Success(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetNumber("guestbook/redis-master/rc/num")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if val != 1 {
		t.Errorf("Unexpected value: '%v', was expecting '%v'", val, 1)
	}
}

func Test_Multi_GetNumber_IncorrectPath(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetNumber("guestbook/redis/num")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != 0 {
		t.Errorf("Unexpected value: '%v', was expecting 0", val)
	}
}

func Test_Multi_GetNumber_IncorrectType(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetNumber("guestbook/frontend/rc/images/image")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != 0 {
		t.Errorf("Unexpected value: '%v', was expecting 0", val)
	}
}

func Test_Multi_GetBool_TopLayer_Success(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetBool("guestbook/someotherbool")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if val != true {
		t.Errorf("Unexpected value: '%v', was expecting '%v'", val, true)
	}
}

func Test_Multi_GetBool_BottomLayer_Success(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetBool("guestbook/somebool")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if val != true {
		t.Errorf("Unexpected value: '%v', was expecting '%v'", val, true)
	}
}

func Test_Multi_GetBool_IncorrectPath(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetBool("guestbook/abool")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != false {
		t.Errorf("Unexpected value: '%v', was expecting false", val)
	}
}

func Test_Multi_GetBool_IncorrectType(t *testing.T) {
	multiDigger := multi()
	val, err := multiDigger.GetBool("guestbook/frontend/rc/images/image")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != false {
		t.Errorf("Unexpected value: '%v', was expecting false", val)
	}
}

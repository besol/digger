package digger

import "testing"

var data = `
{
  "guestbook": {
    "redis-master": {
      "rc": {
        "num": 3
      }
    },
    "redis-slave": {
      "rc": {
        "num": 5
      }
    },
    "frontend": {
      "rc": {
        "images": {
          "image": "google.io/guestbook"
        }
      }
    },
    "somebool": true
  }
}
`

func Test_GetString_Success(t *testing.T) {
	myDigger, err := NewDigger([]byte(data), ".")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, err := myDigger.GetString("guestbook.frontend.rc.images.image")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if val != "google.io/guestbook" {
		t.Errorf("Unexpected value: '%v', was expecting '%v'", val, "google.io/guestbook")
	}
}

func Test_GetString_IncorrectPath(t *testing.T) {
	myDigger, err := NewDigger([]byte(data), ".")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, err := myDigger.GetString("guestbook.frontend.image.prop")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != "" {
		t.Errorf("Unexpected value: '%v', was expecting an empty string", val)
	}
}

func Test_GetString_IncorrectType(t *testing.T) {
	myDigger, err := NewDigger([]byte(data), ".")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, err := myDigger.GetString("guestbook.redis-master.rc.num")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != "" {
		t.Errorf("Unexpected value: '%v', was expecting an empty string", val)
	}
}

func Test_GetNumber_Success(t *testing.T) {
	myDigger, err := NewDigger([]byte(data), ".")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, err := myDigger.GetNumber("guestbook.redis-master.rc.num")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if val != 3 {
		t.Errorf("Unexpected value: '%v', was expecting '%v'", val, 3)
	}
}

func Test_GetNumber_IncorrectPath(t *testing.T) {
	myDigger, err := NewDigger([]byte(data), ".")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, err := myDigger.GetNumber("guestbook.redis.num")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != 0 {
		t.Errorf("Unexpected value: '%v', was expecting 0", val)
	}
}

func Test_GetNumber_IncorrectType(t *testing.T) {
	myDigger, err := NewDigger([]byte(data), ".")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, err := myDigger.GetNumber("guestbook.frontend.rc.images.image")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != 0 {
		t.Errorf("Unexpected value: '%v', was expecting 0", val)
	}
}

func Test_GetBoolean_Success(t *testing.T) {
	myDigger, err := NewDigger([]byte(data), ".")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, err := myDigger.GetBoolean("guestbook.somebool")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if val != true {
		t.Errorf("Unexpected value: '%v', was expecting '%v'", val, 3)
	}
}

func Test_GetBoolean_IncorrectPath(t *testing.T) {
	myDigger, err := NewDigger([]byte(data), ".")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, err := myDigger.GetBoolean("guestbook.abool")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != false {
		t.Errorf("Unexpected value: '%v', was expecting false", val)
	}
}

func Test_GetBoolean_IncorrectType(t *testing.T) {
	myDigger, err := NewDigger([]byte(data), ".")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, err := myDigger.GetBoolean("guestbook.frontend.rc.images.image")
	if err == nil {
		t.Errorf("Expected error but it didnt happen")
	}
	if val != false {
		t.Errorf("Unexpected value: '%v', was expecting false", val)
	}
}

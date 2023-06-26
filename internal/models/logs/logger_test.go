package logs

import (
	"testing"

	"github.com/Owicca/chan/internal/infra"
)

func TestLogger(t *testing.T) {
	_, conn, _ := infra.Setup("./config_test.json")
	defer conn.Close()

	t.Run("success", func(t *testing.T) {
		lg := NewActionLog(conn)
		data := []Entry{
			Entry{
				Action:      "insert",
				Subject:     int64(1),
				Object:      int64(1),
				Object_type: "boards",
				Data:        "data",
			},
			Entry{
				Action:      "delete",
				Subject:     int64(1),
				Object:      int64(1),
				Object_type: "boards",
				Data:        "",
			},
		}

		for _, v := range data {
			err := lg.Write(v)

			if err != nil {
				t.Error("Success run did not complete ", err)
			}
		}
	})
	t.Run("field values are not checked", func(t *testing.T) {
		lg := NewActionLog(conn)
		data := map[string]Entry{
			"Action": Entry{
				Action:      "",
				Subject:     int64(1),
				Object:      int64(1),
				Object_type: "boards",
				Data:        "",
			},
			"Subject": Entry{
				Action:      "delete",
				Subject:     int64(0),
				Object:      int64(1),
				Object_type: "boards",
				Data:        "",
			},
			"Object": Entry{
				Action:      "delete",
				Subject:     int64(1),
				Object:      int64(0),
				Object_type: "boards",
				Data:        "",
			},
			"Object_type": Entry{
				Action:      "delete",
				Subject:     int64(1),
				Object:      int64(1),
				Object_type: "",
				Data:        "",
			},
			"Data_on_insert": Entry{
				Action:      "insert",
				Subject:     int64(1),
				Object:      int64(1),
				Object_type: "boards",
				Data:        "",
			},
			"Data_on_update": Entry{
				Action:      "update",
				Subject:     1,
				Object:      1,
				Object_type: "boards",
				Data:        "",
			},
		}

		for k, v := range data {
			err := lg.Write(v)
			t.Log(k, v, err)

			if err == nil {
				t.Errorf("invalid value '%s' is not checked properly", k)
			}
		}
	})
}

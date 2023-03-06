package booking

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/danielbcnicode/timeslot/internal"
	"github.com/stretchr/testify/assert"
)

func TestRequstAPISerialization(t *testing.T) {
	serialized := `[
		{ 
			"request_id":"bookata_XY123", 
			"check_in":"2020-01-01", 
			"nights":5,
      		"selling_rate":200,
      		"margin":20
    	},
		{
			"request_id":"kayete_PP234",
			"check_in":"2020-01-04",
			"nights":4,
			"selling_rate":156,
			"margin":22
		}
	]`

	want := []RequestAPI{
		{
			RequestID:   "bookata_XY123",
			CheckIn:     "2020-01-01",
			Nights:      5,
			SellingRate: 200,
			Margin:      20,
		},
		{
			RequestID:   "kayete_PP234",
			CheckIn:     "2020-01-04",
			Nights:      4,
			SellingRate: 156,
			Margin:      22,
		},
	}

	request := []RequestAPI{}

	err := json.Unmarshal([]byte(serialized), &request)
	if err != nil {
		t.Errorf("error unmarshalling json: %s\n", err)
	}

	assert.Equal(t, want, request)
	fmt.Printf("request := %#v \n", request)

}

func TestRequestFromRequestAPI(t *testing.T) {
	type args struct {
		req RequestAPI
	}
	tests := []struct {
		name    string
		args    args
		want    Request
		wantErr bool
		error   error
	}{
		{
			name: "Happy path in constructor",
			args: args{
				req: RequestAPI{
					RequestID:   "id-test",
					CheckIn:     "2022-02-01",
					Nights:      5,
					SellingRate: 200,
					Margin:      20,
				},
			},
			want: Request{
				"id-test",
				200,
				20,
				float32(200) * (float32(20) / float32(100)) / float32(5),
				float32(200) * (float32(20) / float32(100)),
				internal.NewDaySlot(
					time.Date(2022, time.Month(2), 1, 12, 0, 0, 0, time.UTC),
					time.Date(2022, time.Month(2), 6, 12, 0, 0, 0, time.UTC),
				),
			},
			wantErr: false,
			error:   nil,
		},
		{
			name: "error when data is wrong",
			args: args{
				req: RequestAPI{
					RequestID:   "id-test",
					CheckIn:     "20sd-02-01",
					Nights:      5,
					SellingRate: 200,
					Margin:      20,
				},
			},
			want:    Request{},
			wantErr: true,
			error:   ErrorDateFormatWrong,
		},
		{
			name: "error when Nights is zero",
			args: args{
				req: RequestAPI{
					RequestID:   "id-test",
					CheckIn:     "2022-02-01",
					Nights:      0,
					SellingRate: 200,
					Margin:      20,
				},
			},
			want:    Request{},
			wantErr: true,
			error:   ErrorNightsCantBeZero,
		},
		{
			name: "error when SellingRate is zero",
			args: args{
				req: RequestAPI{
					RequestID:   "id-test",
					CheckIn:     "2022-02-01",
					Nights:      23,
					SellingRate: 0,
					Margin:      20,
				},
			},
			want:    Request{},
			wantErr: true,
			error:   ErrorSellingRateCantBeZero,
		},
		{
			name: "error when id is empty",
			args: args{
				req: RequestAPI{
					CheckIn:     "2022-02-01",
					Nights:      23,
					SellingRate: 0,
					Margin:      20,
				},
			},
			want:    Request{},
			wantErr: true,
			error:   ErrorIDCantBeEmpty,
		},
		{
			name: "error when margin is empty",
			args: args{
				req: RequestAPI{
					RequestID:   "id-test",
					CheckIn:     "2022-02-01",
					Nights:      23,
					SellingRate: 0,
				},
			},
			want:    Request{},
			wantErr: true,
			error:   ErrorMarginCantBeZero,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RequestFromRequestAPI(tt.args.req)
			if (err != nil) != tt.wantErr && err != tt.error {
				t.Errorf("RequestFromRequestAPI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestFromRequestAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

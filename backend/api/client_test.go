package api

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_ValidateInputFile(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		want    Input
		wantErr bool
	}{
		{
			name: "success_test",
			file: "./jsontests/success.json",
			want: Input{
				AdCampaignId: 1,
				CustomerId:   2,
				GameName:     "halo",
				ImageName:    "haloImage1",
				ValidAccount: true,
			},
			wantErr: false,
		},
		{
			name: "fail_on_bool_test",
			file: "./jsontests/failOnBool.json",
			want: Input{
				AdCampaignId: 1,
				CustomerId:   2,
				GameName:     "halo",
				ImageName:    "haloImage1",
				ValidAccount: false,
			},
			wantErr: true,
		},
		{
			name: "fail_on_int_test",
			file: "./jsontests/failOnInt.json",
			want: Input{
				AdCampaignId: 0,
				CustomerId:   0,
				GameName:     "halo",
				ImageName:    "haloImage1",
				ValidAccount: true,
			},
			wantErr: true,
		},
		{
			name: "fail_on_string_test",
			file: "./jsontests/failOnString.json",
			want: Input{
				AdCampaignId: 1,
				CustomerId:   2,
				GameName:     "",
				ImageName:    "",
				ValidAccount: true,
			},
			wantErr: true,
		},
		{
			name:    "fail_on_invalid_json_test",
			file:    "./jsontests/invalid.json",
			want:    Input{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ValidateInputFile(tt.file)
			assert.Equal(t, tt.want, got)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("error don't match, expect error: %t, got: %t", gotErr != nil, tt.wantErr)
			}
		})
	}
}

func Test_ValidateInputRawJson(t *testing.T) {
	tests := []struct {
		name    string
		rawJson string
		want    Input
		wantErr bool
	}{
		{
			name:    "success_test",
			rawJson: `{"adCampaignId": 1, "customerID": 2, "gameName": "halo", "imageName": "haloImage1", "validAccount": true}`,
			want: Input{
				AdCampaignId: 1,
				CustomerId:   2,
				GameName:     "halo",
				ImageName:    "haloImage1",
				ValidAccount: true,
			},
			wantErr: false,
		},
		{
			name:    "fail_on_bool_test",
			rawJson: `{"adCampaignId": 1, "customerID": 2, "gameName": "halo", "imageName": "haloImage1", "validAccount": "true"}`,
			want: Input{
				AdCampaignId: 1,
				CustomerId:   2,
				GameName:     "halo",
				ImageName:    "haloImage1",
				ValidAccount: false,
			},
			wantErr: true,
		},
		{
			name:    "fail_on_int_test",
			rawJson: `{"adCampaignId": "1", "customerID": "2", "gameName": "halo", "imageName": "haloImage1", "validAccount": true}`,
			want: Input{
				AdCampaignId: 0,
				CustomerId:   0,
				GameName:     "halo",
				ImageName:    "haloImage1",
				ValidAccount: true,
			},
			wantErr: true,
		},
		{
			name:    "fail_on_string_test",
			rawJson: `{"adCampaignId": 1, "customerID": 2, "gameName": 3, "imageName": 4, "validAccount": true}`,
			want: Input{
				AdCampaignId: 1,
				CustomerId:   2,
				GameName:     "",
				ImageName:    "",
				ValidAccount: true,
			},
			wantErr: true,
		},
		{
			name:    "fail_on_invalid_json_test",
			rawJson: `"adCampaignId": 1, "customerID": 2, "gameName": "halo", "imageName": "haloImage1", "validAccount": "true"`,
			want:    Input{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ValidateInputRawJson(tt.rawJson)
			assert.Equal(t, tt.want, got)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("error don't match, expect error: %t, got: %t", gotErr != nil, tt.wantErr)
			}
		})
	}
}

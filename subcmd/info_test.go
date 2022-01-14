package subcmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestInfo(t *testing.T) {
	assert := assert.New(t)

	info := &subcmd.InfoCmd{
		Path: "//1",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGetFromPageNum = func(postNum int) (*model.Post, error) {
		assert.Equal(1, postNum)

		return &model.Post{
			Category: "foo/bar",
			Name:     "name",
			BodyMd:   "body",
			BodyHTML: "html",
		}, nil
	}

	err := info.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`{
  "number": 0,
  "name": "name",
  "full_name": "",
  "wip": false,
  "created_at": "0001-01-01T00:00:00Z",
  "message": "",
  "url": "",
  "updated_at": "0001-01-01T00:00:00Z",
  "tags": null,
  "category": "foo/bar",
  "revision_number": 0,
  "created_by": {
    "myself": false,
    "name": "",
    "screen_name": "",
    "icon": ""
  },
  "updated_by": {
    "myself": false,
    "name": "",
    "screen_name": "",
    "icon": ""
  }
}
`, printer.String())
}

func TestInfo_URL(t *testing.T) {
	assert := assert.New(t)

	info := &subcmd.InfoCmd{
		Path: "https://docs.esa.io/posts/1",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGetFromPageNum = func(postNum int) (*model.Post, error) {
		assert.Equal(1, postNum)

		return &model.Post{
			Category: "foo/bar",
			Name:     "name",
			BodyMd:   "body",
			BodyHTML: "html",
		}, nil
	}

	err := info.Run(&kasa.Context{
		Team:   "docs",
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`{
  "number": 0,
  "name": "name",
  "full_name": "",
  "wip": false,
  "created_at": "0001-01-01T00:00:00Z",
  "message": "",
  "url": "",
  "updated_at": "0001-01-01T00:00:00Z",
  "tags": null,
  "category": "foo/bar",
  "revision_number": 0,
  "created_by": {
    "myself": false,
    "name": "",
    "screen_name": "",
    "icon": ""
  },
  "updated_by": {
    "myself": false,
    "name": "",
    "screen_name": "",
    "icon": ""
  }
}
`, printer.String())
}

func TestInfo_Path(t *testing.T) {
	assert := assert.New(t)

	info := &subcmd.InfoCmd{
		Path: "foo/bar/zoo",
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockGet = func(path string) (*model.Post, error) {
		assert.Equal("foo/bar/zoo", path)

		return &model.Post{
			Category: "foo/bar",
			Name:     "name",
			BodyMd:   "body",
			BodyHTML: "html",
		}, nil
	}

	err := info.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`{
  "number": 0,
  "name": "name",
  "full_name": "",
  "wip": false,
  "created_at": "0001-01-01T00:00:00Z",
  "message": "",
  "url": "",
  "updated_at": "0001-01-01T00:00:00Z",
  "tags": null,
  "category": "foo/bar",
  "revision_number": 0,
  "created_by": {
    "myself": false,
    "name": "",
    "screen_name": "",
    "icon": ""
  },
  "updated_by": {
    "myself": false,
    "name": "",
    "screen_name": "",
    "icon": ""
  }
}
`, printer.String())
}

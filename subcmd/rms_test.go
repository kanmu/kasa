package subcmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/subcmd"
)

func TestRms_Dir(t *testing.T) {
	assert := assert.New(t)

	rms := &subcmd.RmCmd{
		Path:   "foo/bar/",
		Search: true,
		Force:  true,
		Page:   1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockSearch = func(query string, postNum int) ([]*model.Post, bool, error) {
		assert.Equal("foo/bar/", query)
		assert.Equal(1, postNum)

		return []*model.Post{
			{
				Number:   1,
				Name:     "zoo",
				Category: "foo/bar",
			},
			{
				Number:   2,
				Name:     "baz",
				Category: "foo/bar",
			},
		}, false, nil
	}

	driver.MockDelete = func(postNum int) error {
		assert.Contains([]int{1, 2}, postNum)

		return nil
	}

	err := rms.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`rm 'foo/bar/zoo'
rm 'foo/bar/baz'
`, printer.String())
}

func TestRms_HasMore(t *testing.T) {
	assert := assert.New(t)

	rms := &subcmd.RmCmd{
		Path:   "foo/bar/",
		Search: true,
		Force:  true,
		Page:   1,
	}

	driver := NewMockDriver(t)
	printer := &MockPrinterImpl{}

	driver.MockSearch = func(query string, postNum int) ([]*model.Post, bool, error) {
		assert.Equal("foo/bar/", query)
		assert.Equal(1, postNum)

		return []*model.Post{
			{
				Number:   1,
				Name:     "zoo",
				Category: "foo/bar",
			},
			{
				Number:   2,
				Name:     "baz",
				Category: "foo/bar",
			},
		}, true, nil
	}

	driver.MockDelete = func(postNum int) error {
		assert.Contains([]int{1, 2}, postNum)

		return nil
	}

	err := rms.Run(&kasa.Context{
		Driver: driver,
		Fmt:    printer,
	})

	assert.NoError(err)

	assert.Equal(`rm 'foo/bar/zoo'
rm 'foo/bar/baz'
(has more pages. current page is 1, try '-p 2')
`, printer.String())
}

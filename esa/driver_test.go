package esa

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/kasa/esa/model"
)

func TestGetFromPageNumOk(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"number": 1,
		"name": "hi!",
		"full_name": "日報/2015/05/09/hi! #api #dev",
		"wip": true,
		"body_md": "# Getting Started",
		"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
		"created_at": "2015-05-09T11:54:50+09:00",
		"message": "Add Getting Started section",
		"url": "https://docs.esa.io/posts/1",
		"updated_at": "2015-05-09T11:54:51+09:00",
		"tags": [
			"api",
			"dev"
		],
		"category": "日報/2015/05/09",
		"revision_number": 1,
		"created_by": {
			"myself": true,
			"name": "Atsuo Fukaya",
			"screen_name": "fukayatsu",
			"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
		},
		"updated_by": {
			"myself": true,
			"name": "Atsuo Fukaya",
			"screen_name": "fukayatsu",
			"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
		},
		"kind": "flow",
		"comments_count": 1,
		"tasks_count": 1,
		"done_tasks_count": 1,
		"stargazers_count": 1,
		"watchers_count": 1,
		"star": true,
		"watch": true
	}`

	httpmock.RegisterResponder(http.MethodGet, "https://api.esa.io/v1/teams/example/posts/1",
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	post, err := driver.GetFromPageNum(1)
	assert.NoError(err)

	expected := &model.Post{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected, post)
}

func TestGetOk(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBodyPost := `{
		"number": 1,
		"name": "hi!",
		"full_name": "日報/2015/05/09/hi! #api #dev",
		"wip": true,
		"body_md": "# Getting Started",
		"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
		"created_at": "2015-05-09T11:54:50+09:00",
		"message": "Add Getting Started section",
		"url": "https://docs.esa.io/posts/1",
		"updated_at": "2015-05-09T11:54:51+09:00",
		"tags": [
			"api",
			"dev"
		],
		"category": "日報/2015/05/09",
		"revision_number": 1,
		"created_by": {
			"myself": true,
			"name": "Atsuo Fukaya",
			"screen_name": "fukayatsu",
			"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
		},
		"updated_by": {
			"myself": true,
			"name": "Atsuo Fukaya",
			"screen_name": "fukayatsu",
			"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
		},
		"kind": "flow",
		"comments_count": 1,
		"tasks_count": 1,
		"done_tasks_count": 1,
		"stargazers_count": 1,
		"watchers_count": 1,
		"star": true,
		"watch": true
	}`

	params := map[string]string{"page": "1", "per_page": "50", "q": `hi! on:"日報/2015/05/09"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, `{"posts":[`+resBodyPost+`]}`))

	driver := NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	post, err := driver.Get("日報/2015/05/09/hi!")
	assert.NoError(err)

	expected := &model.Post{}
	json.Unmarshal([]byte(resBodyPost), expected)
	assert.Equal(expected, post)
}

func TestListOk(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
					"api",
					"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
		],
		"prev_page": null,
		"next_page": null,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "2", "per_page": "50", "q": ` in:"日報/2015/05/09/"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	posts, hasMore, err := driver.List(newGlobPath("日報/2015/05/09/hi!"), 2, true)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts, posts)
	assert.False(hasMore)
}

func TestListHasMore(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
					"api",
					"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
		],
		"prev_page": null,
		"next_page": 2,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "2", "per_page": "50", "q": ` in:"日報/2015/05/09/"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	posts, hasMore, err := driver.List(newGlobPath("日報/2015/05/09/hi!"), 2, true)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts, posts)
	assert.True(hasMore)
}

func TestListNoRecursive(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	resBody := `{
		"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
					"api",
					"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"myself": true,
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
		],
		"prev_page": null,
		"next_page": null,
		"total_count": 1,
		"page": 1,
		"per_page": 20,
		"max_per_page": 100
	}`

	params := map[string]string{"page": "3", "per_page": "50", "q": ` on:"日報/2015/05/09/"`}
	httpmock.RegisterResponderWithQuery(http.MethodGet, "https://api.esa.io/v1/teams/example/posts", params,
		httpmock.NewStringResponder(http.StatusOK, resBody))

	driver := NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	posts, hasMore, err := driver.List(newGlobPath("日報/2015/05/09/hi!"), 3, false)
	assert.NoError(err)

	expected := &model.Posts{}
	json.Unmarshal([]byte(resBody), expected)
	assert.Equal(expected.Posts, posts)
	assert.False(hasMore)
}

func TestDriverDeleteOk(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://api.esa.io/v1/teams/example/posts/1",
		httpmock.NewStringResponder(http.StatusNoContent, ""))

	driver := NewDriver("example", "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	err := driver.Delete(1)
	assert.NoError(err)
}

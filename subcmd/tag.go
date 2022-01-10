package subcmd

import (
	"fmt"
	pathpkg "path"
	"sort"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
	"github.com/winebarrel/kasa/utils"
)

type TagCmd struct {
	Path      string   `arg:"" help:"Post name/Post category/Post tag."`
	Tags      []string `short:"t" required:"" help:"Post tags."`
	Override  bool     `short:"o" help:"Override tags."`
	Force     bool     `short:"f" default:"false" help:"Skip confirmation of files to move."`
	Page      int      `short:"p" default:"1" help:"Page number."`
	Recursive bool     `short:"r" default:"true" negatable:"" help:"Recursively list posts."`
}

func (cmd *TagCmd) Run(ctx *kasa.Context) error {
	posts, hasMore, err := ctx.Driver.ListOrTagSearch(cmd.Path, cmd.Page, cmd.Recursive)

	if err != nil {
		return err
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].FullName < posts[j].FullName })

	if err != nil {
		return err
	}

	newPosts := make([]*model.NewPostBody, len(posts))

	for i, v := range posts {
		var tags []string

		if cmd.Override {
			tags = cmd.Tags
		} else {
			tags = append(v.Tags, cmd.Tags...)
		}

		tags = utils.Uniq(tags)

		newPost := &model.NewPostBody{
			Tags: tags,
		}

		newPosts[i] = newPost
	}

	if !cmd.Force {
		for i, oldPost := range posts {
			newPost := newPosts[i]
			tags := "[#" + strings.Join(newPost.Tags, ",#") + "]"
			ctx.Fmt.Printf("tag '%s' '%s'\n", tags, oldPost.FullNameWithoutTags())
		}

		if hasMore {
			ctx.Fmt.Printf("(has more pages. current page is %d, try `-p %d`)\n", cmd.Page, cmd.Page+1)
		}

		approval := prompter.YN("Do you want to tag posts?", false)

		if !approval {
			ctx.Fmt.Println("Tagging cancelled.")
			return nil
		}
	}

	for i, oldPost := range posts {
		newPost := newPosts[i]
		tags := "[#" + strings.Join(newPost.Tags, ",#") + "]"

		if cmd.Force {
			ctx.Fmt.Printf("tag '%s' '%s'\n", tags, oldPost.FullNameWithoutTags())
		}

		url, err := ctx.Driver.Post(newPost, oldPost.Number)

		if err != nil {
			return fmt.Errorf("failed to tag '%s':%w", oldPost.FullNameWithoutTags(), err)
		}

		urlDir := pathpkg.Dir(url)
		ctx.Fmt.Printf("%-*s  %s  %s\n", len(urlDir)+9, url, oldPost.FullNameWithoutTags(), tags)
	}

	if hasMore {
		ctx.Fmt.Printf("(has more pages. current page is %d, try `-p %d`)\n", cmd.Page, cmd.Page+1)
	}

	return nil
}

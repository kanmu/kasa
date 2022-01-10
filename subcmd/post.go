package subcmd

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa/model"
)

type PostCmd struct {
	Name     string   `short:"n" help:"Post title."`
	Body     string   `short:"b" help:"Post body file."`
	PostNum  int      `arg:"" optional:"" help:"Post number to update."`
	Tags     []string `short:"t" help:"Post tags."`
	Category string   `short:"c" help:"Post category."`
	WIP      bool     `default:"false" negatable:"" help:"Post as WIP."`
	Message  string   `short:"m" help:"Post message."`
}

func (cmd *PostCmd) Run(ctx *kasa.Context) error {
	if cmd.PostNum == 0 {
		if cmd.Name == "" {
			return errors.New("missing flags: --name=STRING")
		}

		if cmd.Body == "" {
			return errors.New("missing flags: --body=STRING")
		}
	}

	var bodyMd []byte

	if cmd.Body != "" {
		var file io.ReadCloser

		if cmd.Body == "-" {
			file = os.Stdin
		} else {
			var err error
			file, err = os.OpenFile(cmd.Body, os.O_RDONLY, 0)

			if err != nil {
				return err
			}

			defer file.Close()
		}

		var err error
		bodyMd, err = ioutil.ReadAll(file)

		if err != nil {
			return err
		}
	} else {
		bodyMd = []byte{}
	}

	msg := cmd.Message

	if msg == "" {
		msg = "Posted on " + time.Now().Format(time.RFC3339)
	}

	newPost := &model.NewPostBody{
		Name:     cmd.Name,
		BodyMd:   string(bodyMd),
		Tags:     cmd.Tags,
		Category: cmd.Category,
		WIP:      cmd.WIP,
		Message:  msg,
	}

	url, err := ctx.Driver.Post(newPost, cmd.PostNum)

	if err != nil {
		return err
	}

	ctx.Fmt.Println(url)

	return nil
}

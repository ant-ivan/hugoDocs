// Copyright 2023 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/bep/simplecobra"
	"github.com/gohugoio/hugo/common/hugo"
	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func newGenCommand() *genCommand {
	var (
		// Flags.
		gendocdir string
		genmandir string

		// Chroma flags.
		style          string
		highlightStyle string
		linesStyle     string
	)

	newChromaStyles := func() simplecobra.Commander {
		return &simpleCommand{
			name:  "chromastyles",
			short: "Generate CSS stylesheet for the Chroma code highlighter",
			long: `Generate CSS stylesheet for the Chroma code highlighter for a given style. This stylesheet is needed if markup.highlight.noClasses is disabled in config.

See https://xyproto.github.io/splash/docs/all.html for a preview of the available styles`,

			run: func(ctx context.Context, cd *simplecobra.Commandeer, r *rootCommand, args []string) error {
				builder := styles.Get(style).Builder()
				if highlightStyle != "" {
					builder.Add(chroma.LineHighlight, highlightStyle)
				}
				if linesStyle != "" {
					builder.Add(chroma.LineNumbers, linesStyle)
				}
				style, err := builder.Build()
				if err != nil {
					return err
				}
				formatter := html.New(html.WithAllClasses(true))
				formatter.WriteCSS(os.Stdout, style)
				return nil
			},
			withc: func(cmd *cobra.Command) {
				cmd.PersistentFlags().StringVar(&style, "style", "friendly", "highlighter style (see https://xyproto.github.io/splash/docs/)")
				cmd.PersistentFlags().StringVar(&highlightStyle, "highlightStyle", "bg:#ffffcc", "style used for highlighting lines (see https://github.com/alecthomas/chroma)")
				cmd.PersistentFlags().StringVar(&linesStyle, "linesStyle", "", "style used for line numbers (see https://github.com/alecthomas/chroma)")
			},
		}
	}

	newMan := func() simplecobra.Commander {
		return &simpleCommand{
			name:  "man",
			short: "Generate man pages for the Hugo CLI",
			long: `This command automatically generates up-to-date man pages of Hugo's
	command-line interface.  By default, it creates the man page files
	in the "man" directory under the current directory.`,

			run: func(ctx context.Context, cd *simplecobra.Commandeer, r *rootCommand, args []string) error {
				header := &doc.GenManHeader{
					Section: "1",
					Manual:  "Hugo Manual",
					Source:  fmt.Sprintf("Hugo %s", hugo.CurrentVersion),
				}
				if !strings.HasSuffix(genmandir, helpers.FilePathSeparator) {
					genmandir += helpers.FilePathSeparator
				}
				if found, _ := helpers.Exists(genmandir, hugofs.Os); !found {
					r.Println("Directory", genmandir, "does not exist, creating...")
					if err := hugofs.Os.MkdirAll(genmandir, 0777); err != nil {
						return err
					}
				}
				cd.CobraCommand.Root().DisableAutoGenTag = true

				r.Println("Generating Hugo man pages in", genmandir, "...")
				doc.GenManTree(cd.CobraCommand.Root(), header, genmandir)

				r.Println("Done.")

				return nil
			},
			withc: func(cmd *cobra.Command) {
				cmd.PersistentFlags().StringVar(&genmandir, "dir", "man/", "the directory to write the man pages.")
				// For bash-completion
				cmd.PersistentFlags().SetAnnotation("dir", cobra.BashCompSubdirsInDir, []string{})
			},
		}
	}

	newGen := func() simplecobra.Commander {
		const gendocFrontmatterTemplate = `---
title: "%s"
slug: %s
url: %s
---
`

		return &simpleCommand{
			name:  "doc",
			short: "Generate Markdown documentation for the Hugo CLI.",
			long: `Generate Markdown documentation for the Hugo CLI.
			This command is, mostly, used to create up-to-date documentation
	of Hugo's command-line interface for https://gohugo.io/.
	
	It creates one Markdown file per command with front matter suitable
	for rendering in Hugo.`,
			run: func(ctx context.Context, cd *simplecobra.Commandeer, r *rootCommand, args []string) error {
				cd.CobraCommand.VisitParents(func(c *cobra.Command) {
					// Disable the "Auto generated by spf13/cobra on DATE"
					// as it creates a lot of diffs.
					c.DisableAutoGenTag = true
				})
				if !strings.HasSuffix(gendocdir, helpers.FilePathSeparator) {
					gendocdir += helpers.FilePathSeparator
				}
				if found, _ := helpers.Exists(gendocdir, hugofs.Os); !found {
					r.Println("Directory", gendocdir, "does not exist, creating...")
					if err := hugofs.Os.MkdirAll(gendocdir, 0777); err != nil {
						return err
					}
				}
				prepender := func(filename string) string {
					name := filepath.Base(filename)
					base := strings.TrimSuffix(name, path.Ext(name))
					url := "/commands/" + strings.ToLower(base) + "/"
					return fmt.Sprintf(gendocFrontmatterTemplate, strings.Replace(base, "_", " ", -1), base, url)
				}

				linkHandler := func(name string) string {
					base := strings.TrimSuffix(name, path.Ext(name))
					return "/commands/" + strings.ToLower(base) + "/"
				}
				r.Println("Generating Hugo command-line documentation in", gendocdir, "...")
				doc.GenMarkdownTreeCustom(cd.CobraCommand.Root(), gendocdir, prepender, linkHandler)
				r.Println("Done.")

				return nil
			},
			withc: func(cmd *cobra.Command) {
				cmd.PersistentFlags().StringVar(&gendocdir, "dir", "/tmp/hugodoc/", "the directory to write the doc.")
				// For bash-completion
				cmd.PersistentFlags().SetAnnotation("dir", cobra.BashCompSubdirsInDir, []string{})
			},
		}

	}

	return &genCommand{
		commands: []simplecobra.Commander{
			newChromaStyles(),
			newGen(),
			newMan(),
		},
	}

}

type genCommand struct {
	rootCmd *rootCommand

	commands []simplecobra.Commander
}

func (c *genCommand) Commands() []simplecobra.Commander {
	return c.commands
}

func (c *genCommand) Name() string {
	return "gen"
}

func (c *genCommand) Run(ctx context.Context, cd *simplecobra.Commandeer, args []string) error {
	return nil
}

func (c *genCommand) Init(cd *simplecobra.Commandeer) error {
	cmd := cd.CobraCommand
	cmd.Short = "A collection of several useful generators."
	return nil
}

func (c *genCommand) PreRun(cd, runner *simplecobra.Commandeer) error {
	c.rootCmd = cd.Root.Command.(*rootCommand)
	return nil
}

package nanaparty

import (
	"github.com/mkfsn/mizukinana/cmd/mizukinana/nanaparty/biography"
	"github.com/mkfsn/mizukinana/cmd/mizukinana/nanaparty/blog"
	"github.com/mkfsn/mizukinana/cmd/mizukinana/nanaparty/discography"
	"github.com/mkfsn/mizukinana/cmd/mizukinana/nanaparty/news"
	"github.com/mkfsn/mizukinana/cmd/mizukinana/nanaparty/schedule"
	"github.com/mkfsn/mizukinana/cmd/mizukinana/nanaparty/top"
	"github.com/spf13/cobra"
)

func NewCmdNanaparty() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nanaparty",
		Short: "Crawl the information from Nana-Party",
		Long:  `This is a collection of function for crawling the information from NanaParty website.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Usage()
		},
	}
	cmd.AddCommand(news.NewCmdNews())
	cmd.AddCommand(blog.NewCmdBlog())
	cmd.AddCommand(top.NewCmdTop())
	cmd.AddCommand(biography.NewCmdBiography())
	cmd.AddCommand(schedule.NewCmdSchedule())
	cmd.AddCommand(discography.NewCmdDiscography())
	return cmd
}

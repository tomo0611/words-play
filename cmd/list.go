/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/tomo0611/words-play/database"

	"github.com/spf13/cobra"
)

var limit int
var offset int

type WordCollection struct {
	Words []Word `json:"words"`
}

type Word struct {
	Word_en   string `json:"word_en"`
	Word_ja   string `json:"word_ja"`
	Audio_url string `json:"audio_url"`
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "単語の一覧を表示します。",
	Long:  `単語の一覧を表示します。`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		db, err := config.getDatabase()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		queries := database.New(db)

		// list words
		words, err := queries.ListWords(ctx, database.ListWordsParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		for _, word := range words {
			fmt.Printf("%s: %s\n", word.WordEn, word.WordJa)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().IntVarP(&limit, "limit", "l", 25, "表示する単語の数")
	listCmd.Flags().IntVarP(&offset, "offset", "o", 0, "表示する単語のオフセット")
}

/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

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
		raw, err := os.ReadFile("./private/words_out.json")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		var words WordCollection
		json.Unmarshal(raw, &words)
		for i, word := range words.Words {
			if i < offset {
				continue
			}
			if i >= offset+limit {
				break
			}
			fmt.Printf("%s: %s\n", word.Word_en, word.Word_ja)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().IntVarP(&limit, "limit", "l", 10, "表示する単語の数")
	listCmd.Flags().IntVarP(&offset, "offset", "o", 0, "表示する単語のオフセット")
}

/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/

package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
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

		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		c := mysql.Config{
			DBName:    "WORDS_PLAY",
			User:      "wordsplay",
			Passwd:    "password",
			Addr:      "localhost:3306",
			Net:       "tcp",
			ParseTime: true,
			// 'mysql_native_password': this user requires mysql native password authentication
			AllowNativePasswords: true,
			Collation:            "utf8mb4_unicode_ci",
			Loc:                  jst,
		}

		db, err := sql.Open("mysql", c.FormatDSN())
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

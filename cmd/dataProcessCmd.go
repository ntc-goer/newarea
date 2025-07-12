package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"newarea/pkg/fileReader"
)

func dataProcessCmd() *cobra.Command {
	return &cobra.Command{
		Use: "data-process",
		RunE: func(cmd *cobra.Command, args []string) error {
			records, err := fileReader.Read("./data/province.csv")
			if err != nil {
				logrus.Errorf("Read file fail %v", err)
				return err
			}
			for i, row := range records {
				if i == 0 {
					continue
				}
				fmt.Println(row[0], row[1])
			}
			return nil
		},
	}
}

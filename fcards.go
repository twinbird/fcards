// コマンドライン単語帳
package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

const (
	ConfigurationFileName = ".fcards"
)

type Config struct {
	LastRefRowIdx int // 最後に表示したCardのTSVの行
}

type Card struct {
	en string // 表面
	jp string // 裏面
}

var (
	cards []*Card
)

func main() {
	var (
		reverse  bool
		filename string
		reset bool
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "コマンドラインの単語帳\nTSVファイルは「表面」,「裏面」の2列で作成してください\n")
		flag.PrintDefaults()
	}
	flag.StringVar(&filename, "f", "./fcards.tsv", "単語帳のファイル名を指定する")
	flag.BoolVar(&reverse, "r", false, "表面と裏面を反転して始める")
	flag.BoolVar(&reset, "reset", false, "単語帳を最初から始める")
	flag.Parse()

	// 設定ファイルをロード
	config := &Config{}
	if err := config.load(); err != nil {
		log.Fatal(err)
	}

	// Resetする
	if reset == true {
		config.LastRefRowIdx = 0
	}

	// 単語帳ファイルのロード
	if b := fileExists(filename); b == false {
		fmt.Printf("%sは見つかりません", filename)
		return
	}
	if err := csvLoad(filename); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Type ':q' for save and quit.")

	// メインループ
	if reverse == true {
		config.LastRefRowIdx = jpToEnLoop(config.LastRefRowIdx)
	} else {
		config.LastRefRowIdx = enToJpLoop(config.LastRefRowIdx)
	}

	// 設定ファイルの保存
	if err := config.save(); err != nil {
		log.Fatal(err)
	}
}

func jpToEnLoop(idx int) int {
	var answer string
	var last int
	for i, v := range cards {
		if i < idx {
			continue
		}
		last = i
		fmt.Printf("\n%d / %d\n", i+1, len(cards))
		fmt.Printf("%s\n", v.jp)
		fmt.Printf("> ")
		fmt.Scanln(&answer)
		if answer == ":q" {
			break
		}
		fmt.Printf("%s\n", v.en)
	}
	return last
}

func enToJpLoop(idx int) int {
	var answer string
	var last int
	for i, v := range cards {
		if i < idx {
			continue
		}
		last = i
		fmt.Printf("\n%d / %d\n", i+1, len(cards))
		fmt.Printf("%s\n", v.en)
		fmt.Printf("> ")
		fmt.Scanln(&answer)
		if answer == ":q" {
			break
		}
		fmt.Printf("%s\n", v.jp)
	}
	return last
}

func csvLoad(filename string) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	r.Comma = '\t'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, v := range records {
		card := &Card{v[0], v[1]}
		cards = append(cards, card)
	}
	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func configFilePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	s := filepath.Join(usr.HomeDir, ConfigurationFileName)
	return s, nil
}

func (c *Config) load() error {
	p, err := configFilePath()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) save() error {
	p, err := configFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(p, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

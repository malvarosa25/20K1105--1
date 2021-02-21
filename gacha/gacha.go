package gacha

import (
	"math/rand"
	"time"
)

func init() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())
}

func DrawN(p *Player, n int) ([]*Card, map[Rarity]int) {
	p.draw(n)

	results := make([]*Card, n)
	summary := make(map[Rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw()
		summary[results[i].Rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() *Card {
	num := rand.Intn(100)
	switch {
	case num < 80: // num の値により、Rarity を分ける
		num = rand.Intn(100) // 再度 num の値を乱数で決め代入する
		switch {             // 新しい num の値により、モンスターの種類を決める
		case num < 25:
			return &Card{Rarity: RarityN, Name: "スライム"}
		case num < 50:
			return &Card{Rarity: RarityN, Name: "ドラキー"}
		case num < 75:
			return &Card{Rarity: RarityN, Name: "ベビーパンサー"}
		default:
			return &Card{Rarity: RarityN, Name: "アルミラージ"}
		}

	case num < 95: // num の値により、Rarity を分ける
		num = rand.Intn(90) // 再度 num の値を乱数で決め代入する
		switch {            // 新しい num の値により、モンスターの種類を決める
		case num < 30:
			return &Card{Rarity: RarityR, Name: "オーク"}
		case num < 60:
			return &Card{Rarity: RarityR, Name: "ナイトキャット"}
		default:
			return &Card{Rarity: RarityR, Name: "モーモン"}
		}

	case num < 99: // num の値により、Rarity を分ける
		num = rand.Intn(90) // 再度 num の値を乱数で決め代入する
		switch {            // 新しい num の値により、モンスターの種類を決める
		case num < 30:
			return &Card{Rarity: RaritySR, Name: "ドラゴン"}
		case num < 60:
			return &Card{Rarity: RaritySR, Name: "キラーマシン"}
		default:
			return &Card{Rarity: RaritySR, Name: "ごくらくちょう"}
		}
	default: // num の値により、Rarity を分ける
		num = rand.Intn(100) // 再度 num の値を乱数で決め代入する
		switch {             // 新しい num の値により、モンスターの種類を決める
		case num < 25:
			return &Card{Rarity: RarityXR, Name: "イフリート"}
		case num < 50:
			return &Card{Rarity: RarityN, Name: "妖女イシュダル"}
		case num < 75:
			return &Card{Rarity: RarityN, Name: "はぐれメタル"}
		default:
			return &Card{Rarity: RarityXR, Name: "サンディ"}
		}
	}
}

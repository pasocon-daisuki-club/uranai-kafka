package uranai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	m "github.com/anaregdesign/msproto/go/msp/azure/openai/chat/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"net/http"
	"time"
)

const prompt = `
# あなたの役割
* あなたは占い師です。12星座それぞれについて占いを行います。
* 12星座は以下の通りです。
	- おひつじ座
	- おうし座
	- ふたご座
	- かに座
	- しし座
	- おとめ座
	- てんびん座
	- さそり座
	- いて座
	- やぎ座
	- みずがめ座
	- うお座
* それぞれの星座について、以下の項目を占ってください。またそれぞれの項目は他の星座と重複しないようにしてください。
	- 総合順位：　総合的に1位から12位のランキングを重複なくつけてください。
	- ラッキーアイテム: アクセサリーや食べ物、ペット、日用品など身近なものをひとつ選んでください
	- ラッキーカラー： 色の名前
	- ラッキーサービス: マイクロソフト社の製品やサービスの中から適当に選んでください
* それぞれの項目について、以下の3から10の整数で点数をつけてください。
	- 仕事運=CareerLuck
	- 恋愛運=LoveLuck
	- 健康運=HealthLuck
* 仕事運、恋愛運、健康運のうちどれかを今日の運勢を解説し、今日起こりそうな事について簡単なエピソードを教えてください。またその状況をラッキーアイテムもしくはラッキーサービスを使って改善するアイディアを教えてください

# 出力形式
* 出力形式はJSON形式とし、それ以外は何も出力しないでください
* 出力にはバッククオートを含めないでください。

# 出力例
{
	"results": [
		{
			"rank": 1,
			"name": "おうし座",
			"lucky_item": "画鋲",
			"lucky_color": "ターコイズブルー",
			"lucky_service": "Microsoft EntraID",
			"career_luck": 6,
			"love_luck": 8,
			"health_luck": 10,
			"description": "自重運。誰にでもいい顔をしすぎて、恋人が呆れてしまいそう。あなただけが特別だということをわかりやすく伝えましょう。"
		}
	]
}
`

type FortuneTeller struct {
	client *Client
}

func (t *FortuneTeller) Listen(ctx context.Context) (*ResultSet, error) {
	request := &m.CompletionRequest{}
	request.Messages = []*m.CompletionRequest_Message{
		{
			Role:    "user",
			Content: prompt,
		},
	}
	request.Temperature = 0.7

	response, err := t.client.Get(ctx, request)
	if err != nil {
		return nil, err
	}

	resultString := response.Choices[0].Message.Content
	resultSet := ResultSet{}
	err = json.Unmarshal([]byte(resultString), &resultSet)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	for i := range resultSet.Results {
		resultSet.Results[i].CreatedAt = now
	}
	return &resultSet, nil
}

type Client struct {
	httpClient     *http.Client
	resourceName   string
	deploymentName string
	apiVersion     string
	accessToken    string
}

func (c *Client) endpoint() string {
	return fmt.Sprintf("https://%s.openai.azure.com/openai/deployments/%s/chat/completions?api-version=%s", c.resourceName, c.deploymentName, c.apiVersion)
}

func (c *Client) header() http.Header {
	header := http.Header{}
	header.Add("api-key", c.accessToken)
	header.Add("Content-Type", "application/json")
	return header
}

func (c *Client) Get(ctx context.Context, request *m.CompletionRequest) (*m.CompletionResponse, error) {
	body, err := protojson.Marshal(request)
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequestWithContext(ctx, "POST", c.endpoint(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpRequest.Header = c.header()

	httpResponse, err := c.httpClient.Do(httpRequest)
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode == http.StatusOK {
		response := &m.CompletionResponse{}
		err := protojson.Unmarshal(responseBody, response)
		if err != nil {
			return nil, err
		}
		return response, nil
	} else {
		fmt.Printf(string(responseBody))
		return nil, fmt.Errorf("unexpected status code: %d", httpResponse.StatusCode)
	}
}

package model

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type WebhookRepository interface {
	WebhookEventHandler(c echo.Context, reqBody, resBody []byte)
}

type webhookRepository struct {
	secret    string
	channelId string
	id        string
}

func NewWebhookRepository(secret string, channelId string, id string) WebhookRepository {
	return &webhookRepository{
		secret:    secret,
		channelId: channelId,
		id:        id,
	}
}

func (repo *webhookRepository) WebhookEventHandler(c echo.Context, reqBody, resBody []byte) {
	resApp := new(Application)
	err := json.Unmarshal(resBody, resApp)
	if err != nil {
		return
	}
	var content string

	content = "## 申請書が作成されました" + "\n"
	content += fmt.Sprintf("### [%s](%s/applications/%s)", resApp.LatestApplicationsDetail.Title, "https://jomon.trap.jp", resApp.ID) + "\n"
	content += fmt.Sprintf("- 支払日: %s", resApp.LatestApplicationsDetail.PaidAt.PaidAt.Format("2006-01-02")) + "\n"
	content += fmt.Sprintf("- 支払金額: %s", strconv.Itoa(resApp.LatestApplicationsDetail.Amount)) + "\n"
	content += "\n"
	content += resApp.LatestApplicationsDetail.Remarks

	_ = RequestWebhook(content, repo.secret, repo.channelId, repo.id, 1)
}

func RequestWebhook(message, secret, channelID, webhookID string, embed int) error {
	u, err := url.Parse("https://q.trap.jp/api/v3/webhooks")
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, webhookID)
	query := u.Query()
	query.Set("embed", strconv.Itoa(embed))
	u.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(message))
	if err != nil {
		return err
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	req.Header.Set("X-TRAQ-Signature", calcSignature(message, secret))
	if channelID != "" {
		req.Header.Set("X-TRAQ-Channel-Id", channelID)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		return errors.New(http.StatusText(res.StatusCode))
	}
	return nil

}

func calcSignature(message, secret string) string {
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
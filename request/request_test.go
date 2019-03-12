package request

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormContentType(t *testing.T) {
	req, err := Post("http:....").Form("Key", "Vaue").Form("Key1", "Value").makeRequest()
	require.NoError(t, err)

	require.Equal(t, ContentTypeForm, req.Header.Get(headerContentType))
}

func TestPapagoSMT(t *testing.T) {
	type papagoSMTResp struct {
		Message struct {
			Type    string `json:"@type"`
			Service string `json:"@service"`
			Version string `json:"@version"`
			Result  struct {
				TranslatedText string `json:"translatedText"`
				SrcLangType    string `json:"srcLangType"`
			} `json:"result"`
		} `json:"message"`
	}

	resp, err := Post("https://openapi.naver.com/v1/language/translate").
		Headers(map[string]string{
			"X-Naver-Client-Id":     os.Getenv("NAVER_CLIENT_ID"),
			"X-Naver-Client-Secret": os.Getenv("NAVER_CLIENT_SECRET"),
		}).
		Forms(map[string]string{
			"source": "ko",
			"target": "en",
			"text":   "만나서 반갑습니다.",
		}).
		Do()
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	r := &papagoSMTResp{}
	require.NoError(t, resp.JSON(r))
	defer resp.Body.Close()
	require.Equal(t, "Nice to meet you.", r.Message.Result.TranslatedText)
}

func TestGithubGet(t *testing.T) {
	resp, err := Get("https://api.github.com").Do()
	require.NoError(t, err)

	r := make(map[string]string)
	require.NoError(t, resp.JSON(&r))
	defer resp.Body.Close()
	require.Equal(t, "https://api.github.com/hub", r["hub_url"])
}

func TestGoogleCustomSearch(t *testing.T) {
	key, ok := os.LookupEnv("GOOGLE_API_KEY")
	if !ok {
		t.Skip("GOOGLE_API_KEY missed")
	}

	resp, err := Get("https://www.googleapis.com/customsearch/v1").
		Param("key", key).
		Param("cx", os.Getenv("GOOGLE_cx")).
		Param("q", "request").
		Do()

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// https://developers.google.com/custom-search/json-api/v1/reference/cse/list#response
	var response struct {
		Kind string `json:"kind"`
		URL  struct {
			Type     string `json:"type"`
			Template string `json:"template"`
		} `json:"url"`
		Items []struct {
			Kind  string `json:"kind"`
			Title string `json:"title"`
			Link  string `json:"link"`
		} `json:"items"`
	}
	require.NoError(t, resp.JSON(&response))
	defer resp.Body.Close()
	require.True(t, len(response.Items) > 0)
	log.Printf("link: %s", response.Items[0].Link)
}

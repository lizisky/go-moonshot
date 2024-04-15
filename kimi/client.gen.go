package kimi

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"reflect"
	"sync"
	"text/template"
	"time"
)

const (
	CallerListModels                 = "ListModels"
	CallerEstimateTokenCount         = "EstimateTokenCount"
	CallerCheckBalance               = "CheckBalance"
	CallerCreateChatCompletion       = "CreateChatCompletion"
	CallerCreateChatCompletionStream = "CreateChatCompletionStream"
	CallerUploadFile                 = "UploadFile"
	CallerRetrieveFileContent        = "RetrieveFileContent"
)

type Moonshot struct {
	URL    string
	KEY    string
	CLIENT *http.Client
	LOG    func(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
}

func (m Moonshot) BaseUrl() string      { return m.URL }
func (m Moonshot) Key() string          { return m.KEY }
func (m Moonshot) Client() *http.Client { return m.CLIENT }

func (m Moonshot) Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration) {
	if m.LOG != nil {
		m.LOG(ctx, caller, request, response, elapse)
	}
}

func NewClient[C Caller](Client C) Client[C] {
	return &implClient[C]{__Client: Client}
}

// func NewClient(client interface{}) {
// 	return newClient_impl[Moonshot](client)
// }

// func newClient_impl[C Caller](Client C) Client[C] {
// 	return &implClient[C]{__Client: Client}
// }

type implClient[C Caller] struct {
	__Client C
}

var (
	addrTmplListModels                   = template.Must(template.New("AddressListModels").Parse("{{ $.Client.BaseUrl }}/models"))
	headerTmplListModels                 = template.Must(template.New("HeaderListModels").Parse("Authorization: Bearer {{ $.Client.Key }}\r\n\r\n"))
	addrTmplEstimateTokenCount           = template.Must(template.New("AddressEstimateTokenCount").Parse("{{ $.Client.BaseUrl }}/tokenizers/estimate-token-count"))
	headerTmplEstimateTokenCount         = template.Must(template.New("HeaderEstimateTokenCount").Parse("Authorization: Bearer {{ $.Client.Key }}\r\n\r\n{{ $.request.ToJSON }}"))
	addrTmplCheckBalance                 = template.Must(template.New("AddressCheckBalance").Parse("{{ $.Client.BaseUrl }}/users/me/balance"))
	headerTmplCheckBalance               = template.Must(template.New("HeaderCheckBalance").Parse("Authorization: Bearer {{ $.Client.Key }}\r\n\r\n"))
	addrTmplCreateChatCompletion         = template.Must(template.New("AddressCreateChatCompletion").Parse("{{ $.Client.BaseUrl }}/chat/completions"))
	headerTmplCreateChatCompletion       = template.Must(template.New("HeaderCreateChatCompletion").Parse("Content-Type: application/json\r\nAuthorization: Bearer {{ $.Client.Key }}\r\n\r\n{{ $.request.ToJSON }}"))
	addrTmplCreateChatCompletionStream   = template.Must(template.New("AddressCreateChatCompletionStream").Parse("{{ $.Client.BaseUrl }}/chat/completions"))
	headerTmplCreateChatCompletionStream = template.Must(template.New("HeaderCreateChatCompletionStream").Parse("Content-Type: application/json\r\nAuthorization: Bearer {{ $.Client.Key }}\r\n\r\n{{ $.request.ToJSON }}"))
	addrTmplUploadFile                   = template.Must(template.New("AddressUploadFile").Parse("{{ $.Client.BaseUrl }}/files"))
	headerTmplUploadFile                 = template.Must(template.New("HeaderUploadFile").Parse("Content-Type: {{ $.request.ContentType }}\r\nAuthorization: Bearer {{ $.Client.Key }}\r\n\r\n"))
	addrTmplRetrieveFileContent          = template.Must(template.New("AddressRetrieveFileContent").Parse("{{ $.Client.BaseUrl }}/files/{{ $.fileID }}/content"))
	headerTmplRetrieveFileContent        = template.Must(template.New("HeaderRetrieveFileContent").Parse("Authorization: Bearer {{ $.Client.Key }}\r\n\r\n"))
)

func (__imp *implClient[C]) ListModels(ctx context.Context) (*Models, error) {
	var innerListModels any = __imp.Inner()

	addrListModels := __ClientGetBuffer()
	defer __ClientPutBuffer(addrListModels)
	defer addrListModels.Reset()

	headerListModels := __ClientGetBuffer()
	defer __ClientPutBuffer(headerListModels)
	defer headerListModels.Reset()

	var (
		v0ListModels           = new(Models)
		errListModels          error
		httpResponseListModels *http.Response
		responseListModels     ClientResponseInterface = __imp.response()
	)

	if errListModels = addrTmplListModels.Execute(addrListModels, map[string]any{
		"Client": __imp.Inner(),
		"ctx":    ctx,
	}); errListModels != nil {
		return v0ListModels, fmt.Errorf("error building 'ListModels' url: %w", errListModels)
	}

	if errListModels = headerTmplListModels.Execute(headerListModels, map[string]any{
		"Client": __imp.Inner(),
		"ctx":    ctx,
	}); errListModels != nil {
		return v0ListModels, fmt.Errorf("error building 'ListModels' header: %w", errListModels)
	}
	bufReaderListModels := bufio.NewReader(headerListModels)
	mimeHeaderListModels, errListModels := textproto.NewReader(bufReaderListModels).ReadMIMEHeader()
	if errListModels != nil {
		return v0ListModels, fmt.Errorf("error reading 'ListModels' header: %w", errListModels)
	}

	urlListModels := addrListModels.String()
	requestListModels, errListModels := http.NewRequestWithContext(ctx, "GET", urlListModels, http.NoBody)
	if errListModels != nil {
		return v0ListModels, fmt.Errorf("error building 'ListModels' request: %w", errListModels)
	}

	for kListModels, vvListModels := range mimeHeaderListModels {
		for _, vListModels := range vvListModels {
			requestListModels.Header.Add(kListModels, vListModels)
		}
	}

	startListModels := time.Now()

	if httpClientListModels, okListModels := innerListModels.(interface{ Client() *http.Client }); okListModels {
		httpResponseListModels, errListModels = httpClientListModels.Client().Do(requestListModels)
	} else {
		httpResponseListModels, errListModels = http.DefaultClient.Do(requestListModels)
	}

	if logListModels, okListModels := innerListModels.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okListModels {
		logListModels.Log(ctx, "ListModels", requestListModels, httpResponseListModels, time.Since(startListModels))
	}

	if errListModels != nil {
		return v0ListModels, fmt.Errorf("error sending 'ListModels' request: %w", errListModels)
	}

	if httpResponseListModels.StatusCode < 200 || httpResponseListModels.StatusCode > 299 {
		return v0ListModels, __ClientNewResponseError("ListModels", httpResponseListModels)
	}

	if errListModels = responseListModels.FromResponse("ListModels", httpResponseListModels); errListModels != nil {
		return v0ListModels, fmt.Errorf("error converting 'ListModels' response: %w", errListModels)
	}

	addrListModels.Reset()
	headerListModels.Reset()

	if errListModels = responseListModels.Err(); errListModels != nil {
		return v0ListModels, fmt.Errorf("error returned from 'ListModels' response: %w", errListModels)
	}

	if errListModels = responseListModels.ScanValues(v0ListModels); errListModels != nil {
		return v0ListModels, fmt.Errorf("error scanning value from 'ListModels' response: %w", errListModels)
	}

	return v0ListModels, nil
}

func (__imp *implClient[C]) EstimateTokenCount(ctx context.Context, request *EstimateTokenCountRequest) (*EstimateTokenCount, error) {
	var innerEstimateTokenCount any = __imp.Inner()

	addrEstimateTokenCount := __ClientGetBuffer()
	defer __ClientPutBuffer(addrEstimateTokenCount)
	defer addrEstimateTokenCount.Reset()

	headerEstimateTokenCount := __ClientGetBuffer()
	defer __ClientPutBuffer(headerEstimateTokenCount)
	defer headerEstimateTokenCount.Reset()

	var (
		v0EstimateTokenCount           = new(EstimateTokenCount)
		errEstimateTokenCount          error
		httpResponseEstimateTokenCount *http.Response
		responseEstimateTokenCount     ClientResponseInterface = __imp.response()
	)

	if errEstimateTokenCount = addrTmplEstimateTokenCount.Execute(addrEstimateTokenCount, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errEstimateTokenCount != nil {
		return v0EstimateTokenCount, fmt.Errorf("error building 'EstimateTokenCount' url: %w", errEstimateTokenCount)
	}

	if errEstimateTokenCount = headerTmplEstimateTokenCount.Execute(headerEstimateTokenCount, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errEstimateTokenCount != nil {
		return v0EstimateTokenCount, fmt.Errorf("error building 'EstimateTokenCount' header: %w", errEstimateTokenCount)
	}
	bufReaderEstimateTokenCount := bufio.NewReader(headerEstimateTokenCount)
	mimeHeaderEstimateTokenCount, errEstimateTokenCount := textproto.NewReader(bufReaderEstimateTokenCount).ReadMIMEHeader()
	if errEstimateTokenCount != nil {
		return v0EstimateTokenCount, fmt.Errorf("error reading 'EstimateTokenCount' header: %w", errEstimateTokenCount)
	}

	urlEstimateTokenCount := addrEstimateTokenCount.String()
	requestBodyEstimateTokenCount, errEstimateTokenCount := io.ReadAll(bufReaderEstimateTokenCount)
	if errEstimateTokenCount != nil {
		return v0EstimateTokenCount, fmt.Errorf("error reading 'EstimateTokenCount' request body: %w", errEstimateTokenCount)
	}
	requestEstimateTokenCount, errEstimateTokenCount := http.NewRequestWithContext(ctx, "POST", urlEstimateTokenCount, bytes.NewReader(requestBodyEstimateTokenCount))
	if errEstimateTokenCount != nil {
		return v0EstimateTokenCount, fmt.Errorf("error building 'EstimateTokenCount' request: %w", errEstimateTokenCount)
	}

	for kEstimateTokenCount, vvEstimateTokenCount := range mimeHeaderEstimateTokenCount {
		for _, vEstimateTokenCount := range vvEstimateTokenCount {
			requestEstimateTokenCount.Header.Add(kEstimateTokenCount, vEstimateTokenCount)
		}
	}

	startEstimateTokenCount := time.Now()

	if httpClientEstimateTokenCount, okEstimateTokenCount := innerEstimateTokenCount.(interface{ Client() *http.Client }); okEstimateTokenCount {
		httpResponseEstimateTokenCount, errEstimateTokenCount = httpClientEstimateTokenCount.Client().Do(requestEstimateTokenCount)
	} else {
		httpResponseEstimateTokenCount, errEstimateTokenCount = http.DefaultClient.Do(requestEstimateTokenCount)
	}

	if logEstimateTokenCount, okEstimateTokenCount := innerEstimateTokenCount.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okEstimateTokenCount {
		logEstimateTokenCount.Log(ctx, "EstimateTokenCount", requestEstimateTokenCount, httpResponseEstimateTokenCount, time.Since(startEstimateTokenCount))
	}

	if errEstimateTokenCount != nil {
		return v0EstimateTokenCount, fmt.Errorf("error sending 'EstimateTokenCount' request: %w", errEstimateTokenCount)
	}

	if httpResponseEstimateTokenCount.StatusCode < 200 || httpResponseEstimateTokenCount.StatusCode > 299 {
		return v0EstimateTokenCount, __ClientNewResponseError("EstimateTokenCount", httpResponseEstimateTokenCount)
	}

	if errEstimateTokenCount = responseEstimateTokenCount.FromResponse("EstimateTokenCount", httpResponseEstimateTokenCount); errEstimateTokenCount != nil {
		return v0EstimateTokenCount, fmt.Errorf("error converting 'EstimateTokenCount' response: %w", errEstimateTokenCount)
	}

	addrEstimateTokenCount.Reset()
	headerEstimateTokenCount.Reset()

	if errEstimateTokenCount = responseEstimateTokenCount.Err(); errEstimateTokenCount != nil {
		return v0EstimateTokenCount, fmt.Errorf("error returned from 'EstimateTokenCount' response: %w", errEstimateTokenCount)
	}

	if errEstimateTokenCount = responseEstimateTokenCount.ScanValues(v0EstimateTokenCount); errEstimateTokenCount != nil {
		return v0EstimateTokenCount, fmt.Errorf("error scanning value from 'EstimateTokenCount' response: %w", errEstimateTokenCount)
	}

	return v0EstimateTokenCount, nil
}

func (__imp *implClient[C]) CheckBalance(ctx context.Context) (*Balance, error) {
	var innerCheckBalance any = __imp.Inner()

	addrCheckBalance := __ClientGetBuffer()
	defer __ClientPutBuffer(addrCheckBalance)
	defer addrCheckBalance.Reset()

	headerCheckBalance := __ClientGetBuffer()
	defer __ClientPutBuffer(headerCheckBalance)
	defer headerCheckBalance.Reset()

	var (
		v0CheckBalance           = new(Balance)
		errCheckBalance          error
		httpResponseCheckBalance *http.Response
		responseCheckBalance     ClientResponseInterface = __imp.response()
	)

	if errCheckBalance = addrTmplCheckBalance.Execute(addrCheckBalance, map[string]any{
		"Client": __imp.Inner(),
		"ctx":    ctx,
	}); errCheckBalance != nil {
		return v0CheckBalance, fmt.Errorf("error building 'CheckBalance' url: %w", errCheckBalance)
	}

	if errCheckBalance = headerTmplCheckBalance.Execute(headerCheckBalance, map[string]any{
		"Client": __imp.Inner(),
		"ctx":    ctx,
	}); errCheckBalance != nil {
		return v0CheckBalance, fmt.Errorf("error building 'CheckBalance' header: %w", errCheckBalance)
	}
	bufReaderCheckBalance := bufio.NewReader(headerCheckBalance)
	mimeHeaderCheckBalance, errCheckBalance := textproto.NewReader(bufReaderCheckBalance).ReadMIMEHeader()
	if errCheckBalance != nil {
		return v0CheckBalance, fmt.Errorf("error reading 'CheckBalance' header: %w", errCheckBalance)
	}

	urlCheckBalance := addrCheckBalance.String()
	requestCheckBalance, errCheckBalance := http.NewRequestWithContext(ctx, "GET", urlCheckBalance, http.NoBody)
	if errCheckBalance != nil {
		return v0CheckBalance, fmt.Errorf("error building 'CheckBalance' request: %w", errCheckBalance)
	}

	for kCheckBalance, vvCheckBalance := range mimeHeaderCheckBalance {
		for _, vCheckBalance := range vvCheckBalance {
			requestCheckBalance.Header.Add(kCheckBalance, vCheckBalance)
		}
	}

	startCheckBalance := time.Now()

	if httpClientCheckBalance, okCheckBalance := innerCheckBalance.(interface{ Client() *http.Client }); okCheckBalance {
		httpResponseCheckBalance, errCheckBalance = httpClientCheckBalance.Client().Do(requestCheckBalance)
	} else {
		httpResponseCheckBalance, errCheckBalance = http.DefaultClient.Do(requestCheckBalance)
	}

	if logCheckBalance, okCheckBalance := innerCheckBalance.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okCheckBalance {
		logCheckBalance.Log(ctx, "CheckBalance", requestCheckBalance, httpResponseCheckBalance, time.Since(startCheckBalance))
	}

	if errCheckBalance != nil {
		return v0CheckBalance, fmt.Errorf("error sending 'CheckBalance' request: %w", errCheckBalance)
	}

	if httpResponseCheckBalance.StatusCode < 200 || httpResponseCheckBalance.StatusCode > 299 {
		return v0CheckBalance, __ClientNewResponseError("CheckBalance", httpResponseCheckBalance)
	}

	if errCheckBalance = responseCheckBalance.FromResponse("CheckBalance", httpResponseCheckBalance); errCheckBalance != nil {
		return v0CheckBalance, fmt.Errorf("error converting 'CheckBalance' response: %w", errCheckBalance)
	}

	addrCheckBalance.Reset()
	headerCheckBalance.Reset()

	if errCheckBalance = responseCheckBalance.Err(); errCheckBalance != nil {
		return v0CheckBalance, fmt.Errorf("error returned from 'CheckBalance' response: %w", errCheckBalance)
	}

	if errCheckBalance = responseCheckBalance.ScanValues(v0CheckBalance); errCheckBalance != nil {
		return v0CheckBalance, fmt.Errorf("error scanning value from 'CheckBalance' response: %w", errCheckBalance)
	}

	return v0CheckBalance, nil
}

func (__imp *implClient[C]) CreateChatCompletion(ctx context.Context, request *ChatCompletionRequest) (*Completion, error) {
	var innerCreateChatCompletion any = __imp.Inner()

	addrCreateChatCompletion := __ClientGetBuffer()
	defer __ClientPutBuffer(addrCreateChatCompletion)
	defer addrCreateChatCompletion.Reset()

	headerCreateChatCompletion := __ClientGetBuffer()
	defer __ClientPutBuffer(headerCreateChatCompletion)
	defer headerCreateChatCompletion.Reset()

	var (
		v0CreateChatCompletion           = new(Completion)
		errCreateChatCompletion          error
		httpResponseCreateChatCompletion *http.Response
		responseCreateChatCompletion     ClientResponseInterface = __imp.response()
	)

	if errCreateChatCompletion = addrTmplCreateChatCompletion.Execute(addrCreateChatCompletion, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error building 'CreateChatCompletion' url: %w", errCreateChatCompletion)
	}

	if errCreateChatCompletion = headerTmplCreateChatCompletion.Execute(headerCreateChatCompletion, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error building 'CreateChatCompletion' header: %w", errCreateChatCompletion)
	}
	bufReaderCreateChatCompletion := bufio.NewReader(headerCreateChatCompletion)
	mimeHeaderCreateChatCompletion, errCreateChatCompletion := textproto.NewReader(bufReaderCreateChatCompletion).ReadMIMEHeader()
	if errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error reading 'CreateChatCompletion' header: %w", errCreateChatCompletion)
	}

	urlCreateChatCompletion := addrCreateChatCompletion.String()
	requestBodyCreateChatCompletion, errCreateChatCompletion := io.ReadAll(bufReaderCreateChatCompletion)
	if errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error reading 'CreateChatCompletion' request body: %w", errCreateChatCompletion)
	}
	requestCreateChatCompletion, errCreateChatCompletion := http.NewRequestWithContext(ctx, "POST", urlCreateChatCompletion, bytes.NewReader(requestBodyCreateChatCompletion))
	if errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error building 'CreateChatCompletion' request: %w", errCreateChatCompletion)
	}

	for kCreateChatCompletion, vvCreateChatCompletion := range mimeHeaderCreateChatCompletion {
		for _, vCreateChatCompletion := range vvCreateChatCompletion {
			requestCreateChatCompletion.Header.Add(kCreateChatCompletion, vCreateChatCompletion)
		}
	}

	startCreateChatCompletion := time.Now()

	if httpClientCreateChatCompletion, okCreateChatCompletion := innerCreateChatCompletion.(interface{ Client() *http.Client }); okCreateChatCompletion {
		httpResponseCreateChatCompletion, errCreateChatCompletion = httpClientCreateChatCompletion.Client().Do(requestCreateChatCompletion)
	} else {
		httpResponseCreateChatCompletion, errCreateChatCompletion = http.DefaultClient.Do(requestCreateChatCompletion)
	}

	if logCreateChatCompletion, okCreateChatCompletion := innerCreateChatCompletion.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okCreateChatCompletion {
		logCreateChatCompletion.Log(ctx, "CreateChatCompletion", requestCreateChatCompletion, httpResponseCreateChatCompletion, time.Since(startCreateChatCompletion))
	}

	if errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error sending 'CreateChatCompletion' request: %w", errCreateChatCompletion)
	}

	if httpResponseCreateChatCompletion.StatusCode < 200 || httpResponseCreateChatCompletion.StatusCode > 299 {
		return v0CreateChatCompletion, __ClientNewResponseError("CreateChatCompletion", httpResponseCreateChatCompletion)
	}

	if errCreateChatCompletion = responseCreateChatCompletion.FromResponse("CreateChatCompletion", httpResponseCreateChatCompletion); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error converting 'CreateChatCompletion' response: %w", errCreateChatCompletion)
	}

	addrCreateChatCompletion.Reset()
	headerCreateChatCompletion.Reset()

	if errCreateChatCompletion = responseCreateChatCompletion.Err(); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error returned from 'CreateChatCompletion' response: %w", errCreateChatCompletion)
	}

	if errCreateChatCompletion = responseCreateChatCompletion.ScanValues(v0CreateChatCompletion); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error scanning value from 'CreateChatCompletion' response: %w", errCreateChatCompletion)
	}

	return v0CreateChatCompletion, nil
}

func (__imp *implClient[C]) CreateChatCompletionStream(ctx context.Context, request *ChatCompletionStreamRequest) (*Stream, error) {
	var innerCreateChatCompletionStream any = __imp.Inner()

	addrCreateChatCompletionStream := __ClientGetBuffer()
	defer __ClientPutBuffer(addrCreateChatCompletionStream)
	defer addrCreateChatCompletionStream.Reset()

	headerCreateChatCompletionStream := __ClientGetBuffer()
	defer __ClientPutBuffer(headerCreateChatCompletionStream)
	defer headerCreateChatCompletionStream.Reset()

	var (
		v0CreateChatCompletionStream           = new(Stream)
		errCreateChatCompletionStream          error
		httpResponseCreateChatCompletionStream *http.Response
		responseCreateChatCompletionStream     ClientResponseInterface = __imp.response()
	)

	if errCreateChatCompletionStream = addrTmplCreateChatCompletionStream.Execute(addrCreateChatCompletionStream, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error building 'CreateChatCompletionStream' url: %w", errCreateChatCompletionStream)
	}

	if errCreateChatCompletionStream = headerTmplCreateChatCompletionStream.Execute(headerCreateChatCompletionStream, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error building 'CreateChatCompletionStream' header: %w", errCreateChatCompletionStream)
	}
	bufReaderCreateChatCompletionStream := bufio.NewReader(headerCreateChatCompletionStream)
	mimeHeaderCreateChatCompletionStream, errCreateChatCompletionStream := textproto.NewReader(bufReaderCreateChatCompletionStream).ReadMIMEHeader()
	if errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error reading 'CreateChatCompletionStream' header: %w", errCreateChatCompletionStream)
	}

	urlCreateChatCompletionStream := addrCreateChatCompletionStream.String()
	requestBodyCreateChatCompletionStream, errCreateChatCompletionStream := io.ReadAll(bufReaderCreateChatCompletionStream)
	if errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error reading 'CreateChatCompletionStream' request body: %w", errCreateChatCompletionStream)
	}
	requestCreateChatCompletionStream, errCreateChatCompletionStream := http.NewRequestWithContext(ctx, "POST", urlCreateChatCompletionStream, bytes.NewReader(requestBodyCreateChatCompletionStream))
	if errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error building 'CreateChatCompletionStream' request: %w", errCreateChatCompletionStream)
	}

	for kCreateChatCompletionStream, vvCreateChatCompletionStream := range mimeHeaderCreateChatCompletionStream {
		for _, vCreateChatCompletionStream := range vvCreateChatCompletionStream {
			requestCreateChatCompletionStream.Header.Add(kCreateChatCompletionStream, vCreateChatCompletionStream)
		}
	}

	startCreateChatCompletionStream := time.Now()

	if httpClientCreateChatCompletionStream, okCreateChatCompletionStream := innerCreateChatCompletionStream.(interface{ Client() *http.Client }); okCreateChatCompletionStream {
		httpResponseCreateChatCompletionStream, errCreateChatCompletionStream = httpClientCreateChatCompletionStream.Client().Do(requestCreateChatCompletionStream)
	} else {
		httpResponseCreateChatCompletionStream, errCreateChatCompletionStream = http.DefaultClient.Do(requestCreateChatCompletionStream)
	}

	if logCreateChatCompletionStream, okCreateChatCompletionStream := innerCreateChatCompletionStream.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okCreateChatCompletionStream {
		logCreateChatCompletionStream.Log(ctx, "CreateChatCompletionStream", requestCreateChatCompletionStream, httpResponseCreateChatCompletionStream, time.Since(startCreateChatCompletionStream))
	}

	if errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error sending 'CreateChatCompletionStream' request: %w", errCreateChatCompletionStream)
	}

	if httpResponseCreateChatCompletionStream.StatusCode < 200 || httpResponseCreateChatCompletionStream.StatusCode > 299 {
		return v0CreateChatCompletionStream, __ClientNewResponseError("CreateChatCompletionStream", httpResponseCreateChatCompletionStream)
	}

	if errCreateChatCompletionStream = responseCreateChatCompletionStream.FromResponse("CreateChatCompletionStream", httpResponseCreateChatCompletionStream); errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error converting 'CreateChatCompletionStream' response: %w", errCreateChatCompletionStream)
	}

	addrCreateChatCompletionStream.Reset()
	headerCreateChatCompletionStream.Reset()

	if errCreateChatCompletionStream = responseCreateChatCompletionStream.Err(); errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error returned from 'CreateChatCompletionStream' response: %w", errCreateChatCompletionStream)
	}

	if errCreateChatCompletionStream = responseCreateChatCompletionStream.ScanValues(v0CreateChatCompletionStream); errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error scanning value from 'CreateChatCompletionStream' response: %w", errCreateChatCompletionStream)
	}

	return v0CreateChatCompletionStream, nil
}

func (__imp *implClient[C]) UploadFile(ctx context.Context, request *UploadFileRequest) (*File, error) {
	var innerUploadFile any = __imp.Inner()

	addrUploadFile := __ClientGetBuffer()
	defer __ClientPutBuffer(addrUploadFile)
	defer addrUploadFile.Reset()

	headerUploadFile := __ClientGetBuffer()
	defer __ClientPutBuffer(headerUploadFile)
	defer headerUploadFile.Reset()

	var (
		v0UploadFile           = new(File)
		errUploadFile          error
		httpResponseUploadFile *http.Response
		responseUploadFile     ClientResponseInterface = __imp.response()
	)

	if errUploadFile = addrTmplUploadFile.Execute(addrUploadFile, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error building 'UploadFile' url: %w", errUploadFile)
	}

	if errUploadFile = headerTmplUploadFile.Execute(headerUploadFile, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error building 'UploadFile' header: %w", errUploadFile)
	}
	bufReaderUploadFile := bufio.NewReader(headerUploadFile)
	mimeHeaderUploadFile, errUploadFile := textproto.NewReader(bufReaderUploadFile).ReadMIMEHeader()
	if errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error reading 'UploadFile' header: %w", errUploadFile)
	}

	urlUploadFile := addrUploadFile.String()
	requestUploadFile, errUploadFile := http.NewRequestWithContext(ctx, "POST", urlUploadFile, request)
	if errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error building 'UploadFile' request: %w", errUploadFile)
	}

	for kUploadFile, vvUploadFile := range mimeHeaderUploadFile {
		for _, vUploadFile := range vvUploadFile {
			requestUploadFile.Header.Add(kUploadFile, vUploadFile)
		}
	}

	startUploadFile := time.Now()

	if httpClientUploadFile, okUploadFile := innerUploadFile.(interface{ Client() *http.Client }); okUploadFile {
		httpResponseUploadFile, errUploadFile = httpClientUploadFile.Client().Do(requestUploadFile)
	} else {
		httpResponseUploadFile, errUploadFile = http.DefaultClient.Do(requestUploadFile)
	}

	if logUploadFile, okUploadFile := innerUploadFile.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okUploadFile {
		logUploadFile.Log(ctx, "UploadFile", requestUploadFile, httpResponseUploadFile, time.Since(startUploadFile))
	}

	if errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error sending 'UploadFile' request: %w", errUploadFile)
	}

	if httpResponseUploadFile.StatusCode < 200 || httpResponseUploadFile.StatusCode > 299 {
		return v0UploadFile, __ClientNewResponseError("UploadFile", httpResponseUploadFile)
	}

	if errUploadFile = responseUploadFile.FromResponse("UploadFile", httpResponseUploadFile); errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error converting 'UploadFile' response: %w", errUploadFile)
	}

	addrUploadFile.Reset()
	headerUploadFile.Reset()

	if errUploadFile = responseUploadFile.Err(); errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error returned from 'UploadFile' response: %w", errUploadFile)
	}

	if errUploadFile = responseUploadFile.ScanValues(v0UploadFile); errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error scanning value from 'UploadFile' response: %w", errUploadFile)
	}

	return v0UploadFile, nil
}

func (__imp *implClient[C]) RetrieveFileContent(ctx context.Context, fileID string) ([]byte, error) {
	var innerRetrieveFileContent any = __imp.Inner()

	addrRetrieveFileContent := __ClientGetBuffer()
	defer __ClientPutBuffer(addrRetrieveFileContent)
	defer addrRetrieveFileContent.Reset()

	headerRetrieveFileContent := __ClientGetBuffer()
	defer __ClientPutBuffer(headerRetrieveFileContent)
	defer headerRetrieveFileContent.Reset()

	var (
		v0RetrieveFileContent           = __ClientNew[[]byte]()
		errRetrieveFileContent          error
		httpResponseRetrieveFileContent *http.Response
		responseRetrieveFileContent     ClientResponseInterface = __imp.response()
	)

	if errRetrieveFileContent = addrTmplRetrieveFileContent.Execute(addrRetrieveFileContent, map[string]any{
		"Client": __imp.Inner(),
		"ctx":    ctx,
		"fileID": fileID,
	}); errRetrieveFileContent != nil {
		return v0RetrieveFileContent, fmt.Errorf("error building 'RetrieveFileContent' url: %w", errRetrieveFileContent)
	}

	if errRetrieveFileContent = headerTmplRetrieveFileContent.Execute(headerRetrieveFileContent, map[string]any{
		"Client": __imp.Inner(),
		"ctx":    ctx,
		"fileID": fileID,
	}); errRetrieveFileContent != nil {
		return v0RetrieveFileContent, fmt.Errorf("error building 'RetrieveFileContent' header: %w", errRetrieveFileContent)
	}
	bufReaderRetrieveFileContent := bufio.NewReader(headerRetrieveFileContent)
	mimeHeaderRetrieveFileContent, errRetrieveFileContent := textproto.NewReader(bufReaderRetrieveFileContent).ReadMIMEHeader()
	if errRetrieveFileContent != nil {
		return v0RetrieveFileContent, fmt.Errorf("error reading 'RetrieveFileContent' header: %w", errRetrieveFileContent)
	}

	urlRetrieveFileContent := addrRetrieveFileContent.String()
	requestRetrieveFileContent, errRetrieveFileContent := http.NewRequestWithContext(ctx, "GET", urlRetrieveFileContent, http.NoBody)
	if errRetrieveFileContent != nil {
		return v0RetrieveFileContent, fmt.Errorf("error building 'RetrieveFileContent' request: %w", errRetrieveFileContent)
	}

	for kRetrieveFileContent, vvRetrieveFileContent := range mimeHeaderRetrieveFileContent {
		for _, vRetrieveFileContent := range vvRetrieveFileContent {
			requestRetrieveFileContent.Header.Add(kRetrieveFileContent, vRetrieveFileContent)
		}
	}

	startRetrieveFileContent := time.Now()

	if httpClientRetrieveFileContent, okRetrieveFileContent := innerRetrieveFileContent.(interface{ Client() *http.Client }); okRetrieveFileContent {
		httpResponseRetrieveFileContent, errRetrieveFileContent = httpClientRetrieveFileContent.Client().Do(requestRetrieveFileContent)
	} else {
		httpResponseRetrieveFileContent, errRetrieveFileContent = http.DefaultClient.Do(requestRetrieveFileContent)
	}

	if logRetrieveFileContent, okRetrieveFileContent := innerRetrieveFileContent.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okRetrieveFileContent {
		logRetrieveFileContent.Log(ctx, "RetrieveFileContent", requestRetrieveFileContent, httpResponseRetrieveFileContent, time.Since(startRetrieveFileContent))
	}

	if errRetrieveFileContent != nil {
		return v0RetrieveFileContent, fmt.Errorf("error sending 'RetrieveFileContent' request: %w", errRetrieveFileContent)
	}

	if httpResponseRetrieveFileContent.StatusCode < 200 || httpResponseRetrieveFileContent.StatusCode > 299 {
		return v0RetrieveFileContent, __ClientNewResponseError("RetrieveFileContent", httpResponseRetrieveFileContent)
	}

	if errRetrieveFileContent = responseRetrieveFileContent.FromResponse("RetrieveFileContent", httpResponseRetrieveFileContent); errRetrieveFileContent != nil {
		return v0RetrieveFileContent, fmt.Errorf("error converting 'RetrieveFileContent' response: %w", errRetrieveFileContent)
	}

	addrRetrieveFileContent.Reset()
	headerRetrieveFileContent.Reset()

	if errRetrieveFileContent = responseRetrieveFileContent.Err(); errRetrieveFileContent != nil {
		return v0RetrieveFileContent, fmt.Errorf("error returned from 'RetrieveFileContent' response: %w", errRetrieveFileContent)
	}

	if errRetrieveFileContent = responseRetrieveFileContent.ScanValues(&v0RetrieveFileContent); errRetrieveFileContent != nil {
		return v0RetrieveFileContent, fmt.Errorf("error scanning value from 'RetrieveFileContent' response: %w", errRetrieveFileContent)
	}

	return v0RetrieveFileContent, nil
}

func (__imp *implClient[C]) Inner() C {
	return __imp.__Client
}

func (*implClient[C]) response() *ResponseHandler {
	return new(ResponseHandler)
}

var __ClientBufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func __ClientGetBuffer() *bytes.Buffer {
	return __ClientBufferPool.Get().(*bytes.Buffer)
}

func __ClientPutBuffer(buffer *bytes.Buffer) {
	__ClientBufferPool.Put(buffer)
}

type ClientResponseInterface interface {
	Err() error
	ScanValues(...any) error
	FromResponse(string, *http.Response) error
	Break() bool
}

func __ClientNewType(typ reflect.Type) reflect.Value {
	switch typ.Kind() {
	case reflect.Slice:
		return reflect.MakeSlice(typ, 0, 0)
	case reflect.Map:
		return reflect.MakeMap(typ)
	case reflect.Chan:
		return reflect.MakeChan(typ, 0)
	case reflect.Func:
		return reflect.MakeFunc(typ, func(_ []reflect.Value) (results []reflect.Value) {
			results = make([]reflect.Value, typ.NumOut())
			for i := 0; i < typ.NumOut(); i++ {
				results[i] = __ClientNewType(typ.Out(i))
			}
			return results
		})
	case reflect.Pointer:
		return reflect.New(typ.Elem())
	default:
		return reflect.Zero(typ)
	}
}

func __ClientNew[T any]() (v T) {
	val := reflect.ValueOf(&v).Elem()
	switch val.Kind() {
	case reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Pointer:
		val.Set(__ClientNewType(val.Type()))
	}
	return v
}

// ClientResponseErrorInterface represents future Response error interface which would
// be used in next major version of defc, who may cause breaking changes.
//
// generated with --features=api/future
type ClientResponseErrorInterface interface {
	error
	Response() *http.Response
}

func __ClientNewResponseError(caller string, response *http.Response) ClientResponseErrorInterface {
	return &__ClientImplResponseError{
		caller:   caller,
		response: response,
	}
}

type __ClientImplResponseError struct {
	caller   string
	response *http.Response
}

func (e *__ClientImplResponseError) Error() string {
	return fmt.Sprintf("response status code %d for '%s'", e.response.StatusCode, e.caller)
}

func (e *__ClientImplResponseError) Response() *http.Response {
	return e.response
}
